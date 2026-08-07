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

// AddInManager 定义 service 与 runtime 插件管理器共用的操作。
// 生命周期约束和并发保证由具体管理器决定。
type AddInManager interface {
	AddInProvider

	// Install 安装插件并返回状态信息；name 为空时由实现根据实例类型生成名称。
	Install(addInFace iface.FaceAny, name ...string) AddInStatus
	// Uninstall 按名称卸载插件；插件不存在时不执行任何操作。
	Uninstall(name string)
	// GetStatusByName 按名称查询当前由管理器持有的插件状态。
	GetStatusByName(name string) (AddInStatus, bool)
	// GetStatusById 按 ID 查询当前由管理器持有的插件状态。
	GetStatusById(id uint64) (AddInStatus, bool)
	// ListStatuses 返回当前由管理器持有的全部插件状态。
	ListStatuses() []AddInStatus
}
