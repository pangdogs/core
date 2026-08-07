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

//go:generate stringer -type RunningEvent
package runtime

// RunningEvent 标识运行时工作循环、插件和实体生命周期事件。
type RunningEvent int32

const (
	RunningEvent_Birth                              RunningEvent = iota // 运行时对象创建完成。
	RunningEvent_Starting                                               // 运行时开始启动，预装插件在回调前激活。
	RunningEvent_Started                                                // 运行时启动完成。
	RunningEvent_FrameLoopBegin                                         // 一次完整帧循环开始。
	RunningEvent_FrameUpdateBegin                                       // 帧更新阶段开始。
	RunningEvent_FrameUpdateEnd                                         // 帧更新阶段结束。
	RunningEvent_FrameLoopEnd                                           // 一次完整帧循环结束。
	RunningEvent_RunCallBegin                                           // 开始执行一个普通异步调用任务。
	RunningEvent_RunCallEnd                                             // 普通异步调用任务执行结束。
	RunningEvent_RunGCBegin                                             // 运行时 GC 开始。
	RunningEvent_RunGCEnd                                               // 运行时 GC 结束。
	RunningEvent_Terminating                                            // 运行时开始停止。
	RunningEvent_Terminated                                             // 运行时等待组清空、插件停用完成，主体生命周期结束。
	RunningEvent_AddInActivating                                        // 插件开始激活。
	RunningEvent_AddInActivationAborted                                 // 插件激活被中止。
	RunningEvent_AddInActivated                                         // 插件激活完成。
	RunningEvent_AddInDeactivating                                      // 插件开始停用。
	RunningEvent_AddInDeactivated                                       // 插件停用完成。
	RunningEvent_EntityActivating                                       // 实体开始激活。
	RunningEvent_EntityActivationAborted                                // 实体激活被中止。
	RunningEvent_EntityActivated                                        // 实体激活完成。
	RunningEvent_EntityDeactivating                                     // 实体开始停用。
	RunningEvent_EntityDeactivated                                      // 实体停用完成。
	RunningEvent_EntityComponentsActivating                             // 实体开始激活一批新增组件。
	RunningEvent_EntityComponentsActivationAborted                      // 新增组件的激活流程被中止。
	RunningEvent_EntityComponentsActivated                              // 新增组件激活完成。
	RunningEvent_EntityComponentDeactivating                            // 实体开始停用即将删除的组件。
	RunningEvent_EntityComponentDeactivationAborted                     // 组件的停用回调流程被中止；组件删除仍会完成。
	RunningEvent_EntityComponentDeactivated                             // 组件停用完成，随后将从实体删除。
)
