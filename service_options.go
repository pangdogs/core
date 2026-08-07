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

// ServiceOptions 定义创建服务时使用的选项。
type ServiceOptions struct {
	InstanceFace iface.Face[Service] // 自定义服务实例及其接口缓存。
}

type _ServiceOption struct{}

// Default 返回服务选项的默认设置。
func (_ServiceOption) Default() option.Setting[ServiceOptions] {
	return func(options *ServiceOptions) {
		With.Service.InstanceFace(iface.Face[Service]{}).Apply(options)
	}
}

// InstanceFace 设置用于扩展服务能力的自定义实例。
func (_ServiceOption) InstanceFace(face iface.Face[Service]) option.Setting[ServiceOptions] {
	return func(options *ServiceOptions) {
		options.InstanceFace = face
	}
}
