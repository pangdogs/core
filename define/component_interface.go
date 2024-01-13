package define

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/types"
)

// DefineComponentInterface 定义组件接口
func DefineComponentInterface[COMP_IFACE any](compLib pt.ComponentLib) ComponentInterface {
	if compLib == nil {
		panic(fmt.Errorf("%w: %w: compLib is nil", exception.ErrGolaxy, exception.ErrArgs))
	}
	return _ComponentInterface{
		name:    types.FullName[COMP_IFACE](),
		compLib: compLib,
	}.ComponentInterface()
}

// ComponentInterface 组件接口
type ComponentInterface struct {
	Name     string                             // 组件接口名称
	Register generic.Func1[any, pt.ComponentPT] // 注册组件原型
}

type _ComponentInterface struct {
	name    string
	compLib pt.ComponentLib
}

func (c _ComponentInterface) register() generic.Func1[any, pt.ComponentPT] {
	return func(comp any) pt.ComponentPT {
		return c.compLib.Register(comp, c.name)
	}
}

// ComponentInterface 生成组件接口定义
func (c _ComponentInterface) ComponentInterface() ComponentInterface {
	return ComponentInterface{
		Name:     c.name,
		Register: c.register(),
	}
}
