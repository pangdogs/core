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

// EntityState 实体状态
type EntityState int8

const (
	EntityState_Birth     EntityState = iota // 出生
	EntityState_Enter                        // 进入容器
	EntityState_Awake                        // 唤醒
	EntityState_Start                        // 开始
	EntityState_Alive                        // 活跃
	EntityState_Leave                        // 离开容器
	EntityState_Shut                         // 结束
	EntityState_Death                        // 死亡
	EntityState_Destroyed                    // 已销毁
)
