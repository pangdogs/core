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

package service

import (
	_ "unsafe"

	"git.golaxy.org/core/utils/corectx"
)

//go:linkname getServiceContext git.golaxy.org/core/runtime.getServiceContext
func getServiceContext(provider corectx.ConcurrentContextProvider) Context

// Current 返回 provider 所属运行时的服务上下文。
// provider 为 nil 时会 panic；provider 未关联运行时上下文时返回 nil。
func Current(provider corectx.ConcurrentContextProvider) Context {
	return getServiceContext(provider)
}
