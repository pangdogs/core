package golaxy

import (
	"errors"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/runtime"
	"sync/atomic"
)

func Await(ctxResolver ec.ContextResolver, asyncWait func(ctx runtime.Context, ret runtime.Ret), asyncRet runtime.AsyncRet) {
	ctx := runtime.Get(ctxResolver)

	if asyncWait == nil {
		panic("nil asyncWait")
	}

	if asyncRet == nil {
		return
	}

	go func() {
		defer func() {
			recover()
		}()

		ret, ok := <-asyncRet
		if !ok {
			ret = runtime.NewRet(errors.New("asyncRet closed"), nil)
		}

		ctx.AsyncCallNoRet(func() {
			asyncWait(ctx, ret)
		})
	}()
}

func AwaitAny(ctxResolver ec.ContextResolver, asyncWait func(ctx runtime.Context, ret runtime.Ret), asyncRets ...runtime.AsyncRet) {
	ctx := runtime.Get(ctxResolver)

	if asyncWait == nil {
		panic("nil asyncWait")
	}

	if len(asyncRets) <= 0 {
		return
	}

	var b atomic.Bool

	for _, asyncRet := range asyncRets {
		if asyncRet == nil {
			continue
		}

		go func() {
			defer func() {
				recover()
			}()

			ret, ok := <-asyncRet
			if !ok {
				return
			}

			if !ret.OK() {
				return
			}

			if !b.CompareAndSwap(false, true) {
				return
			}

			ctx.AsyncCallNoRet(func() {
				asyncWait(ctx, ret)
			})
		}()
	}
}

func AwaitAll(ctxResolver ec.ContextResolver, asyncWait func(ctx runtime.Context, ret runtime.Ret), asyncRets ...runtime.AsyncRet) {
	ctx := runtime.Get(ctxResolver)

	if asyncWait == nil {
		panic("nil asyncWait")
	}

	if len(asyncRets) <= 0 {
		return
	}

	for _, asyncRet := range asyncRets {
		if asyncRet == nil {
			continue
		}

		go func() {
			defer func() {
				recover()
			}()

			ret, ok := <-asyncRet
			if !ok {
				ret = runtime.NewRet(errors.New("asyncRet closed"), nil)
			}

			ctx.AsyncCallNoRet(func() {
				asyncWait(ctx, ret)
			})
		}()
	}
}

func Async(ctxResolver ec.ContextResolver, segment func(ctx runtime.Context) runtime.Ret) runtime.AsyncRet {
	ctx := runtime.Get(ctxResolver)

	if segment == nil {
		panic("nil segment")
	}

	return ctx.AsyncCall(func() runtime.Ret {
		return segment(ctx)
	})
}

func AsyncVoid(ctxResolver ec.ContextResolver, segment func(ctx runtime.Context)) runtime.AsyncRet {
	ctx := runtime.Get(ctxResolver)

	if segment == nil {
		panic("nil segment")
	}

	return ctx.AsyncCall(func() runtime.Ret {
		segment(ctx)
		return runtime.NewRet(nil, nil)
	})
}
