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
	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/types"
)

// PluginInterface 定义通用插件接口，支持运行时和服务上下文，通常用于为同类插件的不同实现提供统一的接口
func PluginInterface[PLUGIN_IFACE any]() PluginInterfaceDefinition[PLUGIN_IFACE] {
	return definePluginInterface[PLUGIN_IFACE]()
}

// PluginInterfaceDefinition 通用插件接口定义
type PluginInterfaceDefinition[PLUGIN_IFACE any] struct {
	Name  string                                                // 插件名称
	Using generic.Func1[extension.PluginProvider, PLUGIN_IFACE] // 使用插件
}

func definePluginInterface[PLUGIN_IFACE any]() PluginInterfaceDefinition[PLUGIN_IFACE] {
	name := types.FullNameT[PLUGIN_IFACE]()

	return PluginInterfaceDefinition[PLUGIN_IFACE]{
		Name: name,
		Using: func(provider extension.PluginProvider) PLUGIN_IFACE {
			return extension.Using[PLUGIN_IFACE](provider, name)
		},
	}
}
