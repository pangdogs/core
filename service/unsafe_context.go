package service

import "github.com/galaxy-kit/galaxy-go/internal"

func UnsafeContext(ctx Context) _UnsafeContext {
	return _UnsafeContext{
		Context: ctx,
	}
}

type _UnsafeContext struct {
	Context
}

func (uc _UnsafeContext) Init(opts *ContextOptions) {
	uc.init(opts)
}

func (uc _UnsafeContext) GetOptions() *ContextOptions {
	return uc.getOptions()
}

func (uc _UnsafeContext) MarkRunning() bool {
	return internal.UnsafeRunningMark(uc.Context).MarkRunning()
}

func (uc _UnsafeContext) MarkShutdown() bool {
	return internal.UnsafeRunningMark(uc.Context).MarkShutdown()
}
