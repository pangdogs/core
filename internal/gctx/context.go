package gctx

import (
	"context"
	"sync"
	"sync/atomic"
)

// Context 上下文
type Context interface {
	iContext
	context.Context

	// GetParentContext 获取父上下文
	GetParentContext() context.Context
	// GetAutoRecover panic时是否自动恢复
	GetAutoRecover() bool
	// GetReportError 在开启panic时自动恢复时，将会恢复并将错误写入此error channel
	GetReportError() chan error
	// GetWaitGroup 获取等待组
	GetWaitGroup() *sync.WaitGroup
	// Terminate 停止
	Terminate() <-chan struct{}
	// TerminatedChan 已停止chan
	TerminatedChan() <-chan struct{}
}

type iContext interface {
	init(parentCtx context.Context, autoRecover bool, reportError chan error)
	setPaired(v bool) bool
	getPaired() bool
	getTerminatedChan() chan struct{}
}

// ContextBehavior 上下文行为
type ContextBehavior struct {
	context.Context
	parentCtx      context.Context
	autoRecover    bool
	reportError    chan error
	cancel         context.CancelFunc
	terminatedChan chan struct{}
	wg             sync.WaitGroup
	paired         atomic.Bool
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

// Terminate 停止
func (ctx *ContextBehavior) Terminate() <-chan struct{} {
	ctx.cancel()
	return ctx.terminatedChan
}

// TerminatedChan 已停止chan
func (ctx *ContextBehavior) TerminatedChan() <-chan struct{} {
	return ctx.terminatedChan
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
	ctx.terminatedChan = make(chan struct{})
}

func (ctx *ContextBehavior) setPaired(v bool) bool {
	return ctx.paired.CompareAndSwap(!v, v)
}

func (ctx *ContextBehavior) getPaired() bool {
	return ctx.paired.Load()
}

func (ctx *ContextBehavior) getTerminatedChan() chan struct{} {
	return ctx.terminatedChan
}
