package define

import (
	"github.com/galaxy-kit/galaxy-go/pt"
	"github.com/galaxy-kit/galaxy-go/util"
)

type _Component[COMP any] struct {
	_name, _path string
}

func (c _Component[COMP]) name() string {
	return c._name
}

func (c _Component[COMP]) path() string {
	return c._path
}

// Component 组件
type Component struct {
	Name string
	Path string
}

// Component 生成组件定义
func (c _Component[COMP]) Component(descr ...string) Component {
	pt.RegisterComponent[COMP](c.name(), descr...)
	return Component{
		Name: c.name(),
		Path: c.path(),
	}
}

// DefineComponent 定义组件
func DefineComponent[COMP_IFACE, COMP any]() _Component[COMP] {
	return _Component[COMP]{
		_name: util.TypeFullName[COMP_IFACE](),
		_path: util.TypeFullName[COMP](),
	}
}
