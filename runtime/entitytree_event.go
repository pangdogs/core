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
package runtime

import (
	"git.golaxy.org/core/ec"
)

// EventEntityTreeAddNode 事件：新增实体树节点
// +event-gen:export=0
type EventEntityTreeAddNode interface {
	OnEntityTreeAddNode(entityTree EntityTree, parent, child ec.Entity)
}

// EventEntityTreeRemoveNode 事件：删除实体树节点
// +event-gen:export=0
type EventEntityTreeRemoveNode interface {
	OnEntityTreeRemoveNode(entityTree EntityTree, parent, child ec.Entity)
}
