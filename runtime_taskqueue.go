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

package core

import (
	"context"
	"fmt"
	"sync/atomic"

	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/generic"
)

var (
	ErrTaskQueueClosed = fmt.Errorf("%w: task queue is closed", ErrRuntime) // 任务队列已关闭。
	ErrTaskQueueFull   = fmt.Errorf("%w: task queue is full", ErrRuntime)   // 任务队列已满。
)

type _TaskQueueStats struct {
	accepted       atomic.Int64
	queued         atomic.Int64
	running        atomic.Int64
	completed      atomic.Int64
	canceled       atomic.Int64
	panicked       atomic.Int64
	rejectedClosed atomic.Int64
	rejectedFull   atomic.Int64
}

type _TaskQueue struct {
	barrier       generic.Barrier
	boundedChan   chan _Task
	unboundedChan *generic.UnboundedChannel[_Task]
	stats         [taskTypeCount]_TaskQueueStats
}

func (q *_TaskQueue) init(unbounded bool, capacity int) {
	if unbounded {
		q.unboundedChan = generic.NewUnboundedChannel[_Task]()
	} else {
		q.boundedChan = make(chan _Task, capacity)
	}
}

func (q *_TaskQueue) enqueueSubmit(
	executorID async.ExecutorID,
	fun generic.FuncVar1[runtime.Context, any, async.Result],
	action generic.ActionVar1[runtime.Context, any],
	delegate generic.DelegateVar1[runtime.Context, any, async.Result],
	delegateVoid generic.DelegateVoidVar1[runtime.Context, any],
	args []any,
) async.Future {
	promise, future := async.NewPromise(executorID)
	task := _Task{
		typ:          TaskType_Submit,
		fun:          fun,
		action:       action,
		delegate:     delegate,
		delegateVoid: delegateVoid,
		args:         args,
		promise:      promise,
	}
	if err := q.tryEnqueue(task); err != nil {
		promise.Resolve(async.NewResult(nil, err))
	}
	return future
}

func (q *_TaskQueue) enqueuePost(
	action generic.ActionVar1[runtime.Context, any],
	delegateVoid generic.DelegateVoidVar1[runtime.Context, any],
	args []any,
) error {
	return q.tryEnqueue(_Task{
		typ:          TaskType_Post,
		action:       action,
		delegateVoid: delegateVoid,
		args:         args,
	})
}

func (q *_TaskQueue) enqueueFrame(ctx context.Context, action generic.ActionVar1[runtime.Context, any], done chan struct{}) bool {
	task := _Task{typ: TaskType_Frame, action: action, done: done}
	stats := &q.stats[TaskType_Frame]

	if q.boundedChan != nil {
		stats.queued.Add(1)
		select {
		case q.boundedChan <- task:
			stats.accepted.Add(1)
			select {
			case <-done:
				return true
			case <-ctx.Done():
				stats.canceled.Add(1)
				return false
			}
		case <-ctx.Done():
			stats.queued.Add(-1)
			stats.rejectedClosed.Add(1)
			return false
		}
	}

	if q.unboundedChan != nil {
		stats.queued.Add(1)
		q.unboundedChan.In() <- task
		stats.accepted.Add(1)
		select {
		case <-done:
			return true
		case <-ctx.Done():
			stats.canceled.Add(1)
			return false
		}
	}

	stats.rejectedClosed.Add(1)
	return false
}

func (q *_TaskQueue) tryEnqueue(task _Task) error {
	stats := &q.stats[task.typ]
	if !q.barrier.Join(1) {
		stats.rejectedClosed.Add(1)
		return ErrTaskQueueClosed
	}
	defer q.barrier.Done()

	if q.boundedChan != nil {
		stats.queued.Add(1)
		select {
		case q.boundedChan <- task:
			stats.accepted.Add(1)
			return nil
		default:
			stats.queued.Add(-1)
			stats.rejectedFull.Add(1)
			return ErrTaskQueueFull
		}
	}

	if q.unboundedChan != nil {
		stats.queued.Add(1)
		q.unboundedChan.In() <- task
		stats.accepted.Add(1)
		return nil
	}

	stats.rejectedClosed.Add(1)
	return ErrTaskQueueClosed
}

func (q *_TaskQueue) out() <-chan _Task {
	if q.boundedChan != nil {
		return q.boundedChan
	}
	if q.unboundedChan != nil {
		return q.unboundedChan.Out()
	}
	return nil
}

func (q *_TaskQueue) start(typ TaskType) {
	q.stats[typ].queued.Add(-1)
	q.stats[typ].running.Add(1)
}

func (q *_TaskQueue) complete(typ TaskType, panicked bool) {
	q.stats[typ].running.Add(-1)
	q.stats[typ].completed.Add(1)
	if panicked {
		q.stats[typ].panicked.Add(1)
	}
}

func (q *_TaskQueue) close() {
	q.barrier.Close()
	q.barrier.Wait()
	if q.boundedChan != nil {
		close(q.boundedChan)
	}
	if q.unboundedChan != nil {
		q.unboundedChan.Close()
	}
}
