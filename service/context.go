package service

import (
	"context"
	"fmt"
	"git.golaxy.org/core/internal/concurrent"
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/util/iface"
	"git.golaxy.org/core/util/option"
	"git.golaxy.org/core/util/uid"
)

// NewContext 创建服务上下文
func NewContext(settings ...option.Setting[ContextOptions]) Context {
	return UnsafeNewContext(option.Make(Option{}.Default(), settings...))
}

// Deprecated: UnsafeNewContext 内部创建服务上下文
func UnsafeNewContext(options ContextOptions) Context {
	if !options.CompositeFace.IsNil() {
		options.CompositeFace.Iface.init(options)
		return options.CompositeFace.Iface
	}

	ctx := &ContextBehavior{}
	ctx.init(options)

	return ctx.opts.CompositeFace.Iface
}

// Context 服务上下文
type Context interface {
	_Context
	concurrent.Context
	Caller
	plugin.PluginProvider
	pt.EntityPTProvider
	fmt.Stringer

	// GetName 获取名称
	GetName() string
	// GetId 获取服务Id
	GetId() uid.Id
	// GetEntityMgr 获取实体管理器
	GetEntityMgr() EntityMgr
}

type _Context interface {
	init(opts ContextOptions)
	getOptions() *ContextOptions
	changeRunningState(state RunningState)
}

// ContextBehavior 服务上下文行为，在需要扩展服务上下文能力时，匿名嵌入至服务上下文结构体中
type ContextBehavior struct {
	concurrent.ContextBehavior
	opts      ContextOptions
	entityMgr _EntityMgrBehavior
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
func (ctx *ContextBehavior) GetEntityMgr() EntityMgr {
	return &ctx.entityMgr
}

// String implements fmt.Stringer
func (ctx *ContextBehavior) String() string {
	return fmt.Sprintf(`{"id":%q, "name":%q}`, ctx.GetId(), ctx.GetName())
}

func (ctx *ContextBehavior) init(opts ContextOptions) {
	ctx.opts = opts

	if ctx.opts.CompositeFace.IsNil() {
		ctx.opts.CompositeFace = iface.MakeFace[Context](ctx)
	}

	if ctx.opts.Context == nil {
		ctx.opts.Context = context.Background()
	}

	if ctx.opts.PersistId.IsNil() {
		ctx.opts.PersistId = uid.New()
	}

	concurrent.UnsafeContext(&ctx.ContextBehavior).Init(ctx.opts.Context, ctx.opts.AutoRecover, ctx.opts.ReportError)
	ctx.entityMgr.init(ctx.opts.CompositeFace.Iface)
}

func (ctx *ContextBehavior) getOptions() *ContextOptions {
	return &ctx.opts
}

func (ctx *ContextBehavior) changeRunningState(state RunningState) {
	ctx.opts.RunningHandler.Call(ctx.GetAutoRecover(), ctx.GetReportError(), nil, ctx.opts.CompositeFace.Iface, state)
}
