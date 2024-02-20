package pt

import (
	"git.golaxy.org/core/ec"
	"reflect"
)

// ComponentPT 组件原型
type ComponentPT struct {
	Name   string // 组件名称
	tfComp reflect.Type
}

// Construct 创建组件
func (pt ComponentPT) Construct() ec.Component {
	vfComp := reflect.New(pt.tfComp)

	comp := vfComp.Interface().(ec.Component)
	ec.UnsafeComponent(comp).SetReflected(vfComp)

	return comp
}
