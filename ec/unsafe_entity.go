package ec

import (
	"kit.golaxy.org/golaxy/localevent"
	"kit.golaxy.org/golaxy/uid"
	"kit.golaxy.org/golaxy/util"
	"kit.golaxy.org/golaxy/util/container"
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
func (ue _UnsafeEntity) Init(opts *EntityOptions) {
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
func (ue _UnsafeEntity) SetContext(ctx util.IfaceCache) {
	ue.setContext(ctx)
}

// GetChangedVersion 获取组件列表变化版本号
func (ue _UnsafeEntity) GetChangedVersion() int64 {
	return ue.getChangedVersion()
}

// SetGCCollector 设置GC收集器
func (ue _UnsafeEntity) SetGCCollector(gcCollector container.GCCollector) {
	ue.setGCCollector(gcCollector)
}

// GetGCCollector 获取GC收集器
func (ue _UnsafeEntity) GetGCCollector() container.GCCollector {
	return ue.getGCCollector()
}

// SetParent 设置父实体
func (ue _UnsafeEntity) SetParent(parent Entity) {
	ue.setParent(parent)
}

// SetState 设置状态
func (ue _UnsafeEntity) SetState(state EntityState) {
	ue.setState(state)
}

// EventEntityDestroySelf 事件：实体销毁自身
func (ue _UnsafeEntity) EventEntityDestroySelf() localevent.IEvent {
	return ue.eventEntityDestroySelf()
}
