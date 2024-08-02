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

// EventTreeNodeAddChild 事件：实体节点添加子实体
// +event-gen:export=0
type EventTreeNodeAddChild interface {
	OnTreeNodeAddChild(self, child Entity)
}

// EventTreeNodeRemoveChild 事件：实体节点删除子实体
// +event-gen:export=0
type EventTreeNodeRemoveChild interface {
	OnTreeNodeRemoveChild(self, child Entity)
}

// EventTreeNodeEnterParent 事件：实体加入父实体节点
// +event-gen:export=0
type EventTreeNodeEnterParent interface {
	OnTreeNodeEnterParent(self, parent Entity)
}

// EventTreeNodeLeaveParent 事件：实体离开父实体节点
// +event-gen:export=0
type EventTreeNodeLeaveParent interface {
	OnTreeNodeLeaveParent(self, parent Entity)
}
