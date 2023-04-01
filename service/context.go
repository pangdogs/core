package service

import (
	"context"
	"fmt"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/util"
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
	if !options.CompositeFace.IsNil() {
		options.CompositeFace.Iface.init(&options)
		return options.CompositeFace.Iface
	}

	ctx := &ContextBehavior{}
	ctx.init(&options)

	return ctx.opts.CompositeFace.Iface
}

// Context 服务上下文
type Context interface {
	_Context
	internal.Context
	internal.RunningMark
	plugin.PluginResolver
	pt.PtResolver
	_Call

	// GetPrototype 获取原型名称
	GetPrototype() string
	// GenSerialNo 生成流水号（运行时唯一）
	GenSerialNo() int64
	// GenPersistID 生成持久化ID（全局唯一）
	GenPersistID() ec.ID
	// GetEntityMgr 获取实体管理器
	GetEntityMgr() IEntityMgr
	// String 字符串化
	String() string
}

type _Context interface {
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

// GetPrototype 获取原型名称
func (ctx *ContextBehavior) GetPrototype() string {
	return ctx.opts.Prototype
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

// String 字符串化
func (ctx *ContextBehavior) String() string {
	return fmt.Sprintf("[Ptr:0x%x Prototype:%s]", ctx.opts.CompositeFace.Cache[1], ctx.GetPrototype())
}

func (ctx *ContextBehavior) init(opts *ContextOptions) {
	if opts == nil {
		panic("nil opts")
	}

	ctx.opts = *opts

	if ctx.opts.CompositeFace.IsNil() {
		ctx.opts.CompositeFace = util.NewFace[Context](ctx)
	}

	if ctx.opts.Context == nil {
		ctx.opts.Context = context.Background()
	}

	internal.UnsafeContext(&ctx.ContextBehavior).Init(ctx.opts.Context, ctx.opts.AutoRecover, ctx.opts.ReportError)
	ctx.entityMgr.init(ctx.opts.CompositeFace.Iface)
}

func (ctx *ContextBehavior) getOptions() *ContextOptions {
	return &ctx.opts
}
