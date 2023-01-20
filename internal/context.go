package internal

import (
	"context"
	"sync"
	"sync/atomic"
)

// Context 上下文
type Context interface {
	context.Context
	// GetParentContext 获取父上下文
	GetParentContext() context.Context
	// GetAutoRecover panic时是否自动恢复
	GetAutoRecover() bool
	// GetReportError 在开启panic时自动恢复时，将会恢复并将错误写入此error channel
	GetReportError() chan error
	// GetWaitGroup 获取等待组
	GetWaitGroup() *sync.WaitGroup
	// GetCancelFunc 获取取消运行函数
	GetCancelFunc() context.CancelFunc

	init(parentCtx context.Context, autoRecover bool, reportError chan error)
	paired() bool
}

// ContextBehavior 上下文行为
type ContextBehavior struct {
	context.Context
	parentCtx   context.Context
	autoRecover bool
	reportError chan error
	cancel      context.CancelFunc
	wg          sync.WaitGroup
	_paired     atomic.Bool
}

// GetParentContext 获取父上下文
func (ctx *ContextBehavior) GetParentContext() context.Context {
	return ctx.parentCtx
}

// GetAutoRecover panic时是否自动恢复
func (ctx *ContextBehavior) GetAutoRecover() bool {
	return ctx.autoRecover
}

// GetReportError 在开启panic时自动恢复时，将会恢复并将错误写入此error channel
func (ctx *ContextBehavior) GetReportError() chan error {
	return ctx.reportError
}

// GetWaitGroup 获取等待组
func (ctx *ContextBehavior) GetWaitGroup() *sync.WaitGroup {
	return &ctx.wg
}

// GetCancelFunc 获取取消运行函数
func (ctx *ContextBehavior) GetCancelFunc() context.CancelFunc {
	return ctx.cancel
}

func (ctx *ContextBehavior) init(parentCtx context.Context, autoRecover bool, reportError chan error) {
	if parentCtx == nil {
		ctx.parentCtx = context.Background()
	} else {
		ctx.parentCtx = parentCtx
	}

	ctx.autoRecover = autoRecover
	ctx.reportError = reportError

	ctx.Context, ctx.cancel = context.WithCancel(ctx.parentCtx)
}

func (ctx *ContextBehavior) paired() bool {
	return ctx._paired.CompareAndSwap(false, true)
}
