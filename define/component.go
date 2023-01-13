package define

import (
	"github.com/golaxy-kit/golaxy/util"
)

type _Component struct {
	_ifaceName, _name string
}

func (c _Component) ifaceName() string {
	return c._ifaceName
}

func (c _Component) name() string {
	return c._name
}

// Component 组件
type Component struct {
	IfaceName string // 组件接口名
	Name      string // 组件名
}

// Component 生成组件定义
func (c _Component) Component() Component {
	return Component{
		IfaceName: c.ifaceName(),
		Name:      c.name(),
	}
}

// DefineComponent 定义组件
func DefineComponent[COMP_IFACE, COMP any](descr ...string) Component {
	DefineComponentInterface[COMP_IFACE]().Register(util.Zero[COMP](), descr...)
	return _Component{
		_ifaceName: util.TypeFullName[COMP_IFACE](),
		_name:      util.TypeFullName[COMP](),
	}.Component()
}
