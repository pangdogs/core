package runtime

import (
	"fmt"
	"kit.golaxy.org/golaxy/event"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/internal/errors"
	"kit.golaxy.org/golaxy/plugin"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/container"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/uid"
)

// NewContext 创建运行时上下文
func NewContext(serviceCtx service.Context, options ...ContextOption) Context {
	opts := ContextOptions{}
	_ContextOption{}.Default()(&opts)

	for i := range options {
		options[i](&opts)
	}

	return UnsafeNewContext(serviceCtx, opts)
}

// Deprecated: UnsafeNewContext 内部创建运行时上下文
func UnsafeNewContext(serviceCtx service.Context, options ContextOptions) Context {
	if !options.CompositeFace.IsNil() {
		options.CompositeFace.Iface.init(serviceCtx, &options)
		return options.CompositeFace.Iface
	}

	ctx := &ContextBehavior{}
	ctx.init(serviceCtx, &options)

	return ctx.opts.CompositeFace.Iface
}

// Context 运行时上下文接口
type Context interface {
	_Context
	internal.ContextResolver
	container.GCCollector
	internal.Context
	internal.RunningState
	plugin.PluginResolver
	Caller
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
	init(serviceCtx service.Context, opts *ContextOptions)
	getOptions() *ContextOptions
	setFrame(frame Frame)
	setCallee(callee Callee)
	getServiceCtx() service.Context
	gc()
}

// ContextBehavior 运行时上下文行为，在需要扩展运行时上下文能力时，匿名嵌入至运行时上下文结构体中
type ContextBehavior struct {
	internal.ContextBehavior
	internal.RunningStateBehavior
	opts       ContextOptions
	serviceCtx service.Context
	frame      Frame
	entityMgr  _EntityMgr
	ecTree     _ECTree
	callee     Callee
	gcList     []container.GC
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

func (ctx *ContextBehavior) init(serviceCtx service.Context, opts *ContextOptions) {
	if serviceCtx == nil {
		panic(fmt.Errorf("%w: %w: serviceCtx is nil", ErrContext, errors.ErrArgs))
	}

	if opts == nil {
		panic(fmt.Errorf("%w: %w: opts is nil", ErrContext, errors.ErrArgs))
	}

	ctx.opts = *opts

	if ctx.opts.CompositeFace.IsNil() {
		ctx.opts.CompositeFace = iface.NewFace[Context](ctx)
	}

	if ctx.opts.Context == nil {
		ctx.opts.Context = serviceCtx
	}

	if ctx.opts.PersistId.IsNil() {
		ctx.opts.PersistId = uid.New()
	}

	internal.UnsafeContext(&ctx.ContextBehavior).Init(ctx.opts.Context, ctx.opts.AutoRecover, ctx.opts.ReportError)
	ctx.serviceCtx = serviceCtx
	ctx.entityMgr.Init(ctx.opts.CompositeFace.Iface)
	ctx.ecTree.Init(ctx.opts.CompositeFace.Iface)
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
	return ctx.serviceCtx
}
