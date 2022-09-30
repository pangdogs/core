package internal

import (
	"context"
	"sync"
)

// Context 上下文
type Context interface {
	context.Context

	// GetParentCtx 获取父上下文
	GetParentCtx() context.Context

	// GetReportError 在打开服务或运行时的AutoRecover选项时，panic时将会恢复并将错误写入error channel，
	//此函数可以获取error channel
	GetReportError() chan error

	// GetWaitGroup 获取等待组
	GetWaitGroup() *sync.WaitGroup

	// GetCancelFunc 获取取消运行函数
	GetCancelFunc() context.CancelFunc
}

// ContextBehavior 上下文行为
type ContextBehavior struct {
	context.Context
	parentCtx   context.Context
	reportError chan error
	cancel      context.CancelFunc
	wg          sync.WaitGroup
}

// Init 初始化
func (ctx *ContextBehavior) Init(parentCtx context.Context, reportError chan error) {
	if parentCtx == nil {
		ctx.parentCtx = context.Background()
	} else {
		ctx.parentCtx = parentCtx
	}

	ctx.reportError = reportError

	ctx.Context, ctx.cancel = context.WithCancel(ctx.parentCtx)
}

// GetParentCtx 获取父上下文
func (ctx *ContextBehavior) GetParentCtx() context.Context {
	return ctx.parentCtx
}

// GetReportError 在打开服务或运行时的AutoRecover选项时，panic时将会恢复并将错误写入error channel，
// 此函数可以获取error channel
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
