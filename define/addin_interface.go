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

// AddInInterface 创建不绑定构造函数的通用插件接口定义。
// 它适合用同一名称约定访问同类插件的不同实现。
func AddInInterface[ADDIN_IFACE any](name ...string) AddInInterfaceDefinition[ADDIN_IFACE] {
	return defineAddInInterface[ADDIN_IFACE](pie.First(name))
}

// AddInInterfaceDefinition 封装通用插件接口的标识和访问操作。
type AddInInterfaceDefinition[ADDIN_IFACE any] struct {
	ID      uint64                                                        // 由 Name 生成的插件 ID。
	Name    string                                                        // 插件注册名称。
	Require generic.Func1[extension.AddInProvider, ADDIN_IFACE]           // 获取正在运行的插件，不可用时 panic。
	Lookup  generic.FuncPair1[extension.AddInProvider, ADDIN_IFACE, bool] // 查询管理器当前持有的插件。
}

func defineAddInInterface[ADDIN_IFACE any](name string) AddInInterfaceDefinition[ADDIN_IFACE] {
	if name == "" {
		name = extension.GenAddInNameT[ADDIN_IFACE]()
	}
	if name == "" {
		exception.Panicf("%w: anonymous add-in not allowed", extension.ErrExtension)
	}
	id := extension.GenAddInID(name)

	return AddInInterfaceDefinition[ADDIN_IFACE]{
		ID:   id,
		Name: name,
		Require: func(provider extension.AddInProvider) ADDIN_IFACE {
			return extension.Require[ADDIN_IFACE](provider, id)
		},
		Lookup: func(provider extension.AddInProvider) (ADDIN_IFACE, bool) {
			return extension.Lookup[ADDIN_IFACE](provider, id)
		},
	}
}
