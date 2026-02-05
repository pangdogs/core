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

// AddInProvider 插件提供者
type AddInProvider interface {
	// AddInManager 获取插件管理器
	AddInManager() AddInManager
}

// Resolve 解析插件
func Resolve[T any](provider AddInProvider, name string) T {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrExtension, exception.ErrArgs)
	}

	status, ok := provider.AddInManager().Get(name)
	if !ok {
		exception.Panicf("%w: addIn %q not installed", ErrExtension, name)
	}

	if status.State() != AddInState_Running {
		exception.Panicf("%w: addIn %q not actived", ErrExtension, name)
	}

	return iface.Cache2Iface[T](status.InstanceFace().Cache)
}

// Lookup 查找插件
func Lookup[T any](provider AddInProvider, name string) (T, bool) {
	if provider == nil {
		return types.ZeroT[T](), false
	}

	status, ok := provider.AddInManager().Get(name)
	if !ok {
		return types.ZeroT[T](), false
	}

	return iface.Cache2Iface[T](status.InstanceFace().Cache), true
}

// Install 安装插件
func Install[T any](provider AddInProvider, addIn T, name ...string) AddInStatus {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrExtension, exception.ErrArgs)
	}
	return provider.AddInManager().Install(iface.NewFaceAny(addIn), name...)
}

// Uninstall 卸载插件
func Uninstall(provider AddInProvider, name string) {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrExtension, exception.ErrArgs)
	}
	provider.AddInManager().Uninstall(name)
}
