package core

import (
	"github.com/pangdogs/galaxy/core/container"
)

// RuntimeContext 运行时上下文接口
type RuntimeContext interface {
	container.GCCollector
	_InnerGC
	_Context
	_RunnableMark
	_RuntimeContextEntityMgr
	_SafeCall

	init(servCtx ServiceContext, opts *RuntimeContextOptions)

	getOptions() *RuntimeContextOptions

	// GetServiceCtx 获取服务上下文（Service Context）
	GetServiceCtx() ServiceContext

	setFrame(frame Frame)

	// GetFrame 获取帧
	GetFrame() Frame

	// GetECTree 获取主EC树
	GetECTree() IECTree
}

// RuntimeContextGetOptions 获取运行时上下文创建选项，线程安全
func RuntimeContextGetOptions(runtimeCtx RuntimeContext) RuntimeContextOptions {
	return *runtimeCtx.getOptions()
}

// NewRuntimeContext 创建运行时上下文，线程安全
func NewRuntimeContext(servCtx ServiceContext, optSetterFuncs ...RuntimeContextOptionSetterFunc) RuntimeContext {
	opts := RuntimeContextOptions{}
	RuntimeContextOptionSetter.Default()(&opts)

	for i := range optSetterFuncs {
		optSetterFuncs[i](&opts)
	}

	return NewRuntimeContextWithOpts(servCtx, opts)
}

// NewRuntimeContextWithOpts 创建运行时上下文并传入参数，线程安全
func NewRuntimeContextWithOpts(servCtx ServiceContext, opts RuntimeContextOptions) RuntimeContext {
	if !opts.Inheritor.IsNil() {
		opts.Inheritor.IFace.init(servCtx, &opts)
		return opts.Inheritor.IFace
	}

	runtimeCtx := &_RuntimeContextBehavior{}
	runtimeCtx.init(servCtx, &opts)

	return runtimeCtx.opts.Inheritor.IFace
}

type _RuntimeCtxEntityInfo struct {
	Element *container.Element[FaceAny]
	Hooks   [2]Hook
}

type _RuntimeContextBehavior struct {
	_ContextBehavior
	_RunnableMarkBehavior
	opts                                    RuntimeContextOptions
	servCtx                                 ServiceContext
	entityMap                               map[int64]_RuntimeCtxEntityInfo
	entityList                              container.List[FaceAny]
	frame                                   Frame
	ecTree                                  ECTree
	callee                                  _Callee
	eventEntityMgrAddEntity                 Event
	eventEntityMgrRemoveEntity              Event
	eventEntityMgrEntityAddComponents       Event
	eventEntityMgrEntityRemoveComponent     Event
	_eventEntityMgrNotifyECTreeRemoveEntity Event
	gcList                                  []container.GC
	gc                                      _RuntimeContextBehaviorGC
}

func (runtimeCtx *_RuntimeContextBehavior) init(servCtx ServiceContext, opts *RuntimeContextOptions) {
	if servCtx == nil {
		panic("nil servCtx")
	}

	if opts == nil {
		panic("nil opts")
	}

	runtimeCtx.opts = *opts

	if runtimeCtx.opts.Inheritor.IsNil() {
		runtimeCtx.opts.Inheritor = NewFace[RuntimeContext](runtimeCtx)
	}

	runtimeCtx.gc._RuntimeContextBehavior = runtimeCtx

	runtimeCtx._ContextBehavior.init(servCtx, runtimeCtx.opts.ReportError)
	runtimeCtx.servCtx = servCtx

	runtimeCtx.entityList.Init(runtimeCtx.opts.FaceCache, runtimeCtx.opts.Inheritor.IFace)
	runtimeCtx.entityMap = map[int64]_RuntimeCtxEntityInfo{}

	runtimeCtx.eventEntityMgrAddEntity.Init(false, nil, EventRecursion_Discard, runtimeCtx.opts.HookCache, runtimeCtx.opts.Inheritor.IFace)
	runtimeCtx.eventEntityMgrRemoveEntity.Init(false, nil, EventRecursion_Discard, runtimeCtx.opts.HookCache, runtimeCtx.opts.Inheritor.IFace)
	runtimeCtx.eventEntityMgrEntityAddComponents.Init(false, nil, EventRecursion_Discard, runtimeCtx.opts.HookCache, runtimeCtx.opts.Inheritor.IFace)
	runtimeCtx.eventEntityMgrEntityRemoveComponent.Init(false, nil, EventRecursion_Discard, runtimeCtx.opts.HookCache, runtimeCtx.opts.Inheritor.IFace)
	runtimeCtx._eventEntityMgrNotifyECTreeRemoveEntity.Init(false, nil, EventRecursion_Discard, runtimeCtx.opts.HookCache, runtimeCtx.opts.Inheritor.IFace)

	runtimeCtx.ecTree.init(runtimeCtx.opts.Inheritor.IFace, true)
}

func (runtimeCtx *_RuntimeContextBehavior) getOptions() *RuntimeContextOptions {
	return &runtimeCtx.opts
}

// GetServiceCtx 获取服务上下文（Service Context）
func (runtimeCtx *_RuntimeContextBehavior) GetServiceCtx() ServiceContext {
	return runtimeCtx.servCtx
}

func (runtimeCtx *_RuntimeContextBehavior) setFrame(frame Frame) {
	runtimeCtx.frame = frame
}

// GetFrame 获取帧
func (runtimeCtx *_RuntimeContextBehavior) GetFrame() Frame {
	return runtimeCtx.frame
}

// GetECTree 获取主EC树
func (runtimeCtx *_RuntimeContextBehavior) GetECTree() IECTree {
	return &runtimeCtx.ecTree
}
