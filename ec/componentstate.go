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

// ComponentState 组件状态
type ComponentState int8

const (
	ComponentState_Birth     ComponentState = iota // 出生
	ComponentState_Attach                          // 附着
	ComponentState_Awake                           // 唤醒
	ComponentState_Enable                          // 启用
	ComponentState_Idle                            // 空闲
	ComponentState_Start                           // 开始
	ComponentState_Alive                           // 活跃
	ComponentState_Detach                          // 脱离
	ComponentState_Shut                            // 结束
	ComponentState_Disable                         // 禁用
	ComponentState_Death                           // 死亡
	ComponentState_Destroyed                       // 已销毁
)
