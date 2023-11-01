package service

import "kit.golaxy.org/golaxy/internal"

// Deprecated: UnsafeContext 访问服务上下文内部方法
func UnsafeContext(ctx Context) _UnsafeContext {
	return _UnsafeContext{
		Context: ctx,
	}
}

type _UnsafeContext struct {
	Context
}

// Init 初始化
func (uc _UnsafeContext) Init(opts ContextOptions) {
	uc.Context.init(opts)
}

// GetOptions 获取服务上下文所有选项
func (uc _UnsafeContext) GetOptions() *ContextOptions {
	return uc.getOptions()
}

// MarkRunning 标记服务已经开始运行
func (uc _UnsafeContext) MarkRunning(v bool) bool {
	return internal.UnsafeRunningState(uc.Context).MarkRunning(v)
}

// MarkPaired 标记与服务已经配对
func (uc _UnsafeContext) MarkPaired(v bool) bool {
	return internal.UnsafeContext(uc.Context).MarkPaired(v)
}
