package define

import (
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/util/types"
)

// DefineComponentInterface 定义组件接口
func DefineComponentInterface[COMP_IFACE any]() ComponentInterface {
	return _ComponentInterface{
		name: types.FullName[COMP_IFACE](),
	}.ComponentInterface()
}

// ComponentInterface 组件接口
type ComponentInterface struct {
	Name     string                          // 组件接口名
	Register func(comp any, descr ...string) // 注册组件原型
}

type _ComponentInterface struct {
	name string
}

func (c _ComponentInterface) register() func(comp any, descr ...string) {
	return func(comp any, descr ...string) {
		pt.RegisterComponent(c.name, comp, descr...)
	}
}

// ComponentInterface 生成组件接口定义
func (c _ComponentInterface) ComponentInterface() ComponentInterface {
	return ComponentInterface{
		Name:     c.name,
		Register: c.register(),
	}
}
