package runtime

import (
	"fmt"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/internal/concurrent"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/util/container"
	"git.golaxy.org/core/util/iface"
	"git.golaxy.org/core/util/option"
	"git.golaxy.org/core/util/uid"
)

// NewContext 创建运行时上下文
func NewContext(servCtx service.Context, settings ...option.Setting[ContextOptions]) Context {
	return UnsafeNewContext(servCtx, option.Make(_ContextOption{}.Default(), settings...))
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
	_Context
	concurrent.CurrentContextProvider
	concurrent.Context
	concurrent.Caller
	plugin.PluginProvider
	container.GCCollector
	fmt.Stringer

	// GetName 获取名称
	GetName() string
	// GetId 获取运行时Id
	GetId() uid.Id
	// GetFrame 获取帧
	GetFrame() Frame
	// GetEntityMgr 获取实体管理器
	GetEntityMgr() EntityMgr
	// GetECTree 获取主EC树
	GetECTree() ECTree
	// GetFaceAnyAllocator 获取FaceAny内存分配器
	GetFaceAnyAllocator() container.Allocator[iface.FaceAny]
	// GetHookAllocator 获取Hook内存分配器
	GetHookAllocator() container.Allocator[event.Hook]
	// ActivateEvent 启用事件
	ActivateEvent(event event.IEventCtrl, recursion event.EventRecursion)
	// ManagedHooks 托管hook，在运行时停止时自动解绑定
	ManagedHooks(hooks ...event.Hook)
}

type _Context interface {
	init(servCtx service.Context, opts ContextOptions)
	getOptions() *ContextOptions
	setFrame(frame Frame)
	setCallee(callee Callee)
	getServiceCtx() service.Context
	changeRunningState(state RunningState)
	gc()
}

// ContextBehavior 运行时上下文行为，在需要扩展运行时上下文能力时，匿名嵌入至运行时上下文结构体中
type ContextBehavior struct {
	concurrent.ContextBehavior
	servCtx   service.Context
	opts      ContextOptions
	frame     Frame
	entityMgr _EntityMgrBehavior
	ecTree    _ECTreeBehavior
	callee    Callee
	hooks     []event.Hook
	gcList    []container.GC
}

// GetName 获取名称
func (ctx *ContextBehavior) GetName() string {
	return ctx.opts.Name
}

// GetId 获取运行时Id
func (ctx *ContextBehavior) GetId() uid.Id {
	return ctx.opts.PersistId
}

// GetFrame 获取帧
func (ctx *ContextBehavior) GetFrame() Frame {
	return ctx.frame
}

// GetEntityMgr 获取实体管理器
func (ctx *ContextBehavior) GetEntityMgr() EntityMgr {
	return &ctx.entityMgr
}

// GetECTree 获取主EC树
func (ctx *ContextBehavior) GetECTree() ECTree {
	return &ctx.ecTree
}

// GetFaceAnyAllocator 获取FaceAny内存分配器
func (ctx *ContextBehavior) GetFaceAnyAllocator() container.Allocator[iface.FaceAny] {
	return ctx.opts.FaceAnyAllocator
}

// GetHookAllocator 获取Hook内存分配器
func (ctx *ContextBehavior) GetHookAllocator() container.Allocator[event.Hook] {
	return ctx.opts.HookAllocator
}

// ActivateEvent 启用事件
func (ctx *ContextBehavior) ActivateEvent(event event.IEventCtrl, recursion event.EventRecursion) {
	if event == nil {
		panic(fmt.Errorf("%w: %w: event is nil", ErrContext, exception.ErrArgs))
	}
	event.Init(ctx.GetAutoRecover(), ctx.GetReportError(), recursion, ctx.GetHookAllocator(), ctx.opts.CompositeFace.Iface)
}

// ManagedHooks 托管hook，在运行时停止时自动解绑定
func (ctx *ContextBehavior) ManagedHooks(hooks ...event.Hook) {
	for i := len(ctx.hooks) - 1; i >= 0; i-- {
		if !ctx.hooks[i].IsBound() {
			ctx.hooks = append(ctx.hooks[:i], ctx.hooks[i+1:]...)
		}
	}
	ctx.hooks = append(ctx.hooks, hooks...)
}

// GetCurrentContext 获取当前上下文
func (ctx *ContextBehavior) GetCurrentContext() iface.Cache {
	return iface.Iface2Cache[Context](ctx.opts.CompositeFace.Iface)
}

// GetConcurrentContext 获取多线程安全的上下文
func (ctx *ContextBehavior) GetConcurrentContext() iface.Cache {
	return iface.Iface2Cache[Context](ctx.opts.CompositeFace.Iface)
}

// CollectGC 收集GC
func (ctx *ContextBehavior) CollectGC(gc container.GC) {
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
		ctx.opts.CompositeFace = iface.MakeFace[Context](ctx)
	}

	if ctx.opts.Context == nil {
		ctx.opts.Context = servCtx
	}

	if ctx.opts.PersistId.IsNil() {
		ctx.opts.PersistId = uid.New()
	}

	concurrent.UnsafeContext(&ctx.ContextBehavior).Init(ctx.opts.Context, ctx.opts.AutoRecover, ctx.opts.ReportError)
	ctx.servCtx = servCtx
	ctx.entityMgr.init(ctx.opts.CompositeFace.Iface)
	ctx.ecTree.init(ctx.opts.CompositeFace.Iface)
}

func (ctx *ContextBehavior) getOptions() *ContextOptions {
	return &ctx.opts
}

func (ctx *ContextBehavior) setFrame(frame Frame) {
	ctx.frame = frame
}

func (ctx *ContextBehavior) setCallee(callee Callee) {
	ctx.callee = callee
}

func (ctx *ContextBehavior) getServiceCtx() service.Context {
	return ctx.servCtx
}

func (ctx *ContextBehavior) changeRunningState(state RunningState) {
	ctx.entityMgr.changeRunningState(state)
	ctx.ecTree.changeRunningState(state)
	ctx.opts.RunningHandler.Call(ctx.GetAutoRecover(), ctx.GetReportError(), nil, ctx.opts.CompositeFace.Iface, state)

	switch state {
	case RunningState_Terminated:
		ctx.cleanHooks()
	}
}

func (ctx *ContextBehavior) cleanHooks() {
	for i := range ctx.hooks {
		ctx.hooks[i].Unbind()
	}
	ctx.hooks = nil
}
