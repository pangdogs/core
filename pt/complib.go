package pt

import (
	"fmt"
	"github.com/pangdogs/galaxy/ec"
	"reflect"
	"sync"
)

var componentLib _ComponentLib

func init() {
	componentLib.init()
}

// RegisterComponent 注册组件原型，共有RegisterComp()与RegisterCreator()两个注册方法，
// 二者选其一使用即可。一般在init()函数中使用，线程安全。
//
//	@param name 组件实现的接口名称，实体将通过接口名称来获取组件，多个组件可以实现同一个接口。
//	@param descr 组件功能的描述说明。
//	@param comp 组件对象。
func RegisterComponent(name, descr string, comp interface{}) {
	componentLib.RegisterComponent(name, descr, comp)
}

// RegisterComponentCreator 注册组件构建函数，共有RegisterComp()与RegisterCreator()两个注册方法，
// 二者选其一使用即可。一般在init()函数中使用，线程安全。
//
//	@param name 组件实现的接口名称，实体将通过接口名称来获取组件，多个组件可以实现同一个接口。
//	@param descr 组件功能的描述说明。
//	@param creator 组件构建函数。
func RegisterComponentCreator(name, descr string, creator func() ec.Component) {
	componentLib.RegisterCreator(name, descr, creator)
}

// UnregisterComponentPt 取消注册组件原型，线程安全。
//
//	@param tag 组件标签，格式为组件所在包路径+组件名，例如：`github.com/pangdogs/galaxy/ec/comps/helloworld/HelloWorldComp`。
func UnregisterComponentPt(tag string) {
	componentLib.UnregisterComponentPt(tag)
}

// GetComponentPt 获取组件原型，线程安全。
//
//	@param tag 组件标签，格式为组件所在包路径+组件名，例如：`github.com/pangdogs/galaxy/ec/comps/helloworld/HelloWorldComp`。
//	@return 组件原型，可以用于创建组件。
//	@return 是否存在。
func GetComponentPt(tag string) (ComponentPt, bool) {
	return componentLib.Get(tag)
}

// RangeComponentPts 遍历所有已注册的组件原型，线程安全。
//
//	@param fun 遍历函数。
func RangeComponentPts(fun func(compPt ComponentPt) bool) {
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

func (lib *_ComponentLib) RegisterComponent(name, descr string, comp interface{}) {
	if name == "" {
		panic("empty name")
	}

	if comp == nil {
		panic("nil comp")
	}

	lib.register(name, descr, _CompConstructType_Reflect, reflect.TypeOf(comp), nil)
}

func (lib *_ComponentLib) RegisterCreator(name, descr string, creator func() ec.Component) {
	if name == "" {
		panic("empty name")
	}

	if creator == nil {
		panic("nil creator")
	}

	lib.register(name, descr, _CompConstructType_Creator, nil, creator)
}

func (lib *_ComponentLib) UnregisterComponentPt(tag string) {
	lib.mutex.Lock()
	defer lib.mutex.Unlock()

	delete(lib.compPtMap, tag)
}

func (lib *_ComponentLib) Get(tag string) (ComponentPt, bool) {
	lib.mutex.RLock()
	defer lib.mutex.RUnlock()

	compPt, ok := lib.compPtMap[tag]
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

func (lib *_ComponentLib) register(compName, descr string, constructType _CompConstructType, tfComp reflect.Type, creator func() ec.Component) {
	lib.mutex.Lock()
	defer lib.mutex.Unlock()

	var _tfComp reflect.Type

	switch constructType {
	case _CompConstructType_Reflect:
		_tfComp = tfComp
	case _CompConstructType_Creator:
		_tfComp = reflect.TypeOf(creator())
	default:
		panic("not support construct type")
	}

	for _tfComp.Kind() == reflect.Pointer || _tfComp.Kind() == reflect.Interface {
		_tfComp = _tfComp.Elem()
	}

	if _tfComp.Name() == "" {
		panic("register anonymous component not allowed")
	}

	tag := _tfComp.PkgPath() + "/" + _tfComp.Name()

	if !reflect.PointerTo(_tfComp).Implements(reflect.TypeOf((*ec.Component)(nil)).Elem()) {
		panic(fmt.Errorf("component '%s' not implement ec.Component", tag))
	}

	_, ok := lib.compPtMap[tag]
	if ok {
		panic(fmt.Errorf("component '%s' is already registered", tag))
	}

	lib.compPtMap[tag] = ComponentPt{
		Interface:     compName,
		Tag:           tag,
		Description:   descr,
		constructType: constructType,
		tfComp:        tfComp,
		creator:       creator,
	}
}
