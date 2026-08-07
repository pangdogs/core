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
	"github.com/elliotchance/pie/v2"
)

// ServiceAddIn 创建服务插件定义，其 Require 和 Lookup 直接接受 service.Context。
func ServiceAddIn[ADDIN_IFACE, SETTING any](creator generic.FuncVar0[SETTING, ADDIN_IFACE], name ...string) ServiceAddInDefinition[ADDIN_IFACE, SETTING] {
	addIn := defineAddIn[ADDIN_IFACE, SETTING](creator, pie.First(name))

	return ServiceAddInDefinition[ADDIN_IFACE, SETTING]{
		Id:        addIn.Id,
		Name:      addIn.Name,
		Install:   addIn.Install,
		Uninstall: addIn.Uninstall,
		Require:   func(svcCtx service.Context) ADDIN_IFACE { return addIn.Require(svcCtx) },
		Lookup:    func(svcCtx service.Context) (ADDIN_IFACE, bool) { return addIn.Lookup(svcCtx) },
	}
}

// ServiceAddInDefinition 封装服务插件的标识、构造和访问操作。
type ServiceAddInDefinition[ADDIN_IFACE, SETTING any] struct {
	Id        uint64                                                // 由 Name 生成的插件 ID。
	Name      string                                                // 插件注册名称。
	Install   generic.ActionVar1[extension.AddInProvider, SETTING]  // 构造插件并安装到给定提供者。
	Uninstall generic.Action1[extension.AddInProvider]              // 从给定提供者卸载插件。
	Require   generic.Func1[service.Context, ADDIN_IFACE]           // 从服务获取正在运行的插件，不可用时 panic。
	Lookup    generic.FuncPair1[service.Context, ADDIN_IFACE, bool] // 查询服务管理器当前持有的插件。
}
