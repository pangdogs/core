package pt

import (
	"github.com/pangdogs/galaxy/core"
	"reflect"
)

type _CompConstructType int32

const (
	_CompConstructType_Reflect _CompConstructType = iota
	_CompConstructType_Creator
)

// ComponentPt 组件原型
type ComponentPt struct {
	Interface     string
	Tag           string
	Description   string
	constructType _CompConstructType
	tfComp        reflect.Type
	creator       func() core.Component
}

// New 创建组件
func (pt *ComponentPt) New() core.Component {
	switch pt.constructType {
	case _CompConstructType_Reflect:
		return reflect.New(pt.tfComp).Interface().(core.Component)
	case _CompConstructType_Creator:
		return pt.creator()
	default:
		panic("not support construct type")
	}
}
