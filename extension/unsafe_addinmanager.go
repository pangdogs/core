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

import "git.golaxy.org/core/utils/generic"

// Deprecated: UnsafeAddInManager 访问插件管理器的内部方法
func UnsafeAddInManager(addInManager AddInManager) _UnsafeAddInManager {
	return _UnsafeAddInManager{
		AddInManager: addInManager,
	}
}

type _UnsafeAddInManager struct {
	AddInManager
}

// SetCallback 设置回调
func (u _UnsafeAddInManager) SetCallback(installCB, uninstallCB generic.Action1[AddInStatus]) {
	u.setCallback(installCB, uninstallCB)
}
