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

package plugin

import (
	"fmt"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
)

// PluginProvider 插件提供者
type PluginProvider interface {
	// GetPluginBundle 获取插件包
	GetPluginBundle() PluginBundle
}

// Using 使用插件
func Using[T any](provider PluginProvider, name string) T {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrPlugin, exception.ErrArgs))
	}

	status, ok := provider.GetPluginBundle().Get(name)
	if !ok {
		panic(fmt.Errorf("%w: plugin %q not installed", ErrPlugin, name))
	}

	if status.State() != PluginState_Active {
		panic(fmt.Errorf("%w: plugin %q not actived", ErrPlugin, name))
	}

	return iface.Cache2Iface[T](status.InstanceFace().Cache)
}

// Install 安装插件
func Install[T any](provider PluginProvider, plugin T, name ...string) {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrPlugin, exception.ErrArgs))
	}
	provider.GetPluginBundle().Install(iface.MakeFaceAny(plugin), name...)
}

// Uninstall 卸载插件
func Uninstall(provider PluginProvider, name string) {
	if provider == nil {
		panic(fmt.Errorf("%w: %w: provider is nil", ErrPlugin, exception.ErrArgs))
	}
	provider.GetPluginBundle().Uninstall(name)
}
