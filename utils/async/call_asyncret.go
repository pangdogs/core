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
