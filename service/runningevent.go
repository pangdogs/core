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
package service

// RunningEvent 运行事件
type RunningEvent int32

const (
	RunningEvent_Birth               RunningEvent = iota // 出生
	RunningEvent_Starting                                // 开始启动
	RunningEvent_Started                                 // 已启动
	RunningEvent_Heartbeat                               // 心跳
	RunningEvent_Terminating                             // 开始停止
	RunningEvent_Terminated                              // 已停止
	RunningEvent_AddInActivating                         // 开始激活插件
	RunningEvent_AddInActivated                          // 已激活插件
	RunningEvent_AddInDeactivating                       // 开始去激活插件
	RunningEvent_AddInDeactivated                        // 已去激活插件
	RunningEvent_EntityPTDeclared                        // 实体原型已声明
	RunningEvent_ComponentPTDeclared                     // 组件原型已声明
	RunningEvent_EntityRegistered                        // 实体已注册
	RunningEvent_EntityUnregistered                      // 实体已注销
)
