package golaxy

import "kit.golaxy.org/golaxy/runtime"

// Deprecated: UnsafeRuntime 访问运行时内部方法
func UnsafeRuntime(runtime Runtime) _UnsafeRuntime {
	return _UnsafeRuntime{
		Runtime: runtime,
	}
}

type _UnsafeRuntime struct {
	Runtime
}

// Init 初始化
func (ur _UnsafeRuntime) Init(ctx runtime.Context, opts RuntimeOptions) {
	ur.init(ctx, opts)
}

// GetOptions 获取运行时所有选项
func (ur _UnsafeRuntime) GetOptions() *RuntimeOptions {
	return ur.getOptions()
}
