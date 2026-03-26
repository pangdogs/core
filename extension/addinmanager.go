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

	// Install 安装插件
	Install(addInFace iface.FaceAny, name ...string) AddInStatus
	// Uninstall 卸载插件
	Uninstall(name string)
	// GetStatusByName 使用名称查询插件状态信息
	GetStatusByName(name string) (AddInStatus, bool)
	// GetStatusById 使用Id查询插件状态信息
	GetStatusById(id uint64) (AddInStatus, bool)
	// ListStatuses 获取所有插件状态信息
	ListStatuses() []AddInStatus
}
