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

package plugin

import "git.golaxy.org/core/utils/generic"

// Deprecated: UnsafePluginBundle 访问插件包的内部方法
func UnsafePluginBundle(pluginBundle PluginBundle) _UnsafePluginBundle {
	return _UnsafePluginBundle{
		PluginBundle: pluginBundle,
	}
}

type _UnsafePluginBundle struct {
	PluginBundle
}

// SetInstallCB 设置安装插件回调
func (up _UnsafePluginBundle) SetInstallCB(cb generic.Action1[PluginStatus]) {
	up.setInstallCB(cb)
}

// SetUninstallCB 设置卸载插件回调
func (up _UnsafePluginBundle) SetUninstallCB(cb generic.Action1[PluginStatus]) {
	up.setUninstallCB(cb)
}
