package define

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/util/types"
)

// DefineComponent 定义组件
func DefineComponent[COMP any](compLib ...pt.ComponentLib) Component {
	_compLib := pt.DefaultComponentLib()

	if len(compLib) > 0 {
		_compLib = compLib[0]
	}

	if _compLib == nil {
		panic(fmt.Errorf("%w: %w: compLib is nil", exception.ErrCore, exception.ErrArgs))
	}

	comp := _compLib.Register(types.Zero[COMP]())
	return _Component{
		name: comp.Name,
	}.Component()
}

// DefineComponentWithInterface 定义有接口的组件，接口名称将作为组件名
func DefineComponentWithInterface[COMP, COMP_IFACE any](compLib ...pt.ComponentLib) Component {
	_compLib := pt.DefaultComponentLib()

	if len(compLib) > 0 {
		_compLib = compLib[0]
	}

	if _compLib == nil {
		panic(fmt.Errorf("%w: %w: compLib is nil", exception.ErrCore, exception.ErrArgs))
	}

	comp, ifaceName := DefineComponentInterface[COMP_IFACE](_compLib).Register(types.Zero[COMP]())
	return _Component{
		name:          comp.Name,
		interfaceName: ifaceName,
	}.Component()
}

// Component 组件
type Component struct {
	Name          string // 组件名称
	InterfaceName string // 组件接口名称
}

type _Component struct {
	name, interfaceName string
}

// Component 生成组件定义
func (c _Component) Component() Component {
	return Component{
		Name:          c.name,
		InterfaceName: c.interfaceName,
	}
}
