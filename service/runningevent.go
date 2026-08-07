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

// RunningEvent 标识服务生命周期及服务级资源变化事件。
type RunningEvent int32

const (
	RunningEvent_Birth               RunningEvent = iota // 服务对象创建完成。
	RunningEvent_Starting                                // 服务开始启动，插件管理器在回调前冻结并激活插件。
	RunningEvent_Started                                 // 服务启动完成。
	RunningEvent_Heartbeat                               // 服务运行期间的每秒心跳。
	RunningEvent_Terminating                             // 服务开始停止。
	RunningEvent_Terminated                              // 服务等待组清空、插件停用完成，主体生命周期结束。
	RunningEvent_EntityPTDeclared                        // 实体原型已声明。
	RunningEvent_ComponentPTDeclared                     // 组件原型已声明。
	RunningEvent_EntityRegistered                        // 全局实体已注册。
	RunningEvent_EntityDeregistered                      // 全局实体已注销。
)
