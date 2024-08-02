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
package ec

// EventComponentMgrAddComponents 事件：实体的组件管理器添加组件
// +event-gen:export=0
type EventComponentMgrAddComponents interface {
	OnComponentMgrAddComponents(entity Entity, components []Component)
}

// EventComponentMgrRemoveComponent 事件：实体的组件管理器删除组件
// +event-gen:export=0
type EventComponentMgrRemoveComponent interface {
	OnComponentMgrRemoveComponent(entity Entity, component Component)
}

// EventComponentMgrFirstAccessComponent 事件：实体的组件管理器首次访问组件
// +event-gen:export=0
type EventComponentMgrFirstAccessComponent interface {
	OnComponentMgrFirstAccessComponent(entity Entity, component Component)
}
