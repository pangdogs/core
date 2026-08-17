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

//go:generate go run git.golaxy.org/core/event/eventc event
//go:generate go run git.golaxy.org/core/event/eventc eventtab --name=entityManagerEventTab
package runtime

import "git.golaxy.org/core/ec"

// EventEntityManagerAddEntity 在实体写入本地索引并进入 Entered 后派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventEntityManagerAddEntity interface {
	OnEntityManagerAddEntity(entityManager EntityManager, entity ec.Entity)
}

// EventEntityManagerRemoveEntity 在实体进入 Leaving 且树关系已移除、但本地与全局索引尚未删除时派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventEntityManagerRemoveEntity interface {
	OnEntityManagerRemoveEntity(entityManager EntityManager, entity ec.Entity)
}

// EventEntityManagerEntityAddComponents 在新增组件完成 Runtime 身份与事件配置后派发。
// Core Runtime 仅在 Entity 处于 Awaking 至 Alive 时推进这些组件的生命周期。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventEntityManagerEntityAddComponents interface {
	OnEntityManagerEntityAddComponents(entityManager EntityManager, entity ec.Entity, components []ec.Component)
}

// EventEntityManagerEntityRemoveComponent 在组件进入 Detaching 且仍位于实体组件表中时派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventEntityManagerEntityRemoveComponent interface {
	OnEntityManagerEntityRemoveComponent(entityManager EntityManager, entity ec.Entity, component ec.Component)
}

// EventEntityManagerEntityComponentEnableChanged 在组件启用标记改变且完成 Runtime 事件转发后派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventEntityManagerEntityComponentEnableChanged interface {
	OnEntityManagerEntityComponentEnableChanged(entityManager EntityManager, entity ec.Entity, component ec.Component, enable bool)
}

// EventEntityManagerEntityFirstTouchComponent 在受管实体中，处于 Attached 的组件被访问时
// 同步派发。Runtime 通过此事件使目标组件提前进入并执行 Awake。
// 事件允许递归派发，使 Awake 内访问的其他组件可以继续建立依赖顺序。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventEntityManagerEntityFirstTouchComponent interface {
	OnEntityManagerEntityFirstTouchComponent(entityManager EntityManager, entity ec.Entity, component ec.Component)
}
