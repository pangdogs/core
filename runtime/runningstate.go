//go:generate stringer -type RunningState
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

package runtime

// RunningState 运行状态
type RunningState int32

const (
	RunningState_Birth            RunningState = iota // 出生
	RunningState_Starting                             // 开始启动
	RunningState_Started                              // 已启动
	RunningState_FrameLoopBegin                       // 帧循环开始
	RunningState_FrameUpdateBegin                     // 帧更新开始
	RunningState_FrameUpdateEnd                       // 帧更新结束
	RunningState_FrameLoopEnd                         // 帧循环结束
	RunningState_RunCallBegin                         // Call开始执行
	RunningState_RunCallEnd                           // Call结束执行
	RunningState_RunGCBegin                           // GC开始执行
	RunningState_RunGCEnd                             // GC结束执行
	RunningState_Terminating                          // 开始停止
	RunningState_Terminated                           // 已停止
)
