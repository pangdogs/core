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

// EventEntityManagerAddEntity 在实体加入本地管理器后派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventEntityManagerAddEntity interface {
	OnEntityManagerAddEntity(entityManager EntityManager, entity ec.Entity)
}

// EventEntityManagerRemoveEntity 在实体离开本地管理器前派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventEntityManagerRemoveEntity interface {
	OnEntityManagerRemoveEntity(entityManager EntityManager, entity ec.Entity)
}

// EventEntityManagerEntityAddComponents 在受管实体添加组件后派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventEntityManagerEntityAddComponents interface {
	OnEntityManagerEntityAddComponents(entityManager EntityManager, entity ec.Entity, components []ec.Component)
}

// EventEntityManagerEntityRemoveComponent 在受管实体删除组件前派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventEntityManagerEntityRemoveComponent interface {
	OnEntityManagerEntityRemoveComponent(entityManager EntityManager, entity ec.Entity, component ec.Component)
}

// EventEntityManagerEntityComponentEnableChanged 在受管实体的组件启用状态改变后派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventEntityManagerEntityComponentEnableChanged interface {
	OnEntityManagerEntityComponentEnableChanged(entityManager EntityManager, entity ec.Entity, component ec.Component, enable bool)
}

// EventEntityManagerEntityFirstTouchComponent 在受管实体首次访问延迟唤醒组件时派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventEntityManagerEntityFirstTouchComponent interface {
	OnEntityManagerEntityFirstTouchComponent(entityManager EntityManager, entity ec.Entity, component ec.Component)
}
