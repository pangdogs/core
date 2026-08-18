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
	"github.com/elliotchance/pie/v2"
)

// ServiceAddInInterface 创建不绑定构造函数的服务插件接口定义。
// 它适合通过 service.Context 访问同类插件的不同实现。
func ServiceAddInInterface[ADDIN_IFACE any](name ...string) ServiceAddInInterfaceDefinition[ADDIN_IFACE] {
	addIn := defineAddInInterface[ADDIN_IFACE](pie.First(name))

	return ServiceAddInInterfaceDefinition[ADDIN_IFACE]{
		ID:      addIn.ID,
		Name:    addIn.Name,
		Require: func(svcCtx service.Context) ADDIN_IFACE { return addIn.Require(svcCtx) },
		Lookup:  func(svcCtx service.Context) (ADDIN_IFACE, bool) { return addIn.Lookup(svcCtx) },
	}
}

// ServiceAddInInterfaceDefinition 封装服务插件接口的标识和访问操作。
type ServiceAddInInterfaceDefinition[ADDIN_IFACE any] struct {
	ID      uint64                                                // 由 Name 生成的插件 ID。
	Name    string                                                // 插件注册名称。
	Require generic.Func1[service.Context, ADDIN_IFACE]           // 从服务获取正在运行的插件，不可用时 panic。
	Lookup  generic.FuncPair1[service.Context, ADDIN_IFACE, bool] // 查询服务管理器当前持有的插件。
}
