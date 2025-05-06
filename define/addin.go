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

// AddIn 定义通用插件，支持运行时和服务上下文
func AddIn[ADDIN_IFACE, SETTING any](creator generic.FuncVar0[SETTING, ADDIN_IFACE]) AddInDefinition[ADDIN_IFACE, SETTING] {
	return defineAddIn[ADDIN_IFACE, SETTING](creator)
}

// AddInDefinition 通用插件定义
type AddInDefinition[ADDIN_IFACE, SETTING any] struct {
	Name      string                                               // 插件名称
	Install   generic.ActionVar1[extension.AddInProvider, SETTING] // 向插件管理器安装
	Uninstall generic.Action1[extension.AddInProvider]             // 从插件管理器卸载
	Using     generic.Func1[extension.AddInProvider, ADDIN_IFACE]  // 使用插件
}

func defineAddIn[ADDIN_IFACE, SETTING any](creator generic.FuncVar0[SETTING, ADDIN_IFACE]) AddInDefinition[ADDIN_IFACE, SETTING] {
	if creator == nil {
		exception.Panicf("%w: %w: creator is nil", exception.ErrCore, exception.ErrArgs)
	}

	name := types.FullNameT[ADDIN_IFACE]()

	return AddInDefinition[ADDIN_IFACE, SETTING]{
		Name: name,
		Install: func(provider extension.AddInProvider, options ...SETTING) {
			extension.Install[ADDIN_IFACE](provider, creator(options...), name)
		},
		Uninstall: func(provider extension.AddInProvider) {
			extension.Uninstall(provider, name)
		},
		Using: func(provider extension.AddInProvider) ADDIN_IFACE {
			return extension.Using[ADDIN_IFACE](provider, name)
		},
	}
}
