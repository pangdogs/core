package pt

import (
	"fmt"
	"github.com/galaxy-kit/galaxy/ec"
	"reflect"
	"sync"
)

var pluginLib _PluginLib

func init() {
	pluginLib.init()
}

// RegisterPlugin 注册插件，一般在init()函数中使用，线程安全。
//
//	@param pluginName 插件名称，一般是插件实现的接口名称，实体将通过接口名称来获取插件，多个插件可以实现同一个接口。
//	@param descr 插件功能的描述说明。
//	@param plugin 插件对象。
func RegisterPlugin(pluginName, descr string, plugin any) {
	pluginLib.RegisterPlugin(pluginName, descr, plugin)
}

// DeregisterPluginPt 取消注册插件原型，线程安全。
//
//	@param pluginPath 插件路径，格式为插件所在包路径+插件名，例如：`github.com/galaxy-kit/galaxy/ec/plugin/helloworld/HelloWorldComp`。
func DeregisterPluginPt(pluginPath string) {
	pluginLib.DeregisterPluginPt(pluginPath)
}

// GetPluginPt 获取插件原型，线程安全。
//
//	@param pluginPath 插件路径，格式为插件所在包路径+插件名，例如：`github.com/galaxy-kit/galaxy/ec/plugin/helloworld/HelloWorldComp`。
//	@return 插件原型，可以用于创建插件。
//	@return 是否存在。
func GetPluginPt(pluginPath string) (PluginPt, bool) {
	return pluginLib.Get(pluginPath)
}

// RangePluginPts 遍历所有已注册的插件原型，线程安全。
//
//	@param fun 遍历函数。
func RangePluginPts(fun func(pluginPt PluginPt) bool) {
	pluginLib.Range(fun)
}

type _PluginLib struct {
	pluginPtMap map[string]PluginPt
	mutex       sync.RWMutex
}

func (lib *_PluginLib) init() {
	if lib.pluginPtMap == nil {
		lib.pluginPtMap = map[string]PluginPt{}
	}
}

func (lib *_PluginLib) RegisterPlugin(pluginName, descr string, plugin any) {
	if plugin == nil {
		panic("nil plugin")
	}

	if tfComp, ok := plugin.(reflect.Type); ok {
		lib.register(pluginName, descr, _CompConstructType_Reflect, tfComp, nil)
	} else {
		lib.register(pluginName, descr, _CompConstructType_Reflect, reflect.TypeOf(plugin), nil)
	}
}

func (lib *_PluginLib) RegisterCreator(pluginName, descr string, creator func() ec.Plugin) {
	if creator == nil {
		panic("nil creator")
	}

	lib.register(pluginName, descr, _CompConstructType_Creator, nil, creator)
}

func (lib *_PluginLib) DeregisterPluginPt(pluginPath string) {
	lib.mutex.Lock()
	defer lib.mutex.Unlock()

	delete(lib.pluginPtMap, pluginPath)
}

func (lib *_PluginLib) Get(pluginPath string) (PluginPt, bool) {
	lib.mutex.RLock()
	defer lib.mutex.RUnlock()

	pluginPt, ok := lib.pluginPtMap[pluginPath]
	return pluginPt, ok
}

func (lib *_PluginLib) Range(fun func(pluginPt PluginPt) bool) {
	lib.mutex.RLock()
	defer lib.mutex.RUnlock()

	if fun == nil {
		return
	}

	for _, pluginPt := range lib.pluginPtMap {
		if !fun(pluginPt) {
			return
		}
	}
}

func (lib *_PluginLib) register(pluginName, descr string, constructType _CompConstructType, tfComp reflect.Type, creator func() ec.Plugin) {
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
		panic("register anonymous plugin not allowed")
	}

	pluginPath := _tfComp.PkgPath() + "/" + _tfComp.Name()

	if !reflect.PointerTo(_tfComp).Implements(reflect.TypeOf((*ec.Plugin)(nil)).Elem()) {
		panic(fmt.Errorf("plugin '%s' not implement ec.Plugin", pluginPath))
	}

	_, ok := lib.pluginPtMap[pluginPath]
	if ok {
		panic(fmt.Errorf("plugin '%s' is already registered", pluginPath))
	}

	lib.pluginPtMap[pluginPath] = PluginPt{
		Name:          pluginName,
		Path:          pluginPath,
		Description:   descr,
		constructType: constructType,
		tfComp:        tfComp,
		creator:       creator,
	}
}
