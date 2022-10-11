package plugin

import (
	"fmt"
	"github.com/pangdogs/galaxy/util"
)

// PluginLib 组件库
type PluginLib interface {
	// Register 注册插件
	Register(pluginName string, pluginFace util.FaceAny)

	// Unregister 取消注册插件
	Unregister(pluginName string)

	// Get 获取插件
	Get(pluginName string) util.FaceAny

	// Range 遍历所有已注册的插件
	Range(fun func(pluginName string, pluginFace util.FaceAny) bool)
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
		panic(fmt.Errorf("repeated register plugin '%s' invalid", pluginName))
	}

	lib.pluginMap[pluginName] = pluginFace
}

func (lib *_PluginLib) Unregister(pluginName string) {
	delete(lib.pluginMap, pluginName)
}

func (lib *_PluginLib) Get(pluginName string) util.FaceAny {
	pluginFace, ok := lib.pluginMap[pluginName]
	if !ok {
		panic(fmt.Errorf("plugin '%s' not registered invalid", pluginName))
	}

	return pluginFace
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
