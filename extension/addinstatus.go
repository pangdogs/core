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

package extension

import (
	"fmt"
	"reflect"

	"git.golaxy.org/core/utils/iface"
)

// AddInStatus 描述一个已安装或曾经安装的插件实例及其生命周期状态。
type AddInStatus interface {
	fmt.Stringer

	// Id 返回由插件名称生成的 ID。
	Id() uint64
	// Name 返回插件注册名称。
	Name() string
	// InstanceFace 返回插件实例及其接口缓存。
	InstanceFace() iface.FaceAny
	// Reflected 返回插件实例的反射值。
	Reflected() reflect.Value
	// State 返回插件当前的生命周期状态。
	State() AddInState
}
