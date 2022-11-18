package pt

import (
	"github.com/galaxy-kit/galaxy-go/ec"
	"reflect"
)

// ComponentPt 组件原型
type ComponentPt struct {
	Name        string // 组件名称
	Path        string // 组件路径
	Description string // 组件描述信息
	tfComp      reflect.Type
}

// New 创建组件
func (pt *ComponentPt) New() ec.Component {
	vfComp := reflect.New(pt.tfComp)

	comp := vfComp.Interface().(ec.Component)
	ec.UnsafeComponent(comp).SetReflectValue(vfComp)

	return comp
}
