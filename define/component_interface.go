package define

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/types"
)

// DefineComponentInterface 定义组件接口
func DefineComponentInterface[COMP_IFACE any](compLib ...pt.ComponentLib) ComponentInterface {
	_compLib := pt.DefaultComponentLib()

	if len(compLib) > 0 {
		_compLib = compLib[0]
	}

	if _compLib == nil {
		panic(fmt.Errorf("%w: %w: compLib is nil", exception.ErrCore, exception.ErrArgs))
	}

	return _ComponentInterface{
		name:    types.FullName[COMP_IFACE](),
		compLib: _compLib,
	}.ComponentInterface()
}

// ComponentInterface 组件接口
type ComponentInterface struct {
	Name    string                                         // 组件接口名称
	Declare generic.PairFunc1[any, pt.ComponentPT, string] // 声明组件原型
}

type _ComponentInterface struct {
	name    string
	compLib pt.ComponentLib
}

func (c _ComponentInterface) declare() generic.PairFunc1[any, pt.ComponentPT, string] {
	return func(comp any) (pt.ComponentPT, string) {
		return c.compLib.Declare(comp, c.name), c.name
	}
}

// ComponentInterface 生成组件接口定义
func (c _ComponentInterface) ComponentInterface() ComponentInterface {
	return ComponentInterface{
		Name:    c.name,
		Declare: c.declare(),
	}
}
