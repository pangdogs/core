package core

import "github.com/pangdogs/galaxy/core/container"

type Entity interface {
	container.GC
	container.GCCollector
	ComponentMgr
	ComponentMgrEvents
	init(opts *EntityOptions)
	getOptions() *EntityOptions
	setID(id uint64)
	GetID() uint64
	setRuntimeCtx(runtimeCtx RuntimeContext)
	GetRuntimeCtx() RuntimeContext
	setParent(parent Entity)
	GetParent() (Entity, bool)
	setInitialing(v bool)
	getInitialing() bool
	setShutting(v bool)
	getShutting() bool
	DestroySelf()
	eventEntityDestroySelf() IEvent
}

func EntityGetOptions(e Entity) EntityOptions {
	return *e.getOptions()
}

func EntityGetInheritor(e Entity) Face[Entity] {
	return e.getOptions().Inheritor
}

func EntityGetInheritorIFace[T any](e Entity) T {
	return Cache2IFace[T](e.getOptions().Inheritor.Cache)
}

func EntityGetInitialing(e Entity) bool {
	return e.getInitialing()
}

func NewEntity(optFuncs ...NewEntityOptionFunc) Entity {
	opts := &EntityOptions{}
	NewEntityOption.Default()(opts)

	for i := range optFuncs {
		optFuncs[i](opts)
	}

	if !opts.Inheritor.IsNil() {
		opts.Inheritor.IFace.init(opts)
		return opts.Inheritor.IFace
	}

	e := &EntityBehavior{}
	e.init(opts)

	return e.opts.Inheritor.IFace
}

type EntityBehavior struct {
	id                          uint64
	opts                        EntityOptions
	runtimeCtx                  RuntimeContext
	parent                      Entity
	componentList               container.List[FaceAny]
	componentMap                map[string]*container.Element[FaceAny]
	componentByIDMap            map[uint64]*container.Element[FaceAny]
	initialing, shutting        bool
	_eventEntityDestroySelf     Event
	eventCompMgrAddComponents   Event
	eventCompMgrRemoveComponent Event
	gcMark, gcCollected         bool
}

func (entity *EntityBehavior) GC() {
	if !entity.gcMark {
		return
	}
	entity.gcMark = false
	entity.gcCollected = false

	entity.componentList.GC()
	entity._eventEntityDestroySelf.GC()
	entity.eventCompMgrAddComponents.GC()
	entity.eventCompMgrRemoveComponent.GC()
}

func (entity *EntityBehavior) NeedGC() bool {
	return entity.gcMark
}

func (entity *EntityBehavior) CollectGC(gc container.GC) {
	if gc == nil || !gc.NeedGC() {
		return
	}

	entity.gcMark = true

	if entity.runtimeCtx != nil && !entity.gcCollected {
		entity.gcCollected = true
		entity.runtimeCtx.CollectGC(entity.opts.Inheritor.IFace)
	}
}

func (entity *EntityBehavior) init(opts *EntityOptions) {
	if opts == nil {
		panic("nil opts")
	}

	entity.opts = *opts

	if entity.opts.Inheritor.IsNil() {
		entity.opts.Inheritor = NewFace[Entity](entity)
	}

	entity.componentList.Init(entity.opts.FaceCache, entity.opts.Inheritor.IFace)

	if entity.opts.EnableFastGetComponent {
		entity.componentMap = map[string]*container.Element[FaceAny]{}
	}

	if entity.opts.EnableFastGetComponentByID {
		entity.componentByIDMap = map[uint64]*container.Element[FaceAny]{}
	}

	entity._eventEntityDestroySelf.Init(false, nil, EventRecursion_Discard, opts.HookCache, entity.opts.Inheritor.IFace)
	entity.eventCompMgrAddComponents.Init(false, nil, EventRecursion_Discard, opts.HookCache, entity.opts.Inheritor.IFace)
	entity.eventCompMgrRemoveComponent.Init(false, nil, EventRecursion_Discard, opts.HookCache, entity.opts.Inheritor.IFace)
}

func (entity *EntityBehavior) getOptions() *EntityOptions {
	return &entity.opts
}

func (entity *EntityBehavior) setID(id uint64) {
	entity.id = id
}

func (entity *EntityBehavior) GetID() uint64 {
	return entity.id
}

func (entity *EntityBehavior) setRuntimeCtx(runtimeCtx RuntimeContext) {
	entity.runtimeCtx = runtimeCtx
}

func (entity *EntityBehavior) GetRuntimeCtx() RuntimeContext {
	return entity.runtimeCtx
}

func (entity *EntityBehavior) setParent(parent Entity) {
	entity.parent = parent
}

func (entity *EntityBehavior) GetParent() (Entity, bool) {
	return entity.parent, entity.parent != nil
}

func (entity *EntityBehavior) setInitialing(v bool) {
	entity.initialing = v
}

func (entity *EntityBehavior) getInitialing() bool {
	return entity.initialing
}

func (entity *EntityBehavior) setShutting(v bool) {
	entity.shutting = v
}

func (entity *EntityBehavior) getShutting() bool {
	return entity.shutting
}

func (entity *EntityBehavior) DestroySelf() {
	emitEventEntityDestroySelf(&entity._eventEntityDestroySelf, entity.opts.Inheritor.IFace)
}

func (entity *EntityBehavior) eventEntityDestroySelf() IEvent {
	return &entity._eventEntityDestroySelf
}
