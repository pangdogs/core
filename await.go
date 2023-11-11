package golaxy

import (
	"context"
	"fmt"
	"kit.golaxy.org/golaxy/internal/concurrent"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/generic"
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
func (ad AwaitDirector) Any(ctxResolver service.ContextResolver, fun generic.ActionVar2[runtime.Context, runtime.Ret, any], va ...any) {
	if ad.context == nil {
		ad.context = context.Background()
	}

	if len(ad.asyncRets) <= 0 {
		return
	}

	runtimeCtx := getContext(ctxResolver)

	var b atomic.Bool
	ctx, cancel := context.WithCancel(ad.context)

	for i := range ad.asyncRets {
		asyncRet := ad.asyncRets[i]
		if asyncRet == nil {
			continue
		}

		go func(b *atomic.Bool, ctx context.Context, cancel context.CancelFunc,
			asyncRet runtime.AsyncRet, runtimeCtx runtime.Context, fun generic.ActionVar2[runtime.Context, runtime.Ret, any], va []any) {

			ret := asyncRet.Wait(ctx)

			if !b.CompareAndSwap(false, true) {
				return
			}

			cancel()

			runtimeCtx.CallVoid(func(va ...any) {
				runtimeCtx := va[0].(runtime.Context)
				fun := va[1].(generic.ActionVar2[runtime.Context, runtime.Ret, any])
				ret := va[2].(runtime.Ret)
				funVa := va[3].([]any)
				fun.Exec(runtimeCtx, ret, funVa...)
			}, runtimeCtx, fun, ret, va)
		}(&b, ctx, cancel, asyncRet, runtimeCtx, fun, va)
	}
}

// AnyOK 异步等待任意一个结果成功返回
func (ad AwaitDirector) AnyOK(ctxResolver service.ContextResolver, fun generic.ActionVar2[runtime.Context, runtime.Ret, any], va ...any) {
	if ad.context == nil {
		ad.context = context.Background()
	}

	if len(ad.asyncRets) <= 0 {
		return
	}

	runtimeCtx := getContext(ctxResolver)

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
			asyncRet runtime.AsyncRet, runtimeCtx runtime.Context, fun generic.ActionVar2[runtime.Context, runtime.Ret, any], va []any) {
			defer wg.Done()

			ret := asyncRet.Wait(ctx)
			if !ret.OK() {
				return
			}

			if !b.CompareAndSwap(false, true) {
				return
			}

			cancel()

			runtimeCtx.CallVoid(func(va ...any) {
				runtimeCtx := va[0].(runtime.Context)
				fun := va[1].(generic.ActionVar2[runtime.Context, runtime.Ret, any])
				ret := va[2].(runtime.Ret)
				funVa := va[3].([]any)
				fun.Exec(runtimeCtx, ret, funVa...)
			}, runtimeCtx, fun, ret, va)
		}(&wg, &b, ctx, cancel, asyncRet, runtimeCtx, fun, va)
	}

	go func(wg *sync.WaitGroup, b *atomic.Bool,
		runtimeCtx runtime.Context, fun generic.ActionVar2[runtime.Context, runtime.Ret, any], va []any) {
		wg.Wait()

		if b.Load() {
			return
		}

		runtimeCtx.CallVoid(func(va ...any) {
			runtimeCtx := va[0].(runtime.Context)
			fun := va[1].(generic.ActionVar2[runtime.Context, runtime.Ret, any])
			funVa := va[2].([]any)
			fun.Exec(runtimeCtx, concurrent.MakeRet(nil, ErrAllFailures), funVa...)
		}, runtimeCtx, fun, va)
	}(&wg, &b, runtimeCtx, fun, va)
}

// All 异步等待所有结果返回
func (ad AwaitDirector) All(ctxResolver service.ContextResolver, fun generic.ActionVar2[runtime.Context, []runtime.Ret, any], va ...any) {
	if ad.context == nil {
		ad.context = context.Background()
	}

	if len(ad.asyncRets) <= 0 {
		return
	}

	runtimeCtx := getContext(ctxResolver)

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

	go func(wg *sync.WaitGroup, runtimeCtx runtime.Context, fun generic.ActionVar2[runtime.Context, []runtime.Ret, any], rets []runtime.Ret, va []any) {
		wg.Wait()

		runtimeCtx.CallVoid(func(va ...any) {
			runtimeCtx := va[0].(runtime.Context)
			fun := va[1].(generic.ActionVar2[runtime.Context, []runtime.Ret, any])
			rets := va[2].([]runtime.Ret)
			funVa := va[3].([]any)
			fun.Exec(runtimeCtx, rets, funVa...)
		}, runtimeCtx, fun, rets, va)
	}(&wg, runtimeCtx, fun, rets, va)
}
