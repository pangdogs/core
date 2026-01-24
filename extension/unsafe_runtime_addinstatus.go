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

package extension

import "git.golaxy.org/core/event"

// Deprecated: UnsafeRuntimeAddInStatus 访问运行时插件状态信息的内部方法
func UnsafeRuntimeAddInStatus(status RuntimeAddInStatus) _UnsafeRuntimeAddInStatus {
	return _UnsafeRuntimeAddInStatus{
		RuntimeAddInStatus: status,
	}
}

type _UnsafeRuntimeAddInStatus struct {
	RuntimeAddInStatus
}

// SetState 修改状态
func (u _UnsafeRuntimeAddInStatus) SetState(state AddInState) {
	u.setState(state)
}

// ManagedRuntimeRunningEventHandle 托管运行时运行事件句柄
func (u _UnsafeRuntimeAddInStatus) ManagedRuntimeRunningEventHandle(runtimeRunningEventHandle event.Handle) {
	u.managedRuntimeRunningEventHandle(runtimeRunningEventHandle)
}

// ManagedUnbindRuntimeHandles 解绑运行时句柄
func (u _UnsafeRuntimeAddInStatus) ManagedUnbindRuntimeHandles() {
	u.managedUnbindRuntimeHandles()
}
