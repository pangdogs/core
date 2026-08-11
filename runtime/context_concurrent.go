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
	"git.golaxy.org/core/utils/async"
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
	// ExecutorID 返回进程内 Runtime 执行器 ID。
	ExecutorID() async.ExecutorID
	// BlockedFutureID 返回当前阻塞等待的 Future ID。
	BlockedFutureID() uint64
	// LastWaitRejectID 返回最近一次被 Runtime 等待规则拒绝的 Future ID。
	LastWaitRejectID() uint64
}

type iConcurrentContext interface {
	getInstance() Context
}

// ConcurrentContextCache 返回可跨 goroutine 使用的上下文接口缓存。
func (ctx *ContextBehavior) ConcurrentContextCache() iface.Cache {
	return iface.Iface2Cache[Context](ctx.options.InstanceFace.Iface)
}

// ExecutorID 返回用于 Future 完成归属和 Runtime 自等待检测的进程内执行器 ID。
func (ctx *ContextBehavior) ExecutorID() async.ExecutorID {
	return ctx.executorID
}

// BlockedFutureID 返回当前阻塞等待的 Future ID。
func (ctx *ContextBehavior) BlockedFutureID() uint64 {
	return ctx.blockedFuture.Load()
}

// LastWaitRejectID 返回最近一次被 Runtime 等待规则拒绝的 Future ID。
func (ctx *ContextBehavior) LastWaitRejectID() uint64 {
	return ctx.lastWaitReject.Load()
}

// String 实现 fmt.Stringer，返回包含运行时 ID 和名称的 JSON 文本。
func (ctx *ContextBehavior) String() string {
	if cached := ctx.stringerCache.Load(); cached != nil {
		return *cached
	}

	value := fmt.Sprintf(`{"id":%q,"name":%q}`, ctx.Id(), ctx.Name())
	if ctx.stringerCache.CompareAndSwap(nil, &value) {
		return value
	}
	return *ctx.stringerCache.Load()
}

func (ctx *ContextBehavior) getInstance() Context {
	return ctx.options.InstanceFace.Iface
}

// Concurrent 从 provider 获取可跨 goroutine 使用的运行时上下文；provider 为 nil 时会 panic。
func Concurrent(provider corectx.ConcurrentContextProvider) ConcurrentContext {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrContext, exception.ErrArgs)
	}
	return iface.Cache2Iface[Context](provider.ConcurrentContextCache())
}

func getServiceContext(provider corectx.ConcurrentContextProvider) service.Context {
	if provider == nil {
		exception.Panicf("%w: %w: provider is nil", ErrContext, exception.ErrArgs)
	}
	ctx := iface.Cache2Iface[Context](provider.ConcurrentContextCache())
	if ctx == nil {
		return nil
	}
	return ctx.getServiceContext()
}
