// Package plugin 插件，用于开发一些需要使用单例模式设计的功能，例如服务发现、消息队列与日志等，服务与运行时上下文均支持安装插件，注意服务类插件需要支持多线程访问，运行时类插件仅支持单线程访问即可。
package plugin

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/util/iface"
	"sync"
)

// PluginInfo 插件信息
type PluginInfo struct {
	Name   string        // 插件名
	Face   iface.FaceAny // 插件Face
	Active bool          // 是否激活
}

// PluginBundle 插件包
type PluginBundle interface {
	// Install 安装插件。
	//
	//	@param name 插件名称。
	//	@param plugin 插件Face。
	Install(name string, pluginFace iface.FaceAny)

	// Uninstall 卸载插件。
	//
	//	@param name 插件名称。
	Uninstall(name string)

	// Get 获取插件。
	//
	//	@param name 插件名称。
	//	@return 插件信息。
	//	@return 是否存在。
	Get(name string) (PluginInfo, bool)

	// Range 遍历所有已注册的插件
	//
	//	@param fun 遍历函数。
	Range(fun func(info PluginInfo) bool)

	// ReverseRange 反向遍历所有已注册的插件
	//
	//	@param fun 遍历函数。
	ReverseRange(fun func(info PluginInfo) bool)

	activate(name string, b bool)
}

// Install 安装插件。
//
//	@param pluginBundle 插件包。
//	@param name 插件名称。
//	@param plugin 插件。
func Install[T any](pluginBundle PluginBundle, name string, plugin T) {
	if pluginBundle == nil {
		panic(fmt.Errorf("%w: %w: pluginBundle is nil", ErrPlugin, internal.ErrArgs))
	}
	pluginBundle.Install(name, iface.NewFacePair[any](plugin, plugin))
}

// Uninstall 卸载插件。
//
//	@param pluginBundle 插件包。
func Uninstall(pluginBundle PluginBundle, name string) {
	if pluginBundle == nil {
		panic(fmt.Errorf("%w: %w: pluginBundle is nil", ErrPlugin, internal.ErrArgs))
	}
	pluginBundle.Uninstall(name)
}

// NewPluginBundle 创建插件包
func NewPluginBundle() PluginBundle {
	pluginBundle := &_PluginBundle{}
	pluginBundle.init()
	return pluginBundle
}

type _PluginBundle struct {
	sync.RWMutex
	pluginMap  map[string]*PluginInfo
	pluginList []*PluginInfo
}

func (bundle *_PluginBundle) init() {
	bundle.pluginMap = map[string]*PluginInfo{}
}

// Install 安装插件。
//
//	@param name 插件名称。
//	@param plugin 插件Face。
func (bundle *_PluginBundle) Install(name string, pluginFace iface.FaceAny) {
	if pluginFace.IsNil() {
		panic(fmt.Errorf("%w: %w: pluginFace is nil", ErrPlugin, internal.ErrArgs))
	}

	bundle.Lock()
	defer bundle.Unlock()

	_, ok := bundle.pluginMap[name]
	if ok {
		panic(fmt.Errorf("%w: %q is already installed", ErrPlugin, name))
	}

	pluginInfo := &PluginInfo{
		Name:   name,
		Face:   pluginFace,
		Active: false,
	}

	bundle.pluginList = append(bundle.pluginList, pluginInfo)
	bundle.pluginMap[name] = pluginInfo
}

// Uninstall 卸载插件。
//
//	@param name 插件名称。
func (bundle *_PluginBundle) Uninstall(name string) {
	bundle.Lock()
	defer bundle.Unlock()

	delete(bundle.pluginMap, name)

	for i, pluginInfo := range bundle.pluginList {
		if pluginInfo.Name == name {
			bundle.pluginList = append(bundle.pluginList[:i], bundle.pluginList[i+1:]...)
			return
		}
	}
}

// Get 获取插件。
//
//	@param name 插件名称。
//	@return 插件Face。
//	@return 是否存在。
func (bundle *_PluginBundle) Get(name string) (PluginInfo, bool) {
	bundle.RLock()
	defer bundle.RUnlock()

	pluginInfo, ok := bundle.pluginMap[name]
	if !ok {
		return PluginInfo{}, false
	}

	return *pluginInfo, ok
}

// Range 遍历所有已注册的插件
//
//	@param fun 遍历函数。
func (bundle *_PluginBundle) Range(fun func(info PluginInfo) bool) {
	if fun == nil {
		return
	}

	bundle.RLock()
	pluginList := append(make([]*PluginInfo, 0, len(bundle.pluginList)), bundle.pluginList...)
	bundle.RUnlock()

	for i := range pluginList {
		if !fun(*pluginList[i]) {
			return
		}
	}
}

// ReverseRange 反向遍历所有已注册的插件
//
//	@param fun 遍历函数。
func (bundle *_PluginBundle) ReverseRange(fun func(info PluginInfo) bool) {
	if fun == nil {
		return
	}

	bundle.RLock()
	pluginList := append(make([]*PluginInfo, 0, len(bundle.pluginList)), bundle.pluginList...)
	bundle.RUnlock()

	for i := len(pluginList) - 1; i >= 0; i-- {
		if !fun(*pluginList[i]) {
			return
		}
	}
}

func (bundle *_PluginBundle) activate(name string, b bool) {
	bundle.Lock()
	defer bundle.Unlock()

	pluginInfo, ok := bundle.pluginMap[name]
	if !ok {
		return
	}

	pluginInfo.Active = b
}
