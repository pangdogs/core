package ec

import (
	"github.com/pangdogs/galaxy/localevent"
	"github.com/pangdogs/galaxy/util"
	"github.com/pangdogs/galaxy/util/container"
)

// Entity 实体接口
type Entity interface {
	_InnerGC
	_InnerGCCollector
	_ComponentMgr
	ContextHolder

	init(opts *EntityOptions)

	getOptions() *EntityOptions

	setID(id int64)

	// GetID 获取实体全局唯一ID
	GetID() int64

	setSerialNo(sn int64)

	// GetSerialNo 获取序列号
	GetSerialNo() int64

	// GetPrototype 获取实体原型
	GetPrototype() string

	setContext(ctx util.IfaceCache)

	setGCCollector(gcCollect container.GCCollector)

	getGCCollector() container.GCCollector

	setParent(parent Entity)

	// GetParent 获取在运行时上下文的主EC树上的父实体
	GetParent() (Entity, bool)

	setInitialing(v bool)

	getInitialing() bool

	setShutting(v bool)

	getShutting() bool

	setAdding(v bool)

	getAdding() bool

	setRemoving(v bool)

	getRemoving() bool

	// DestroySelf 销毁自身，注意在生命周期[Init,InitFin,Shut,ShutFin]中调用无效
	DestroySelf()

	eventEntityDestroySelf() localevent.IEvent
}

// NewEntity 创建实体
func NewEntity(options ...EntityOptionSetter) Entity {
	opts := EntityOptions{}
	EntityOption.Default()(&opts)

	for i := range options {
		options[i](&opts)
	}

	return UnsafeNewEntity(opts)
}

func UnsafeNewEntity(options EntityOptions) Entity {
	if !options.Inheritor.IsNil() {
		options.Inheritor.Iface.init(&options)
		return options.Inheritor.Iface
	}

	e := &EntityBehavior{}
	e.init(&options)

	return e.opts.Inheritor.Iface
}

// EntityBehavior 实体行为，在需要拓展实体能力时，匿名嵌入至实体结构体中
type EntityBehavior struct {
	id, serialNo                     int64
	opts                             EntityOptions
	context                          util.IfaceCache
	gcCollector                      container.GCCollector
	parent                           Entity
	componentList                    container.List[util.FaceAny]
	adding, removing                 bool
	initialing, shutting             bool
	_eventEntityDestroySelf          localevent.Event
	eventCompMgrAddComponents        localevent.Event
	eventCompMgrRemoveComponent      localevent.Event
	eventCompMgrFirstAccessComponent localevent.Event
	innerGC                          _EntityInnerGC
}

func (entity *EntityBehavior) init(opts *EntityOptions) {
	if opts == nil {
		panic("nil opts")
	}

	entity.opts = *opts

	if entity.opts.Inheritor.IsNil() {
		entity.opts.Inheritor = util.NewFace[Entity](entity)
	}

	entity.innerGC.Init(entity)

	entity.id = entity.opts.PersistID
	entity.componentList.Init(entity.opts.FaceCache, &entity.innerGC)

	entity._eventEntityDestroySelf.Init(false, nil, localevent.EventRecursion_NotEmit, opts.HookCache, &entity.innerGC)
	entity.eventCompMgrAddComponents.Init(false, nil, localevent.EventRecursion_Discard, opts.HookCache, &entity.innerGC)
	entity.eventCompMgrRemoveComponent.Init(false, nil, localevent.EventRecursion_Discard, opts.HookCache, &entity.innerGC)
	entity.eventCompMgrFirstAccessComponent.Init(false, nil, localevent.EventRecursion_Discard, opts.HookCache, &entity.innerGC)
}

func (entity *EntityBehavior) getOptions() *EntityOptions {
	return &entity.opts
}

func (entity *EntityBehavior) setID(id int64) {
	entity.id = id
}

// GetID 获取实体全局唯一ID
func (entity *EntityBehavior) GetID() int64 {
	return entity.id
}

func (entity *EntityBehavior) setSerialNo(sn int64) {
	entity.serialNo = sn
}

// GetSerialNo 获取序列号
func (entity *EntityBehavior) GetSerialNo() int64 {
	return entity.serialNo
}

// GetPrototype 获取实体原型
func (entity *EntityBehavior) GetPrototype() string {
	return entity.opts.Prototype
}

func (entity *EntityBehavior) setContext(ctx util.IfaceCache) {
	entity.context = ctx
}

func (entity *EntityBehavior) getContext() util.IfaceCache {
	return entity.context
}

func (entity *EntityBehavior) setGCCollector(gcCollect container.GCCollector) {
	entity.gcCollector = gcCollect
}

func (entity *EntityBehavior) getGCCollector() container.GCCollector {
	return entity.gcCollector
}

func (entity *EntityBehavior) setParent(parent Entity) {
	entity.parent = parent
}

// GetParent 获取在运行时上下文的主EC树上的父实体
func (entity *EntityBehavior) GetParent() (Entity, bool) {
	return entity.parent, entity.parent != nil
}

func (entity *EntityBehavior) setAdding(v bool) {
	entity.adding = v
}

func (entity *EntityBehavior) getAdding() bool {
	return entity.adding
}

func (entity *EntityBehavior) setRemoving(v bool) {
	entity.removing = v
}

func (entity *EntityBehavior) getRemoving() bool {
	return entity.removing
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

// DestroySelf 销毁自身，注意在生命周期[Init,InitFin,Shut,ShutFin]中调用无效
func (entity *EntityBehavior) DestroySelf() {
	emitEventEntityDestroySelf(&entity._eventEntityDestroySelf, entity.opts.Inheritor.Iface)
	entity._eventEntityDestroySelf.Close()
}

func (entity *EntityBehavior) eventEntityDestroySelf() localevent.IEvent {
	return &entity._eventEntityDestroySelf
}

func (entity *EntityBehavior) getInnerGC() container.GC {
	return &entity.innerGC
}

func (entity *EntityBehavior) getInnerGCCollector() container.GCCollector {
	return &entity.innerGC
}
