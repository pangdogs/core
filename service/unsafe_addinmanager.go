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

// Deprecated: UnsafeAddInManager 暴露服务插件管理器的生命周期操作，仅供 core 使用。
func UnsafeAddInManager(mgr AddInManager) _UnsafeAddInManager {
	return _UnsafeAddInManager{AddInManager: mgr}
}

type _UnsafeAddInManager struct {
	AddInManager
}

// Freeze 冻结安装与卸载入口，并按安装顺序返回插件状态。
func (mgr _UnsafeAddInManager) Freeze() []AddInStatus {
	return mgr.freeze()
}

// List 按安装顺序返回当前服务插件状态的副本。
func (mgr _UnsafeAddInManager) List() []AddInStatus {
	return mgr.getList()
}
