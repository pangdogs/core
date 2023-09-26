package define

import (
	"kit.golaxy.org/golaxy/util/types"
)

type _Component struct {
	_name, _implementation string
}

func (c _Component) name() string {
	return c._name
}

func (c _Component) implementation() string {
	return c._implementation
}

// Component 组件
type Component struct {
	Name           string // 组件名
	Implementation string // 组件实现
}

// Component 生成组件定义
func (c _Component) Component() Component {
	return Component{
		Name:           c.name(),
		Implementation: c.implementation(),
	}
}

// DefineComponent 定义组件
func DefineComponent[COMP_IFACE, COMP any](descr ...string) Component {
	compIface := DefineComponentInterface[COMP_IFACE]()
	compIface.Register(types.Zero[COMP](), descr...)
	return _Component{
		_name:           compIface.Name,
		_implementation: types.FullName[COMP](),
	}.Component()
}
