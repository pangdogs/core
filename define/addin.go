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
	"github.com/elliotchance/pie/v2"
)

// AddIn 创建可用于任意 extension.AddInProvider 的通用插件定义。
// ADDIN_IFACE 必须是接口类型。未指定名称时使用 ADDIN_IFACE 的完整限定名；
// creator 为 nil 或名称无效时会 panic。
func AddIn[ADDIN_IFACE, SETTING any](creator generic.FuncVar0[SETTING, ADDIN_IFACE], name ...string) AddInDefinition[ADDIN_IFACE, SETTING] {
	return defineAddIn[ADDIN_IFACE, SETTING](creator, pie.First(name))
}

// AddInDefinition 封装通用插件的标识、构造和访问操作。
type AddInDefinition[ADDIN_IFACE, SETTING any] struct {
	Id        uint64                                                        // 由 Name 生成的插件 ID。
	Name      string                                                        // 插件注册名称。
	Install   generic.ActionVar1[extension.AddInProvider, SETTING]          // 构造插件并安装到给定提供者。
	Uninstall generic.Action1[extension.AddInProvider]                      // 从给定提供者卸载插件。
	Require   generic.Func1[extension.AddInProvider, ADDIN_IFACE]           // 获取正在运行的插件，不可用时 panic。
	Lookup    generic.FuncPair1[extension.AddInProvider, ADDIN_IFACE, bool] // 查询管理器当前持有的插件。
}

func defineAddIn[ADDIN_IFACE, SETTING any](creator generic.FuncVar0[SETTING, ADDIN_IFACE], name string) AddInDefinition[ADDIN_IFACE, SETTING] {
	if creator == nil {
		exception.Panicf("%w: %w: creator is nil", exception.ErrCore, exception.ErrArgs)
	}
	if name == "" {
		name = extension.GenAddInNameT[ADDIN_IFACE]()
	}
	if name == "" {
		exception.Panicf("%w: anonymous add-in not allowed", extension.ErrExtension)
	}
	id := extension.GenAddInId(name)

	return AddInDefinition[ADDIN_IFACE, SETTING]{
		Id:   id,
		Name: name,
		Install: func(provider extension.AddInProvider, settings ...SETTING) {
			extension.Install[ADDIN_IFACE](provider, creator(settings...), name)
		},
		Uninstall: func(provider extension.AddInProvider) {
			extension.Uninstall(provider, name)
		},
		Require: func(provider extension.AddInProvider) ADDIN_IFACE {
			return extension.Require[ADDIN_IFACE](provider, id)
		},
		Lookup: func(provider extension.AddInProvider) (ADDIN_IFACE, bool) {
			return extension.Lookup[ADDIN_IFACE](provider, id)
		},
	}
}
