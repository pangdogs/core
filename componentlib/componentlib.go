package componentlib

import (
	"fmt"
	"github.com/pangdogs/galaxy/core"
	"reflect"
)

var componentLib _ComponentLib

func RegisterPt(api, descr string, compPt interface{}) {
	componentLib.RegisterPt(api, descr, compPt)
}

func RegisterBuilder(api, descr string, compBuilder func() core.Component) {
	componentLib.RegisterBuilder(api, descr, compBuilder)
}

func New(tag string) (api string, comp core.Component) {
	return componentLib.New(tag)
}

func Range(fun func(info ComponentInfo) bool) {
	componentLib.Range(fun)
}

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
	compInfoMap map[string]*_ComponentInfo
}

func (lib *_ComponentLib) init() {
	if lib.compInfoMap == nil {
		lib.compInfoMap = map[string]*_ComponentInfo{}
	}
}

func (lib *_ComponentLib) RegisterPt(api, descr string, compPt interface{}) {
	lib.init()

	if api == "" {
		panic("empty api")
	}

	if compPt == nil {
		panic("nil compPt")
	}

	lib.register(api, descr, _ComponentConstructType_Pt, reflect.TypeOf(compPt), nil)
}

func (lib *_ComponentLib) RegisterBuilder(api, descr string, compBuilder func() core.Component) {
	lib.init()

	if api == "" {
		panic("empty api")
	}

	if compBuilder == nil {
		panic("nil compBuilder")
	}

	lib.register(api, descr, _ComponentConstructType_Builder, reflect.TypeOf(compBuilder()), compBuilder)
}

func (lib *_ComponentLib) register(api, descr string, constructType _ComponentConstructType, tfCompPt reflect.Type, compBuilder func() core.Component) {
	for tfCompPt.Kind() == reflect.Pointer {
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

func (lib *_ComponentLib) New(tag string) (api string, comp core.Component) {
	lib.init()

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

func (lib *_ComponentLib) Range(fun func(info ComponentInfo) bool) {
	lib.init()

	if fun == nil {
		return
	}

	for _, info := range lib.compInfoMap {
		if !fun(info.ComponentInfo) {
			return
		}
	}
}
