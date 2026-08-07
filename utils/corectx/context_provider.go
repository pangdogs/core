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

package corectx

import (
	"git.golaxy.org/core/utils/iface"
)

// CurrentContextProvider 提供只能在所属执行协程中使用的当前上下文。
type CurrentContextProvider interface {
	ConcurrentContextProvider
	// CurrentContext 返回当前上下文的接口缓存。
	CurrentContext() iface.Cache
}

// ConcurrentContextProvider 提供可跨协程使用的并发上下文。
type ConcurrentContextProvider interface {
	// ConcurrentContext 返回并发上下文的接口缓存。
	ConcurrentContext() iface.Cache
}
