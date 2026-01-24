package ec

import (
	"reflect"

	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/meta"
	"git.golaxy.org/core/utils/option"
)

// EntityPT 实体原型接口
type EntityPT interface {
	// Prototype 实体原型名称
	Prototype() string
	// InstanceRT 实体实例反射类型
	InstanceRT() reflect.Type
	// Scope 可访问作用域
	Scope() Scope
	// ComponentAwakeOnFirstTouch 当实体组件首次被访问时，生命周期是否进入唤醒（Awake）
	ComponentAwakeOnFirstTouch() bool
	// ComponentUniqueID 是否为实体组件分配唯一Id
	ComponentUniqueID() bool
	// Meta 原型Meta信息
	Meta() meta.Meta
	// CountComponents // 组件数量
	CountComponents() int
	// GetComponent 获取组件
	GetComponent(idx int) BuiltinComponent
	// ListComponents 获取所有组件
	ListComponents() []BuiltinComponent
	// Construct 创建实体
	Construct(settings ...option.Setting[EntityOptions]) Entity
}

// BuiltinComponent 实体原型中的组件信息
type BuiltinComponent struct {
	PT        ComponentPT // 组件原型
	Offset    int         // 组件位置
	Name      string      // 组件名称
	Removable bool        // 可以删除
	Meta      meta.Meta   // 原型Meta信息
}

// ComponentPT 组件原型接口
type ComponentPT interface {
	// Prototype 组件原型名称
	Prototype() string
	// InstanceRT 组件实例反射类型
	InstanceRT() reflect.Type
	// Construct 创建组件
	Construct() Component
}

var (
	noneEntityPT         = &_NoneEntityPT{}
	noneComponentPT      = &_NoneComponentPT{}
	noneBuiltinComponent = &BuiltinComponent{PT: noneComponentPT, Offset: -1}
)

type _NoneEntityPT struct{}

// Prototype 实体原型名称
func (_NoneEntityPT) Prototype() string {
	return ""
}

// InstanceRT 实体实例反射类型
func (_NoneEntityPT) InstanceRT() reflect.Type {
	return nil
}

// Scope 可访问作用域
func (_NoneEntityPT) Scope() Scope {
	return Scope_Global
}

// ComponentAwakeOnFirstTouch 当实体组件首次被访问时，生命周期是否进入唤醒（Awake）
func (_NoneEntityPT) ComponentAwakeOnFirstTouch() bool {
	return false
}

// ComponentUniqueID 是否为实体组件分配唯一Id
func (_NoneEntityPT) ComponentUniqueID() bool {
	return false
}

// Meta 原型Meta信息
func (_NoneEntityPT) Meta() meta.Meta {
	return nil
}

// CountComponents // 组件数量
func (_NoneEntityPT) CountComponents() int {
	return 0
}

// GetComponent 获取组件
func (_NoneEntityPT) GetComponent(idx int) BuiltinComponent {
	exception.Panicf("%w: %w: idx out of range", ErrEC, exception.ErrArgs)
	panic("unreachable")
}

// ListComponents 获取所有组件
func (_NoneEntityPT) ListComponents() []BuiltinComponent {
	return nil
}

// Construct 创建实体
func (_NoneEntityPT) Construct(settings ...option.Setting[EntityOptions]) Entity {
	exception.Panicf("%w: %w: none prototype", ErrEC, exception.ErrArgs)
	panic("unreachable")
}

type _NoneComponentPT struct{}

// Prototype 组件原型名称
func (_NoneComponentPT) Prototype() string {
	return ""
}

// InstanceRT 组件实例反射类型
func (_NoneComponentPT) InstanceRT() reflect.Type {
	return nil
}

// Construct 创建组件
func (_NoneComponentPT) Construct() Component {
	exception.Panicf("%w: %w: none prototype", ErrEC, exception.ErrArgs)
	panic("unreachable")
}
