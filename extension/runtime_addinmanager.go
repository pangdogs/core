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
	"reflect"

	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/types"
	"github.com/elliotchance/pie/v2"
)

// RuntimeAddInManager 运行时插件管理器
type RuntimeAddInManager interface {
	AddInManager
	IRuntimeAddInManagerEventTab
}

// NewRuntimeAddInManager 创建运行时插件管理器
func NewRuntimeAddInManager() RuntimeAddInManager {
	return &_RuntimeAddInManager{
		addInNameIndex: map[string]int{},
	}
}

type _RuntimeAddInManager struct {
	addInNameIndex map[string]int
	addInList      generic.FreeList[*_RuntimeAddInStatus]

	runtimeAddInManagerEventTab
}

// GetAddInManager 获取插件管理器
func (mgr *_RuntimeAddInManager) GetAddInManager() AddInManager {
	return mgr
}

// Install 安装插件，不设置插件名称时，将会使用插件实例名称作为插件名称
func (mgr *_RuntimeAddInManager) Install(addInFace iface.FaceAny, name ...string) AddInStatus {
	if addInFace.IsNil() {
		exception.Panicf("%w: %w: addInFace is nil", ErrExtension, exception.ErrArgs)
	}

	addInName := pie.First(name)
	if addInName == "" {
		addInName = types.FullName(addInFace.Iface)
	}

	if _, ok := mgr.addInNameIndex[addInName]; ok {
		exception.Panicf("%w: addIn %q is already installed", ErrExtension, addInName)
	}

	status := &_RuntimeAddInStatus{
		mgr:          mgr,
		name:         addInName,
		instanceFace: addInFace,
		reflected:    reflect.ValueOf(addInFace.Iface),
	}
	slot := mgr.addInList.PushBack(status)
	status.idx = slot.Index()
	status.ver = slot.Version()
	mgr.addInNameIndex[addInName] = slot.Index()

	status.setState(AddInState_Loaded)

	if status.state == AddInState_Loaded {
		_EmitEventRuntimeInstallAddIn(mgr, status)
	}

	return status
}

// Uninstall 卸载插件
func (mgr *_RuntimeAddInManager) Uninstall(name string) {
	statusIdx, ok := mgr.addInNameIndex[name]
	if !ok {
		return
	}
	status := mgr.addInList.Get(statusIdx).V
	status.Uninstall()
}

// Get 获取插件
func (mgr *_RuntimeAddInManager) Get(name string) (AddInStatus, bool) {
	statusIdx, ok := mgr.addInNameIndex[name]
	if !ok {
		return nil, false
	}
	return mgr.addInList.Get(statusIdx).V, true
}

// List 获取所有插件
func (mgr *_RuntimeAddInManager) List() []AddInStatus {
	status := make([]AddInStatus, 0, mgr.addInList.Len())

	mgr.addInList.TraversalEach(func(slot *generic.FreeSlot[*_RuntimeAddInStatus]) {
		status = append(status, slot.V)
	})

	return status
}

func (mgr *_RuntimeAddInManager) uninstallIfVersion(idx int, ver int64) {
	slot := mgr.addInList.Get(idx)
	if slot == nil || slot.Version() != ver {
		return
	}

	status := slot.V

	status.managedUnbindRuntimeHandles()

	if status.state == AddInState_Running {
		_EmitEventRuntimeUninstallAddIn(mgr, status)
	}

	delete(mgr.addInNameIndex, status.name)
	mgr.addInList.ReleaseIfVersion(idx, ver)

	status.setState(AddInState_Unloaded)
}
