package define

import (
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/pt"
	"github.com/pangdogs/galaxy/util"
)

type _Component[T any] struct {
	name string
}

// Name 生成组件名称
func (c _Component[T]) Name() string {
	return c.name
}

// Register 生成注册组件原型函数
func (c _Component[T]) Register() func(descr string, comp any) {
	return func(descr string, comp any) {
		pt.RegisterComponent(c.Name(), descr, comp)
	}
}

// RegisterCreator 生成注册组件构件函数的函数
func (c _Component[T]) RegisterCreator() func(descr string, creator func() ec.Component) {
	return func(descr string, creator func() ec.Component) {
		pt.RegisterComponentCreator(c.Name(), descr, creator)
	}
}

// Component 用于定义组件
func Component[T any]() _Component[T] {
	return _Component[T]{
		name: util.TypeFullName[T](),
	}
}
