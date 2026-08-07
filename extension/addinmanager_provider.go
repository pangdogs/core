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
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/types"
)

// AddInProvider 提供插件管理器。
type AddInProvider interface {
	// AddInManager 返回关联的插件管理器。
	AddInManager() AddInManager
}

// Install 将 addIn 安装到 provider，并返回插件状态。
// provider 为 nil 时会 panic；具体生命周期和并发约束由管理器决定。
func Install[T any](provider AddInProvider, addIn T, name ...string) AddInStatus {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrExtension, exception.ErrArgs)
	}
	return provider.AddInManager().Install(iface.NewFaceAny(addIn), name...)
}

// Uninstall 从 provider 卸载指定名称的插件。
func Uninstall(provider AddInProvider, name string) {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrExtension, exception.ErrArgs)
	}
	provider.AddInManager().Uninstall(name)
}

// Require 返回指定 ID 且处于 Running 状态的插件实例。
// 插件不存在、尚未运行、已经卸载或 provider 为 nil 时会 panic。
// T 必须是插件实际实例实现的接口类型。
func Require[T any](provider AddInProvider, id uint64) T {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrExtension, exception.ErrArgs)
	}

	status, ok := provider.AddInManager().GetStatusById(id)
	if !ok {
		exception.Panicf("%w: add-in id %d not installed", ErrExtension, id)
	}

	if status.State() != AddInState_Running {
		exception.Panicf("%w: add-in id %d not actived", ErrExtension, id)
	}

	return iface.Cache2Iface[T](status.InstanceFace().Cache)
}

// Lookup 查询指定 ID 的插件实例。
// 只要管理器仍持有该插件就会返回 true，不要求插件处于 Running 状态。
// T 必须是插件实际实例实现的接口类型。
func Lookup[T any](provider AddInProvider, id uint64) (T, bool) {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrExtension, exception.ErrArgs)
	}

	status, ok := provider.AddInManager().GetStatusById(id)
	if !ok {
		return types.Zero[T](), false
	}

	return iface.Cache2Iface[T](status.InstanceFace().Cache), true
}
