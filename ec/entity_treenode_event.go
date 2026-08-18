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

// EventTreeNodeAddChild 在实体节点加入直接子实体时同步派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventTreeNodeAddChild interface {
	OnTreeNodeAddChild(entity Entity, childID uid.ID)
}

// EventTreeNodeRemoveChild 在直接子实体离开实体节点时同步派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventTreeNodeRemoveChild interface {
	OnTreeNodeRemoveChild(entity Entity, childID uid.ID)
}

// EventTreeNodeAttachParent 在实体接入父节点时同步派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventTreeNodeAttachParent interface {
	OnTreeNodeAttachParent(entity Entity, parentID uid.ID)
}

// EventTreeNodeDetachParent 在实体脱离父节点时同步派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventTreeNodeDetachParent interface {
	OnTreeNodeDetachParent(entity Entity, parentID uid.ID)
}

// EventTreeNodeMoveTo 在实体从一个父节点移动到另一个父节点时同步派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventTreeNodeMoveTo interface {
	OnTreeNodeMoveTo(entity Entity, fromParentID, toParentID uid.ID)
}
