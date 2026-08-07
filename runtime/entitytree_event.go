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
//go:generate go run git.golaxy.org/core/event/eventc eventtab --name=entityTreeEventTab
package runtime

import (
	"git.golaxy.org/core/utils/uid"
)

// EventEntityTreeAddNode 在节点成功加入实体树后派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventEntityTreeAddNode interface {
	OnEntityTreeAddNode(entityTree EntityTree, parentId, childId uid.Id)
}

// EventEntityTreeRemoveNode 在节点的树关系移除前派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventEntityTreeRemoveNode interface {
	OnEntityTreeRemoveNode(entityTree EntityTree, parentId, childId uid.Id)
}

// EventEntityTreeMoveNode 在节点切换父节点后派发。
// +event-gen:export_emit=0
// +event-tab-gen:recursion=allow
type EventEntityTreeMoveNode interface {
	OnEntityTreeMoveNode(entityTree EntityTree, childId, fromParentId, toParentId uid.Id)
}
