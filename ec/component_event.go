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
//go:generate go run git.golaxy.org/core/event/eventc eventtab --name=componentEventTab
package ec

// EventComponentEnableChanged 在组件的启用标记发生变化时同步派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventComponentEnableChanged interface {
	OnComponentEnableChanged(comp Component, enable bool)
}

// EventComponentDestroy 描述 Component.Destroy 在移除流程末尾尝试派发的通知。
// 正常移除会先进入 Dead 并禁用组件事件表，因此受管组件的订阅者通常不会收到它；
// 需要观察移除流程时应订阅所属 Entity 的 EventComponentManagerRemoveComponent。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventComponentDestroy interface {
	OnComponentDestroy(comp Component)
}
