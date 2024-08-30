/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

package plugin

import (
	"fmt"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/types"
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
	Get(name string) (PluginStatus, bool)
	// Range 遍历所有已注册的插件
	Range(fun generic.Func1[PluginStatus, bool])
	// ReversedRange 反向遍历所有已注册的插件
	ReversedRange(fun generic.Func1[PluginStatus, bool])
}

type iPluginBundle interface {
	setPluginState(name string, state PluginState)
	setInstallCB(cb generic.Action1[PluginStatus])
	setUninstallCB(cb generic.Action1[PluginStatus])
}

// NewPluginBundle 创建插件包
func NewPluginBundle() PluginBundle {
	return &_PluginBundle{
		pluginIdx: map[string]*PluginStatus{},
	}
}

// PluginStatus 插件状态信息
type PluginStatus struct {
	Name         string        // 插件名
	InstanceFace iface.FaceAny // 插件实例
	Reflected    reflect.Value // 插件反射值
	State        PluginState   // 状态
}

type _PluginBundle struct {
	sync.RWMutex
	pluginIdx              map[string]*PluginStatus
	pluginList             []*PluginStatus
	installCB, uninstallCB generic.Action1[PluginStatus]
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
	pluginStatus, ok := bundle.uninstall(name)
	if !ok {
		return
	}
	bundle.uninstallCB.Exec(pluginStatus)
}

// Get 获取插件
func (bundle *_PluginBundle) Get(name string) (PluginStatus, bool) {
	bundle.RLock()
	defer bundle.RUnlock()

	pluginStatus, ok := bundle.pluginIdx[name]
	if !ok {
		return PluginStatus{}, false
	}

	return *pluginStatus, ok
}

// Range 遍历所有已注册的插件
func (bundle *_PluginBundle) Range(fun generic.Func1[PluginStatus, bool]) {
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
func (bundle *_PluginBundle) ReversedRange(fun generic.Func1[PluginStatus, bool]) {
	bundle.RLock()
	copied := slices.Clone(bundle.pluginList)
	bundle.RUnlock()

	for i := len(copied) - 1; i >= 0; i-- {
		if !fun.Exec(*copied[i]) {
			return
		}
	}
}

func (bundle *_PluginBundle) setPluginState(name string, state PluginState) {
	bundle.Lock()
	defer bundle.Unlock()

	pluginStatus, ok := bundle.pluginIdx[name]
	if !ok {
		return
	}

	pluginStatus.State = state
}

func (bundle *_PluginBundle) setInstallCB(cb generic.Action1[PluginStatus]) {
	bundle.installCB = cb
}

func (bundle *_PluginBundle) setUninstallCB(cb generic.Action1[PluginStatus]) {
	bundle.uninstallCB = cb
}

func (bundle *_PluginBundle) install(pluginFace iface.FaceAny, name ...string) PluginStatus {
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

	pluginStatus := &PluginStatus{
		Name:         _name,
		InstanceFace: pluginFace,
		Reflected:    reflect.ValueOf(pluginFace.Iface),
		State:        PluginState_Loaded,
	}

	bundle.pluginList = append(bundle.pluginList, pluginStatus)
	bundle.pluginIdx[_name] = pluginStatus

	return *pluginStatus
}

func (bundle *_PluginBundle) uninstall(name string) (PluginStatus, bool) {
	bundle.Lock()
	defer bundle.Unlock()

	pluginStatus, ok := bundle.pluginIdx[name]
	if !ok {
		return PluginStatus{}, false
	}

	delete(bundle.pluginIdx, name)

	bundle.pluginList = slices.DeleteFunc(bundle.pluginList, func(pluginStatus *PluginStatus) bool {
		return pluginStatus.Name == name
	})

	return *pluginStatus, true
}
