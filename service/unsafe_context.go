package service

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
