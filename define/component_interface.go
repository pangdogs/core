package define

import (
	"github.com/galaxy-kit/galaxy-go/pt"
	"github.com/galaxy-kit/galaxy-go/util"
)

type _ComponentInterface struct {
	_name string
}

func (c _ComponentInterface) name() string {
	return c._name
}

func (c _ComponentInterface) register() func(comp any, descr ...string) {
	return func(comp any, descr ...string) {
		pt.RegisterComponent(c.name(), comp, descr...)
	}
}

// ComponentInterface 组件接口
type ComponentInterface struct {
	Name     string                          // 组件名
	Register func(comp any, descr ...string) // 注册组件原型
}

// ComponentInterface 生成组件接口定义
func (c _ComponentInterface) ComponentInterface() ComponentInterface {
	return ComponentInterface{
		Name:     c.name(),
		Register: c.register(),
	}
}

// DefineComponentInterface 定义组件接口
func DefineComponentInterface[COMP_IFACE any]() _ComponentInterface {
	return _ComponentInterface{
		_name: util.TypeFullName[COMP_IFACE](),
	}
}
