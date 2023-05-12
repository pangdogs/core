package golaxy

import (
	"kit.golaxy.org/golaxy/ec"
	. "kit.golaxy.org/golaxy/runtime"
)

func Async(ctxResolver ec.ContextResolver, awaitRet AsyncRet, asyncWait func(ctx Context, ret Ret)) {
	ctx := Get(ctxResolver)

	if awaitRet == nil {
		panic("nil awaitRet")
	}

	if asyncWait == nil {
		panic("nil asyncWait")
	}

	go func() {
		defer func() {
			recover()
		}()

		ret, ok := <-awaitRet
		if !ok {
			return
		}

		ctx.AsyncCallNoRet(func() {
			asyncWait(ctx, ret)
		})
	}()
}

func Await(ctxResolver ec.ContextResolver, segment func(ctx Context) Ret) AsyncRet {
	ctx := Get(ctxResolver)

	if segment == nil {
		panic("nil segment")
	}

	return ctx.AsyncCall(func() Ret {
		return segment(ctx)
	})
}
