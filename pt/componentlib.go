package pt

import (
	"fmt"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/util"
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
func RegisterComponent(compName string, comp any, descr ...string) {
	var _descr string
	if len(descr) > 0 {
		_descr = descr[0]
	}
	componentLib.RegisterComponent(compName, comp, _descr)
}

// DeregisterComponent 取消注册组件原型，线程安全。
//
//	@param compImpl 组件实现，格式为组件所在包路径+组件名，例如：`kit.golaxy.org/components/helloworld/HelloWorld`。
func DeregisterComponent(compImpl string) {
	componentLib.DeregisterComponent(compImpl)
}

// GetComponent 获取组件原型，线程安全。
//
//	@param compImpl 组件实现，格式为组件所在包路径+组件名，例如：`kit.golaxy.org/components/helloworld/HelloWorld`。
//	@return 组件原型，可以用于创建组件。
//	@return 是否存在。
func GetComponent(compImpl string) (ComponentPt, bool) {
	return componentLib.Get(compImpl)
}

// RangeComponent 遍历所有已注册的组件原型，线程安全。
//
//	@param fun 遍历函数。
func RangeComponent(fun func(compPt ComponentPt) bool) {
	componentLib.Range(fun)
}

type _ComponentLib struct {
	compPtMap map[string]ComponentPt
	mutex     sync.RWMutex
}

func (lib *_ComponentLib) init() {
	lib.compPtMap = map[string]ComponentPt{}
}

// RegisterComponent 注册组件原型，一般在init()函数中使用，线程安全。
//
//	@param compName 组件名称，一般是组件实现的接口名称，实体将通过接口名称来获取组件，多个组件可以实现同一个接口。
//	@param COMP 组件对象。
//	@param descr 组件功能的描述说明。
func (lib *_ComponentLib) RegisterComponent(compName string, comp any, descr string) {
	if comp == nil {
		panic("nil comp")
	}

	if tfComp, ok := comp.(reflect.Type); ok {
		lib.register(compName, tfComp, descr)
	} else {
		lib.register(compName, reflect.TypeOf(comp), descr)
	}
}

// DeregisterComponent 取消注册组件原型，线程安全。
//
//	@param compImpl 组件实现，格式为组件所在包路径+组件名，例如：`kit.golaxy.org/components/helloworld/HelloWorld`。
func (lib *_ComponentLib) DeregisterComponent(compImpl string) {
	lib.mutex.Lock()
	defer lib.mutex.Unlock()

	delete(lib.compPtMap, compImpl)
}

// Get 获取组件原型，线程安全。
//
//	@param compImpl 组件实现，格式为组件所在包路径+组件名，例如：`kit.golaxy.org/components/helloworld/HelloWorld`。
//	@return 组件原型，可以用于创建组件。
//	@return 是否存在。
func (lib *_ComponentLib) Get(compImpl string) (ComponentPt, bool) {
	lib.mutex.RLock()
	defer lib.mutex.RUnlock()

	compPt, ok := lib.compPtMap[compImpl]
	return compPt, ok
}

// Range 遍历所有已注册的组件原型，线程安全。
//
//	@param fun 遍历函数。
func (lib *_ComponentLib) Range(fun func(compPt ComponentPt) bool) {
	if fun == nil {
		return
	}

	lib.mutex.RLock()
	defer lib.mutex.RUnlock()

	for _, compPt := range lib.compPtMap {
		if !fun(compPt) {
			return
		}
	}
}

func (lib *_ComponentLib) register(compName string, tfComp reflect.Type, descr string) {
	lib.mutex.Lock()
	defer lib.mutex.Unlock()

	for tfComp.Kind() == reflect.Pointer || tfComp.Kind() == reflect.Interface {
		tfComp = tfComp.Elem()
	}

	if tfComp.Name() == "" {
		panic("register anonymous component not allowed")
	}

	compImpl := util.TypeOfAnyFullName(tfComp)

	if !reflect.PointerTo(tfComp).Implements(reflect.TypeOf((*ec.Component)(nil)).Elem()) {
		panic(fmt.Errorf("component %q not implement ec.Component", compImpl))
	}

	_, ok := lib.compPtMap[compImpl]
	if ok {
		panic(fmt.Errorf("component %q is already registered", compImpl))
	}

	lib.compPtMap[compImpl] = ComponentPt{
		Name:           compName,
		Implementation: compImpl,
		Description:    descr,
		tfComp:         tfComp,
	}
}
