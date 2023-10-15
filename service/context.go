package service

import (
	"context"
	"fmt"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/uid"
)

// NewContext 创建服务上下文
func NewContext(options ...ContextOption) Context {
	opts := ContextOptions{}
	Option{}.Default()(&opts)

	for i := range options {
		options[i](&opts)
	}

	return UnsafeNewContext(opts)
}

// Deprecated: UnsafeNewContext 内部创建服务上下文
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
	internal.RunningState
	plugin.PluginResolver
	pt.EntityPtResolver
	Caller
	fmt.Stringer

	// GetName 获取名称
	GetName() string
	// GetId 获取服务Id
	GetId() uid.Id
	// GetEntityMgr 获取实体管理器
	GetEntityMgr() IEntityMgr
}

type _Context interface {
	init(opts *ContextOptions)
	getOptions() *ContextOptions
}

// ContextBehavior 服务上下文行为，在需要扩展服务上下文能力时，匿名嵌入至服务上下文结构体中
type ContextBehavior struct {
	internal.ContextBehavior
	internal.RunningStateBehavior
	opts      ContextOptions
	entityMgr _EntityMgr
}

// GetName 获取名称
func (ctx *ContextBehavior) GetName() string {
	return ctx.opts.Name
}

// GetId 获取服务Id
func (ctx *ContextBehavior) GetId() uid.Id {
	return ctx.opts.PersistId
}

// GetEntityMgr 获取实体管理器
func (ctx *ContextBehavior) GetEntityMgr() IEntityMgr {
	return &ctx.entityMgr
}

// String implements fmt.Stringer
func (ctx *ContextBehavior) String() string {
	return fmt.Sprintf(`{"id":%q "name":%q}`, ctx.GetId(), ctx.GetName())
}

func (ctx *ContextBehavior) init(opts *ContextOptions) {
	if opts == nil {
		panic(fmt.Errorf("%w: %w: opts is nil", ErrContext, internal.ErrArgs))
	}

	ctx.opts = *opts

	if ctx.opts.CompositeFace.IsNil() {
		ctx.opts.CompositeFace = iface.NewFace[Context](ctx)
	}

	if ctx.opts.Context == nil {
		ctx.opts.Context = context.Background()
	}

	if ctx.opts.PersistId.IsNil() {
		ctx.opts.PersistId = uid.New()
	}

	internal.UnsafeContext(&ctx.ContextBehavior).Init(ctx.opts.Context, ctx.opts.AutoRecover, ctx.opts.ReportError)
	ctx.entityMgr.init(ctx.opts.CompositeFace.Iface)
}

func (ctx *ContextBehavior) getOptions() *ContextOptions {
	return &ctx.opts
}
