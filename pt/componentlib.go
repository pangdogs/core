package pt

import (
	"fmt"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/internal/errors"
	"kit.golaxy.org/golaxy/util/types"
	"reflect"
	"sync"
)

var componentLib _ComponentLib

func init() {
	componentLib.init()
}

// RegisterComponent 注册组件原型，一般在init()函数中使用，线程安全。
//
//	@param compName 组件名称，一般是组件实现的接口名称，实体将通过接口名称来获取组件，多个组件可以实现同一个接口。
//	@param COMP 组件对象。
//	@param descr 组件功能的描述说明。
func RegisterComponent(name string, comp any, descr ...string) {
	var _descr string
	if len(descr) > 0 {
		_descr = descr[0]
	}
	componentLib.RegisterComponent(name, comp, _descr)
}

// DeregisterComponent 取消注册组件原型，线程安全。
//
//	@param compImpl 组件实现，格式为组件所在包路径+组件名，例如：`kit.golaxy.org/components/helloworld.HelloWorld`。
func DeregisterComponent(impl string) {
	componentLib.DeregisterComponent(impl)
}

// AccessComponent 访问组件原型，线程安全。
//
//	@param compImpl 组件实现，格式为组件所在包路径+组件名，例如：`kit.golaxy.org/components/helloworld.HelloWorld`。
//	@return 组件原型，可以用于创建组件。
//	@return 是否存在。
func AccessComponent(impl string) (ComponentPt, bool) {
	return componentLib.Get(impl)
}

// RangeComponent 遍历所有已注册的组件原型，线程安全。
//
//	@param fun 遍历函数。
func RangeComponent(fun func(compPt ComponentPt) bool) {
	componentLib.Range(fun)
}

type _ComponentLib struct {
	sync.RWMutex
	compPtMap  map[string]*ComponentPt
	compPtList []*ComponentPt
}

func (lib *_ComponentLib) init() {
	lib.compPtMap = map[string]*ComponentPt{}
}

// RegisterComponent 注册组件原型，一般在init()函数中使用，线程安全。
//
//	@param compName 组件名称，一般是组件实现的接口名称，实体将通过接口名称来获取组件，多个组件可以实现同一个接口。
//	@param COMP 组件对象。
//	@param descr 组件功能的描述说明。
func (lib *_ComponentLib) RegisterComponent(name string, comp any, descr string) {
	if comp == nil {
		panic(fmt.Errorf("%w: %w: comp is nil", ErrPt, errors.ErrArgs))
	}

	if tfComp, ok := comp.(reflect.Type); ok {
		lib.register(name, tfComp, descr)
	} else {
		lib.register(name, reflect.TypeOf(comp), descr)
	}
}

// DeregisterComponent 取消注册组件原型，线程安全。
//
//	@param compImpl 组件实现，格式为组件所在包路径+组件名，例如：`kit.golaxy.org/components/helloworld.HelloWorld`。
func (lib *_ComponentLib) DeregisterComponent(impl string) {
	lib.Lock()
	defer lib.Unlock()

	delete(lib.compPtMap, impl)

	for i, compPt := range lib.compPtList {
		if compPt.Implementation == impl {
			lib.compPtList = append(lib.compPtList[:i], lib.compPtList[i+1:]...)
			return
		}
	}
}

// Get 获取组件原型，线程安全。
//
//	@param compImpl 组件实现，格式为组件所在包路径+组件名，例如：`kit.golaxy.org/components/helloworld.HelloWorld`。
//	@return 组件原型，可以用于创建组件。
//	@return 是否存在。
func (lib *_ComponentLib) Get(impl string) (ComponentPt, bool) {
	lib.RLock()
	defer lib.RUnlock()

	compPt, ok := lib.compPtMap[impl]
	if !ok {
		return ComponentPt{}, false
	}

	return *compPt, ok
}

// Range 遍历所有已注册的组件原型，线程安全。
//
//	@param fun 遍历函数。
func (lib *_ComponentLib) Range(fun func(compPt ComponentPt) bool) {
	if fun == nil {
		return
	}

	lib.RLock()
	compPtList := append(make([]*ComponentPt, 0, len(lib.compPtList)), lib.compPtList...)
	lib.RUnlock()

	for _, compPt := range compPtList {
		if !fun(*compPt) {
			return
		}
	}
}

func (lib *_ComponentLib) register(name string, tfComp reflect.Type, descr string) {
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

	_, ok := lib.compPtMap[compImpl]
	if ok {
		panic(fmt.Errorf("%w: component %q is already registered", ErrPt, compImpl))
	}

	compPt := &ComponentPt{
		Name:           name,
		Implementation: compImpl,
		Description:    descr,
		tfComp:         tfComp,
	}

	lib.compPtMap[compImpl] = compPt
	lib.compPtList = append(lib.compPtList, compPt)
}
