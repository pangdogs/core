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
	"context"
	"log"
	"testing"
	"time"

	"git.golaxy.org/core"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
)

type ComponentTestFrameUpdate struct {
	ec.ComponentBehavior
}

func (c *ComponentTestFrameUpdate) Update() {
	frame := runtime.Current(c).Frame()
	log.Printf("Component %s.%s Update, fps: %.2f", c.Entity().Id(), c.Name(), frame.CurFPS())
}

func (c *ComponentTestFrameUpdate) LateUpdate() {
	log.Printf("Component %s.%s LateUpdate", c.Entity().Id(), c.Name())
}

func Test_CreateEntityFrameUpdate(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	svcCtx := service.NewContext(
		service.With.Context(ctx),
		service.With.RunningEventCB(func(ctx service.Context, runningEvent service.RunningEvent, args ...any) {
			switch runningEvent {
			case service.RunningEvent_Birth:
				core.BuildEntityPT(ctx, "Test1").
					AddComponent(ComponentTestFrameUpdate{}).
					Declare()
			case service.RunningEvent_Started:
				core.NewRuntime(
					runtime.NewContext(ctx,
						runtime.With.RunningEventCB(func(ctx runtime.Context, runningEvent runtime.RunningEvent, args ...any) {
							switch runningEvent {
							case runtime.RunningEvent_Started:
								for range 10 {
									core.BuildEntity(ctx, "Test1").New()
								}
							}
							log.Println("runtime event:", runningEvent, args)
						}),
					),
					core.With.Runtime.AutoRun(true),
				)
			}
			log.Println("service event:", runningEvent, args)
		}),
	)

	<-core.NewService(svcCtx).Run().Done()
}

type ComponentTestStressFrameUpdate struct {
	ec.ComponentBehavior
	count int
}

func (c *ComponentTestStressFrameUpdate) Update() {
	c.count++
}

func (c *ComponentTestStressFrameUpdate) LateUpdate() {
	c.count++
}

func Test_CreateEntityStressFrameUpdate(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	svcCtx := service.NewContext(
		service.With.Context(ctx),
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
							case runtime.RunningEvent_FrameLoopBegin:
								for range 200 {
									core.BuildEntity(ctx, "Test1").New()
								}
							case runtime.RunningEvent_RunGCBegin:
								log.Printf("fps: %.2f, running_elapse_time: %.3f, last_loop_elapse_time: %.3f, entities: %d",
									ctx.Frame().CurFPS(),
									ctx.Frame().RunningElapseTime().Seconds(),
									ctx.Frame().LastLoopElapseTime().Seconds(),
									ctx.EntityManager().CountEntities())
							}
						}),
					),
					core.With.Runtime.AutoRun(true),
				)
			}
		}),
	)

	<-core.NewService(svcCtx).Run().Done()
}
