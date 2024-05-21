package plugin

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/iface"
	"git.golaxy.org/core/util/types"
	"reflect"
	"slices"
	"sync"
)

// PluginBundle 插件包
type PluginBundle interface {
	iPluginBundle
	PluginProvider

	// Install 安装插件，不设置插件名称时，将会使用插件实例名称作为插件名称
	Install(pluginFace iface.FaceAny, name ...string)
	// Uninstall 卸载插件
	Uninstall(name string)
	// Get 获取插件
	Get(name string) (PluginInfo, bool)
	// Range 遍历所有已注册的插件
	Range(fun generic.Func1[PluginInfo, bool])
	// ReversedRange 反向遍历所有已注册的插件
	ReversedRange(fun generic.Func1[PluginInfo, bool])
}

type iPluginBundle interface {
	setActive(name string, b bool)
	setInstallCB(cb generic.Action1[PluginInfo])
	setUninstallCB(cb generic.Action1[PluginInfo])
}

// NewPluginBundle 创建插件包
func NewPluginBundle() PluginBundle {
	return &_PluginBundle{
		pluginIdx: map[string]*PluginInfo{},
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
	pluginIdx              map[string]*PluginInfo
	pluginList             []*PluginInfo
	installCB, uninstallCB generic.Action1[PluginInfo]
}

// GetPluginBundle 获取插件包
func (bundle *_PluginBundle) GetPluginBundle() PluginBundle {
	return bundle
}

// Install 安装插件，不设置插件名称时，将会使用插件实例名称作为插件名称
func (bundle *_PluginBundle) Install(pluginFace iface.FaceAny, name ...string) {
	bundle.installCB.Exec(bundle.install(pluginFace, name...))
}

// Uninstall 卸载插件
func (bundle *_PluginBundle) Uninstall(name string) {
	pluginInfo, ok := bundle.uninstall(name)
	if !ok {
		return
	}
	bundle.uninstallCB.Exec(pluginInfo)
}

// Get 获取插件
func (bundle *_PluginBundle) Get(name string) (PluginInfo, bool) {
	bundle.RLock()
	defer bundle.RUnlock()

	pluginInfo, ok := bundle.pluginIdx[name]
	if !ok {
		return PluginInfo{}, false
	}

	return *pluginInfo, ok
}

// Range 遍历所有已注册的插件
func (bundle *_PluginBundle) Range(fun generic.Func1[PluginInfo, bool]) {
	bundle.RLock()
	copied := slices.Clone(bundle.pluginList)
	bundle.RUnlock()

	for i := range copied {
		if !fun.Exec(*copied[i]) {
			return
		}
	}
}

// ReversedRange 反向遍历所有已注册的插件
func (bundle *_PluginBundle) ReversedRange(fun generic.Func1[PluginInfo, bool]) {
	bundle.RLock()
	copied := slices.Clone(bundle.pluginList)
	bundle.RUnlock()

	for i := len(copied) - 1; i >= 0; i-- {
		if !fun.Exec(*copied[i]) {
			return
		}
	}
}

func (bundle *_PluginBundle) setActive(name string, b bool) {
	bundle.Lock()
	defer bundle.Unlock()

	pluginInfo, ok := bundle.pluginIdx[name]
	if !ok {
		return
	}

	pluginInfo.Active = b
}

func (bundle *_PluginBundle) setInstallCB(cb generic.Action1[PluginInfo]) {
	bundle.installCB = cb
}

func (bundle *_PluginBundle) setUninstallCB(cb generic.Action1[PluginInfo]) {
	bundle.uninstallCB = cb
}

func (bundle *_PluginBundle) install(pluginFace iface.FaceAny, name ...string) PluginInfo {
	if pluginFace.IsNil() {
		panic(fmt.Errorf("%w: %w: pluginFace is nil", ErrPlugin, exception.ErrArgs))
	}

	bundle.Lock()
	defer bundle.Unlock()

	var _name string
	if len(name) > 0 {
		_name = name[0]
	} else {
		_name = types.FullName(pluginFace.Iface)
	}

	_, ok := bundle.pluginIdx[_name]
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
	bundle.pluginIdx[_name] = pluginInfo

	return *pluginInfo
}

func (bundle *_PluginBundle) uninstall(name string) (PluginInfo, bool) {
	bundle.Lock()
	defer bundle.Unlock()

	pluginInfo, ok := bundle.pluginIdx[name]
	if !ok {
		return PluginInfo{}, false
	}

	delete(bundle.pluginIdx, name)

	bundle.pluginList = slices.DeleteFunc(bundle.pluginList, func(pluginInfo *PluginInfo) bool {
		return pluginInfo.Name == name
	})

	return *pluginInfo, true
}
