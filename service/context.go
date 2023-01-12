package service

import (
	"context"
	"github.com/golaxy-kit/golaxy/ec"
	"github.com/golaxy-kit/golaxy/internal"
	"github.com/golaxy-kit/golaxy/util"
	"sync/atomic"
)

// NewContext 创建服务上下文
func NewContext(options ...ContextOption) Context {
	opts := ContextOptions{}
	WithContextOption{}.Default()(&opts)

	for i := range options {
		options[i](&opts)
	}

	return UnsafeNewContext(opts)
}

func UnsafeNewContext(options ContextOptions) Context {
	if !options.Inheritor.IsNil() {
		options.Inheritor.Iface.init(&options)
		return options.Inheritor.Iface
	}

	ctx := &ContextBehavior{}
	ctx.init(&options)

	return ctx.opts.Inheritor.Iface
}

// Context 服务上下文
type Context interface {
	internal.Context
	internal.RunningMark
	_SafeCall

	// GenSerialNo 生成流水号（运行时唯一）
	GenSerialNo() int64
	// GenPersistID 生成持久化ID（全局唯一）
	GenPersistID() ec.ID
	// GetEntityMgr 获取实体管理器
	GetEntityMgr() IEntityMgr

	init(opts *ContextOptions)
	getOptions() *ContextOptions
}

// ContextBehavior 服务上下文行为，在需要扩展服务上下文能力时，匿名嵌入至服务上下文结构体中
type ContextBehavior struct {
	internal.ContextBehavior
	internal.RunningMarkBehavior
	opts        ContextOptions
	snGenerator int64
	entityMgr   _EntityMgr
}

// GenSerialNo 生成流水号（运行时唯一）
func (ctx *ContextBehavior) GenSerialNo() int64 {
	return atomic.AddInt64(&ctx.snGenerator, 1)
}

// GenPersistID 生成持久化ID（全局唯一）
func (ctx *ContextBehavior) GenPersistID() ec.ID {
	return ctx.opts.GenPersistID()
}

// GetEntityMgr 获取实体管理器
func (ctx *ContextBehavior) GetEntityMgr() IEntityMgr {
	return &ctx.entityMgr
}

func (ctx *ContextBehavior) init(opts *ContextOptions) {
	if opts == nil {
		panic("nil opts")
	}

	ctx.opts = *opts

	if ctx.opts.Inheritor.IsNil() {
		ctx.opts.Inheritor = util.NewFace[Context](ctx)
	}

	if ctx.opts.Context == nil {
		ctx.opts.Context = context.Background()
	}

	internal.UnsafeContext(&ctx.ContextBehavior).Init(ctx.opts.Context, ctx.opts.AutoRecover, ctx.opts.ReportError)
	ctx.entityMgr.init(ctx.opts.Inheritor.Iface)
}

func (ctx *ContextBehavior) getOptions() *ContextOptions {
	return &ctx.opts
}
