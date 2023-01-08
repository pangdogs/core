package galaxy

import "github.com/golaxy-kit/golaxy/runtime"

func UnsafeRuntime(runtime Runtime) _UnsafeRuntime {
	return _UnsafeRuntime{
		Runtime: runtime,
	}
}

type _UnsafeRuntime struct {
	Runtime
}

func (ur _UnsafeRuntime) Init(runtimeCtx runtime.Context, opts *RuntimeOptions) {
	ur.init(runtimeCtx, opts)
}

func (ur _UnsafeRuntime) GetOptions() *RuntimeOptions {
	return ur.getOptions()
}
