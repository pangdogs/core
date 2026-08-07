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

//go:generate stringer -type EntityState
package ec

// EntityState 表示实体生命周期所处的阶段；状态通常由 Runtime 单向推进。
type EntityState int8

const (
	EntityState_Born      EntityState = iota // EntityState_Born 表示实体已构造、尚未进入 Runtime。
	EntityState_Entered                      // EntityState_Entered 表示实体已加入 Runtime 的实体管理器。
	EntityState_Awakened                     // EntityState_Awakened 表示实体已完成唤醒。
	EntityState_Starting                     // EntityState_Starting 表示实体正在执行启动阶段。
	EntityState_Alive                        // EntityState_Alive 表示实体已启动并参与运行时更新。
	EntityState_Leaving                      // EntityState_Leaving 表示实体正在离开 Runtime。
	EntityState_Shutting                     // EntityState_Shutting 表示实体正在执行关闭阶段。
	EntityState_Dead                         // EntityState_Dead 表示实体已死亡，其上下文已取消。
	EntityState_Destroyed                    // EntityState_Destroyed 表示实体已完成销毁并释放托管句柄。
)
