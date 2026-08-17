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
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/generic"
)

func (rt *RuntimeBehavior) Submit(fun generic.FuncVar1[runtime.Context, any, async.Result], args ...any) async.Future {
	return rt.taskQueue.enqueueSubmit(rt.ctx.ExecutorID(), fun, nil, nil, nil, args)
}

func (rt *RuntimeBehavior) SubmitDelegate(fun generic.DelegateVar1[runtime.Context, any, async.Result], args ...any) async.Future {
	return rt.taskQueue.enqueueSubmit(rt.ctx.ExecutorID(), nil, nil, fun, nil, args)
}

func (rt *RuntimeBehavior) SubmitVoid(fun generic.ActionVar1[runtime.Context, any], args ...any) async.Future {
	return rt.taskQueue.enqueueSubmit(rt.ctx.ExecutorID(), nil, fun, nil, nil, args)
}

func (rt *RuntimeBehavior) SubmitDelegateVoid(fun generic.DelegateVoidVar1[runtime.Context, any], args ...any) async.Future {
	return rt.taskQueue.enqueueSubmit(rt.ctx.ExecutorID(), nil, nil, nil, fun, args)
}

func (rt *RuntimeBehavior) Post(fun generic.ActionVar1[runtime.Context, any], args ...any) error {
	return rt.taskQueue.enqueuePost(fun, nil, args)
}

func (rt *RuntimeBehavior) PostDelegate(fun generic.DelegateVoidVar1[runtime.Context, any], args ...any) error {
	return rt.taskQueue.enqueuePost(nil, fun, args)
}
