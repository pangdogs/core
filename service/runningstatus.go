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
package service

// RunningStatus 运行状态
type RunningStatus int32

const (
	RunningStatus_Birth              RunningStatus = iota // 出生
	RunningStatus_Starting                                // 开始启动
	RunningStatus_Started                                 // 已启动
	RunningStatus_Terminating                             // 开始停止
	RunningStatus_Terminated                              // 已停止
	RunningStatus_ActivatingAddIn                         // 开始激活插件
	RunningStatus_AddInActivated                          // 插件已激活
	RunningStatus_DeactivatingAddIn                       // 开始去激活插件
	RunningStatus_AddInDeactivated                        // 插件已去激活
	RunningStatus_EntityPTDeclared                        // 实体原型已声明
	RunningStatus_EntityPTRedeclared                      // 实体原型已重声明
	RunningStatus_EntityPTUndeclared                      // 实体原型已取消声明
)
