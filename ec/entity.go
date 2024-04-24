package ec

import (
	"fmt"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/internal/concurrent"
	"git.golaxy.org/core/util/container"
	"git.golaxy.org/core/util/iface"
	"git.golaxy.org/core/util/option"
	"git.golaxy.org/core/util/reinterpret"
	"git.golaxy.org/core/util/uid"
	"reflect"
)

// NewEntity 创建实体
func NewEntity(settings ...option.Setting[EntityOptions]) Entity {
	return UnsafeNewEntity(option.Make(With.Default(), settings...))
}

// Deprecated: UnsafeNewEntity 内部创建实体
func UnsafeNewEntity(options EntityOptions) Entity {
	if !options.CompositeFace.IsNil() {
		options.CompositeFace.Iface.init(options)
		return options.CompositeFace.Iface
	}

	e := &EntityBehavior{}
	e.init(options)

	return e.opts.CompositeFace.Iface
}

// Entity 实体接口
type Entity interface {
	iEntity
	iComponentMgr
	concurrent.CurrentContextProvider
	reinterpret.CompositeProvider
	fmt.Stringer

	// GetId 获取实体Id
	GetId() uid.Id
	// GetPrototype 获取实体原型
	GetPrototype() string
	// GetScope 获取可访问作用域
	GetScope() Scope
	// GetTreeNodeState 获取实体树节点状态
	GetTreeNodeState() TreeNodeState
	// GetTreeNodeParent 获取在实体树中的父实体
	GetTreeNodeParent() (Entity, bool)
	// GetState 获取实体状态
	GetState() EntityState
	// GetReflected 获取反射值
	GetReflected() reflect.Value
	// GetMeta 获取Meta信息
	GetMeta() Meta
	// DestroySelf 销毁自身
	DestroySelf()
}

type iEntity interface {
	init(opts EntityOptions)
	getOptions() *EntityOptions
	setId(id uid.Id)
	setContext(ctx iface.Cache)
	getVersion() int64
	setTreeNodeState(state TreeNodeState)
	setTreeNodeParent(parent Entity)
	setState(state EntityState)
	setReflected(v reflect.Value)
	eventEntityDestroySelf() event.IEvent
	cleanManagedHooks()
}

// EntityBehavior 实体行为，在需要扩展实体能力时，匿名嵌入至实体结构体中
type EntityBehavior struct {
	opts                             EntityOptions
	context                          iface.Cache
	treeNodeState                    TreeNodeState
	parent                           Entity
	componentList                    container.List[iface.FaceAny]
	state                            EntityState
	reflected                        reflect.Value
	_eventEntityDestroySelf          event.Event
	eventCompMgrAddComponents        event.Event
	eventCompMgrRemoveComponent      event.Event
	eventCompMgrFirstAccessComponent event.Event
	managedHooks                     []event.Hook
}

// GetId 获取实体Id
func (entity *EntityBehavior) GetId() uid.Id {
	return entity.opts.PersistId
}

// GetPrototype 获取实体原型
func (entity *EntityBehavior) GetPrototype() string {
	return entity.opts.Prototype
}

// GetScope 获取可访问作用域
func (entity *EntityBehavior) GetScope() Scope {
	return entity.opts.Scope
}

// GetTreeNodeState 获取实体树节点状态
func (entity *EntityBehavior) GetTreeNodeState() TreeNodeState {
	return entity.treeNodeState
}

// GetTreeNodeParent 获取在实体树中的父实体
func (entity *EntityBehavior) GetTreeNodeParent() (Entity, bool) {
	return entity.parent, entity.parent != nil
}

// GetState 获取实体状态
func (entity *EntityBehavior) GetState() EntityState {
	return entity.state
}

// GetReflected 获取反射值
func (entity *EntityBehavior) GetReflected() reflect.Value {
	if entity.reflected.IsValid() {
		return entity.reflected
	}
	entity.reflected = reflect.ValueOf(entity.opts.CompositeFace.Iface)
	return entity.reflected
}

// GetMeta 获取Meta信息
func (entity *EntityBehavior) GetMeta() Meta {
	return entity.opts.Meta
}

// DestroySelf 销毁自身
func (entity *EntityBehavior) DestroySelf() {
	switch entity.GetState() {
	case EntityState_Awake, EntityState_Start, EntityState_Living:
		_EmitEventEntityDestroySelf(UnsafeEntity(entity), entity.opts.CompositeFace.Iface)
	}
}

// GetCurrentContext 获取当前上下文
func (entity *EntityBehavior) GetCurrentContext() iface.Cache {
	return entity.context
}

// GetConcurrentContext 解析线程安全的上下文
func (entity *EntityBehavior) GetConcurrentContext() iface.Cache {
	return entity.context
}

// GetCompositeFaceCache 支持重新解释类型
func (entity *EntityBehavior) GetCompositeFaceCache() iface.Cache {
	return entity.opts.CompositeFace.Cache
}

// String implements fmt.Stringer
func (entity *EntityBehavior) String() string {
	return fmt.Sprintf(`{"id":%q, "prototype":%q}`, entity.GetId(), entity.GetPrototype())
}

func (entity *EntityBehavior) init(opts EntityOptions) {
	entity.opts = opts

	if entity.opts.CompositeFace.IsNil() {
		entity.opts.CompositeFace = iface.MakeFace[Entity](entity)
	}

	entity._eventEntityDestroySelf.Init(false, nil, event.EventRecursion_Discard)
	entity.eventCompMgrAddComponents.Init(false, nil, event.EventRecursion_Allow)
	entity.eventCompMgrRemoveComponent.Init(false, nil, event.EventRecursion_Allow)
	entity.eventCompMgrFirstAccessComponent.Init(false, nil, event.EventRecursion_Allow)
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

func (entity *EntityBehavior) getVersion() int64 {
	return entity.componentList.Version()
}

func (entity *EntityBehavior) setTreeNodeState(state TreeNodeState) {
	entity.treeNodeState = state
}

func (entity *EntityBehavior) setTreeNodeParent(parent Entity) {
	entity.parent = parent
}

func (entity *EntityBehavior) setState(state EntityState) {
	if state <= entity.state {
		return
	}

	entity.state = state

	switch entity.state {
	case EntityState_Leave:
		entity._eventEntityDestroySelf.Close()
	case EntityState_Shut:
		entity.eventCompMgrAddComponents.Close()
		entity.eventCompMgrRemoveComponent.Close()
		entity.eventCompMgrFirstAccessComponent.Close()
	}
}

func (entity *EntityBehavior) setReflected(v reflect.Value) {
	entity.reflected = v
}

func (entity *EntityBehavior) eventEntityDestroySelf() event.IEvent {
	return &entity._eventEntityDestroySelf
}
