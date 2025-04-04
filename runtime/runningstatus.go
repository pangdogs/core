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

//go:generate stringer -type RunningStatus
package runtime

// RunningStatus 运行状态
type RunningStatus int32

const (
	RunningStatus_Birth                RunningStatus = iota // 出生
	RunningStatus_Starting                                  // 开始启动
	RunningStatus_Started                                   // 已启动
	RunningStatus_FrameLoopBegin                            // 帧循环开始
	RunningStatus_FrameUpdateBegin                          // 帧更新开始
	RunningStatus_FrameUpdateEnd                            // 帧更新结束
	RunningStatus_FrameLoopEnd                              // 帧循环结束
	RunningStatus_RunCallBegin                              // Call开始执行
	RunningStatus_RunCallEnd                                // Call结束执行
	RunningStatus_RunGCBegin                                // GC开始执行
	RunningStatus_RunGCEnd                                  // GC结束执行
	RunningStatus_Terminating                               // 开始停止
	RunningStatus_Terminated                                // 已停止
	RunningStatus_ActivatingAddIn                           // 开始激活插件
	RunningStatus_AddInActivated                            // 插件已激活
	RunningStatus_DeactivatingAddIn                         // 开始去激活插件
	RunningStatus_AddInDeactivated                          // 插件已去激活
	RunningStatus_ActivatingEntity                          // 开始激活实体
	RunningStatus_EntityActivated                           // 实体已激活
	RunningStatus_DeactivatingEntity                        // 开始去激活实体
	RunningStatus_EntityDeactivated                         // 实体已去激活
	RunningStatus_EntityInitComponents                      // 实体初始化组件
	RunningStatus_EntityAddComponents                       // 实体新增组件
)
