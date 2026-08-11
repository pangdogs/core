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

package ec

import (
	"context"
	"fmt"

	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
)

// ConcurrentEntity 暴露可跨协程安全读取的实体信息与上下文。
//
// 组件管理、实体树和 Destroy 等操作仍须通过所属 Runtime 的运行协程执行。
type ConcurrentEntity interface {
	iConcurrentEntity
	iEntityContext
	corectx.ConcurrentContextProvider
	corectx.AsyncScopeProvider
	fmt.Stringer

	// Id 返回实体 ID。
	Id() uid.Id
	// PT 返回实体原型。
	PT() EntityPT
}

// iEntityContext 将实体生命周期作为 context.Context 暴露，并提供最终销毁完成通知。
// Context 的 Done 在实体进入 Dead 时关闭；Terminated 在实体进入 Destroyed 并完成清理后兑现。
type iEntityContext interface {
	context.Context

	// Terminated 返回实体完成销毁时兑现的 Signal。
	Terminated() async.Signal
}

// runtimeContext 汇集 Entity 绑定所属 Runtime 时需要的最小能力。
// 使用结构接口避免 ec 反向依赖 runtime 包。
type runtimeContext interface {
	context.Context
	corectx.CurrentContextProvider
	async.WaitGuard
}

type iConcurrentEntity interface {
	getInstance() Entity
	setContext(parent runtimeContext)
}

// ConcurrentContextCache 返回实体所属 Runtime 的并发上下文接口缓存。
func (entity *EntityBehavior) ConcurrentContextCache() iface.Cache {
	return entity.runtimeCtxCache
}

// AsyncScope 返回绑定实体生命周期的后台任务作用域。
func (entity *EntityBehavior) AsyncScope() *async.Scope {
	return entity.asyncScope
}

// Terminated 返回实体进入 Destroyed 状态时兑现的 Signal。
func (entity *EntityBehavior) Terminated() async.Signal {
	return entity.terminated.Signal()
}

// BeforeFutureWait 把 Entity 作为等待 Context 时的检查转交给所属 Runtime。
func (entity *EntityBehavior) BeforeFutureWait(futureID uint64, completionExecutorID async.ExecutorID) error {
	if entity.waitGuard == nil {
		return nil
	}
	return entity.waitGuard.BeforeFutureWait(futureID, completionExecutorID)
}

// AfterFutureWait 清理所属 Runtime 的等待诊断状态。
func (entity *EntityBehavior) AfterFutureWait(futureID uint64) {
	if entity.waitGuard != nil {
		entity.waitGuard.AfterFutureWait(futureID)
	}
}

// String 返回包含实体 ID 与原型名的 JSON 文本；首次调用应发生在实体进入 Runtime 后。
func (entity *EntityBehavior) String() string {
	if cached := entity.stringerCache.Load(); cached != nil {
		return *cached
	}

	value := fmt.Sprintf(`{"id":%q,"prototype":%q}`, entity.Id(), entity.PT().Prototype())
	if entity.stringerCache.CompareAndSwap(nil, &value) {
		return value
	}
	return *entity.stringerCache.Load()
}

func (entity *EntityBehavior) getInstance() Entity {
	return entity.options.InstanceFace.Iface
}

func (entity *EntityBehavior) setContext(parent runtimeContext) {
	if entity.asyncScope != nil {
		return
	}
	entity.asyncScope = async.NewScope(parent)
	entity.Context = entity.asyncScope.Context()
	entity.runtimeCtxCache = parent.CurrentContextCache()
	entity.waitGuard = parent
	entity.terminated, _ = async.NewSignal()
	entity.componentList.TraversalEach(func(slot *generic.FreeSlot[Component]) {
		slot.V.setContext(entity.getInstance())
	})
}
