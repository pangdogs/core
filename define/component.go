package define

import (
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/util/types"
)

// DefineComponent 定义组件
func DefineComponent[COMP any](compLib pt.ComponentLib) Component {
	comp := DefineComponentInterface[COMP](compLib).Register(types.Zero[COMP]())
	return _Component{
		name:           comp.Name,
		implementation: comp.Implementation,
	}.Component()
}

// DefineComponentWithInterface 定义有接口的组件，接口名称将作为组件名
func DefineComponentWithInterface[COMP, COMP_IFACE any](compLib pt.ComponentLib) Component {
	comp := DefineComponentInterface[COMP_IFACE](compLib).Register(types.Zero[COMP]())
	return _Component{
		name:           comp.Name,
		implementation: comp.Implementation,
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
