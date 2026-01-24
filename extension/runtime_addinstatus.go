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

	"git.golaxy.org/core/event"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
)

// RuntimeAddInStatus 运行时插件状态信息
type RuntimeAddInStatus interface {
	iRuntimeAddInStatus
	AddInStatus
}

type iRuntimeAddInStatus interface {
	setState(state AddInState)
	managedRuntimeRunningEventHandle(runtimeRunningEventHandle event.Handle)
	managedUnbindRuntimeHandles()
}

const (
	runtimeAddInStatusReentrancyGuard_Uninstall = iota
)

type _RuntimeAddInStatus struct {
	mgr                   *_RuntimeAddInManager
	name                  string
	instanceFace          iface.FaceAny
	reflected             reflect.Value
	state                 AddInState
	processedStateBits    generic.Bits16
	reentrancyGuard       generic.ReentrancyGuardBits8
	idx                   int
	ver                   int64
	managedRuntimeHandles [1]event.Handle
}

// Name 插件名称
func (s *_RuntimeAddInStatus) Name() string {
	return s.name
}

// InstanceFace 插件实例
func (s *_RuntimeAddInStatus) InstanceFace() iface.FaceAny {
	return s.instanceFace
}

// Reflected 插件反射值
func (s *_RuntimeAddInStatus) Reflected() reflect.Value {
	return s.reflected
}

// State 状态
func (s *_RuntimeAddInStatus) State() AddInState {
	return s.state
}

// Uninstall 卸载
func (s *_RuntimeAddInStatus) Uninstall() {
	s.reentrancyGuard.Call(runtimeAddInStatusReentrancyGuard_Uninstall, func() {
		s.mgr.uninstallIfVersion(s.idx, s.ver)
	})
}

func (s *_RuntimeAddInStatus) setState(state AddInState) {
	slot := s.mgr.addInList.Get(s.idx)
	if slot.Version() != s.ver {
		return
	}

	if s.processedStateBits.Is(int(state)) {
		return
	}

	s.state = state
	s.processedStateBits.Set(int(state), true)

	_EmitEventRuntimeAddInStateChanged(s.mgr, s, state)
}

func (s *_RuntimeAddInStatus) managedRuntimeRunningEventHandle(runtimeRunningEventHandle event.Handle) {
	if s.managedRuntimeHandles[0] != runtimeRunningEventHandle {
		s.managedRuntimeHandles[0].Unbind()
	}
	s.managedRuntimeHandles[0] = runtimeRunningEventHandle
}

func (s *_RuntimeAddInStatus) managedUnbindRuntimeHandles() {
	event.UnbindHandles(s.managedRuntimeHandles[:])
}
