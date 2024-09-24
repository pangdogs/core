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

package define

import (
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/generic"
)

// ServicePluginInterface 定义服务插件接口，因为仅有接口没有实现，所以不能用于向插件包安装插件
func ServicePluginInterface[PLUGIN_IFACE any]() ServicePluginInterfaceDefinition[PLUGIN_IFACE] {
	plug := definePluginInterface[PLUGIN_IFACE]()

	return ServicePluginInterfaceDefinition[PLUGIN_IFACE]{
		Name:  plug.Name,
		Using: func(svcCtx service.Context) PLUGIN_IFACE { return plug.Using(svcCtx) },
	}
}

// ServicePluginInterfaceDefinition 服务插件接口定义，只能在服务上下文中使用
type ServicePluginInterfaceDefinition[PLUGIN_IFACE any] struct {
	Name  string                                       // 插件名称
	Using generic.Func1[service.Context, PLUGIN_IFACE] // 使用插件
}
