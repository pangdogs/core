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

package reinterpret

import "git.golaxy.org/core/utils/iface"

// InstanceProvider 提供实际实例的接口缓存，以支持快速重解释。
type InstanceProvider interface {
	// InstanceFaceCache 返回实际实例的接口缓存。
	InstanceFaceCache() iface.Cache
}

// Cast 将 cp 缓存的实际实例重解释为 T。
//
// 该函数不执行类型兼容性检查；cp 必须非 nil 且实例必须实现 T。
func Cast[T any](cp InstanceProvider) T {
	return iface.Cache2Iface[T](cp.InstanceFaceCache())
}
