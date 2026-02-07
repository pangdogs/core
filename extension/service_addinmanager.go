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
	"context"
	"reflect"
	"sync"

	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/types"
	"github.com/elliotchance/pie/v2"
)

// ServiceAddInManager 服务插件管理器
type ServiceAddInManager interface {
	AddInManager

	// WatchEvent 监听插件事件
	WatchEvent(ctx context.Context) <-chan any
}

// NewServiceAddInManager 创建服务插件管理器
func NewServiceAddInManager() ServiceAddInManager {
	return &_ServiceAddInManager{
		addInNameIndex: map[string]int{},
		addInIdIndex:   map[uint64]int{},
	}
}

type _ServiceAddInManager struct {
	sync.RWMutex
	addInNameIndex map[string]int
	addInIdIndex   map[uint64]int
	addInList      generic.FreeList[*_ServiceAddInStatus]
	eventStream    generic.EventStream[any]
}

// AddInManager 获取插件管理器
func (mgr *_ServiceAddInManager) AddInManager() AddInManager {
	return mgr
}

// Install 安装插件，不设置插件名称时，将会使用插件实例名称作为插件名称
func (mgr *_ServiceAddInManager) Install(addInFace iface.FaceAny, name ...string) AddInStatus {
	if addInFace.IsNil() {
		exception.Panicf("%w: %w: addInFace is nil", ErrExtension, exception.ErrArgs)
	}

	mgr.Lock()
	defer mgr.Unlock()

	addInName := pie.First(name)
	if addInName == "" {
		addInName = types.FullName(addInFace.Iface)
	}

	if _, ok := mgr.addInNameIndex[addInName]; ok {
		exception.Panicf("%w: addIn %q is already installed", ErrExtension, addInName)
	}

	id := GenAddInId(addInName)

	if existsIdx, ok := mgr.addInIdIndex[id]; ok {
		exception.Panicf("%w: addIn %q id index %d conflict with %q", ErrExtension, addInName, id, mgr.addInList.Get(existsIdx).V.name)
	}

	status := &_ServiceAddInStatus{
		mgr:          mgr,
		id:           id,
		name:         addInName,
		instanceFace: addInFace,
		reflected:    reflect.ValueOf(addInFace.Iface),
	}
	for i := range status.waitState {
		status.waitState[i] = async.NewAsyncRet()
	}
	slot := mgr.addInList.PushBack(status)
	status.idx = slot.Index()
	status.ver = slot.Version()
	mgr.addInNameIndex[addInName] = slot.Index()
	mgr.addInIdIndex[id] = slot.Index()

	async.YieldBreakT(status.waitState[AddInState_Loaded])

	mgr.eventStream.Publish(&EventServiceInstallAddIn{Status: status})

	return status
}

// Uninstall 卸载插件
func (mgr *_ServiceAddInManager) Uninstall(name string) {
	status, ok := mgr.Get(name)
	if !ok {
		return
	}
	status.Uninstall()
}

// Get 获取插件
func (mgr *_ServiceAddInManager) Get(name string) (AddInStatus, bool) {
	mgr.RLock()
	defer mgr.RUnlock()

	statusIdx, ok := mgr.addInNameIndex[name]
	if !ok {
		return nil, false
	}

	return mgr.addInList.Get(statusIdx).V, true
}

func (mgr *_ServiceAddInManager) GetById(id uint64) (AddInStatus, bool) {
	statusIdx, ok := mgr.addInIdIndex[id]
	if !ok {
		return nil, false
	}
	return mgr.addInList.Get(statusIdx).V, true
}

// List 获取所有插件
func (mgr *_ServiceAddInManager) List() []AddInStatus {
	mgr.RLock()
	defer mgr.RUnlock()

	statuses := make([]AddInStatus, 0, mgr.addInList.Len())

	mgr.addInList.TraversalEach(func(slot *generic.FreeSlot[*_ServiceAddInStatus]) {
		statuses = append(statuses, slot.V)
	})

	return statuses
}

// WatchEvent 监听插件事件
func (mgr *_ServiceAddInManager) WatchEvent(ctx context.Context) <-chan any {
	if ctx == nil {
		ctx = context.Background()
	}

	mgr.Lock()
	defer mgr.Unlock()

	statuses := make([]AddInStatus, 0, mgr.addInList.Len())

	mgr.addInList.TraversalEach(func(slot *generic.FreeSlot[*_ServiceAddInStatus]) {
		statuses = append(statuses, slot.V)
	})

	return mgr.eventStream.Subscribe(ctx, &EventServiceAddInSnapshot{Statuses: statuses})
}

func (mgr *_ServiceAddInManager) uninstallIfVersion(idx int, ver int64) {
	mgr.Lock()
	defer mgr.Unlock()

	slot := mgr.addInList.Get(idx)
	if slot == nil || slot.Version() != ver {
		return
	}

	status := slot.V

	delete(mgr.addInNameIndex, status.name)
	delete(mgr.addInIdIndex, status.id)
	mgr.addInList.Release(idx)
}
