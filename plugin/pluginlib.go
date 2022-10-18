// Package plugin 插件。
package plugin

import (
	"fmt"
	"github.com/pangdogs/galaxy/util"
)

// PluginLib 插件库
type PluginLib interface {
	// Register 注册插件。
	//
	//	@param pluginName 插件名称。
	//	@param pluginFace 插件Face。
	Register(pluginName string, pluginFace util.FaceAny)

	// Unregister 取消注册插件。
	//
	//	@param pluginName 插件名称。
	Unregister(pluginName string)

	// Get 获取插件。
	//
	//	@param pluginName 插件名称。
	//	@return 插件Face。
	//	@return 是否存在。
	Get(pluginName string) (util.FaceAny, bool)

	// Range 遍历所有已注册的插件
	//
	//	@param fun 遍历函数。
	Range(fun func(pluginName string, pluginFace util.FaceAny) bool)
}

// RegisterPlugin 注册插件。
//
//	@param pluginLib 插件库。
//	@param pluginName 插件名称。
//	@param plugin 插件。
func RegisterPlugin[T any](pluginLib PluginLib, pluginName string, plugin T) {
	if pluginLib == nil {
		panic("nil pluginLib")
	}
	pluginLib.Register(pluginName, util.NewFacePair[any](plugin, plugin))
}

// GetPlugin 获取插件。
//
//	@param pluginLib 插件库。
//	@param pluginName 插件名称。
//	@return 插件。
//	@return 是否存在。
func GetPlugin[T any](pluginLib PluginLib, pluginName string) (T, bool) {
	if pluginLib == nil {
		panic("nil pluginLib")
	}

	pluginFace, ok := pluginLib.Get(pluginName)
	if !ok {
		return util.Zero[T](), false
	}

	return util.Cache2Iface[T](pluginFace.Cache), true
}

// NewPluginLib 创建插件库
func NewPluginLib() PluginLib {
	lib := &_PluginLib{}
	lib.init()
	return lib
}

type _PluginLib struct {
	pluginMap map[string]util.FaceAny
}

func (lib *_PluginLib) init() {
	lib.pluginMap = map[string]util.FaceAny{}
}

func (lib *_PluginLib) Register(pluginName string, pluginFace util.FaceAny) {
	if pluginFace.IsNil() {
		panic("nil pluginFace")
	}

	_, ok := lib.pluginMap[pluginName]
	if ok {
		panic(fmt.Errorf("plugin '%s' is already registered", pluginName))
	}

	lib.pluginMap[pluginName] = pluginFace
}

func (lib *_PluginLib) Unregister(pluginName string) {
	delete(lib.pluginMap, pluginName)
}

func (lib *_PluginLib) Get(pluginName string) (util.FaceAny, bool) {
	pluginFace, ok := lib.pluginMap[pluginName]
	return pluginFace, ok
}

func (lib *_PluginLib) Range(fun func(pluginName string, pluginFace util.FaceAny) bool) {
	if fun == nil {
		return
	}

	for pluginName, pluginFace := range lib.pluginMap {
		if !fun(pluginName, pluginFace) {
			return
		}
	}
}
