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

package ectx

import (
	"git.golaxy.org/core/utils/iface"
)

// CurrentContextProvider 当前上下文提供者
type CurrentContextProvider interface {
	ConcurrentContextProvider
	// GetCurrentContext 获取当前上下文
	GetCurrentContext() iface.Cache
}

// ConcurrentContextProvider 多线程安全的上下文提供者
type ConcurrentContextProvider interface {
	// GetConcurrentContext 获取多线程安全的上下文
	GetConcurrentContext() iface.Cache
}
