package pt

import (
	"fmt"
	"github.com/galaxy-kit/galaxy-go/ec"
	"github.com/galaxy-kit/galaxy-go/util"
	"reflect"
	"sync"
)

var componentLib _ComponentLib

func init() {
	componentLib.init()
}

// RegisterComponent 注册组件原型，一般在init()函数中使用，线程安全。
//
//	@type_param COMP 组件类型。
//	@param compName 组件名称，一般是组件实现的接口名称，实体将通过接口名称来获取组件，多个组件可以实现同一个接口。
//	@param descr 组件功能的描述说明。
func RegisterComponent[COMP any](compName string, descr ...string) {
	var _descr string
	if len(descr) > 0 {
		_descr = descr[0]
	}
	componentLib.RegisterComponent(compName, util.Zero[COMP](), _descr)
}

// DeregisterComponent 取消注册组件原型，线程安全。
//
//	@param compPath 组件路径，格式为组件所在包路径+组件名，例如：`github.com/galaxy-kit/components-go/helloworld/HelloWorldComp`。
func DeregisterComponent(compPath string) {
	componentLib.DeregisterComponent(compPath)
}

// GetComponent 获取组件原型，线程安全。
//
//	@param compPath 组件路径，格式为组件所在包路径+组件名，例如：`github.com/galaxy-kit/components-go/helloworld/HelloWorldComp`。
//	@return 组件原型，可以用于创建组件。
//	@return 是否存在。
func GetComponent(compPath string) (ComponentPt, bool) {
	return componentLib.Get(compPath)
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
	if lib.compPtMap == nil {
		lib.compPtMap = map[string]ComponentPt{}
	}
}

func (lib *_ComponentLib) RegisterComponent(compName string, comp any, descr string) {
	if comp == nil {
		panic("nil comp")
	}

	lib.register(compName, reflect.TypeOf(comp), descr)
}

func (lib *_ComponentLib) DeregisterComponent(compPath string) {
	lib.mutex.Lock()
	defer lib.mutex.Unlock()

	delete(lib.compPtMap, compPath)
}

func (lib *_ComponentLib) Get(compPath string) (ComponentPt, bool) {
	lib.mutex.RLock()
	defer lib.mutex.RUnlock()

	compPt, ok := lib.compPtMap[compPath]
	return compPt, ok
}

func (lib *_ComponentLib) Range(fun func(compPt ComponentPt) bool) {
	lib.mutex.RLock()
	defer lib.mutex.RUnlock()

	if fun == nil {
		return
	}

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

	compPath := tfComp.PkgPath() + "/" + tfComp.Name()

	if !reflect.PointerTo(tfComp).Implements(reflect.TypeOf((*ec.Component)(nil)).Elem()) {
		panic(fmt.Errorf("component '%s' not implement ec.Component", compPath))
	}

	_, ok := lib.compPtMap[compPath]
	if ok {
		panic(fmt.Errorf("component '%s' is already registered", compPath))
	}

	lib.compPtMap[compPath] = ComponentPt{
		Name:        compName,
		Path:        compPath,
		Description: descr,
		tfComp:      tfComp,
	}
}
