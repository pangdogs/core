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

package extension

import (
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/types"
	"github.com/elliotchance/pie/v2"
	"reflect"
	"slices"
	"sync"
)

// AddInManager 插件管理器
type AddInManager interface {
	iAddInManager
	AddInProvider

	// Install 安装插件，不设置插件名称时，将会使用插件实例名称作为插件名称
	Install(addInFace iface.FaceAny, name ...string)
	// Uninstall 卸载插件
	Uninstall(name string)
	// Get 获取插件
	Get(name string) (AddInStatus, bool)
	// Range 遍历所有已注册的插件
	Range(fun generic.Func1[AddInStatus, bool])
	// ReversedRange 反向遍历所有已注册的插件
	ReversedRange(fun generic.Func1[AddInStatus, bool])
}

type iAddInManager interface {
	setCallback(installCB, uninstallCB generic.Action1[AddInStatus])
}

// NewAddInManager 创建插件管理器
func NewAddInManager() AddInManager {
	return &_AddInManager{
		addInIdx: map[string]*_AddInStatus{},
	}
}

type _AddInManager struct {
	sync.RWMutex
	addInIdx               map[string]*_AddInStatus
	addInList              []*_AddInStatus
	installCB, uninstallCB generic.Action1[AddInStatus]
}

// GetAddInManager 获取插件管理器
func (mgr *_AddInManager) GetAddInManager() AddInManager {
	return mgr
}

// Install 安装插件，不设置插件名称时，将会使用插件实例名称作为插件名称
func (mgr *_AddInManager) Install(addInFace iface.FaceAny, name ...string) {
	mgr.installCB.Exec(mgr.install(addInFace, name...))
}

// Uninstall 卸载插件
func (mgr *_AddInManager) Uninstall(name string) {
	status, ok := mgr.uninstall(name)
	if !ok {
		return
	}
	mgr.uninstallCB.Exec(status)
}

// Get 获取插件
func (mgr *_AddInManager) Get(name string) (AddInStatus, bool) {
	mgr.RLock()
	defer mgr.RUnlock()

	status, ok := mgr.addInIdx[name]
	if !ok {
		return nil, false
	}

	return status, ok
}

// Range 遍历所有已注册的插件
func (mgr *_AddInManager) Range(fun generic.Func1[AddInStatus, bool]) {
	mgr.RLock()
	copied := slices.Clone(mgr.addInList)
	mgr.RUnlock()

	for i := range copied {
		if !fun.Exec(copied[i]) {
			return
		}
	}
}

// ReversedRange 反向遍历所有已注册的插件
func (mgr *_AddInManager) ReversedRange(fun generic.Func1[AddInStatus, bool]) {
	mgr.RLock()
	copied := slices.Clone(mgr.addInList)
	mgr.RUnlock()

	for i := len(copied) - 1; i >= 0; i-- {
		if !fun.Exec(copied[i]) {
			return
		}
	}
}

func (mgr *_AddInManager) setCallback(installCB, uninstallCB generic.Action1[AddInStatus]) {
	mgr.Lock()
	defer mgr.Unlock()

	mgr.installCB = installCB
	mgr.uninstallCB = uninstallCB
}

func (mgr *_AddInManager) install(addInFace iface.FaceAny, name ...string) *_AddInStatus {
	if addInFace.IsNil() {
		exception.Panicf("%w: %w: addInFace is nil", ErrExtension, exception.ErrArgs)
	}

	mgr.Lock()
	defer mgr.Unlock()

	addInName := pie.First(name)
	if addInName == "" {
		addInName = types.FullName(addInFace.Iface)
	}

	if _, ok := mgr.addInIdx[addInName]; ok {
		exception.Panicf("%w: addIn %q is already installed", ErrExtension, addInName)
	}

	status := &_AddInStatus{
		name:         addInName,
		instanceFace: addInFace,
		reflected:    reflect.ValueOf(addInFace.Iface),
	}
	status.state.Store(int32(AddInState_Loaded))

	mgr.addInList = append(mgr.addInList, status)
	mgr.addInIdx[addInName] = status

	return status
}

func (mgr *_AddInManager) uninstall(name string) (*_AddInStatus, bool) {
	mgr.Lock()
	defer mgr.Unlock()

	status, ok := mgr.addInIdx[name]
	if !ok {
		return nil, false
	}

	delete(mgr.addInIdx, name)

	mgr.addInList = slices.DeleteFunc(mgr.addInList, func(status *_AddInStatus) bool {
		return status.name == name
	})

	return status, true
}
