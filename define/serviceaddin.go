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

// ServiceAddIn 定义服务插件，支持安装至服务上下文
func ServiceAddIn[ADDIN_IFACE, SETTING any](creator generic.FuncVar0[SETTING, ADDIN_IFACE]) ServiceAddInDefinition[ADDIN_IFACE, SETTING] {
	addIn := defineAddIn[ADDIN_IFACE, SETTING](creator)

	return ServiceAddInDefinition[ADDIN_IFACE, SETTING]{
		Id:        addIn.Id,
		Name:      addIn.Name,
		Install:   addIn.Install,
		Uninstall: addIn.Uninstall,
		Resolve:   func(svcCtx service.Context) ADDIN_IFACE { return addIn.Resolve(svcCtx) },
		Lookup:    func(svcCtx service.Context) (ADDIN_IFACE, bool) { return addIn.Lookup(svcCtx) },
	}
}

// ServiceAddInDefinition 服务插件定义
type ServiceAddInDefinition[ADDIN_IFACE, SETTING any] struct {
	Id        uint64                                                // 插件Id
	Name      string                                                // 插件名称
	Install   generic.ActionVar1[extension.AddInProvider, SETTING]  // 向插件管理器安装
	Uninstall generic.Action1[extension.AddInProvider]              // 从插件管理器卸载
	Resolve   generic.Func1[service.Context, ADDIN_IFACE]           // 解析插件
	Lookup    generic.FuncPair1[service.Context, ADDIN_IFACE, bool] // 查找插件
}
