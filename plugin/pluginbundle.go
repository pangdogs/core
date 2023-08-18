// Package plugin 插件，用于开发一些需要使用单例模式设计的功能，例如服务发现、消息队列与日志等，服务与运行时上下文均支持安装插件，注意服务类插件需要支持多线程访问，运行时类插件仅支持单线程访问即可。
package plugin

import (
	"fmt"
	"kit.golaxy.org/golaxy/util"
	"sync"
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

	// ReverseRange 反向遍历所有已注册的插件
	//
	//	@param fun 遍历函数。
	ReverseRange(fun func(pluginName string, pluginFace util.FaceAny) bool)
}

// InstallPlugin 安装插件。
//
//	@param pluginBundle 插件包。
//	@param pluginName 插件名称。
//	@param plugin 插件。
func InstallPlugin[T any](pluginBundle PluginBundle, pluginName string, plugin T) {
	if pluginBundle == nil {
		panic("nil pluginBundle")
	}
	pluginBundle.Install(pluginName, util.NewFacePair[any](plugin, plugin))
}

// UninstallPlugin 卸载插件。
//
//	@param pluginBundle 插件包。
func UninstallPlugin(pluginBundle PluginBundle, pluginName string) {
	if pluginBundle == nil {
		panic("nil pluginBundle")
	}
	pluginBundle.Uninstall(pluginName)
}

// NewPluginBundle 创建插件包
func NewPluginBundle() PluginBundle {
	pluginBundle := &_PluginBundle{}
	pluginBundle.init()
	return pluginBundle
}

type _PluginInfo struct {
	Name   string
	Plugin util.FaceAny
}

type _PluginBundle struct {
	pluginMap  map[string]util.FaceAny
	pluginList []_PluginInfo
	mutex      sync.RWMutex
}

func (bundle *_PluginBundle) init() {
	bundle.pluginMap = map[string]util.FaceAny{}
}

// Install 安装插件。
//
//	@param pluginName 插件名称。
//	@param pluginFace 插件Face。
func (bundle *_PluginBundle) Install(pluginName string, pluginFace util.FaceAny) {
	if pluginFace.IsNil() {
		panic("nil pluginFace")
	}

	bundle.mutex.Lock()
	defer bundle.mutex.Unlock()

	_, ok := bundle.pluginMap[pluginName]
	if ok {
		panic(fmt.Errorf("plugin %q is already installed", pluginName))
	}

	bundle.pluginMap[pluginName] = pluginFace
	bundle.pluginList = append(bundle.pluginList, _PluginInfo{Name: pluginName, Plugin: pluginFace})
}

// Uninstall 卸载插件。
//
//	@param pluginName 插件名称。
func (bundle *_PluginBundle) Uninstall(pluginName string) {
	bundle.mutex.Lock()
	defer bundle.mutex.Unlock()

	delete(bundle.pluginMap, pluginName)

	for i, pluginInfo := range bundle.pluginList {
		if pluginInfo.Name == pluginName {
			bundle.pluginList = append(bundle.pluginList[:i], bundle.pluginList[i+1:]...)
			return
		}
	}
}

// Get 获取插件。
//
//	@param pluginName 插件名称。
//	@return 插件Face。
//	@return 是否存在。
func (bundle *_PluginBundle) Get(pluginName string) (util.FaceAny, bool) {
	bundle.mutex.RLock()
	defer bundle.mutex.RUnlock()

	pluginFace, ok := bundle.pluginMap[pluginName]
	return pluginFace, ok
}

// Range 遍历所有已注册的插件
//
//	@param fun 遍历函数。
func (bundle *_PluginBundle) Range(fun func(pluginName string, pluginFace util.FaceAny) bool) {
	if fun == nil {
		return
	}

	bundle.mutex.RLock()
	defer bundle.mutex.RUnlock()

	pluginList := append(make([]_PluginInfo, 0, len(bundle.pluginList)), bundle.pluginList...)

	for i := range pluginList {
		if !fun(pluginList[i].Name, pluginList[i].Plugin) {
			return
		}
	}
}

// ReverseRange 反向遍历所有已注册的插件
//
//	@param fun 遍历函数。
func (bundle *_PluginBundle) ReverseRange(fun func(pluginName string, pluginFace util.FaceAny) bool) {
	if fun == nil {
		return
	}

	bundle.mutex.RLock()
	defer bundle.mutex.RUnlock()

	pluginList := append(make([]_PluginInfo, 0, len(bundle.pluginList)), bundle.pluginList...)

	for i := len(pluginList) - 1; i >= 0; i-- {
		if !fun(pluginList[i].Name, pluginList[i].Plugin) {
			return
		}
	}
}
