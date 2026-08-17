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

// EventComponentManagerAddComponents 在一批组件写入实体组件表并进入 Attached 后同步派发。
// Runtime 通过此事件初始化受管组件身份；仅当 Entity 处于 Awaking 至 Alive 时推进后续生命周期。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventComponentManagerAddComponents interface {
	OnComponentManagerAddComponents(entity Entity, components []Component)
}

// EventComponentManagerRemoveComponent 在组件进入 Detaching、但尚未从实体组件表移除时同步派发。
// Runtime 通过此事件执行组件停用与销毁回调。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventComponentManagerRemoveComponent interface {
	OnComponentManagerRemoveComponent(entity Entity, component Component)
}

// EventComponentManagerComponentEnableChanged 在组件自身的启用标记事件派发后同步派发，
// Runtime 通过此事件推进对应的启用或禁用生命周期。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventComponentManagerComponentEnableChanged interface {
	OnComponentManagerComponentEnableChanged(entity Entity, component Component, enable bool)
}

// EventComponentManagerFirstTouchComponent 在启用 ComponentAwakeOnFirstTouch 后，处于 Attached
// 的组件被访问时同步派发。Runtime 通过此事件使目标组件提前进入并执行 Awake。
// 事件允许递归派发，使 Awake 内访问的其他组件可以继续建立依赖顺序。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventComponentManagerFirstTouchComponent interface {
	OnComponentManagerFirstTouchComponent(entity Entity, component Component)
}
