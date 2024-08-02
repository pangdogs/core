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

type (
	AsyncRet = AsyncRetT[any]
)

var (
	MakeAsyncRet = MakeAsyncRetT[any]
)

// MakeAsyncRetT 创建异步调用结果
func MakeAsyncRetT[T any]() chan RetT[T] {
	return make(chan RetT[T], 1)
}

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
