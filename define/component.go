package define

import "kit.golaxy.org/golaxy/util/types"

// DefineComponent 定义组件
func DefineComponent[COMP_IFACE, COMP any]() Component {
	compIface := DefineComponentInterface[COMP_IFACE]()
	compIface.Register(types.Zero[COMP]())
	return _Component{
		name:           compIface.Name,
		implementation: types.FullName[COMP](),
	}.Component()
}

// Component 组件
type Component struct {
	Name           string // 组件名称
	Implementation string // 组件实现名称
}

type _Component struct {
	name, implementation string
}

// Component 生成组件定义
func (c _Component) Component() Component {
	return Component{
		Name:           c.name,
		Implementation: c.implementation,
	}
}
