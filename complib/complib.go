// Package complib 用于注册开发的组件（Component），提供给实体原型（Entity Prototype）创建实体（Entity）时使用。
package complib

import (
	"fmt"
	"github.com/pangdogs/galaxy/core"
	"reflect"
	"sync"
	"sync/atomic"
)

var componentLib _ComponentLib

func init() {
	componentLib.init()
}

// RegisterPt 注册组件原型（Component Prototype），共有RegisterPt()与RegisterBuilder()两个注册方法，
//二者选其一使用即可。一般在init()函数中使用。线程安全。
//	参数：
//		api：组件实现的api名称，实体将通过api名称来获取组件，多个组件可以实现同一个api。
//		descr：组件功能的描述说明。
//		compPt：组件原型。
func RegisterPt(api, descr string, compPt interface{}) {
	componentLib.RegisterPt(api, descr, compPt)
}

// RegisterBuilder 注册组件构建函数（Component Builder），共有RegisterPt()与RegisterBuilder()两个注册方法，
//二者选其一使用即可。一般在init()函数中使用。线程安全。
//	参数：
//		api：组件实现的api名称，实体将通过api名称来获取组件，多个组件可以实现同一个api。
//		descr：组件功能的描述说明。
//		compBuilder：组件构建函数。
func RegisterBuilder(api, descr string, compBuilder func() core.Component) {
	componentLib.RegisterBuilder(api, descr, compBuilder)
}

// New 创建组件对象，线程安全。
//	参数：
//		tag：组件标签，用于查询组件，格式为组件所在包路径+组件名，例如：`github.com/pangdogs/galaxy/comps/helloworld/HelloWorldComp`。
//	返回值：
//		api：组件实现的实现的api名称。
//		comp：组件对象。
func New(tag string) (api string, comp core.Component) {
	return componentLib.New(tag)
}

// Range 遍历所有已注册的组件信息，线程安全
func Range(fun func(info ComponentInfo) bool) {
	componentLib.Range(fun)
}

// ComponentInfo 组件信息
type ComponentInfo struct {
	Api, Descr, Tag string
}

type _ComponentConstructType int32

const (
	_ComponentConstructType_Pt _ComponentConstructType = iota
	_ComponentConstructType_Builder
)

type _ComponentInfo struct {
	ComponentInfo
	ConstructType _ComponentConstructType
	CompPt        reflect.Type
	CompBuilder   func() core.Component
}

type _ComponentLib struct {
	compInfoMap      map[string]*_ComponentInfo
	registerDisabled int32
	registerMutex    sync.Mutex
}

func (lib *_ComponentLib) init() {
	if lib.compInfoMap == nil {
		lib.compInfoMap = map[string]*_ComponentInfo{}
	}
}

// RegisterPt ...
func (lib *_ComponentLib) RegisterPt(api, descr string, compPt interface{}) {
	if api == "" {
		panic("empty api")
	}

	if compPt == nil {
		panic("nil compPt")
	}

	lib.register(api, descr, _ComponentConstructType_Pt, reflect.TypeOf(compPt), nil)
}

// RegisterBuilder ...
func (lib *_ComponentLib) RegisterBuilder(api, descr string, compBuilder func() core.Component) {
	if api == "" {
		panic("empty api")
	}

	if compBuilder == nil {
		panic("nil compBuilder")
	}

	lib.register(api, descr, _ComponentConstructType_Builder, reflect.TypeOf(compBuilder()), compBuilder)
}

func (lib *_ComponentLib) register(api, descr string, constructType _ComponentConstructType, tfCompPt reflect.Type, compBuilder func() core.Component) {
	lib.registerMutex.Lock()
	defer lib.registerMutex.Unlock()

	if atomic.LoadInt32(&lib.registerDisabled) != 0 {
		panic("register comp disabled")
	}

	for tfCompPt.Kind() == reflect.Pointer || tfCompPt.Kind() == reflect.Interface {
		tfCompPt = tfCompPt.Elem()
	}

	if tfCompPt.Name() == "" {
		panic("register anonymous comp not allowed")
	}

	tag := tfCompPt.PkgPath() + "/" + tfCompPt.Name()

	if !reflect.PointerTo(tfCompPt).Implements(reflect.TypeOf((*core.Component)(nil)).Elem()) {
		panic(fmt.Errorf("comp '%s' not implement core.Component invalid", tag))
	}

	_, ok := lib.compInfoMap[tag]
	if ok {
		panic(fmt.Errorf("repeated register comp '%s' invalid", tag))
	}

	lib.compInfoMap[tag] = &_ComponentInfo{
		ComponentInfo: ComponentInfo{
			Api:   api,
			Descr: descr,
			Tag:   tag,
		},
		ConstructType: constructType,
		CompPt:        tfCompPt,
		CompBuilder:   compBuilder,
	}
}

// New ...
func (lib *_ComponentLib) New(tag string) (api string, comp core.Component) {
	atomic.StoreInt32(&lib.registerDisabled, 1)

	info, ok := lib.compInfoMap[tag]
	if !ok {
		panic(fmt.Errorf("comp '%s' not registered invalid", tag))
	}

	switch info.ConstructType {
	case _ComponentConstructType_Pt:
		return info.Api, reflect.New(info.CompPt).Interface().(core.Component)
	case _ComponentConstructType_Builder:
		return info.Api, info.CompBuilder()
	default:
		panic(fmt.Errorf("comp '%s' construct type not support", tag))
	}
}

// Range ...
func (lib *_ComponentLib) Range(fun func(info ComponentInfo) bool) {
	atomic.StoreInt32(&lib.registerDisabled, 1)

	if fun == nil {
		return
	}

	for _, info := range lib.compInfoMap {
		if !fun(info.ComponentInfo) {
			return
		}
	}
}
