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
	"github.com/elliotchance/pie/v2"
)

// RuntimeAddInManager 运行时插件管理器
type RuntimeAddInManager interface {
	iRuntimeAddInManager
	AddInManager

	IRuntimeAddInManagerEventTab
}

type iRuntimeAddInManager interface {
	getList() []RuntimeAddInStatus
}

// NewRuntimeAddInManager 创建运行时插件管理器
func NewRuntimeAddInManager() RuntimeAddInManager {
	return &_RuntimeAddInManager{
		addInNameIndex: map[string]int{},
		addInIdIndex:   map[uint64]int{},
	}
}

type _RuntimeAddInManager struct {
	addInNameIndex map[string]int
	addInIdIndex   map[uint64]int
	addInList      generic.FreeList[*_RuntimeAddInStatus]

	runtimeAddInManagerEventTab
}

// AddInManager 获取插件管理器
func (mgr *_RuntimeAddInManager) AddInManager() AddInManager {
	return mgr
}

// Install 安装插件
func (mgr *_RuntimeAddInManager) Install(addInFace iface.FaceAny, name ...string) AddInStatus {
	if addInFace.IsNil() {
		exception.Panicf("%w: %w: addInFace is nil", ErrExtension, exception.ErrArgs)
	}

	addInName := pie.First(name)
	if addInName == "" {
		addInName = GenAddInName(addInFace.Iface)
	}

	if addInName == "" {
		exception.Panicf("%w: anonymous add-in not allowed", ErrExtension)
	}

	if _, ok := mgr.addInNameIndex[addInName]; ok {
		exception.Panicf("%w: add-in %q is already installed", ErrExtension, addInName)
	}

	id := GenAddInId(addInName)

	if existsIdx, ok := mgr.addInIdIndex[id]; ok {
		exception.Panicf("%w: add-in %q id %d conflict with %q, rename required", ErrExtension, addInName, id, mgr.addInList.Get(existsIdx).V.Name())
	}

	status := &_RuntimeAddInStatus{
		mgr:          mgr,
		id:           id,
		name:         addInName,
		instanceFace: addInFace,
		reflected:    reflect.ValueOf(addInFace.Iface),
	}
	slot := mgr.addInList.PushBack(status)
	status.idx = slot.Index()
	status.ver = slot.Version()
	mgr.addInNameIndex[addInName] = slot.Index()
	mgr.addInIdIndex[id] = slot.Index()

	_EmitEventRuntimeAddInStateChanged(mgr, status, AddInState_Loaded)

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
	status.uninstall()
}

// GetByName 使用名称查询插件状态信息
func (mgr *_RuntimeAddInManager) GetByName(name string) (AddInStatus, bool) {
	statusIdx, ok := mgr.addInNameIndex[name]
	if !ok {
		return nil, false
	}
	return mgr.addInList.Get(statusIdx).V, true
}

// GetById 使用Id查询插件状态信息
func (mgr *_RuntimeAddInManager) GetById(id uint64) (AddInStatus, bool) {
	statusIdx, ok := mgr.addInIdIndex[id]
	if !ok {
		return nil, false
	}
	return mgr.addInList.Get(statusIdx).V, true
}

// List 获取所有插件状态信息
func (mgr *_RuntimeAddInManager) List() []AddInStatus {
	statuses := make([]AddInStatus, 0, mgr.addInList.Len())

	mgr.addInList.TraversalEach(func(slot *generic.FreeSlot[*_RuntimeAddInStatus]) {
		statuses = append(statuses, slot.V)
	})

	return statuses
}

func (mgr *_RuntimeAddInManager) getList() []RuntimeAddInStatus {
	statuses := make([]RuntimeAddInStatus, 0, mgr.addInList.Len())

	mgr.addInList.TraversalEach(func(slot *generic.FreeSlot[*_RuntimeAddInStatus]) {
		statuses = append(statuses, slot.V)
	})

	return statuses
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
	delete(mgr.addInIdIndex, status.id)
	mgr.addInList.ReleaseIfVersion(idx, ver)

	status.setState(AddInState_Unloaded)
}
