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

// Deprecated: UnsafeServiceAddInStatus 访问服务插件状态信息的内部方法
func UnsafeServiceAddInStatus(status ServiceAddInStatus) _UnsafeServiceAddInStatus {
	return _UnsafeServiceAddInStatus{
		ServiceAddInStatus: status,
	}
}

type _UnsafeServiceAddInStatus struct {
	ServiceAddInStatus
}

// Started 已启动
func (u _UnsafeServiceAddInStatus) Started() {
	u.started()
}

// Stopped 已停止
func (u _UnsafeServiceAddInStatus) Stopped() {
	u.stopped()
}

// DoInstallOnce 执行安装函数一次
func (u _UnsafeServiceAddInStatus) DoInstallOnce(fun func()) {
	u.doInstallOnce(fun)
}

// DoUninstallOnce 执行卸载函数一次
func (u _UnsafeServiceAddInStatus) DoUninstallOnce(fun func()) {
	u.doUninstallOnce(fun)
}
