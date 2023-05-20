package service

import "kit.golaxy.org/golaxy/internal"

func UnsafeContext(ctx Context) _UnsafeContext {
	return _UnsafeContext{
		Context: ctx,
	}
}

type _UnsafeContext struct {
	Context
}

func (uc _UnsafeContext) Init(opts *Options) {
	uc.Context.init(opts)
}

func (uc _UnsafeContext) GetOptions() *Options {
	return uc.getOptions()
}

func (uc _UnsafeContext) MarkRunning(v bool) bool {
	return internal.UnsafeRunningState(uc.Context).MarkRunning(v)
}
