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
	"fmt"

	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/uid"
)

// ConcurrentContextProvider 提供可跨 goroutine 使用的运行时上下文缓存。
type ConcurrentContextProvider = corectx.ConcurrentContextProvider

// ConcurrentContext 暴露可跨 goroutine 安全调用的运行时能力。
// 需要读写实体或其他运行时局部状态时，应通过 Caller 调度回运行时 goroutine。
type ConcurrentContext interface {
	iConcurrentContext
	corectx.Context
	corectx.ConcurrentContextProvider
	Caller
	fmt.Stringer

	// Name 返回运行时名称。
	Name() string
	// Id 返回运行时的持久化 ID。
	Id() uid.Id
}

type iConcurrentContext interface {
	getContext() Context
}

func (ctx *ContextBehavior) getContext() Context {
	return ctx.options.InstanceFace.Iface
}

// Concurrent 从 provider 获取可跨 goroutine 使用的运行时上下文；provider 为 nil 时会 panic。
func Concurrent(provider corectx.ConcurrentContextProvider) ConcurrentContext {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrContext, exception.ErrArgs)
	}
	return iface.Cache2Iface[Context](provider.ConcurrentContext())
}

func getServiceContext(provider corectx.ConcurrentContextProvider) service.Context {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrContext, exception.ErrArgs)
	}
	ctx := iface.Cache2Iface[Context](provider.ConcurrentContext())
	if ctx == nil {
		return nil
	}
	return ctx.getServiceContext()
}
