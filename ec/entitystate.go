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
	EntityState_Awaking                      // EntityState_Awaking 表示实体正在唤醒阶段，Awake 在此状态执行。
	EntityState_Starting                     // EntityState_Starting 表示实体已进入启动阶段，Start 在此状态执行。
	EntityState_Alive                        // EntityState_Alive 表示实体已启动并参与运行时更新。
	EntityState_Leaving                      // EntityState_Leaving 表示实体正在离开 Runtime，后续新增组件不再进入 Runtime 生命周期。
	EntityState_Shutting                     // EntityState_Shutting 表示实体已进入关闭阶段，Shut 在此状态执行。
	EntityState_Dead                         // EntityState_Dead 表示实体上下文及事件已关闭，正在完成销毁回调。
	EntityState_Destroyed                    // EntityState_Destroyed 表示实体已释放托管句柄并完成 Terminated 通知。
)
