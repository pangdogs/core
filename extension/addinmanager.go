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

import (
	"git.golaxy.org/core/utils/iface"
)

// AddInManager 插件管理器
type AddInManager interface {
	AddInProvider

	// Install 安装插件，不设置插件名称时，将会使用插件实例名称作为插件名称
	Install(addInFace iface.FaceAny, name ...string) AddInStatus
	// Uninstall 卸载插件
	Uninstall(name string)
	// Get 获取插件
	Get(name string) (AddInStatus, bool)
	// List 获取所有插件
	List() []AddInStatus
}
