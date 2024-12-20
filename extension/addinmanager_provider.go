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
)

// AddInProvider 插件提供者
type AddInProvider interface {
	// GetAddInManager 获取插件管理器
	GetAddInManager() AddInManager
}

// Using 使用插件
func Using[T any](provider AddInProvider, name string) T {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrExtension, exception.ErrArgs)
	}

	status, ok := provider.GetAddInManager().Get(name)
	if !ok {
		exception.Panicf("%w: addIn %q not installed", ErrExtension, name)
	}

	if status.State() != AddInState_Active {
		exception.Panicf("%w: addIn %q not actived", ErrExtension, name)
	}

	return iface.Cache2Iface[T](status.InstanceFace().Cache)
}

// Install 安装插件
func Install[T any](provider AddInProvider, addIn T, name ...string) {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrExtension, exception.ErrArgs)
	}
	provider.GetAddInManager().Install(iface.MakeFaceAny(addIn), name...)
}

// Uninstall 卸载插件
func Uninstall(provider AddInProvider, name string) {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrExtension, exception.ErrArgs)
	}
	provider.GetAddInManager().Uninstall(name)
}
