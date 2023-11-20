package ec

import (
	"kit.golaxy.org/golaxy/event"
	"kit.golaxy.org/golaxy/util/container"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/uid"
)

// Deprecated: UnsafeEntity 访问实体内部函数
func UnsafeEntity(entity Entity) _UnsafeEntity {
	return _UnsafeEntity{
		Entity: entity,
	}
}

type _UnsafeEntity struct {
	Entity
}

// Init 初始化
func (ue _UnsafeEntity) Init(opts EntityOptions) {
	ue.init(opts)
}

// GetOptions 获取实体所有选项
func (ue _UnsafeEntity) GetOptions() *EntityOptions {
	return ue.getOptions()
}

// SetId 设置Id
func (ue _UnsafeEntity) SetId(id uid.Id) {
	ue.setId(id)
}

// SetContext 设置上下文
func (ue _UnsafeEntity) SetContext(ctx iface.Cache) {
	ue.setContext(ctx)
}

// GetVersion 获取组件列表变化版本号
func (ue _UnsafeEntity) GetVersion() int32 {
	return ue.getVersion()
}

// SetGCCollector 设置GC收集器
func (ue _UnsafeEntity) SetGCCollector(gcCollector container.GCCollector) {
	ue.setGCCollector(gcCollector)
}

// SetECNodeState 设置EC节点状态
func (ue _UnsafeEntity) SetECNodeState(state ECNodeState) {
	ue.setECNodeState(state)
}

// SetECParent 设置在EC树中的父实体
func (ue _UnsafeEntity) SetECParent(parent Entity) {
	ue.setECParent(parent)
}

// SetState 设置状态
func (ue _UnsafeEntity) SetState(state EntityState) {
	ue.setState(state)
}

// EventEntityDestroySelf 事件：实体销毁自身
func (ue _UnsafeEntity) EventEntityDestroySelf() event.IEvent {
	return ue.eventEntityDestroySelf()
}

func (uc _UnsafeEntity) CleanHooks() {
	uc.cleanHooks()
}
