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
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/generic"
)

// ServicePlugin 定义服务插件，支持服务上下文
func ServicePlugin[PLUGIN_IFACE, OPTION any](creator generic.FuncVar0[OPTION, PLUGIN_IFACE]) ServicePluginDefinition[PLUGIN_IFACE, OPTION] {
	plug := definePlugin[PLUGIN_IFACE, OPTION](creator)

	return ServicePluginDefinition[PLUGIN_IFACE, OPTION]{
		Name:      plug.Name,
		Install:   plug.Install,
		Uninstall: plug.Uninstall,
		Using:     func(svcCtx service.Context) PLUGIN_IFACE { return plug.Using(svcCtx) },
	}
}

// ServicePluginDefinition 服务插件定义
type ServicePluginDefinition[PLUGIN_IFACE, OPTION any] struct {
	Name      string                                               // 插件名称
	Install   generic.ActionVar1[extension.PluginProvider, OPTION] // 向插件包安装
	Uninstall generic.Action1[extension.PluginProvider]            // 从插件包卸载
	Using     generic.Func1[service.Context, PLUGIN_IFACE]         // 使用插件
}
