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
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/types"
)

// Plugin 定义通用插件，支持运行时和服务上下文
func Plugin[PLUGIN_IFACE, OPTION any](creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) PluginDefinition[PLUGIN_IFACE, OPTION] {
	return definePlugin[PLUGIN_IFACE, OPTION](creator)
}

// PluginDefinition 通用插件定义
type PluginDefinition[PLUGIN_IFACE, OPTION any] struct {
	Name      string                                                // 插件名称
	Install   generic.ActionVar1[extension.PluginProvider, OPTION]  // 向插件包安装
	Uninstall generic.Action1[extension.PluginProvider]             // 从插件包卸载
	Using     generic.Func1[extension.PluginProvider, PLUGIN_IFACE] // 使用插件
}

func definePlugin[PLUGIN_IFACE, OPTION any](creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) PluginDefinition[PLUGIN_IFACE, OPTION] {
	if creator == nil {
		exception.Panicf("%w: %w: creator is nil", exception.ErrCore, exception.ErrArgs)
	}

	name := types.FullNameT[PLUGIN_IFACE]()

	return PluginDefinition[PLUGIN_IFACE, OPTION]{
		Name: name,
		Install: func(provider extension.PluginProvider, options ...OPTION) {
			extension.Install[PLUGIN_IFACE](provider, creator(options...), name)
		},
		Uninstall: func(provider extension.PluginProvider) {
			extension.Uninstall(provider, name)
		},
		Using: func(provider extension.PluginProvider) PLUGIN_IFACE {
			return extension.Using[PLUGIN_IFACE](provider, name)
		},
	}
}
