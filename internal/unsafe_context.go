package internal

import "context"

// Deprecated: UnsafeContext 访问上下文内部方法
func UnsafeContext(ctx Context) _UnsafeContext {
	return _UnsafeContext{
		Context: ctx,
	}
}

type _UnsafeContext struct {
	Context
}

// Init 初始化
func (uc _UnsafeContext) Init(parentCtx context.Context, autoRecover bool, reportError chan error) {
	uc.init(parentCtx, autoRecover, reportError)
}

// MarkPaired 标记已经配对
func (uc _UnsafeContext) MarkPaired(v bool) bool {
	return uc.markPaired(v)
}
