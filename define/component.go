package define

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/util/types"
)

// Component 定义组件
func Component[COMP any](compLib ...pt.ComponentLib) ComponentDefinition {
	_compLib := pt.DefaultComponentLib()

	if len(compLib) > 0 {
		_compLib = compLib[0]
	}

	if _compLib == nil {
		panic(fmt.Errorf("%w: %w: compLib is nil", exception.ErrCore, exception.ErrArgs))
	}

	comp := _compLib.Declare(types.Zero[COMP]())
	return _DefineComponent{
		name: comp.Name,
	}.Component()
}

// ComponentWithInterface 定义有接口的组件，接口名称将作为组件名
func ComponentWithInterface[COMP, COMP_IFACE any](compLib ...pt.ComponentLib) ComponentDefinition {
	_compLib := pt.DefaultComponentLib()

	if len(compLib) > 0 {
		_compLib = compLib[0]
	}

	if _compLib == nil {
		panic(fmt.Errorf("%w: %w: compLib is nil", exception.ErrCore, exception.ErrArgs))
	}

	comp, ifaceName := ComponentInterface[COMP_IFACE](_compLib).Declare(types.Zero[COMP]())
	return _DefineComponent{
		name:          comp.Name,
		interfaceName: ifaceName,
	}.Component()
}

// ComponentDefinition 组件定义
type ComponentDefinition struct {
	Name          string // 组件名称
	InterfaceName string // 组件接口名称
}

type _DefineComponent struct {
	name, interfaceName string
}

// Component 生成组件定义
func (d _DefineComponent) Component() ComponentDefinition {
	return ComponentDefinition{
		Name:          d.name,
		InterfaceName: d.interfaceName,
	}
}
