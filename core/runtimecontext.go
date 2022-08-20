package core

import "github.com/pangdogs/galaxy/core/container"

type RuntimeContext interface {
	container.GC
	container.GCCollector
	Context
	_RunnableMark
	EntityMgr
	EntityMgrEvents
	EntityReverseQuery
	EntityCountQuery
	SafeCall
	init(servCtx ServiceContext, opts *RuntimeContextOptions)
	getOptions() *RuntimeContextOptions
	GetServiceCtx() ServiceContext
	setFrame(frame Frame)
	GetFrame() Frame
	GetECTree() IECTree
}

func RuntimeContextGetOptions(runtimeCtx RuntimeContext) RuntimeContextOptions {
	return *runtimeCtx.getOptions()
}

func RuntimeContextGetInheritor(runtimeCtx RuntimeContext) Face[RuntimeContext] {
	return runtimeCtx.getOptions().Inheritor
}

func RuntimeContextGetInheritorIFace[T any](runtimeCtx RuntimeContext) T {
	return Cache2IFace[T](runtimeCtx.getOptions().Inheritor.Cache)
}

func NewRuntimeContext(servCtx ServiceContext, optFuncs ...NewRuntimeContextOptionFunc) RuntimeContext {
	opts := &RuntimeContextOptions{}
	NewRuntimeContextOption.Default()(opts)

	for i := range optFuncs {
		optFuncs[i](opts)
	}

	if !opts.Inheritor.IsNil() {
		opts.Inheritor.IFace.init(servCtx, opts)
		return opts.Inheritor.IFace
	}

	runtimeCtx := &RuntimeContextBehavior{}
	runtimeCtx.init(servCtx, opts)

	return runtimeCtx.opts.Inheritor.IFace
}

type _RuntimeCtxEntityInfo struct {
	Element *container.Element[FaceAny]
	Hooks   [2]Hook
}

type RuntimeContextBehavior struct {
	_ContextBehavior
	_RunnableMarkBehavior
	opts                                    RuntimeContextOptions
	servCtx                                 ServiceContext
	entityMap                               map[uint64]_RuntimeCtxEntityInfo
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
}

func (runtimeCtx *RuntimeContextBehavior) GC() {
	for i := range runtimeCtx.gcList {
		runtimeCtx.gcList[i].GC()
	}
	runtimeCtx.gcList = runtimeCtx.gcList[:0]
}

func (runtimeCtx *RuntimeContextBehavior) NeedGC() bool {
	return len(runtimeCtx.gcList) > 0
}

func (runtimeCtx *RuntimeContextBehavior) CollectGC(gc container.GC) {
	if gc == nil || !gc.NeedGC() {
		return
	}

	runtimeCtx.gcList = append(runtimeCtx.gcList, gc)
}

func (runtimeCtx *RuntimeContextBehavior) init(servCtx ServiceContext, opts *RuntimeContextOptions) {
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

	runtimeCtx._ContextBehavior.init(servCtx, runtimeCtx.opts.ReportError)
	runtimeCtx.servCtx = servCtx

	runtimeCtx.entityList.Init(runtimeCtx.opts.FaceCache, runtimeCtx.opts.Inheritor.IFace)
	runtimeCtx.entityMap = map[uint64]_RuntimeCtxEntityInfo{}

	runtimeCtx.eventEntityMgrAddEntity.Init(false, nil, EventRecursion_Discard, runtimeCtx.opts.HookCache, runtimeCtx.opts.Inheritor.IFace)
	runtimeCtx.eventEntityMgrRemoveEntity.Init(false, nil, EventRecursion_Discard, runtimeCtx.opts.HookCache, runtimeCtx.opts.Inheritor.IFace)
	runtimeCtx.eventEntityMgrEntityAddComponents.Init(false, nil, EventRecursion_Discard, runtimeCtx.opts.HookCache, runtimeCtx.opts.Inheritor.IFace)
	runtimeCtx.eventEntityMgrEntityRemoveComponent.Init(false, nil, EventRecursion_Discard, runtimeCtx.opts.HookCache, runtimeCtx.opts.Inheritor.IFace)
	runtimeCtx._eventEntityMgrNotifyECTreeRemoveEntity.Init(false, nil, EventRecursion_Discard, runtimeCtx.opts.HookCache, runtimeCtx.opts.Inheritor.IFace)

	runtimeCtx.ecTree.init(runtimeCtx.opts.Inheritor.IFace, true)
}

func (runtimeCtx *RuntimeContextBehavior) getOptions() *RuntimeContextOptions {
	return &runtimeCtx.opts
}

func (runtimeCtx *RuntimeContextBehavior) GetServiceCtx() ServiceContext {
	return runtimeCtx.servCtx
}

func (runtimeCtx *RuntimeContextBehavior) setFrame(frame Frame) {
	runtimeCtx.frame = frame
}

func (runtimeCtx *RuntimeContextBehavior) GetFrame() Frame {
	return runtimeCtx.frame
}

func (runtimeCtx *RuntimeContextBehavior) GetECTree() IECTree {
	return &runtimeCtx.ecTree
}
