package runtime

import (
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/service"
)

func UnsafeContext(ctx Context) _UnsafeContext {
	return _UnsafeContext{
		Context: ctx,
	}
}

type _UnsafeContext struct {
	Context
}

func (uc _UnsafeContext) Init(serviceCtx service.Context, opts *ContextOptions) {
	uc.Context.init(serviceCtx, opts)
}

func (uc _UnsafeContext) GetOptions() *ContextOptions {
	return uc.getOptions()
}

func (uc _UnsafeContext) SetFrame(frame Frame) {
	uc.setFrame(frame)
}

func (uc _UnsafeContext) SetCallee(callee Callee) {
	uc.setCallee(callee)
}

func (uc _UnsafeContext) GetServiceCtx() service.Context {
	return uc.getServiceCtx()
}

func (uc _UnsafeContext) GC() {
	uc.gc()
}

func (uc _UnsafeContext) MarkRunning(v bool) bool {
	return internal.UnsafeRunningState(uc.Context).MarkRunning(v)
}

func (uc _UnsafeContext) MarkPaired(v bool) bool {
	return internal.UnsafeContext(uc.Context).MarkPaired(v)
}
