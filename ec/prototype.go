package ec

import (
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/option"
	"reflect"
)

// EntityPT 实体原型接口
type EntityPT interface {
	// Prototype 实体原型名称
	Prototype() string
	// InstanceRT 实体实例反射类型
	InstanceRT() reflect.Type
	// Scope 可访问作用域
	Scope() *Scope
	// ComponentAwakeOnFirstTouch 开启组件被首次访问时，检测并调用Awake()
	ComponentAwakeOnFirstTouch() *bool
	// ComponentUniqueID 开启组件唯一Id
	ComponentUniqueID() *bool
	// Extra 自定义原型属性
	Extra() generic.SliceMap[string, any]
	// CountComponents // 组件数量
	CountComponents() int
	// Component 获取组件
	Component(idx int) BuiltinComponent
	// Components 获取所有组件
	Components() []BuiltinComponent
	// Construct 创建实体
	Construct(settings ...option.Setting[EntityOptions]) Entity
}

// BuiltinComponent 实体原型中的组件信息
type BuiltinComponent struct {
	PT           ComponentPT                   // 组件原型
	Offset       int                           // 组件位置
	Name         string                        // 组件名称
	NonRemovable bool                          // 不可删除
	Extra        generic.SliceMap[string, any] // 自定义原型属性
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
func (_NoneEntityPT) Scope() *Scope {
	return nil
}

// ComponentAwakeOnFirstTouch 开启组件被首次访问时，检测并调用Awake()
func (_NoneEntityPT) ComponentAwakeOnFirstTouch() *bool {
	return nil
}

// ComponentUniqueID 开启组件唯一Id
func (_NoneEntityPT) ComponentUniqueID() *bool {
	return nil
}

// Extra 自定义原型属性
func (_NoneEntityPT) Extra() generic.SliceMap[string, any] {
	return nil
}

// CountComponents // 组件数量
func (_NoneEntityPT) CountComponents() int {
	return 0
}

// Component 获取组件
func (_NoneEntityPT) Component(idx int) BuiltinComponent {
	exception.Panicf("%w: %w: idx out of range", ErrEC, exception.ErrArgs)
	panic("unreachable")
}

// Components 获取所有组件
func (_NoneEntityPT) Components() []BuiltinComponent {
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
