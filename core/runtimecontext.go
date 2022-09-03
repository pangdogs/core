package core

import (
	"github.com/pangdogs/galaxy/core/container"
)

// RuntimeContext ...
type RuntimeContext interface {
	container.GCCollector
	_InnerGC
	_Context
	_RunnableMark
	_RuntimeContextEntityMgr
	_SafeCall
	init(servCtx ServiceContext, opts *RuntimeContextOptions)
	getOptions() *RuntimeContextOptions
	GetServiceCtx() ServiceContext
	setFrame(frame Frame)
	GetFrame() Frame
	GetECTree() IECTree
}

// RuntimeContextGetOptions ...
func RuntimeContextGetOptions(runtimeCtx RuntimeContext) RuntimeContextOptions {
	return *runtimeCtx.getOptions()
}

// NewRuntimeContext ...
func NewRuntimeContext(servCtx ServiceContext, optSetterFuncs ...NewRuntimeContextOptionFunc) RuntimeContext {
	opts := RuntimeContextOptions{}
	NewRuntimeContextOption.Default()(&opts)

	for i := range optSetterFuncs {
		optSetterFuncs[i](&opts)
	}

	return NewRuntimeContextWithOpts(servCtx, opts)
}

// NewRuntimeContextWithOpts ...
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

// GetServiceCtx ...
func (runtimeCtx *_RuntimeContextBehavior) GetServiceCtx() ServiceContext {
	return runtimeCtx.servCtx
}

func (runtimeCtx *_RuntimeContextBehavior) setFrame(frame Frame) {
	runtimeCtx.frame = frame
}

// GetFrame ...
func (runtimeCtx *_RuntimeContextBehavior) GetFrame() Frame {
	return runtimeCtx.frame
}

// GetECTree ...
func (runtimeCtx *_RuntimeContextBehavior) GetECTree() IECTree {
	return &runtimeCtx.ecTree
}
