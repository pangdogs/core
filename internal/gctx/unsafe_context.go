package gctx

import (
	"context"
)

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

// SetPaired 设置配对标记
func (uc _UnsafeContext) SetPaired(v bool) bool {
	return uc.setPaired(v)
}

// GetPaired 获取配对标记
func (uc _UnsafeContext) GetPaired() bool {
	return uc.getPaired()
}

func (uc _UnsafeContext) GetTerminatedChan() chan struct{} {
	return uc.getTerminatedChan()
}
