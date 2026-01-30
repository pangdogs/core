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
	"sync/atomic"

	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
)

// ServiceAddInStatus 服务插件状态信息
type ServiceAddInStatus interface {
	iServiceAddInStatus
	AddInStatus

	// WaitState 等待状态完成
	WaitState(state AddInState) async.AsyncRet
}

type iServiceAddInStatus interface {
	setState(must, state AddInState) bool
	doInstallingOnce() bool
	doUninstallingOnce() bool
}

type _ServiceAddInStatus struct {
	mgr                    *_ServiceAddInManager
	name                   string
	instanceFace           iface.FaceAny
	reflected              reflect.Value
	state                  atomic.Int32
	idx                    int
	ver                    int64
	waitState              [AddInState_Unloaded + 1]chan async.Ret
	doInstallingOnceMark   atomic.Bool
	doUninstallingOnceMark atomic.Bool
}

// Name 插件名称
func (s *_ServiceAddInStatus) Name() string {
	return s.name
}

// InstanceFace 插件实例
func (s *_ServiceAddInStatus) InstanceFace() iface.FaceAny {
	return s.instanceFace
}

// Reflected 插件反射值
func (s *_ServiceAddInStatus) Reflected() reflect.Value {
	return s.reflected
}

// State 状态
func (s *_ServiceAddInStatus) State() AddInState {
	return AddInState(s.state.Load())
}

// Uninstall 卸载
func (s *_ServiceAddInStatus) Uninstall() {
	s.mgr.eventStream.Publish(&EventServiceUninstallAddIn{Status: s})
}

// WaitState 等待状态完成
func (s *_ServiceAddInStatus) WaitState(state AddInState) async.AsyncRet {
	if state < 0 || int(state) >= len(s.waitState) {
		exception.Panicf("%w: invalid state %q", ErrExtension, state)
	}
	return s.waitState[state]
}

func (s *_ServiceAddInStatus) setState(must, state AddInState) bool {
	if state <= must {
		return false
	}

	if !s.state.CompareAndSwap(int32(must), int32(state)) {
		return false
	}

	if state == AddInState_Unloaded {
		s.mgr.uninstallIfVersion(s.idx, s.ver)
	}

	async.YieldBreakT(s.waitState[state])

	return true
}

func (s *_ServiceAddInStatus) doInstallingOnce() bool {
	return s.doInstallingOnceMark.CompareAndSwap(false, true)
}

func (s *_ServiceAddInStatus) doUninstallingOnce() bool {
	return s.doUninstallingOnceMark.CompareAndSwap(false, true)
}
