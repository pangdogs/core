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

	ListAndWatch(ctx context.Context) <-chan any
}

// NewServiceAddInManager 创建服务插件管理器
func NewServiceAddInManager() ServiceAddInManager {
	return &_ServiceAddInManager{
		addInNameIndex: map[string]int{},
		watchers:       map[*generic.UnboundedChannel[any]]struct{}{},
	}
}

type _ServiceAddInManager struct {
	sync.RWMutex
	addInNameIndex map[string]int
	addInList      generic.FreeList[*_ServiceAddInStatus]
	watchers       map[*generic.UnboundedChannel[any]]struct{}
}

// GetAddInManager 获取插件管理器
func (mgr *_ServiceAddInManager) GetAddInManager() AddInManager {
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

	status := &_ServiceAddInStatus{
		mgr:          mgr,
		name:         addInName,
		instanceFace: addInFace,
		reflected:    reflect.ValueOf(addInFace.Iface),
	}
	for i := range status.waitState {
		status.waitState[i] = async.MakeAsyncRet()
	}
	slot := mgr.addInList.PushBack(status)
	status.idx = slot.Index()
	status.ver = slot.Version()
	mgr.addInNameIndex[addInName] = slot.Index()

	async.YieldBreakT(status.waitState[AddInState_Loaded])

	for watcher := range mgr.watchers {
		watcher.In() <- &EventServiceInstallAddIn{Status: status}
	}

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

// List 获取所有插件
func (mgr *_ServiceAddInManager) List() []AddInStatus {
	mgr.RLock()
	defer mgr.RUnlock()

	status := make([]AddInStatus, 0, mgr.addInList.Len())

	mgr.addInList.TraversalEach(func(slot *generic.FreeSlot[*_ServiceAddInStatus]) {
		status = append(status, slot.V)
	})

	return status
}

// ListAndWatch 获取并观察所有插件变化
func (mgr *_ServiceAddInManager) ListAndWatch(ctx context.Context) <-chan any {
	if ctx == nil {
		ctx = context.Background()
	}

	mgr.Lock()
	defer mgr.Unlock()

	watcher := generic.NewUnboundedChannel[any]()
	mgr.watchers[watcher] = struct{}{}

	statusList := make([]AddInStatus, 0, mgr.addInList.Len())

	mgr.addInList.TraversalEach(func(slot *generic.FreeSlot[*_ServiceAddInStatus]) {
		statusList = append(statusList, slot.V)
	})

	watcher.In() <- &EventServiceAddInSnapshot{
		StatusList: statusList,
	}

	go func() {
		<-ctx.Done()
		mgr.Lock()
		defer mgr.Unlock()
		watcher.Close()
		delete(mgr.watchers, watcher)
	}()

	return watcher.Out()
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
	mgr.addInList.Release(idx)
}
