package define

import (
	"github.com/galaxy-kit/galaxy-go/util"
)

type _ComponentInterface struct {
	_name string
}

func (c _ComponentInterface) name() string {
	return c._name
}

// ComponentInterface 组件
type ComponentInterface struct {
	Name string
}

// Component 生成组件定义
func (c _ComponentInterface) Component() ComponentInterface {
	return ComponentInterface{
		Name: c.name(),
	}
}

// DefineComponentInterface 定义组件
func DefineComponentInterface[COMP_IFACE any]() _ComponentInterface {
	return _ComponentInterface{
		_name: util.TypeFullName[COMP_IFACE](),
	}
}
