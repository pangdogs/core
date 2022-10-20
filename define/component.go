package define

import "github.com/pangdogs/galaxy/util"

type _Component[T any] struct {
	name string
}

// Name 生成组件名称
func (c _Component[T]) Name() string {
	return c.name
}

// Component 用于定义组件
func Component[T any]() _Component[T] {
	return _Component[T]{
		name: util.TypeFullName[T](),
	}
}
