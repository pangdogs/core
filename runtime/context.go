package runtime

import (
	"fmt"
	"kit.golaxy.org/golaxy/event"
	"kit.golaxy.org/golaxy/internal/concurrent"
	"kit.golaxy.org/golaxy/internal/exception"
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/container"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/option"
	"kit.golaxy.org/golaxy/util/uid"
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
	concurrent.CurrentContextResolver
	concurrent.ConcurrentContextResolver
	concurrent.Context
	concurrent.Caller
	plugin.PluginResolver
	container.GCCollector
	fmt.Stringer

	// GetName 获取名称
	GetName() string
	// GetId 获取运行时Id
	GetId() uid.Id
	// GetFrame 获取帧
	GetFrame() Frame
	// GetEntityMgr 获取实体管理器
	GetEntityMgr() IEntityMgr
	// GetECTree 获取主EC树
	GetECTree() IECTree
	// GetFaceAnyAllocator 获取FaceAny内存分配器
	GetFaceAnyAllocator() container.Allocator[iface.FaceAny]
	// GetHookAllocator 获取Hook内存分配器
	GetHookAllocator() container.Allocator[event.Hook]
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
	entityMgr _EntityMgr
	ecTree    _ECTree
	callee    Callee
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
func (ctx *ContextBehavior) GetEntityMgr() IEntityMgr {
	return &ctx.entityMgr
}

// GetECTree 获取主EC树
func (ctx *ContextBehavior) GetECTree() IECTree {
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

// ResolveContext 解析上下文
func (ctx *ContextBehavior) ResolveContext() iface.Cache {
	return iface.Iface2Cache[Context](ctx.opts.CompositeFace.Iface)
}

// ResolveCurrentContext 解析当前上下文
func (ctx *ContextBehavior) ResolveCurrentContext() iface.Cache {
	return ctx.ResolveContext()
}

// ResolveConcurrentContext 解析多线程安全的上下文
func (ctx *ContextBehavior) ResolveConcurrentContext() iface.Cache {
	return ctx.ResolveContext()
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
	return fmt.Sprintf(`{"id":%q "name":%q}`, ctx.GetId(), ctx.GetName())
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
}
