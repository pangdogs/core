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

// Yield 产出异步调用结果
func Yield(ctx context.Context, asyncRet chan Ret, ret Ret) bool {
	return YieldT(ctx, asyncRet, ret)
}

// YieldT 产出异步调用结果
func YieldT[T any](ctx context.Context, asyncRet chan RetT[T], ret RetT[T]) bool {
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

// End 结束产出异步调用结果
func End(asyncRet chan Ret) {
	EndT(asyncRet)
}

// EndT 结束产出异步调用结果
func EndT[T any](asyncRet chan RetT[T]) {
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
		return MakeRetT[T](types.ZeroT[T](), context.Canceled)
	}
}
