package core

import (
	"github.com/pangdogs/galaxy/core/container"
)

// Entity 实体接口
type Entity interface {
	_InnerGCCollector
	_InnerGC
	_EntityComponentMgr

	init(opts *EntityOptions)

	getOptions() *EntityOptions

	setID(id int64)

	// GetID 获取实体（Entity）全局唯一ID，线程安全
	GetID() int64

	// GetPrototype 获取实体（Entity）原型，线程安全
	GetPrototype() string

	setRuntimeCtx(runtimeCtx RuntimeContext)

	// GetRuntimeCtx 获取运行时上下文（Runtime Context），线程安全
	GetRuntimeCtx() RuntimeContext

	// GetServiceCtx 获取服务上下文（Service Context），线程安全
	GetServiceCtx() ServiceContext

	setParent(parent Entity)

	// GetParent 获取在运行时上下文（Runtime Context）的主EC树上的父实体（Entity），非线程安全
	GetParent() (Entity, bool)

	setInitialing(v bool)

	getInitialing() bool

	setShutting(v bool)

	getShutting() bool

	// DestroySelf 销毁自身，注意在生命周期[Init,InitFin,Shut,ShutFin]中调用无效，非线程安全
	DestroySelf()

	eventEntityDestroySelf() IEvent
}

// EntityGetOptions 获取实体创建选项，线程安全
func EntityGetOptions(e Entity) EntityOptions {
	return *e.getOptions()
}

// EntitySetPersistID 实体（Entity）设置持久化ID，需要在实体加入运行时上下文（Runtime Context）前设置，通常用于实体持久化
func EntitySetPersistID(entity Entity, persistID int64) {
	if persistID <= 0 {
		panic("persistID not invalid")
	}

	if entity.GetRuntimeCtx() != nil {
		panic("entity already added in runtime context")
	}

	entity.setID(persistID)
}

// EntityGetInheritor 获取实体的继承者，线程安全
func EntityGetInheritor[T any](e Entity) T {
	return Cache2IFace[T](e.getOptions().Inheritor.Cache)
}

// EntityGetInitialing 获取实体是否正在初始化，非线程安全
func EntityGetInitialing(e Entity) bool {
	return e.getInitialing()
}

// EntityGetShutting 获取实体是否正在销毁，非线程安全
func EntityGetShutting(e Entity) bool {
	return e.getShutting()
}

// NewEntity 创建实体，线程安全
func NewEntity(optSetterFuncs ...EntityOptionSetterFunc) Entity {
	opts := EntityOptions{}
	EntityOptionSetter.Default()(&opts)

	for i := range optSetterFuncs {
		optSetterFuncs[i](&opts)
	}

	return NewEntityWithOpts(opts)
}

// NewEntityWithOpts 创建实体并传入参数，线程安全
func NewEntityWithOpts(opts EntityOptions) Entity {
	if !opts.Inheritor.IsNil() {
		opts.Inheritor.IFace.init(&opts)
		return opts.Inheritor.IFace
	}

	e := &EntityBehavior{}
	e.init(&opts)

	return e.opts.Inheritor.IFace
}

// EntityBehavior 实体行为，需要在拓展实体能力时，匿名嵌入至实体结构体中，一般情况下无需使用
type EntityBehavior struct {
	id                          int64
	opts                        EntityOptions
	runtimeCtx                  RuntimeContext
	parent                      Entity
	componentList               container.List[FaceAny]
	initialing, shutting        bool
	_eventEntityDestroySelf     Event
	eventCompMgrAddComponents   Event
	eventCompMgrRemoveComponent Event
	gc                          _EntityBehaviorGC
}

func (entity *EntityBehavior) init(opts *EntityOptions) {
	if opts == nil {
		panic("nil opts")
	}

	entity.opts = *opts

	if entity.opts.Inheritor.IsNil() {
		entity.opts.Inheritor = NewFace[Entity](entity)
	}

	entity.gc.EntityBehavior = entity

	entity.componentList.Init(entity.opts.FaceCache, entity.getGCCollector())

	entity._eventEntityDestroySelf.Init(false, nil, EventRecursion_Discard, opts.HookCache, entity.getGCCollector())
	entity.eventCompMgrAddComponents.Init(false, nil, EventRecursion_Discard, opts.HookCache, entity.getGCCollector())
	entity.eventCompMgrRemoveComponent.Init(false, nil, EventRecursion_Discard, opts.HookCache, entity.getGCCollector())
}

func (entity *EntityBehavior) getOptions() *EntityOptions {
	return &entity.opts
}

func (entity *EntityBehavior) setID(id int64) {
	entity.id = id
}

// GetID 获取实体（Entity）全局唯一ID，线程安全
func (entity *EntityBehavior) GetID() int64 {
	return entity.id
}

// GetPrototype 获取实体原型（Entity Prototype），线程安全
func (entity *EntityBehavior) GetPrototype() string {
	return entity.opts.Prototype
}

func (entity *EntityBehavior) setRuntimeCtx(runtimeCtx RuntimeContext) {
	entity.runtimeCtx = runtimeCtx
}

// GetRuntimeCtx 获取运行时上下文（Runtime Context），线程安全
func (entity *EntityBehavior) GetRuntimeCtx() RuntimeContext {
	return entity.runtimeCtx
}

// GetServiceCtx 获取服务上下文（Service Context），线程安全
func (entity *EntityBehavior) GetServiceCtx() ServiceContext {
	return entity.runtimeCtx.GetServiceCtx()
}

func (entity *EntityBehavior) setParent(parent Entity) {
	entity.parent = parent
}

// GetParent 获取在运行时上下文（Runtime Context）的主EC树上的父实体（Entity），非线程安全
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

// DestroySelf 销毁自身，注意在生命周期[Init,InitFin,Shut,ShutFin]中调用无效，非线程安全
func (entity *EntityBehavior) DestroySelf() {
	emitEventEntityDestroySelf(&entity._eventEntityDestroySelf, entity.opts.Inheritor.IFace)
}

func (entity *EntityBehavior) eventEntityDestroySelf() IEvent {
	return &entity._eventEntityDestroySelf
}
