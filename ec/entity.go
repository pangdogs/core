package ec

import (
	"fmt"
	"github.com/segmentio/ksuid"
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
)

// NewEntity 创建实体
func NewEntity(options ...EntityOption) Entity {
	opts := EntityOptions{}
	WithEntityOption{}.Default()(&opts)

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

// Entity 实体接口
type Entity interface {
	_Entity
	_InnerGC
	_InnerGCCollector
	_ComponentMgr
	ContextResolver

	// GetID 获取实体ID
	GetID() ID
	// GetSerialNo 获取序列号
	GetSerialNo() int64
	// GetPrototype 获取实体原型
	GetPrototype() string
	// GetParent 获取在运行时上下文的主EC树上的父实体
	GetParent() (Entity, bool)
	// GetState 获取实体状态
	GetState() EntityState
	// DestroySelf 销毁自身
	DestroySelf()
	// String 字符串化
	String() string
}

type _Entity interface {
	init(opts *EntityOptions)
	getOptions() *EntityOptions
	setID(id ID)
	setSerialNo(sn int64)
	setContext(ctx util.IfaceCache)
	getChangedVersion() int64
	setGCCollector(gcCollect container.GCCollector)
	getGCCollector() container.GCCollector
	setParent(parent Entity)
	setState(state EntityState)
	eventEntityDestroySelf() localevent.IEvent
}

// EntityBehavior 实体行为，在需要扩展实体能力时，匿名嵌入至实体结构体中
type EntityBehavior struct {
	id                               ID
	serialNo                         int64
	opts                             EntityOptions
	context                          util.IfaceCache
	gcCollector                      container.GCCollector
	parent                           Entity
	componentList                    container.List[util.FaceAny]
	changedVersion                   int64
	state                            EntityState
	_eventEntityDestroySelf          localevent.Event
	eventCompMgrAddComponents        localevent.Event
	eventCompMgrRemoveComponent      localevent.Event
	eventCompMgrFirstAccessComponent localevent.Event
	innerGC                          _EntityInnerGC
}

// GetID 获取实体ID
func (entity *EntityBehavior) GetID() ID {
	return entity.id
}

// GetSerialNo 获取序列号
func (entity *EntityBehavior) GetSerialNo() int64 {
	return entity.serialNo
}

// GetPrototype 获取实体原型
func (entity *EntityBehavior) GetPrototype() string {
	return entity.opts.Prototype
}

// GetParent 获取在运行时上下文的主EC树上的父实体
func (entity *EntityBehavior) GetParent() (Entity, bool) {
	return entity.parent, entity.parent != nil
}

// GetState 获取实体状态
func (entity *EntityBehavior) GetState() EntityState {
	return entity.state
}

// DestroySelf 销毁自身
func (entity *EntityBehavior) DestroySelf() {
	switch entity.GetState() {
	case EntityState_Init, EntityState_Start, EntityState_Living:
		emitEventEntityDestroySelf(&entity._eventEntityDestroySelf, entity.opts.Inheritor.Iface)
	}
}

// String 字符串化
func (entity *EntityBehavior) String() string {
	var parentID string
	if parent, ok := entity.GetParent(); ok {
		parentID = ksuid.KSUID(parent.GetID()).String()
	}

	return fmt.Sprintf("[ID:%s SerialNo:%d Prototype:%s Parent:%s State:%s]",
		ksuid.KSUID(entity.GetID()).String(),
		entity.GetSerialNo(),
		entity.GetPrototype(),
		parentID,
		entity.GetState())
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
	entity.eventCompMgrAddComponents.Init(false, nil, localevent.EventRecursion_Allow, opts.HookCache, &entity.innerGC)
	entity.eventCompMgrRemoveComponent.Init(false, nil, localevent.EventRecursion_Allow, opts.HookCache, &entity.innerGC)
	entity.eventCompMgrFirstAccessComponent.Init(false, nil, localevent.EventRecursion_Allow, opts.HookCache, &entity.innerGC)
}

func (entity *EntityBehavior) getOptions() *EntityOptions {
	return &entity.opts
}

func (entity *EntityBehavior) setID(id ID) {
	entity.id = id
}

func (entity *EntityBehavior) setSerialNo(sn int64) {
	entity.serialNo = sn
}

func (entity *EntityBehavior) setContext(ctx util.IfaceCache) {
	entity.context = ctx
}

func (entity *EntityBehavior) getContext() util.IfaceCache {
	return entity.context
}

func (entity *EntityBehavior) getChangedVersion() int64 {
	return entity.changedVersion
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

func (entity *EntityBehavior) setState(state EntityState) {
	if state <= entity.state {
		return
	}
	entity.state = state
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
