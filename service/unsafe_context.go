package service

import "reflect"

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

// ChangeRunningState 修改运行状态
func (uc _UnsafeContext) ChangeRunningState(state RunningState) {
	uc.changeRunningState(state)
}

// GetReflected 获取反射值
func (uc _UnsafeContext) GetReflected() reflect.Value {
	return uc.getReflected()
}
