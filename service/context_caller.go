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

package service

import (
	"fmt"
	_ "unsafe"

	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/uid"
)

// Caller 通过全局 Entity ID 把任务投递到实体所属 Runtime。
type Caller interface {
	Submit(entityID uid.Id, fun generic.FuncVar1[ec.Entity, any, async.Result], args ...any) async.Future
	SubmitDelegate(entityID uid.Id, fun generic.DelegateVar1[ec.Entity, any, async.Result], args ...any) async.Future
	SubmitVoid(entityID uid.Id, fun generic.ActionVar1[ec.Entity, any], args ...any) async.Future
	SubmitDelegateVoid(entityID uid.Id, fun generic.DelegateVoidVar1[ec.Entity, any], args ...any) async.Future
	Post(entityID uid.Id, fun generic.ActionVar1[ec.Entity, any], args ...any) error
	PostDelegate(entityID uid.Id, fun generic.DelegateVoidVar1[ec.Entity, any], args ...any) error
}

//go:linkname submit git.golaxy.org/core/runtime.submit
func submit(entity ec.ConcurrentEntity, fun generic.FuncVar1[ec.Entity, any, async.Result], args ...any) async.Future

//go:linkname submitDelegate git.golaxy.org/core/runtime.submitDelegate
func submitDelegate(entity ec.ConcurrentEntity, fun generic.DelegateVar1[ec.Entity, any, async.Result], args ...any) async.Future

//go:linkname submitVoid git.golaxy.org/core/runtime.submitVoid
func submitVoid(entity ec.ConcurrentEntity, fun generic.ActionVar1[ec.Entity, any], args ...any) async.Future

//go:linkname submitDelegateVoid git.golaxy.org/core/runtime.submitDelegateVoid
func submitDelegateVoid(entity ec.ConcurrentEntity, fun generic.DelegateVoidVar1[ec.Entity, any], args ...any) async.Future

//go:linkname post git.golaxy.org/core/runtime.post
func post(entity ec.ConcurrentEntity, fun generic.ActionVar1[ec.Entity, any], args ...any) error

//go:linkname postDelegate git.golaxy.org/core/runtime.postDelegate
func postDelegate(entity ec.ConcurrentEntity, fun generic.DelegateVoidVar1[ec.Entity, any], args ...any) error

func (ctx *ContextBehavior) Submit(entityID uid.Id, fun generic.FuncVar1[ec.Entity, any, async.Result], args ...any) async.Future {
	entity, err := ctx.getEntity(entityID)
	if err != nil {
		return async.Rejected(err)
	}
	return submit(entity, fun, args...)
}

func (ctx *ContextBehavior) SubmitDelegate(entityID uid.Id, fun generic.DelegateVar1[ec.Entity, any, async.Result], args ...any) async.Future {
	entity, err := ctx.getEntity(entityID)
	if err != nil {
		return async.Rejected(err)
	}
	return submitDelegate(entity, fun, args...)
}

func (ctx *ContextBehavior) SubmitVoid(entityID uid.Id, fun generic.ActionVar1[ec.Entity, any], args ...any) async.Future {
	entity, err := ctx.getEntity(entityID)
	if err != nil {
		return async.Rejected(err)
	}
	return submitVoid(entity, fun, args...)
}

func (ctx *ContextBehavior) SubmitDelegateVoid(entityID uid.Id, fun generic.DelegateVoidVar1[ec.Entity, any], args ...any) async.Future {
	entity, err := ctx.getEntity(entityID)
	if err != nil {
		return async.Rejected(err)
	}
	return submitDelegateVoid(entity, fun, args...)
}

// Post 只报告实体查询和入队错误；任务执行前实体已经失活时会静默丢弃。
func (ctx *ContextBehavior) Post(entityID uid.Id, fun generic.ActionVar1[ec.Entity, any], args ...any) error {
	entity, err := ctx.getEntity(entityID)
	if err != nil {
		return err
	}
	return post(entity, fun, args...)
}

// PostDelegate 是 Post 的 DelegateVoid 版本。
func (ctx *ContextBehavior) PostDelegate(entityID uid.Id, fun generic.DelegateVoidVar1[ec.Entity, any], args ...any) error {
	entity, err := ctx.getEntity(entityID)
	if err != nil {
		return err
	}
	return postDelegate(entity, fun, args...)
}

func (ctx *ContextBehavior) getEntity(id uid.Id) (ec.ConcurrentEntity, error) {
	entity, ok := ctx.entityManager.GetEntity(id)
	if !ok {
		return nil, fmt.Errorf("%w: entity not exist", ErrContext)
	}
	return entity, nil
}
