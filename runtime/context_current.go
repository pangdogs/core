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

package runtime

import (
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
)

// CurrentContextProvider 提供仅限所属运行时 goroutine 使用的上下文缓存。
type CurrentContextProvider = corectx.CurrentContextProvider

// Current 从 provider 获取当前运行时上下文；provider 为 nil 时会 panic。
func Current(provider corectx.CurrentContextProvider) Context {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrContext, exception.ErrArgs)
	}
	return iface.Cache2Iface[Context](provider.CurrentContextCache())
}

func getRuntimeContext(provider corectx.CurrentContextProvider) Context {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrContext, exception.ErrArgs)
	}
	return iface.Cache2Iface[Context](provider.CurrentContextCache())
}
