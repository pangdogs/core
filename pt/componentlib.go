package pt

import (
	"fmt"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/internal/exception"
	"kit.golaxy.org/golaxy/util/generic"
	"kit.golaxy.org/golaxy/util/types"
	"reflect"
	"sync"
)

// ComponentLib 组件原型库
type ComponentLib interface {
	// Register 注册组件原型，不设置组件名称时，将会使用组件实例名称作为组件名称
	Register(comp any, name ...string) ComponentPT
	// Deregister 取消注册组件原型
	Deregister(impl string)
	// Get 获取组件原型
	Get(impl string) (ComponentPT, bool)
	// Range 遍历所有已注册的组件原型
	Range(fun generic.Func1[ComponentPT, bool])
}

var compLib = NewComponentLib()

// DefaultComponentLib 默认组件库
func DefaultComponentLib() ComponentLib {
	return compLib
}

// NewComponentLib 创建组件原型库
func NewComponentLib() ComponentLib {
	return &_ComponentLib{
		compMap: map[string]*ComponentPT{},
	}
}

type _ComponentLib struct {
	sync.RWMutex
	compMap  map[string]*ComponentPT
	compList []*ComponentPT
}

// Register 注册组件原型，不设置组件名称时，将会使用组件实例名称作为组件名称
func (lib *_ComponentLib) Register(comp any, name ...string) ComponentPT {
	if comp == nil {
		panic(fmt.Errorf("%w: %w: comp is nil", ErrPt, exception.ErrArgs))
	}

	var _name string
	if len(name) > 0 {
		_name = name[0]
	}

	if tfComp, ok := comp.(reflect.Type); ok {
		return lib.register(tfComp, _name)
	} else {
		return lib.register(reflect.TypeOf(comp), _name)
	}
}

// Deregister 取消注册组件原型
func (lib *_ComponentLib) Deregister(impl string) {
	lib.Lock()
	defer lib.Unlock()

	delete(lib.compMap, impl)

	for i, comp := range lib.compList {
		if comp.Implementation == impl {
			lib.compList = append(lib.compList[:i], lib.compList[i+1:]...)
			return
		}
	}
}

// Get 获取组件原型
func (lib *_ComponentLib) Get(impl string) (ComponentPT, bool) {
	lib.RLock()
	defer lib.RUnlock()

	comp, ok := lib.compMap[impl]
	if !ok {
		return ComponentPT{}, false
	}

	return *comp, ok
}

// Range 遍历所有已注册的组件原型
func (lib *_ComponentLib) Range(fun generic.Func1[ComponentPT, bool]) {
	lib.RLock()
	compList := append(make([]*ComponentPT, 0, len(lib.compList)), lib.compList...)
	lib.RUnlock()

	for _, comp := range compList {
		if !fun.Exec(*comp) {
			return
		}
	}
}

func (lib *_ComponentLib) register(tfComp reflect.Type, name string) ComponentPT {
	lib.Lock()
	defer lib.Unlock()

	for tfComp.Kind() == reflect.Pointer || tfComp.Kind() == reflect.Interface {
		tfComp = tfComp.Elem()
	}

	if tfComp.Name() == "" {
		panic(fmt.Errorf("%w: anonymous component not allowed", ErrPt))
	}

	compImpl := types.AnyFullName(tfComp)

	if !reflect.PointerTo(tfComp).Implements(reflect.TypeOf((*ec.Component)(nil)).Elem()) {
		panic(fmt.Errorf("%w: component %q not implement ec.Component", ErrPt, compImpl))
	}

	if name == "" {
		name = compImpl
	}

	comp, ok := lib.compMap[compImpl]
	if ok {
		if comp.Name != name {
			panic(fmt.Errorf("%w: component %q has already been registered with name %q", ErrPt, compImpl, comp.Name))
		}
		return *comp
	}

	comp = &ComponentPT{
		Name:           name,
		Implementation: compImpl,
		tfComp:         tfComp,
	}

	lib.compMap[compImpl] = comp
	lib.compList = append(lib.compList, comp)

	return *comp
}
