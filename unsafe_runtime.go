package golaxy

import "kit.golaxy.org/golaxy/runtime"

func UnsafeRuntime(runtime Runtime) _UnsafeRuntime {
	return _UnsafeRuntime{
		Runtime: runtime,
	}
}

type _UnsafeRuntime struct {
	Runtime
}

func (ur _UnsafeRuntime) Init(ctx runtime.Context, opts *RuntimeOptions) {
	ur.init(ctx, opts)
}

func (ur _UnsafeRuntime) GetOptions() *RuntimeOptions {
	return ur.getOptions()
}
