package core

import (
	"context"
	"fmt"
	"git.golaxy.org/core/internal/concurrent"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/util/generic"
	"sync"
	"sync/atomic"
)

var (
	ErrAllFailures = fmt.Errorf("%w: all of async result failures", ErrGolaxy)
)

// Await 异步等待结果返回
func Await(ctx context.Context, asyncRet ...runtime.AsyncRet) AwaitDirector {
	return AwaitDirector{
		context:   ctx,
		asyncRets: asyncRet,
	}
}

// AwaitDirector 异步等待分发器
type AwaitDirector struct {
	context   context.Context
	asyncRets []runtime.AsyncRet
}

// Any 异步等待任意一个结果返回
func (ad AwaitDirector) Any(ctxProvider service.ContextProvider, fun generic.ActionVar2[runtime.Context, runtime.Ret, any], va ...any) {
	if ad.context == nil {
		ad.context = context.Background()
	}

	if len(ad.asyncRets) <= 0 {
		return
	}

	rtCtx := getRuntimeContext(ctxProvider)

	var b atomic.Bool
	ctx, cancel := context.WithCancel(ad.context)

	for i := range ad.asyncRets {
		asyncRet := ad.asyncRets[i]
		if asyncRet == nil {
			continue
		}

		go func(b *atomic.Bool, ctx context.Context, cancel context.CancelFunc,
			asyncRet runtime.AsyncRet, rtCtx runtime.Context, fun generic.ActionVar2[runtime.Context, runtime.Ret, any], va []any) {

			ret := asyncRet.Wait(ctx)

			if !b.CompareAndSwap(false, true) {
				return
			}

			cancel()

			rtCtx.CallVoid(func(va ...any) {
				rtCtx := va[0].(runtime.Context)
				fun := va[1].(generic.ActionVar2[runtime.Context, runtime.Ret, any])
				ret := va[2].(runtime.Ret)
				funVa := va[3].([]any)
				fun.Exec(rtCtx, ret, funVa...)
			}, rtCtx, fun, ret, va)
		}(&b, ctx, cancel, asyncRet, rtCtx, fun, va)
	}
}

// AnyOK 异步等待任意一个结果成功返回
func (ad AwaitDirector) AnyOK(ctxProvider service.ContextProvider, fun generic.ActionVar2[runtime.Context, runtime.Ret, any], va ...any) {
	if ad.context == nil {
		ad.context = context.Background()
	}

	if len(ad.asyncRets) <= 0 {
		return
	}

	rtCtx := getRuntimeContext(ctxProvider)

	var wg sync.WaitGroup
	var b atomic.Bool
	ctx, cancel := context.WithCancel(ad.context)

	for i := range ad.asyncRets {
		asyncRet := ad.asyncRets[i]
		if asyncRet == nil {
			continue
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup, b *atomic.Bool, ctx context.Context, cancel context.CancelFunc,
			asyncRet runtime.AsyncRet, rtCtx runtime.Context, fun generic.ActionVar2[runtime.Context, runtime.Ret, any], va []any) {
			defer wg.Done()

			ret := asyncRet.Wait(ctx)
			if !ret.OK() {
				return
			}

			if !b.CompareAndSwap(false, true) {
				return
			}

			cancel()

			rtCtx.CallVoid(func(va ...any) {
				rtCtx := va[0].(runtime.Context)
				fun := va[1].(generic.ActionVar2[runtime.Context, runtime.Ret, any])
				ret := va[2].(runtime.Ret)
				funVa := va[3].([]any)
				fun.Exec(rtCtx, ret, funVa...)
			}, rtCtx, fun, ret, va)
		}(&wg, &b, ctx, cancel, asyncRet, rtCtx, fun, va)
	}

	go func(wg *sync.WaitGroup, b *atomic.Bool,
		rtCtx runtime.Context, fun generic.ActionVar2[runtime.Context, runtime.Ret, any], va []any) {
		wg.Wait()

		if b.Load() {
			return
		}

		rtCtx.CallVoid(func(va ...any) {
			rtCtx := va[0].(runtime.Context)
			fun := va[1].(generic.ActionVar2[runtime.Context, runtime.Ret, any])
			funVa := va[2].([]any)
			fun.Exec(rtCtx, concurrent.MakeRet(nil, ErrAllFailures), funVa...)
		}, rtCtx, fun, va)
	}(&wg, &b, rtCtx, fun, va)
}

// All 异步等待所有结果返回
func (ad AwaitDirector) All(ctxProvider service.ContextProvider, fun generic.ActionVar2[runtime.Context, []runtime.Ret, any], va ...any) {
	if ad.context == nil {
		ad.context = context.Background()
	}

	if len(ad.asyncRets) <= 0 {
		return
	}

	rtCtx := getRuntimeContext(ctxProvider)

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(ad.context)
	rets := make([]runtime.Ret, len(ad.asyncRets))

	for i := range ad.asyncRets {
		asyncRet := ad.asyncRets[i]
		if asyncRet == nil {
			continue
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc, ret *runtime.Ret, asyncRet runtime.AsyncRet) {
			defer wg.Done()
			*ret = asyncRet.Wait(ctx)
		}(&wg, ctx, cancel, &rets[i], asyncRet)
	}

	go func(wg *sync.WaitGroup, rtCtx runtime.Context, fun generic.ActionVar2[runtime.Context, []runtime.Ret, any], rets []runtime.Ret, va []any) {
		wg.Wait()

		rtCtx.CallVoid(func(va ...any) {
			rtCtx := va[0].(runtime.Context)
			fun := va[1].(generic.ActionVar2[runtime.Context, []runtime.Ret, any])
			rets := va[2].([]runtime.Ret)
			funVa := va[3].([]any)
			fun.Exec(rtCtx, rets, funVa...)
		}, rtCtx, fun, rets, va)
	}(&wg, rtCtx, fun, rets, va)
}

// Pipe 异步等待管道返回
func (ad AwaitDirector) Pipe(ctxProvider service.ContextProvider, fun generic.ActionVar2[runtime.Context, runtime.Ret, any], va ...any) {
	if ad.context == nil {
		ad.context = context.Background()
	}

	if len(ad.asyncRets) <= 0 {
		return
	}

	rtCtx := getRuntimeContext(ctxProvider)

	for i := range ad.asyncRets {
		asyncRet := ad.asyncRets[i]
		if asyncRet == nil {
			continue
		}

		go func(ctx context.Context, rtCtx runtime.Context, asyncRet runtime.AsyncRet, fun generic.ActionVar2[runtime.Context, runtime.Ret, any], va []any) {
			for {
				select {
				case ret, ok := <-asyncRet:
					if !ok {
						return
					}
					rtCtx.CallVoid(func(va ...any) {
						rtCtx := va[0].(runtime.Context)
						fun := va[1].(generic.ActionVar2[runtime.Context, runtime.Ret, any])
						ret := va[2].(runtime.Ret)
						funVa := va[3].([]any)
						fun.Exec(rtCtx, ret, funVa...)
					}, rtCtx, fun, ret, va)
				case <-ctx.Done():
					return
				}
			}
		}(ad.context, rtCtx, asyncRet, fun, va)
	}
}
