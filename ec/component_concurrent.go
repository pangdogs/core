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
	"fmt"

	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
)

// ConcurrentComponent 暴露可跨 goroutine 使用的组件身份、字符串表示、并发上下文和生命周期 Scope。
//
// 组件随 Entity 成功加入 Runtime，或动态加入运行中 Entity 后，才能跨 goroutine 使用
// 此接口；提前调用依赖 Runtime Context 的方法属于未定义行为。AsyncScope 在所属
// Entity Context 绑定前返回 nil，String 在 Runtime 完成组件身份初始化前返回空字符串。
//
// 该视图不暴露 State、Enabled、Entity、Destroy 等 Runtime 局部能力。需要读取或修改
// 这些状态时，应通过 Submit、Post 或 ContinueOn 回到组件所属 Runtime。
type ConcurrentComponent interface {
	iConcurrentComponent
	corectx.ConcurrentContextProvider
	corectx.AsyncScopeProvider
	fmt.Stringer

	// ID 返回组件 ID；未启用组件唯一 ID 时通常与 Entity ID 相同。
	ID() uid.ID
	// Name 返回组件在 Entity 中的名称。
	Name() string
}

type iConcurrentComponent interface {
	getInstance() Component
}

// componentAsyncScopeState 发布后不可变；nil 指针表示尚未创建且仍可用。
type componentAsyncScopeState struct {
	scope  *async.Scope
	closed bool
}

var closedComponentAsyncScopeState = &componentAsyncScopeState{closed: true}

// ConcurrentContextCache 返回所属 Entity 的并发 Runtime 上下文接口缓存。
func (comp *ComponentBehavior) ConcurrentContextCache() iface.Cache {
	return comp.entity.ConcurrentContextCache()
}

// AsyncScope 返回绑定组件 Lifetime 的后台任务作用域，并在首次访问时懒创建。
// Scope 在组件从 Entity 移除或随 Entity 销毁时关闭；SetEnabled(false) 不会关闭它。
// 所属 Entity 尚未绑定 Runtime Context 时返回 nil；组件已关闭后首次访问会返回已关闭的 Scope。
func (comp *ComponentBehavior) AsyncScope() *async.Scope {
	for {
		state := comp.asyncScope.Load()
		if state != nil && state.scope != nil {
			return state.scope
		}

		if comp.entity == nil {
			return nil
		}
		entityScope := comp.entity.AsyncScope()
		if entityScope == nil {
			return nil
		}

		asyncScope := async.NewScope(entityScope.Context())
		closed := state != nil && state.closed
		if closed {
			asyncScope.Close()
		}

		newState := &componentAsyncScopeState{
			scope:  asyncScope,
			closed: closed,
		}
		if comp.asyncScope.CompareAndSwap(state, newState) {
			return asyncScope
		}

		asyncScope.Close()
	}
}

// String 返回包含组件、实体及原型标识的 JSON 文本；组件尚未完成 Runtime 初始化时返回空字符串。
func (comp *ComponentBehavior) String() string {
	if cached := comp.stringerCache.Load(); cached != nil {
		return *cached
	}

	if comp.entity == nil || comp.id.IsNil() {
		return ""
	}
	entityScope := comp.entity.AsyncScope()
	if entityScope == nil {
		return ""
	}

	value := fmt.Sprintf(`{"id":%q,"entity_id":%q,"name":%q,"prototype":%q}`, comp.ID(), comp.Entity().ID(), comp.Name(), comp.Builtin().PT.Prototype())
	if comp.stringerCache.CompareAndSwap(nil, &value) {
		return value
	}
	return *comp.stringerCache.Load()
}

func (comp *ComponentBehavior) getInstance() Component {
	return comp.instance
}

func (comp *ComponentBehavior) closeAsyncScope() {
	for {
		state := comp.asyncScope.Load()
		if state == nil {
			if comp.asyncScope.CompareAndSwap(nil, closedComponentAsyncScopeState) {
				return
			}
			continue
		}
		if state.closed {
			return
		}

		// 先关闭实际 Scope 再发布关闭状态，避免读取方看到关闭标记时 Scope 仍可接收任务。
		state.scope.Close()
		closedState := &componentAsyncScopeState{
			scope:  state.scope,
			closed: true,
		}
		if comp.asyncScope.CompareAndSwap(state, closedState) {
			return
		}
	}
}
