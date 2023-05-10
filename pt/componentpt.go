package pt

import (
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/uid"
	"reflect"
)

// ComponentPt 组件原型
type ComponentPt struct {
	Name        string // 组件名称
	Path        string // 组件路径
	Description string // 组件描述信息
	tfComp      reflect.Type
}

// Construct 创建组件
func (pt *ComponentPt) Construct(id uid.Id) ec.Component {
	vfComp := reflect.New(pt.tfComp)

	comp := vfComp.Interface().(ec.Component)
	ec.UnsafeComponent(comp).SetReflectValue(vfComp)
	ec.UnsafeComponent(comp).SetId(id)

	return comp
}
