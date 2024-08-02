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

package core

import (
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
)

// ServiceOptions 创建服务的所有选项
type ServiceOptions struct {
	CompositeFace iface.Face[Service] // 扩展者，在扩展服务自身功能时使用
}

type _ServiceOption struct{}

// Default 默认值
func (_ServiceOption) Default() option.Setting[ServiceOptions] {
	return func(o *ServiceOptions) {
		With.Service.CompositeFace(iface.Face[Service]{})(o)
	}
}

// CompositeFace 扩展者，在扩展服务自身功能时使用
func (_ServiceOption) CompositeFace(face iface.Face[Service]) option.Setting[ServiceOptions] {
	return func(o *ServiceOptions) {
		o.CompositeFace = face
	}
}
