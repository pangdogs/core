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

//go:generate stringer -type ComponentState
package ec

// ComponentState 表示组件生命周期所处的阶段；状态通常由所属实体的 Runtime 推进。
type ComponentState int8

const (
	ComponentState_Born      ComponentState = iota // ComponentState_Born 表示组件已构造、尚未依附实体。
	ComponentState_Attached                        // ComponentState_Attached 表示组件已加入实体。
	ComponentState_Awakened                        // ComponentState_Awakened 表示组件已完成唤醒。
	ComponentState_Enabling                        // ComponentState_Enabling 表示组件正在执行启用阶段。
	ComponentState_Idle                            // ComponentState_Idle 表示组件已唤醒但当前未启用。
	ComponentState_Starting                        // ComponentState_Starting 表示组件正在执行启动阶段。
	ComponentState_Alive                           // ComponentState_Alive 表示组件已启动并处于活动状态。
	ComponentState_Detaching                       // ComponentState_Detaching 表示组件正在脱离实体。
	ComponentState_Shutting                        // ComponentState_Shutting 表示组件正在执行关闭阶段。
	ComponentState_Disabling                       // ComponentState_Disabling 表示组件正在执行禁用阶段。
	ComponentState_Dead                            // ComponentState_Dead 表示组件已死亡，事件表已停用。
	ComponentState_Destroyed                       // ComponentState_Destroyed 表示组件已完成销毁并释放托管句柄。
)
