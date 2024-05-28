package define

import (
	"fmt"
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/types"
)

// Component 定义组件
func Component[COMP any](compLib ...pt.ComponentLib) ComponentDefinition {
	return defineComponent[COMP](getCompLib(compLib...), "")
}

// ComponentWithInterface 定义有接口的组件，接口名称将作为组件名
func ComponentWithInterface[COMP, COMP_IFACE any](compLib ...pt.ComponentLib) ComponentDefinition {
	_compLib := getCompLib(compLib...)
	return defineComponent[COMP](_compLib, ComponentInterface[COMP_IFACE](_compLib).Name)
}

// ComponentDefinition 组件定义
type ComponentDefinition struct {
	Name          string // 组件名称
	InterfaceName string // 组件接口名称
}

func defineComponent[COMP any](compLib pt.ComponentLib, ifaceName string) ComponentDefinition {
	if compLib == nil {
		panic(fmt.Errorf("%w: %w: compLib is nil", exception.ErrCore, exception.ErrArgs))
	}

	return ComponentDefinition{
		Name:          compLib.Declare(types.ZeroT[COMP]()).Name,
		InterfaceName: ifaceName,
	}
}

func getCompLib(compLib ...pt.ComponentLib) pt.ComponentLib {
	if len(compLib) > 0 && compLib[0] != nil {
		return compLib[0]
	}
	return pt.DefaultComponentLib()
}
