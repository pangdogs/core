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

package service

// Deprecated: UnsafeAddInStatus 暴露服务插件状态的生命周期操作，仅供 core 使用。
func UnsafeAddInStatus(status AddInStatus) _UnsafeAddInStatus {
	return _UnsafeAddInStatus{AddInStatus: status}
}

type _UnsafeAddInStatus struct {
	AddInStatus
}

// Started 将已加载插件标记为正在运行。
func (u _UnsafeAddInStatus) Started() {
	u.started()
}

// Stopped 将插件从管理器移除并标记为已卸载。
func (u _UnsafeAddInStatus) Stopped() {
	u.stopped()
}
