package plugin

import (
	"fmt"
	"github.com/pangdogs/galaxy/util"
	"sync"
)

// PluginLib 组件库
type PluginLib interface {
	// Register 注册插件，线程安全
	Register(pluginName string, pluginFace util.FaceAny)

	// Unregister 取消注册插件，线程安全
	Unregister(pluginName string)

	// Get 获取插件，线程安全
	Get(pluginName string) util.FaceAny

	// Range 遍历所有已注册的插件，线程安全
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
	mutex     sync.RWMutex
}

func (lib *_PluginLib) init() {
	lib.pluginMap = map[string]util.FaceAny{}
}

func (lib *_PluginLib) Register(pluginName string, pluginFace util.FaceAny) {
	lib.mutex.Lock()
	defer lib.mutex.Unlock()

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
	lib.mutex.Lock()
	defer lib.mutex.Unlock()

	delete(lib.pluginMap, pluginName)
}

func (lib *_PluginLib) Get(pluginName string) util.FaceAny {
	lib.mutex.RLock()
	defer lib.mutex.RUnlock()

	pluginFace, ok := lib.pluginMap[pluginName]
	if !ok {
		panic(fmt.Errorf("plugin '%s' not registered invalid", pluginName))
	}

	return pluginFace
}

func (lib *_PluginLib) Range(fun func(pluginName string, pluginFace util.FaceAny) bool) {
	lib.mutex.RLock()
	defer lib.mutex.RUnlock()

	if fun == nil {
		return
	}

	for pluginName, pluginFace := range lib.pluginMap {
		if !fun(pluginName, pluginFace) {
			return
		}
	}
}
