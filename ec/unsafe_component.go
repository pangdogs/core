package ec

import (
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/util/container"
	"git.golaxy.org/core/util/uid"
	"reflect"
)

// Deprecated: UnsafeComponent 访问组件内部函数
func UnsafeComponent(comp Component) _UnsafeComponent {
	return _UnsafeComponent{
		Component: comp,
	}
}

type _UnsafeComponent struct {
	Component
}

// Init 初始化
func (uc _UnsafeComponent) Init(name string, entity Entity, composite Component, hookAllocator container.Allocator[event.Hook], gcCollector container.GCCollector) {
	uc.init(name, entity, composite, hookAllocator, gcCollector)
}

// SetId 设置Id
func (uc _UnsafeComponent) SetId(id uid.Id) {
	uc.setId(id)
}

// SetFixed 设置为固定的（不可删除）
func (uc _UnsafeComponent) SetFixed(v bool) {
	uc.setFixed(v)
}

// GetFixed 获取是否为固定的（不可删除）
func (uc _UnsafeComponent) GetFixed() bool {
	return uc.getFixed()
}

// SetState 设置状态
func (uc _UnsafeComponent) SetState(state ComponentState) {
	uc.setState(state)
}

// SetReflected 设置反射值
func (uc _UnsafeComponent) SetReflected(v reflect.Value) {
	uc.setReflected(v)
}

// GetReflected 获取反射值
func (uc _UnsafeComponent) GetReflected() reflect.Value {
	return uc.getReflected()
}

// GetComposite 获取扩展者
func (uc _UnsafeComponent) GetComposite() Component {
	return uc.getComposite()
}

// SetGCCollector 设置GC收集器
func (uc _UnsafeComponent) SetGCCollector(gcCollector container.GCCollector) {
	uc.setGCCollector(gcCollector)
}

// EventComponentDestroySelf 事件：组件销毁自身
func (uc _UnsafeComponent) EventComponentDestroySelf() event.IEvent {
	return uc.eventComponentDestroySelf()
}

func (uc _UnsafeComponent) CleanHooks() {
	uc.cleanHooks()
}
