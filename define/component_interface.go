package define

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/types"
)

// ComponentInterface 定义组件接口
func ComponentInterface[COMP_IFACE any](compLib ...pt.ComponentLib) ComponentInterfaceDefinition {
	_compLib := pt.DefaultComponentLib()

	if len(compLib) > 0 {
		_compLib = compLib[0]
	}

	if _compLib == nil {
		panic(fmt.Errorf("%w: %w: compLib is nil", exception.ErrCore, exception.ErrArgs))
	}

	return _DefineComponentInterface{
		name:    types.FullName[COMP_IFACE](),
		compLib: _compLib,
	}.ComponentInterface()
}

// ComponentInterfaceDefinition 组件接口定义
type ComponentInterfaceDefinition struct {
	Name    string                                         // 组件接口名称
	Declare generic.PairFunc1[any, pt.ComponentPT, string] // 声明组件原型
}

type _DefineComponentInterface struct {
	name    string
	compLib pt.ComponentLib
}

func (d _DefineComponentInterface) declare() generic.PairFunc1[any, pt.ComponentPT, string] {
	return func(comp any) (pt.ComponentPT, string) {
		return d.compLib.Declare(comp, d.name), d.name
	}
}

// ComponentInterface 生成组件接口定义
func (d _DefineComponentInterface) ComponentInterface() ComponentInterfaceDefinition {
	return ComponentInterfaceDefinition{
		Name:    d.name,
		Declare: d.declare(),
	}
}
