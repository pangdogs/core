package internal

import "context"

func UnsafeContext(ctx Context) _UnsafeContext {
	return _UnsafeContext{
		Context: ctx,
	}
}

type _UnsafeContext struct {
	Context
}

func (uc _UnsafeContext) Init(parentCtx context.Context, autoRecover bool, reportError chan error) {
	uc.init(parentCtx, autoRecover, reportError)
}

func (uc _UnsafeContext) Paired() bool {
	return uc.paired()
}
