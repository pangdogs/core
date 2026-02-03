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

type _TaskQueueBehavior struct {
	boundedChan    chan _Task
	unboundedChan  *generic.UnboundedChannel[_Task]
	callEnqueued   atomic.Int64
	callCompleted  atomic.Int64
	frameEnqueued  atomic.Int64
	frameCompleted atomic.Int64
}

func (q *_TaskQueueBehavior) init(unbounded bool, capacity int) {
	if unbounded {
		q.unboundedChan = generic.NewUnboundedChannel[_Task]()
	} else {
		q.boundedChan = make(chan _Task, capacity)
	}
}

func (q *_TaskQueueBehavior) pushCall(fun generic.FuncVar0[any, async.Ret], action generic.ActionVar0[any], delegate generic.DelegateVar0[any, async.Ret], delegateVoid generic.DelegateVoidVar0[any], args []any) (asyncRet async.AsyncRet) {
	task := _Task{
		typ:          _TaskType_Call,
		fun:          fun,
		action:       action,
		delegate:     delegate,
		delegateVoid: delegateVoid,
		args:         args,
		asyncRet:     async.NewAsyncRet(),
	}

	defer func() {
		if panicInfo := recover(); panicInfo != nil {
			asyncRet = async.Return(task.asyncRet, async.NewRet(nil, ErrTaskQueueClosed))
		}
	}()

	if q.boundedChan != nil {
		select {
		case q.boundedChan <- task:
			q.callEnqueued.Add(1)
			return task.asyncRet
		default:
			return async.Return(task.asyncRet, async.NewRet(nil, ErrTaskQueueFull))
		}
	}

	if q.unboundedChan != nil {
		q.unboundedChan.In() <- task
		q.callEnqueued.Add(1)
		return task.asyncRet
	}

	return async.Return(task.asyncRet, async.NewRet(nil, ErrTaskQueueClosed))
}

func (q *_TaskQueueBehavior) pushFrame(ctx context.Context, action generic.ActionVar0[any], done chan struct{}) bool {
	task := _Task{
		typ:    _TaskType_Frame,
		action: action,
		done:   done,
	}

	if q.boundedChan != nil {
		select {
		case q.boundedChan <- task:
			q.frameEnqueued.Add(1)
			select {
			case <-done:
				return true
			case <-ctx.Done():
				return false
			}
		case <-ctx.Done():
			return false
		}
	}

	if q.unboundedChan != nil {
		q.unboundedChan.In() <- task
		q.frameEnqueued.Add(1)
		select {
		case <-done:
			return true
		case <-ctx.Done():
			return false
		}
	}

	return false
}

func (q *_TaskQueueBehavior) out() <-chan _Task {
	if q.boundedChan != nil {
		return q.boundedChan
	}
	if q.unboundedChan != nil {
		return q.unboundedChan.Out()
	}
	return nil
}

func (q *_TaskQueueBehavior) complete(typ _TaskType) {
	switch typ {
	case _TaskType_Call:
		q.callCompleted.Add(1)
	case _TaskType_Frame:
		q.frameCompleted.Add(1)
	}
}

func (q *_TaskQueueBehavior) close() {
	if q.boundedChan != nil {
		close(q.boundedChan)
	}
	if q.unboundedChan != nil {
		q.unboundedChan.Close()
	}
}
