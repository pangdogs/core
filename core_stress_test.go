//go:build stress

/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

package core_test

import (
	"flag"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"git.golaxy.org/core"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
)

const stressEntityBatchSize = int64(200)

var stressDuration = flag.Duration("stress.duration", 120*time.Second, "duration of tagged stress tests")

type stressFrameCounters struct {
	frames       int64
	entities     int64
	updates      int64
	lateUpdates  int64
	peakEntities int64
}

var activeStressFrameCounters atomic.Pointer[stressFrameCounters]

type ComponentTestStressFrameUpdate struct {
	ec.ComponentBehavior
	counters *stressFrameCounters
}

func (c *ComponentTestStressFrameUpdate) Awake() {
	c.counters = activeStressFrameCounters.Load()
}

func (c *ComponentTestStressFrameUpdate) Update() {
	if c.counters != nil {
		c.counters.updates++
	}
}

func (c *ComponentTestStressFrameUpdate) LateUpdate() {
	if c.counters != nil {
		c.counters.lateUpdates++
	}
}

func Test_CreateEntityStressFrameUpdate(t *testing.T) {
	if *stressDuration <= 0 {
		t.Fatalf("stress.duration must be greater than zero: %s", *stressDuration)
	}

	scenario := newCoreTestScenario(*stressDuration + 15*time.Second)
	counters := &stressFrameCounters{}
	if !activeStressFrameCounters.CompareAndSwap(nil, counters) {
		t.Fatal("another frame stress test is already active")
	}
	defer activeStressFrameCounters.CompareAndSwap(counters, nil)

	var (
		startedAt     time.Time
		stopRequested atomic.Bool
	)

	svcCtx := service.NewContext(
		service.With.Context(scenario.ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Birth:
				core.BuildEntityPT(ctx, "Test1").
					AddComponent(ComponentTestStressFrameUpdate{}).
					Declare()
			case service.RunningEvent_Started:
				core.NewRuntime(
					runtime.NewContext(ctx,
						runtime.With.RunningEventCB(func(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
							switch runningEvent {
							case runtime.RunningEvent_Started:
								startedAt = time.Now()
								go func() {
									timer := time.NewTimer(*stressDuration)
									defer timer.Stop()
									select {
									case <-timer.C:
										stopRequested.Store(true)
									case <-scenario.ctx.Done():
									}
								}()
							case runtime.RunningEvent_FrameLoopBegin:
								counters.frames++
								frame := counters.frames
								for i := int64(0); i < stressEntityBatchSize; i++ {
									if _, err := core.BuildEntity(ctx, "Test1").New(); err != nil {
										scenario.complete(fmt.Errorf("create entity in frame %d: %w", frame, err))
										return
									}
									counters.entities++
								}
								managedEntities := int64(ctx.EntityManager().CountEntities())
								if managedEntities > counters.peakEntities {
									counters.peakEntities = managedEntities
								}
							case runtime.RunningEvent_FrameUpdateEnd:
								if stopRequested.Load() {
									scenario.complete(nil)
								}
							}
						}),
					),
					core.With.Runtime.AutoRun(true),
				)
			}
		}),
	)

	scenario.run(t, svcCtx)

	frames := counters.frames
	entities := counters.entities
	updates := counters.updates
	lateUpdates := counters.lateUpdates
	peakEntities := counters.peakEntities

	if frames <= 0 {
		t.Fatal("stress runtime completed without executing a frame")
	}
	if want := frames * stressEntityBatchSize; entities != want {
		t.Errorf("created entities: got %d, want %d", entities, want)
	}
	if peakEntities != entities {
		t.Errorf("peak managed entities: got %d, want %d", peakEntities, entities)
	}
	wantUpdates := stressEntityBatchSize * frames * (frames + 1) / 2
	if updates != wantUpdates {
		t.Errorf("Update callbacks: got %d, want %d", updates, wantUpdates)
	}
	if lateUpdates != wantUpdates {
		t.Errorf("LateUpdate callbacks: got %d, want %d", lateUpdates, wantUpdates)
	}

	elapsed := time.Since(startedAt)
	t.Logf("duration=%s frames=%d avg_fps=%.2f entities=%d updates=%d",
		elapsed.Round(time.Millisecond),
		frames,
		float64(frames)/elapsed.Seconds(),
		entities,
		updates,
	)
}
