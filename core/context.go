package core

import (
	"context"
	"sync"
)

// Context 上下文
type Context interface {
	context.Context

	// GetParentCtx 获取父上下文（Context），线程安全
	GetParentCtx() context.Context

	// GetReportError 在打开服务（Service）或运行时（Runtime）的AutoRecover选项时，panic时将会恢复并将错误写入error channel，
	//此函数可以获取error channel，线程安全
	GetReportError() chan error

	// GetOrSetValue 获取或设置值，线程安全
	GetOrSetValue(key string, value interface{}) (actual interface{}, got bool)

	// SetValue 设置值，线程安全
	SetValue(key string, value interface{})

	// GetValue 获取值，线程安全
	GetValue(key string) (interface{}, bool)

	// GetWaitGroup 获取等待组，用于等待子上下文（Context）所在的服务（Service）或运行时（Runtime）停止运行，线程安全
	GetWaitGroup() *sync.WaitGroup

	// GetCancelFunc 获取取消运行函数，用于停止运行服务（Service）或运行时（Runtime），线程安全
	GetCancelFunc() context.CancelFunc
}

type _ContextBehavior struct {
	context.Context
	parentCtx   context.Context
	reportError chan error
	cancel      context.CancelFunc
	wg          sync.WaitGroup
	valueMap    sync.Map
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

// GetOrSetValue 获取或设置值，线程安全
func (ctx *_ContextBehavior) GetOrSetValue(key string, value interface{}) (actual interface{}, got bool) {
	return ctx.valueMap.LoadOrStore(key, value)
}

// SetValue 设置值，线程安全
func (ctx *_ContextBehavior) SetValue(key string, value interface{}) {
	ctx.valueMap.Store(key, value)
}

// GetValue 获取值，线程安全
func (ctx *_ContextBehavior) GetValue(key string) (interface{}, bool) {
	return ctx.valueMap.Load(key)
}

// GetWaitGroup 获取等待组，用于等待子上下文（Context）所在的服务（Service）或运行时（Runtime）停止运行，线程安全
func (ctx *_ContextBehavior) GetWaitGroup() *sync.WaitGroup {
	return &ctx.wg
}

// GetCancelFunc 获取取消运行函数，用于停止运行服务（Service）或运行时（Runtime），线程安全
func (ctx *_ContextBehavior) GetCancelFunc() context.CancelFunc {
	return ctx.cancel
}
