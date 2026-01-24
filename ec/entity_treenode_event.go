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
//go:generate go run git.golaxy.org/core/event/eventc eventtab --name=entityTreeNodeEventTab
package ec

import "git.golaxy.org/core/utils/uid"

// EventTreeNodeAddChild 事件：实体节点添加子实体
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventTreeNodeAddChild interface {
	OnTreeNodeAddChild(entity Entity, childId uid.Id)
}

// EventTreeNodeRemoveChild 事件：实体节点删除子实体
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventTreeNodeRemoveChild interface {
	OnTreeNodeRemoveChild(entity Entity, childId uid.Id)
}

// EventTreeNodeAttachParent 事件：实体加入父实体节点
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventTreeNodeAttachParent interface {
	OnTreeNodeAttachParent(entity Entity, parentId uid.Id)
}

// EventTreeNodeDetachParent 事件：实体离开父实体节点
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventTreeNodeDetachParent interface {
	OnTreeNodeDetachParent(entity Entity, parentId uid.Id)
}

// EventTreeNodeMoveTo 事件：实体切换父节点
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventTreeNodeMoveTo interface {
	OnTreeNodeMoveTo(entity Entity, fromParentId, toParentId uid.Id)
}
