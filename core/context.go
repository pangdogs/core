package core

import (
	"context"
	"sync"
)

// _Context 上下文
type _Context interface {
	context.Context

	// GetParentCtx 获取父上下文（Context），线程安全
	GetParentCtx() context.Context

	// GetReportError 在打开服务（Service）或运行时（Runtime）的AutoRecover选项时，panic时将会恢复并将错误写入error channel，
	//此函数可以获取error channel，线程安全
	GetReportError() chan error

	getWaitGroup() *sync.WaitGroup

	// GetCancelFunc 获取取消运行函数，用于停止运行服务（Service）或运行时（Runtime），线程安全
	GetCancelFunc() context.CancelFunc
}

type _ContextBehavior struct {
	context.Context
	parentCtx   context.Context
	reportError chan error
	cancel      context.CancelFunc
	wg          sync.WaitGroup
}

func (ctx *_ContextBehavior) init(parentCtx context.Context, reportError chan error) {
	if parentCtx == nil {
		ctx.parentCtx = context.Background()
	} else {
		ctx.parentCtx = parentCtx
	}

	ctx.reportError = reportError

	ctx.Context, ctx.cancel = context.WithCancel(ctx.parentCtx)
}

// GetParentCtx 获取父上下文（Context），线程安全
func (ctx *_ContextBehavior) GetParentCtx() context.Context {
	return ctx.parentCtx
}

// GetReportError 在打开服务（Service）或运行时（Runtime）的AutoRecover选项时，panic时将会恢复并将错误写入error channel，
//此函数可以获取error channel，线程安全
func (ctx *_ContextBehavior) GetReportError() chan error {
	return ctx.reportError
}

func (ctx *_ContextBehavior) getWaitGroup() *sync.WaitGroup {
	return &ctx.wg
}

// GetCancelFunc 获取取消运行函数，用于停止运行服务（Service）或运行时（Runtime），线程安全
func (ctx *_ContextBehavior) GetCancelFunc() context.CancelFunc {
	return ctx.cancel
}
