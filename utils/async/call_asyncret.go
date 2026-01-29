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

package async

import (
	"context"

	"git.golaxy.org/core/utils/types"
)

// MakeAsyncRet 创建异步调用结果
func MakeAsyncRet() chan Ret {
	return MakeAsyncRetT[any]()
}

// MakeAsyncRetT 创建异步调用结果
func MakeAsyncRetT[T any]() chan RetT[T] {
	return make(chan RetT[T], 1)
}

// Return 返回异步调用结果
func Return(asyncRet chan Ret, ret Ret) chan Ret {
	return ReturnT(asyncRet, ret)
}

// ReturnT 返回异步调用结果
func ReturnT[T any](asyncRet chan RetT[T], ret RetT[T]) chan RetT[T] {
	asyncRet <- ret
	close(asyncRet)
	return asyncRet
}

// YieldReturn 产出异步调用结果
func YieldReturn(ctx context.Context, asyncRet chan Ret, ret Ret) bool {
	return YieldReturnT(ctx, asyncRet, ret)
}

// YieldReturnT 产出异步调用结果
func YieldReturnT[T any](ctx context.Context, asyncRet chan RetT[T], ret RetT[T]) bool {
	if ctx == nil {
		ctx = context.Background()
	}

	select {
	case asyncRet <- ret:
		return true
	case <-ctx.Done():
		return false
	}
}

// YieldBreak 结束产出异步调用结果
func YieldBreak(asyncRet chan Ret) {
	YieldBreakT(asyncRet)
}

// YieldBreakT 结束产出异步调用结果
func YieldBreakT[T any](asyncRet chan RetT[T]) {
	close(asyncRet)
}

// AsyncRet 异步调用结果
type AsyncRet = AsyncRetT[any]

// AsyncRetT 异步调用结果
type AsyncRetT[T any] <-chan RetT[T]

// Wait 等待异步调用结果
func (asyncRet AsyncRetT[T]) Wait(ctx context.Context) RetT[T] {
	if ctx == nil {
		ctx = context.Background()
	}

	select {
	case ret, ok := <-asyncRet:
		if !ok {
			return MakeRetT[T](types.ZeroT[T](), ErrAsyncRetClosed)
		}
		return ret
	case <-ctx.Done():
		return MakeRetT[T](types.ZeroT[T](), ctx.Err())
	}
}

// Context 转为上下文，丢弃异步调用结果
func (asyncRet AsyncRetT[T]) Context(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		select {
		case <-ctx.Done():
			return
		case <-asyncRet:
			cancel()
		}
	}()

	return ctx
}
