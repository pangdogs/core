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

package runtime

import (
	"reflect"

	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"github.com/elliotchance/pie/v2"
)

// AddInManager 管理可在运行时热插拔的插件及其生命周期事件。
//
// 管理器按安装顺序保存插件且不提供并发保护；所有操作应串行执行，通常由所属运行时
// goroutine 调用。运行时启动前安装的插件保持 Loaded，启动后安装的插件会由 core
// 同步激活；卸载和停服时按相反顺序停用。
type AddInManager interface {
	iAddInManager
	extension.AddInManager

	IAddInManagerEventTab
}

type iAddInManager interface {
	getListStatuses() []AddInStatus
}

// NewAddInManager 创建一个空的运行时插件管理器。
func NewAddInManager() AddInManager {
	return &_AddInManager{
		addInNameIndex: map[string]int{},
		addInIdIndex:   map[uint64]int{},
	}
}

type _AddInManager struct {
	addInNameIndex map[string]int
	addInIdIndex   map[uint64]int
	addInList      generic.FreeList[*_AddInStatus]

	addInManagerEventTab
}

// AddInManager 返回当前运行时插件管理器的公共接口。
func (mgr *_AddInManager) AddInManager() extension.AddInManager {
	return mgr
}

// Install 安装插件并同步派发状态变化与安装事件。
//
// 未指定名称时使用插件实例类型的完整名称。插件实例为空、无法生成名称、名称重复或
// ID 冲突时会 panic。事件处理器可以在此方法返回前继续激活或卸载插件。
func (mgr *_AddInManager) Install(addInFace iface.FaceAny, name ...string) extension.AddInStatus {
	if addInFace.IsNil() {
		exception.Panicf("%w: %w: addInFace is nil", extension.ErrExtension, exception.ErrArgs)
	}

	addInName := pie.First(name)
	if addInName == "" {
		addInName = extension.GenAddInName(addInFace.Iface)
	}
	if addInName == "" {
		exception.Panicf("%w: anonymous add-in not allowed", extension.ErrExtension)
	}

	if _, ok := mgr.addInNameIndex[addInName]; ok {
		exception.Panicf("%w: add-in %q is already installed", extension.ErrExtension, addInName)
	}

	id := extension.GenAddInId(addInName)
	if existsIdx, ok := mgr.addInIdIndex[id]; ok {
		exception.Panicf("%w: add-in %q id %d conflict with %q, rename required", extension.ErrExtension, addInName, id, mgr.addInList.Get(existsIdx).V.Name())
	}

	status := &_AddInStatus{
		mgr:          mgr,
		id:           id,
		name:         addInName,
		instanceFace: addInFace,
		reflected:    reflect.ValueOf(addInFace.Iface),
	}
	statusSlot := mgr.addInList.PushBack(status)
	status.idx = statusSlot.Index()
	status.ver = statusSlot.Version()
	mgr.addInNameIndex[addInName] = statusSlot.Index()
	mgr.addInIdIndex[id] = statusSlot.Index()

	_EmitEventAddInStateChanged(mgr, status, extension.AddInState_Loaded)

	if status.state == extension.AddInState_Loaded {
		_EmitEventInstallAddIn(mgr, status)
	}

	return status
}

// Uninstall 按名称卸载插件。
// 插件不存在时直接返回；正在运行的插件会先同步派发卸载事件，再从管理器移除。
func (mgr *_AddInManager) Uninstall(name string) {
	statusIdx, ok := mgr.addInNameIndex[name]
	if !ok {
		return
	}
	status := mgr.addInList.Get(statusIdx).V
	status.uninstall()
}

// GetStatusByName 按名称查询当前已安装插件的状态信息。
func (mgr *_AddInManager) GetStatusByName(name string) (extension.AddInStatus, bool) {
	statusIdx, ok := mgr.addInNameIndex[name]
	if !ok {
		return nil, false
	}
	return mgr.addInList.Get(statusIdx).V, true
}

// GetStatusById 按 ID 查询当前已安装插件的状态信息。
func (mgr *_AddInManager) GetStatusById(id uint64) (extension.AddInStatus, bool) {
	statusIdx, ok := mgr.addInIdIndex[id]
	if !ok {
		return nil, false
	}
	return mgr.addInList.Get(statusIdx).V, true
}

// ListStatuses 按安装顺序返回当前插件状态的副本。
func (mgr *_AddInManager) ListStatuses() []extension.AddInStatus {
	statuses := make([]extension.AddInStatus, 0, mgr.addInList.Len())
	mgr.addInList.TraversalEach(func(slot *generic.FreeSlot[*_AddInStatus]) {
		statuses = append(statuses, slot.V)
	})
	return statuses
}

// getListStatuses 按安装顺序返回当前插件状态的内部接口副本。
func (mgr *_AddInManager) getListStatuses() []AddInStatus {
	statuses := make([]AddInStatus, 0, mgr.addInList.Len())
	mgr.addInList.TraversalEach(func(slot *generic.FreeSlot[*_AddInStatus]) {
		statuses = append(statuses, slot.V)
	})
	return statuses
}

// uninstallIfVersion 卸载仍占用指定槽位版本的插件，避免旧状态误删复用后的槽位。
func (mgr *_AddInManager) uninstallIfVersion(idx int, ver int64) {
	slot := mgr.addInList.Get(idx)
	if slot == nil || slot.Version() != ver {
		return
	}

	status := slot.V
	status.managedUnbindRuntimeHandles()

	if status.state == extension.AddInState_Running {
		_EmitEventUninstallAddIn(mgr, status)
	}

	delete(mgr.addInNameIndex, status.name)
	delete(mgr.addInIdIndex, status.id)
	mgr.addInList.ReleaseIfVersion(idx, ver)

	status.setState(extension.AddInState_Unloaded)
}
