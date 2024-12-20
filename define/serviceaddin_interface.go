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

// ServiceAddInInterface 定义服务插件接口，支持服务上下文，通常用于为同类插件的不同实现提供统一的接口
func ServiceAddInInterface[ADDIN_IFACE any]() ServiceAddInInterfaceDefinition[ADDIN_IFACE] {
	plug := defineAddInInterface[ADDIN_IFACE]()

	return ServiceAddInInterfaceDefinition[ADDIN_IFACE]{
		Name:  plug.Name,
		Using: func(svcCtx service.Context) ADDIN_IFACE { return plug.Using(svcCtx) },
	}
}

// ServiceAddInInterfaceDefinition 服务插件接口定义
type ServiceAddInInterfaceDefinition[ADDIN_IFACE any] struct {
	Name  string                                      // 插件名称
	Using generic.Func1[service.Context, ADDIN_IFACE] // 使用插件
}
