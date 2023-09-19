package ec

import (
	"fmt"
	"kit.golaxy.org/golaxy/event"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util/container"
	"kit.golaxy.org/golaxy/util/iface"
)

// NewEntity 创建实体
func NewEntity(options ...EntityOption) Entity {
	opts := EntityOptions{}
	Option{}.Default()(&opts)

	for i := range options {
		options[i](&opts)
	}

	return UnsafeNewEntity(opts)
}

// Deprecated: UnsafeNewEntity 内部创建实体
func UnsafeNewEntity(options EntityOptions) Entity {
	if !options.CompositeFace.IsNil() {
		options.CompositeFace.Iface.init(&options)
		return options.CompositeFace.Iface
	}

	e := &EntityBehavior{}
	e.init(&options)

	return e.opts.CompositeFace.Iface
}

// Entity 实体接口
type Entity interface {
	_Entity
	_ComponentMgr
	internal.ContextResolver
	fmt.Stringer

	// GetId 获取实体Id
	GetId() uid.Id
	// GetPrototype 获取实体原型
	GetPrototype() string
	// GetParent 获取在运行时上下文的主EC树上的父实体
	GetParent() (Entity, bool)
	// GetState 获取实体状态
	GetState() EntityState
	// DestroySelf 销毁自身
	DestroySelf()
}

type _Entity interface {
	init(opts *EntityOptions)
	getOptions() *EntityOptions
	setId(id uid.Id)
	setContext(ctx iface.Cache)
	getVersion() int32
	setGCCollector(gcCollector container.GCCollector)
	getGCCollector() container.GCCollector
	setParent(parent Entity)
	setState(state EntityState)
	eventEntityDestroySelf() event.IEvent
}

// EntityBehavior 实体行为，在需要扩展实体能力时，匿名嵌入至实体结构体中
type EntityBehavior struct {
	opts                             EntityOptions
	context                          iface.Cache
	parent                           Entity
	componentList                    container.List[iface.FaceAny]
	version                          int32
	state                            EntityState
	_eventEntityDestroySelf          event.Event
	eventCompMgrAddComponents        event.Event
	eventCompMgrRemoveComponent      event.Event
	eventCompMgrFirstAccessComponent event.Event
}

// GetId 获取实体Id
func (entity *EntityBehavior) GetId() uid.Id {
	return entity.opts.PersistId
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
	case EntityState_Init, EntityState_Inited, EntityState_Living:
		emitEventEntityDestroySelf(entity.eventEntityDestroySelf(), entity.opts.CompositeFace.Iface)
	}
}

// ResolveContext 解析上下文
func (entity *EntityBehavior) ResolveContext() iface.Cache {
	return entity.context
}

// String implements fmt.Stringer
func (entity *EntityBehavior) String() string {
	var parentInfo string
	if parent, ok := entity.GetParent(); ok {
		parentInfo = parent.GetId().String()
	} else {
		parentInfo = "nil"
	}

	return fmt.Sprintf("{Id:%s Prototype:%s Parent:%s State:%s}", entity.GetId(), entity.GetPrototype(), parentInfo, entity.GetState())
}

func (entity *EntityBehavior) init(opts *EntityOptions) {
	if opts == nil {
		panic(fmt.Errorf("%w: %w: opts is nil", ErrEC, internal.ErrArgs))
	}

	entity.opts = *opts

	if entity.opts.CompositeFace.IsNil() {
		entity.opts.CompositeFace = iface.NewFace[Entity](entity)
	}

	entity.componentList.Init(entity.opts.FaceAnyAllocator, entity.opts.GCCollector)

	entity._eventEntityDestroySelf.Init(false, nil, event.EventRecursion_NotEmit, entity.opts.HookAllocator, entity.opts.GCCollector)
	entity.eventCompMgrAddComponents.Init(false, nil, event.EventRecursion_Allow, entity.opts.HookAllocator, entity.opts.GCCollector)
	entity.eventCompMgrRemoveComponent.Init(false, nil, event.EventRecursion_Allow, entity.opts.HookAllocator, entity.opts.GCCollector)
	entity.eventCompMgrFirstAccessComponent.Init(false, nil, event.EventRecursion_Allow, entity.opts.HookAllocator, entity.opts.GCCollector)
}

func (entity *EntityBehavior) getOptions() *EntityOptions {
	return &entity.opts
}

func (entity *EntityBehavior) setId(id uid.Id) {
	entity.opts.PersistId = id
}

func (entity *EntityBehavior) setContext(ctx iface.Cache) {
	entity.context = ctx
}

func (entity *EntityBehavior) getVersion() int32 {
	return entity.version
}

func (entity *EntityBehavior) setGCCollector(gcCollector container.GCCollector) {
	if entity.opts.GCCollector == gcCollector {
		return
	}

	entity.opts.GCCollector = gcCollector

	entity.componentList.SetGCCollector(gcCollector)
	entity.componentList.Traversal(func(e *container.Element[iface.FaceAny]) bool {
		comp := iface.Cache2Iface[Component](e.Value.Cache)
		comp.setGCCollector(gcCollector)
		return true
	})

	event.UnsafeEvent(&entity._eventEntityDestroySelf).SetGCCollector(gcCollector)
	event.UnsafeEvent(&entity.eventCompMgrAddComponents).SetGCCollector(gcCollector)
	event.UnsafeEvent(&entity.eventCompMgrRemoveComponent).SetGCCollector(gcCollector)
}

func (entity *EntityBehavior) getGCCollector() container.GCCollector {
	return entity.opts.GCCollector
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

func (entity *EntityBehavior) eventEntityDestroySelf() event.IEvent {
	return &entity._eventEntityDestroySelf
}
