package runtime

import (
	"fmt"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/internal/gctx"
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/reinterpret"
	"git.golaxy.org/core/utils/uid"
	"reflect"
)

// NewContext 创建运行时上下文
func NewContext(servCtx service.Context, settings ...option.Setting[ContextOptions]) Context {
	return UnsafeNewContext(servCtx, option.Make(With.Context.Default(), settings...))
}

// Deprecated: UnsafeNewContext 内部创建运行时上下文
func UnsafeNewContext(servCtx service.Context, options ContextOptions) Context {
	if !options.CompositeFace.IsNil() {
		options.CompositeFace.Iface.init(servCtx, options)
		return options.CompositeFace.Iface
	}

	ctx := &ContextBehavior{}
	ctx.init(servCtx, options)

	return ctx.opts.CompositeFace.Iface
}

// Context 运行时上下文接口
type Context interface {
	iContext
	gctx.CurrentContextProvider
	gctx.Context
	async.Caller
	reinterpret.CompositeProvider
	plugin.PluginProvider
	GCCollector
	fmt.Stringer

	// GetName 获取名称
	GetName() string
	// GetId 获取运行时Id
	GetId() uid.Id
	// GetReflected 获取反射值
	GetReflected() reflect.Value
	// GetFrame 获取帧
	GetFrame() Frame
	// GetEntityMgr 获取实体管理器
	GetEntityMgr() EntityMgr
	// GetEntityTree 获取实体树
	GetEntityTree() EntityTree
	// ActivateEvent 启用事件
	ActivateEvent(event event.IEventCtrl, recursion event.EventRecursion)
	// ManagedHooks 托管hook，在运行时停止时自动解绑定
	ManagedHooks(hooks ...event.Hook)
}

type iContext interface {
	init(servCtx service.Context, opts ContextOptions)
	getOptions() *ContextOptions
	setFrame(frame Frame)
	setCallee(callee async.Callee)
	getServiceCtx() service.Context
	changeRunningState(state RunningState)
	gc()
}

// ContextBehavior 运行时上下文行为，在需要扩展运行时上下文能力时，匿名嵌入至运行时上下文结构体中
type ContextBehavior struct {
	gctx.ContextBehavior
	servCtx      service.Context
	opts         ContextOptions
	reflected    reflect.Value
	frame        Frame
	entityMgr    _EntityMgrBehavior
	callee       async.Callee
	managedHooks []event.Hook
	gcList       []GC
}

// GetName 获取名称
func (ctx *ContextBehavior) GetName() string {
	return ctx.opts.Name
}

// GetId 获取运行时Id
func (ctx *ContextBehavior) GetId() uid.Id {
	return ctx.opts.PersistId
}

// GetReflected 获取反射值
func (ctx *ContextBehavior) GetReflected() reflect.Value {
	return ctx.reflected
}

// GetFrame 获取帧
func (ctx *ContextBehavior) GetFrame() Frame {
	return ctx.frame
}

// GetEntityMgr 获取实体管理器
func (ctx *ContextBehavior) GetEntityMgr() EntityMgr {
	return &ctx.entityMgr
}

// GetEntityTree 获取主实体树
func (ctx *ContextBehavior) GetEntityTree() EntityTree {
	return &ctx.entityMgr
}

// ActivateEvent 启用事件
func (ctx *ContextBehavior) ActivateEvent(event event.IEventCtrl, recursion event.EventRecursion) {
	if event == nil {
		panic(fmt.Errorf("%w: %w: event is nil", ErrContext, exception.ErrArgs))
	}
	event.Init(ctx.GetAutoRecover(), ctx.GetReportError(), recursion)
}

// GetCurrentContext 获取当前上下文
func (ctx *ContextBehavior) GetCurrentContext() iface.Cache {
	return iface.Iface2Cache[Context](ctx.opts.CompositeFace.Iface)
}

// GetConcurrentContext 获取多线程安全的上下文
func (ctx *ContextBehavior) GetConcurrentContext() iface.Cache {
	return iface.Iface2Cache[Context](ctx.opts.CompositeFace.Iface)
}

// GetCompositeFaceCache 支持重新解释类型
func (ctx *ContextBehavior) GetCompositeFaceCache() iface.Cache {
	return ctx.opts.CompositeFace.Cache
}

// CollectGC 收集GC
func (ctx *ContextBehavior) CollectGC(gc GC) {
	if gc == nil || !gc.NeedGC() {
		return
	}

	ctx.gcList = append(ctx.gcList, gc)
}

// String implements fmt.Stringer
func (ctx *ContextBehavior) String() string {
	return fmt.Sprintf(`{"id":%q, "name":%q}`, ctx.GetId(), ctx.GetName())
}

func (ctx *ContextBehavior) init(servCtx service.Context, opts ContextOptions) {
	if servCtx == nil {
		panic(fmt.Errorf("%w: %w: servCtx is nil", ErrContext, exception.ErrArgs))
	}

	ctx.opts = opts

	if ctx.opts.CompositeFace.IsNil() {
		ctx.opts.CompositeFace = iface.MakeFaceT[Context](ctx)
	}

	if ctx.opts.Context == nil {
		ctx.opts.Context = servCtx
	}

	if ctx.opts.PersistId.IsNil() {
		ctx.opts.PersistId = uid.New()
	}

	gctx.UnsafeContext(&ctx.ContextBehavior).Init(ctx.opts.Context, ctx.opts.AutoRecover, ctx.opts.ReportError)
	ctx.servCtx = servCtx
	ctx.reflected = reflect.ValueOf(ctx.opts.CompositeFace.Iface)
	ctx.entityMgr.init(ctx.opts.CompositeFace.Iface)
}

func (ctx *ContextBehavior) getOptions() *ContextOptions {
	return &ctx.opts
}

func (ctx *ContextBehavior) setFrame(frame Frame) {
	ctx.frame = frame
}

func (ctx *ContextBehavior) setCallee(callee async.Callee) {
	ctx.callee = callee
}

func (ctx *ContextBehavior) getServiceCtx() service.Context {
	return ctx.servCtx
}

func (ctx *ContextBehavior) changeRunningState(state RunningState) {
	ctx.entityMgr.changeRunningState(state)
	ctx.opts.RunningHandler.Call(ctx.GetAutoRecover(), ctx.GetReportError(), nil, ctx.opts.CompositeFace.Iface, state)

	switch state {
	case RunningState_Terminated:
		ctx.cleanManagedHooks()
	}
}
