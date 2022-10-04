package runtime

import (
	"github.com/pangdogs/galaxy/internal"
	"github.com/pangdogs/galaxy/localevent"
	"github.com/pangdogs/galaxy/service"
	"github.com/pangdogs/galaxy/util"
	"github.com/pangdogs/galaxy/util/container"
)

// Context 运行时上下文接口
type Context interface {
	_InnerGC
	container.GCCollector
	internal.Context
	internal.RunningMark
	_SafeCall

	init(serviceCtx service.Context, opts *ContextOptions)

	getOptions() *ContextOptions

	// GetFaceCache 获取Face缓存
	GetFaceCache() *container.Cache[util.FaceAny]

	// GetHookCache 获取Hook缓存
	GetHookCache() *container.Cache[localevent.Hook]

	// GetServiceCtx 获取服务上下文
	GetServiceCtx() service.Context

	setFrame(frame Frame)

	// GetFrame 获取帧
	GetFrame() Frame

	// GetEntityMgr 获取实体管理器
	GetEntityMgr() IEntityMgr

	// GetECTree 获取主EC树
	GetECTree() IECTree
}

// NewContext 创建运行时上下文
func NewContext(serviceCtx service.Context, optSetter ...ContextOptionSetter) Context {
	opts := ContextOptions{}
	ContextOption.Default()(&opts)

	for i := range optSetter {
		optSetter[i](&opts)
	}

	if !opts.Inheritor.IsNil() {
		opts.Inheritor.Iface.init(serviceCtx, &opts)
		return opts.Inheritor.Iface
	}

	runtimeCtx := &ContextBehavior{}
	runtimeCtx.init(serviceCtx, &opts)

	return runtimeCtx.opts.Inheritor.Iface
}

type _EntityInfo struct {
	Element *container.Element[util.FaceAny]
	Hooks   [3]localevent.Hook
}

// ContextBehavior 运行时上下文行为，在需要拓展运行时上下文能力时，匿名嵌入至运行时上下文结构体中
type ContextBehavior struct {
	internal.ContextBehavior
	internal.RunningMarkBehavior
	opts       ContextOptions
	serviceCtx service.Context
	frame      Frame
	entityMgr  EntityMgr
	ecTree     ECTree
	callee     Callee
	gcList     []container.GC
	innerGC    _ContextInnerGC
}

func (ctx *ContextBehavior) init(serviceCtx service.Context, opts *ContextOptions) {
	if serviceCtx == nil {
		panic("nil serviceCtx")
	}

	if opts == nil {
		panic("nil opts")
	}

	ctx.opts = *opts

	if ctx.opts.Inheritor.IsNil() {
		ctx.opts.Inheritor = util.NewFace[Context](ctx)
	}

	if ctx.opts.ParentContext == nil {
		ctx.opts.ParentContext = serviceCtx
	}

	ctx.innerGC.Init(ctx)

	ctx.ContextBehavior.Init(ctx.opts.ParentContext, ctx.opts.AutoRecover, ctx.opts.ReportError)
	ctx.serviceCtx = serviceCtx
	ctx.entityMgr.Init(ctx.getOptions().Inheritor.Iface)
	ctx.ecTree.init(ctx.opts.Inheritor.Iface, true)
}

func (ctx *ContextBehavior) getOptions() *ContextOptions {
	return &ctx.opts
}

// GetFaceCache 获取Face缓存
func (ctx *ContextBehavior) GetFaceCache() *container.Cache[util.FaceAny] {
	return ctx.opts.FaceCache
}

// GetHookCache 获取Hook缓存
func (ctx *ContextBehavior) GetHookCache() *container.Cache[localevent.Hook] {
	return ctx.opts.HookCache
}

// GetServiceCtx 获取服务上下文
func (ctx *ContextBehavior) GetServiceCtx() service.Context {
	return ctx.serviceCtx
}

func (ctx *ContextBehavior) setFrame(frame Frame) {
	ctx.frame = frame
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

func (ctx *ContextBehavior) CollectGC(gc container.GC) {
	if gc == nil || !gc.NeedGC() {
		return
	}

	ctx.gcList = append(ctx.gcList, gc)
}

func (ctx *ContextBehavior) getInnerGC() container.GC {
	return &ctx.innerGC
}
