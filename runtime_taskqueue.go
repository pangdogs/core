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

	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/generic"
)

var (
	ErrTaskQueueClosed = fmt.Errorf("%w: task queue is closed", ErrRuntime) // 任务处理流水线关闭
	ErrTaskQueueFull   = fmt.Errorf("%w: task queue is full", ErrRuntime)   // 任务处理流水线已满
)

type _TaskQueueStats struct {
	enqueued  atomic.Int64
	pending   atomic.Int64
	rejected  atomic.Int64
	completed atomic.Int64
}

type _TaskQueue struct {
	boundedChan   chan _Task
	unboundedChan *generic.UnboundedChannel[_Task]
	stats         [2]_TaskQueueStats
}

func (q *_TaskQueue) init(unbounded bool, capacity int) {
	if unbounded {
		q.unboundedChan = generic.NewUnboundedChannel[_Task]()
	} else {
		q.boundedChan = make(chan _Task, capacity)
	}
}

func (q *_TaskQueue) enqueueCall(fun generic.FuncVar0[any, async.Ret], action generic.ActionVar0[any], delegate generic.DelegateVar0[any, async.Ret], delegateVoid generic.DelegateVoidVar0[any], args []any) (asyncRet async.AsyncRet) {
	task := _Task{
		typ:          TaskType_Call,
		fun:          fun,
		action:       action,
		delegate:     delegate,
		delegateVoid: delegateVoid,
		args:         args,
		asyncRet:     async.NewAsyncRet(),
	}

	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			q.stats[TaskType_Call].rejected.Add(1)
			asyncRet = async.Return(task.asyncRet, async.NewRet(nil, ErrTaskQueueClosed))
		}
	}()

	q.stats[TaskType_Call].enqueued.Add(1)

	if q.boundedChan != nil {
		select {
		case q.boundedChan <- task:
			q.stats[TaskType_Call].pending.Add(1)
			return task.asyncRet
		default:
			q.stats[TaskType_Call].rejected.Add(1)
			return async.Return(task.asyncRet, async.NewRet(nil, ErrTaskQueueFull))
		}
	}

	if q.unboundedChan != nil {
		q.unboundedChan.In() <- task
		q.stats[TaskType_Call].pending.Add(1)
		return task.asyncRet
	}

	q.stats[TaskType_Call].rejected.Add(1)
	return async.Return(task.asyncRet, async.NewRet(nil, ErrTaskQueueClosed))
}

func (q *_TaskQueue) enqueueFrame(ctx context.Context, action generic.ActionVar0[any], done chan struct{}) bool {
	task := _Task{
		typ:    TaskType_Frame,
		action: action,
		done:   done,
	}

	q.stats[TaskType_Frame].enqueued.Add(1)

	if q.boundedChan != nil {
		select {
		case q.boundedChan <- task:
			q.stats[TaskType_Frame].pending.Add(1)
			select {
			case <-done:
				return true
			case <-ctx.Done():
				q.stats[TaskType_Frame].rejected.Add(1)
				return false
			}
		case <-ctx.Done():
			q.stats[TaskType_Frame].rejected.Add(1)
			return false
		}
	}

	if q.unboundedChan != nil {
		q.unboundedChan.In() <- task
		q.stats[TaskType_Frame].pending.Add(1)
		select {
		case <-done:
			return true
		case <-ctx.Done():
			q.stats[TaskType_Frame].rejected.Add(1)
			return false
		}
	}

	q.stats[TaskType_Frame].rejected.Add(1)
	return false
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

func (q *_TaskQueue) complete(typ TaskType) {
	q.stats[typ].pending.Add(-1)
	q.stats[typ].completed.Add(1)
}

func (q *_TaskQueue) close() {
	if q.boundedChan != nil {
		close(q.boundedChan)
	}
	if q.unboundedChan != nil {
		q.unboundedChan.Close()
	}
}
