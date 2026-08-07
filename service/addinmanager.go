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

package service

import (
	"maps"
	"reflect"
	"slices"
	"sync/atomic"

	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
	"github.com/elliotchance/pie/v2"
)

// AddInManager 管理与服务生命周期绑定的插件。
//
// 插件只能在服务启动前安装或卸载。服务启动时管理器会被冻结，此后 Install 和
// Uninstall 均会 panic。插件按安装顺序启动，并在服务停止时按相反顺序关闭。
//
// 管理器通过不可变快照支持多个 goroutine 并发安装、卸载和查询插件。
type AddInManager interface {
	iAddInManager
	extension.AddInManager
}

type iAddInManager interface {
	freeze() []AddInStatus
	getList() []AddInStatus
}

// NewAddInManager 创建一个未冻结的空服务插件管理器。
func NewAddInManager() AddInManager {
	mgr := &_AddInManager{}
	mgr.snapshot.Store(&_AddInManagerSnapshot{
		addInNameIndex: map[string]*_AddInStatus{},
		addInIdIndex:   map[uint64]*_AddInStatus{},
	})
	return mgr
}

type _AddInManager struct {
	snapshot atomic.Pointer[_AddInManagerSnapshot]
}

// _AddInManagerSnapshot 是插件集合的只读快照。
// 快照发布后不再原地修改；写操作先克隆快照，再通过 CAS 将其整体替换。
type _AddInManagerSnapshot struct {
	frozen         bool
	addInNameIndex map[string]*_AddInStatus
	addInIdIndex   map[uint64]*_AddInStatus
	addInList      []*_AddInStatus
}

// clone 复制当前快照，返回一个可供写操作修改的未发布副本。
func (snapshot *_AddInManagerSnapshot) clone() *_AddInManagerSnapshot {
	return &_AddInManagerSnapshot{
		frozen:         snapshot.frozen,
		addInNameIndex: maps.Clone(snapshot.addInNameIndex),
		addInIdIndex:   maps.Clone(snapshot.addInIdIndex),
		addInList:      slices.Clone(snapshot.addInList),
	}
}

// remove 从未发布的快照副本中移除指定插件状态。
// 指针校验可避免旧状态误删后来安装的同名插件。
func (snapshot *_AddInManagerSnapshot) remove(status *_AddInStatus) {
	if current, ok := snapshot.addInNameIndex[status.name]; !ok || current != status {
		return
	}

	delete(snapshot.addInNameIndex, status.name)
	delete(snapshot.addInIdIndex, status.id)

	idx := slices.Index(snapshot.addInList, status)
	if idx >= 0 {
		snapshot.addInList = slices.Delete(snapshot.addInList, idx, idx+1)
	}
}

// AddInManager 返回当前服务插件管理器的公共接口。
func (mgr *_AddInManager) AddInManager() extension.AddInManager {
	return mgr
}

// Install 安装插件并返回其状态信息。
//
// 未指定名称时使用插件实例类型的完整名称。插件实例为空、无法生成名称、名称重复、
// ID 冲突或管理器已冻结时会 panic。安装成功后，插件状态为 AddInState_Loaded。
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

	id := extension.GenAddInId(addInName)
	status := &_AddInStatus{
		mgr:          mgr,
		id:           id,
		name:         addInName,
		instanceFace: addInFace,
		reflected:    reflect.ValueOf(addInFace.Iface),
	}

	for {
		snapshot := mgr.snapshot.Load()
		mgr.checkMutable(snapshot)

		if _, ok := snapshot.addInNameIndex[addInName]; ok {
			exception.Panicf("%w: add-in %q is already installed", extension.ErrExtension, addInName)
		}
		if exists, ok := snapshot.addInIdIndex[id]; ok {
			exception.Panicf("%w: add-in %q id %d conflict with %q, rename required", extension.ErrExtension, addInName, id, exists.name)
		}

		next := snapshot.clone()
		next.addInNameIndex[addInName] = status
		next.addInIdIndex[id] = status
		next.addInList = append(next.addInList, status)

		if mgr.snapshot.CompareAndSwap(snapshot, next) {
			return status
		}
	}
}

// Uninstall 卸载指定名称的插件。
//
// 在未冻结阶段，插件不存在时直接返回；卸载成功后，其状态变为
// AddInState_Unloaded。管理器冻结后调用此方法会 panic。
func (mgr *_AddInManager) Uninstall(name string) {
	for {
		snapshot := mgr.snapshot.Load()
		mgr.checkMutable(snapshot)

		status, ok := snapshot.addInNameIndex[name]
		if !ok {
			return
		}

		next := snapshot.clone()
		next.remove(status)

		if mgr.snapshot.CompareAndSwap(snapshot, next) {
			status.state.Store(int32(extension.AddInState_Unloaded))
			return
		}
	}
}

// GetStatusByName 按名称查询当前已安装插件的状态信息。
func (mgr *_AddInManager) GetStatusByName(name string) (extension.AddInStatus, bool) {
	status, ok := mgr.snapshot.Load().addInNameIndex[name]
	return status, ok
}

// GetStatusById 按 ID 查询当前已安装插件的状态信息。
func (mgr *_AddInManager) GetStatusById(id uint64) (extension.AddInStatus, bool) {
	status, ok := mgr.snapshot.Load().addInIdIndex[id]
	return status, ok
}

// ListStatuses 按安装顺序返回当前插件的状态信息。
// 返回的切片是独立副本，修改切片不会影响管理器。
func (mgr *_AddInManager) ListStatuses() []extension.AddInStatus {
	addInList := mgr.snapshot.Load().addInList
	statuses := make([]extension.AddInStatus, len(addInList))
	for i, status := range addInList {
		statuses[i] = status
	}
	return statuses
}

// freeze 原子地冻结管理器，并按安装顺序返回插件状态。
// 管理器已经冻结时再次调用会 panic。
func (mgr *_AddInManager) freeze() []AddInStatus {
	for {
		snapshot := mgr.snapshot.Load()
		if snapshot.frozen {
			exception.Panicf("%w: service add-in manager is already frozen", extension.ErrExtension)
		}

		next := &_AddInManagerSnapshot{
			frozen:         true,
			addInNameIndex: snapshot.addInNameIndex,
			addInIdIndex:   snapshot.addInIdIndex,
			addInList:      snapshot.addInList,
		}
		if !mgr.snapshot.CompareAndSwap(snapshot, next) {
			continue
		}

		statuses := make([]AddInStatus, len(next.addInList))
		for i, status := range next.addInList {
			statuses[i] = status
		}
		return statuses
	}
}

// getList 按安装顺序返回当前插件状态的内部接口副本。
func (mgr *_AddInManager) getList() []AddInStatus {
	addInList := mgr.snapshot.Load().addInList
	statuses := make([]AddInStatus, len(addInList))
	for i, status := range addInList {
		statuses[i] = status
	}
	return statuses
}

// checkMutable 确认指定快照仍处于可修改阶段。
func (mgr *_AddInManager) checkMutable(snapshot *_AddInManagerSnapshot) {
	if snapshot.frozen {
		exception.Panicf("%w: service add-in manager is frozen", extension.ErrExtension)
	}
}

// stop 在插件完成停服流程后，将其从管理器中移除并标记为已卸载。
// 此内部清理允许在管理器冻结后执行。
func (mgr *_AddInManager) stop(status *_AddInStatus) {
	for {
		snapshot := mgr.snapshot.Load()
		if current, ok := snapshot.addInNameIndex[status.name]; !ok || current != status {
			return
		}

		next := snapshot.clone()
		next.remove(status)

		if mgr.snapshot.CompareAndSwap(snapshot, next) {
			status.state.Store(int32(extension.AddInState_Unloaded))
			return
		}
	}
}
