// Package plugin 插件，用于开发一些需要使用单例模式设计的功能，例如服务发现、消息队列与日志等，服务与运行时上下文均支持安装插件，注意服务类插件需要支持多线程访问，运行时类插件仅支持单线程访问即可。
package plugin

import (
	"fmt"
	"kit.golaxy.org/golaxy/util"
)

// PluginBundle 插件包
type PluginBundle interface {
	// Install 安装插件。
	//
	//	@param pluginName 插件名称。
	//	@param pluginFace 插件Face。
	Install(pluginName string, pluginFace util.FaceAny)

	// Uninstall 卸载插件。
	//
	//	@param pluginName 插件名称。
	Uninstall(pluginName string)

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

// NewPluginBundle 创建插件包
func NewPluginBundle() PluginBundle {
	pluginBundle := &_PluginBundle{}
	pluginBundle.init()
	return pluginBundle
}

type _PluginBundle struct {
	pluginMap map[string]util.FaceAny
}

func (bundle *_PluginBundle) init() {
	bundle.pluginMap = map[string]util.FaceAny{}
}

func (bundle *_PluginBundle) Install(pluginName string, pluginFace util.FaceAny) {
	if pluginFace.IsNil() {
		panic("nil pluginFace")
	}

	_, ok := bundle.pluginMap[pluginName]
	if ok {
		panic(fmt.Errorf("plugin '%s' is already installed", pluginName))
	}

	bundle.pluginMap[pluginName] = pluginFace
}

func (bundle *_PluginBundle) Uninstall(pluginName string) {
	delete(bundle.pluginMap, pluginName)
}

func (bundle *_PluginBundle) Get(pluginName string) (util.FaceAny, bool) {
	pluginFace, ok := bundle.pluginMap[pluginName]
	return pluginFace, ok
}

func (bundle *_PluginBundle) Range(fun func(pluginName string, pluginFace util.FaceAny) bool) {
	if fun == nil {
		return
	}

	for pluginName, pluginFace := range bundle.pluginMap {
		if !fun(pluginName, pluginFace) {
			return
		}
	}
}
