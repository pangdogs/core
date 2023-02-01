package define

import (
	"github.com/golaxy-kit/golaxy/util"
)

type _Component struct {
	_name, _path string
}

func (c _Component) name() string {
	return c._name
}

func (c _Component) path() string {
	return c._path
}

// Component 组件
type Component struct {
	Name string // 组件名
	Path string // 组件路径
}

// Component 生成组件定义
func (c _Component) Component() Component {
	return Component{
		Name: c.name(),
		Path: c.path(),
	}
}

// DefineComponent 定义组件
func DefineComponent[COMP_IFACE, COMP any](descr ...string) Component {
	compIface := DefineComponentInterface[COMP_IFACE]()
	compIface.Register(util.Zero[COMP](), descr...)
	return _Component{
		_name: compIface.Name,
		_path: util.TypeFullName[COMP](),
	}.Component()
}
