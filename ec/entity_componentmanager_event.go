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
//go:generate go run git.golaxy.org/core/event/eventc eventtab --name=entityComponentManagerEventTab
package ec

// EventComponentManagerAddComponents 在一批组件加入实体后同步派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventComponentManagerAddComponents interface {
	OnComponentManagerAddComponents(entity Entity, components []Component)
}

// EventComponentManagerRemoveComponent 在组件脱离实体时同步派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventComponentManagerRemoveComponent interface {
	OnComponentManagerRemoveComponent(entity Entity, component Component)
}

// EventComponentManagerComponentEnableChanged 在所属组件的启用标记发生变化时同步派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventComponentManagerComponentEnableChanged interface {
	OnComponentManagerComponentEnableChanged(entity Entity, component Component, enable bool)
}

// EventComponentManagerFirstTouchComponent 在启用首次访问唤醒且 Attached 状态组件被访问时同步派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventComponentManagerFirstTouchComponent interface {
	OnComponentManagerFirstTouchComponent(entity Entity, component Component)
}
