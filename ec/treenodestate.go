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

//go:generate stringer -type TreeNodeState
package ec

// TreeNodeState 表示实体在 Runtime 实体树中的过渡状态。
type TreeNodeState int8

const (
	TreeNodeState_Free      TreeNodeState = iota // TreeNodeState_Free 表示实体不在实体树中。
	TreeNodeState_Attaching                      // TreeNodeState_Attaching 表示实体正在加入实体树。
	TreeNodeState_Attached                       // TreeNodeState_Attached 表示实体已在实体树中。
	TreeNodeState_Detaching                      // TreeNodeState_Detaching 表示实体正在脱离实体树。
	TreeNodeState_Moving                         // TreeNodeState_Moving 表示实体正在切换父节点。
)
