package define

import (
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/pt"
	"github.com/pangdogs/galaxy/util"
)

type _ComponentPt[T any] struct {
	name string
}

// Name 生成组件名称
func (c _ComponentPt[T]) Name() string {
	return c.name
}

// Register 生成注册组件原型函数
func (c _ComponentPt[T]) Register() func(descr string, comp any) {
	return func(descr string, comp any) {
		pt.RegisterComponent(c.Name(), descr, comp)
	}
}

// RegisterCreator 生成注册组件构件函数的函数
func (c _ComponentPt[T]) RegisterCreator() func(descr string, creator func() ec.Component) {
	return func(descr string, creator func() ec.Component) {
		pt.RegisterComponentCreator(c.Name(), descr, creator)
	}
}

// ComponentPt 组件原型
type ComponentPt struct {
	Name            string
	Register        func(descr string, comp any)
	RegisterCreator func(descr string, creator func() ec.Component)
}

// ComponentPt 生成组件原型定义
func (c _ComponentPt[T]) ComponentPt() ComponentPt {
	return ComponentPt{
		Name:            c.Name(),
		Register:        c.Register(),
		RegisterCreator: c.RegisterCreator(),
	}
}

// DefineComponentPt 定义组件原型
func DefineComponentPt[T any]() _ComponentPt[T] {
	return _ComponentPt[T]{
		name: util.TypeFullName[T](),
	}
}
