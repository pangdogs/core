package plugin

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/exception"
	"kit.golaxy.org/golaxy/util/generic"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/types"
	"reflect"
	"sync"
)

// PluginBundle 插件包
type PluginBundle interface {
	PluginProvider
	// Install 安装插件，不设置插件名称时，将会使用插件实例名称作为插件名称
	Install(pluginFace iface.FaceAny, name ...string) PluginInfo
	// Uninstall 卸载插件
	Uninstall(name string)
	// Get 获取插件
	Get(name string) (PluginInfo, bool)
	// Range 遍历所有已注册的插件
	Range(fun generic.Func1[PluginInfo, bool])
	// ReverseRange 反向遍历所有已注册的插件
	ReverseRange(fun generic.Func1[PluginInfo, bool])

	activate(name string, b bool)
}

// NewPluginBundle 创建插件包
func NewPluginBundle() PluginBundle {
	return &_PluginBundle{
		pluginMap: map[string]*PluginInfo{},
	}
}

// PluginInfo 插件信息
type PluginInfo struct {
	Name      string        // 插件名
	Face      iface.FaceAny // 插件Face
	Reflected reflect.Value // 插件反射值
	Active    bool          // 是否激活
}

type _PluginBundle struct {
	sync.RWMutex
	pluginMap  map[string]*PluginInfo
	pluginList []*PluginInfo
}

// GetPluginBundle 获取插件包
func (bundle *_PluginBundle) GetPluginBundle() PluginBundle {
	return bundle
}

// Install 安装插件，不设置插件名称时，将会使用插件实例名称作为插件名称
func (bundle *_PluginBundle) Install(pluginFace iface.FaceAny, name ...string) PluginInfo {
	if pluginFace.IsNil() {
		panic(fmt.Errorf("%w: %w: pluginFace is nil", ErrPlugin, exception.ErrArgs))
	}

	bundle.Lock()
	defer bundle.Unlock()

	var _name string
	if len(name) > 0 {
		_name = name[0]
	} else {
		_name = types.AnyFullName(pluginFace.Iface)
	}

	_, ok := bundle.pluginMap[_name]
	if ok {
		panic(fmt.Errorf("%w: plugin %q is already installed", ErrPlugin, name))
	}

	pluginInfo := &PluginInfo{
		Name:      _name,
		Face:      pluginFace,
		Reflected: reflect.ValueOf(pluginFace.Iface),
		Active:    false,
	}

	bundle.pluginList = append(bundle.pluginList, pluginInfo)
	bundle.pluginMap[_name] = pluginInfo

	return *pluginInfo
}

// Uninstall 卸载插件
func (bundle *_PluginBundle) Uninstall(name string) {
	bundle.Lock()
	defer bundle.Unlock()

	pluginInfo, ok := bundle.pluginMap[name]
	if !ok {
		return
	}

	if pluginInfo.Active {
		panic(fmt.Errorf("%w: plugin %q is active, can't uninstall", ErrPlugin, name))
	}

	delete(bundle.pluginMap, name)

	for i, pi := range bundle.pluginList {
		if pi.Name == name {
			bundle.pluginList = append(bundle.pluginList[:i], bundle.pluginList[i+1:]...)
			return
		}
	}
}

// Get 获取插件
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
func (bundle *_PluginBundle) Range(fun generic.Func1[PluginInfo, bool]) {
	bundle.RLock()
	pluginList := append(make([]*PluginInfo, 0, len(bundle.pluginList)), bundle.pluginList...)
	bundle.RUnlock()

	for i := range pluginList {
		if !fun.Exec(*pluginList[i]) {
			return
		}
	}
}

// ReverseRange 反向遍历所有已注册的插件
func (bundle *_PluginBundle) ReverseRange(fun generic.Func1[PluginInfo, bool]) {
	bundle.RLock()
	pluginList := append(make([]*PluginInfo, 0, len(bundle.pluginList)), bundle.pluginList...)
	bundle.RUnlock()

	for i := len(pluginList) - 1; i >= 0; i-- {
		if !fun.Exec(*pluginList[i]) {
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
