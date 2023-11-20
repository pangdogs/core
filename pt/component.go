package pt

import (
	"kit.golaxy.org/golaxy/ec"
	"reflect"
)

// ComponentPT 组件原型
type ComponentPT struct {
	Name           string // 组件名称
	Implementation string // 组件实例名称
	tfComp         reflect.Type
}

// Construct 创建组件
func (pt ComponentPT) Construct() ec.Component {
	vfComp := reflect.New(pt.tfComp)

	comp := vfComp.Interface().(ec.Component)
	ec.UnsafeComponent(comp).SetReflectValue(vfComp)

	return comp
}
