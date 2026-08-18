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
	"fmt"
	"reflect"

	"git.golaxy.org/core/event"
	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
)

// AddInStatus 描述可热插拔运行时插件的状态。
type AddInStatus interface {
	iAddInStatus
	extension.AddInStatus
}

type iAddInStatus interface {
	started()
	managedRuntimeRunningEventHandle(runtimeRunningEventHandle event.Handle)
	managedUnbindRuntimeHandles()
}

const addInStatusReentrancyGuardUninstall = iota

type _AddInStatus struct {
	mgr                   *_AddInManager
	id                    uint64
	name                  string
	instanceFace          iface.FaceAny
	reflected             reflect.Value
	state                 extension.AddInState
	reentrancyGuard       generic.ReentrancyGuardBits8
	idx                   int
	ver                   int64
	managedRuntimeHandles [1]event.Handle
	stringer              string
}

// ID 返回由插件名称生成的 ID。
func (s *_AddInStatus) ID() uint64 {
	return s.id
}

// Name 返回插件注册名称。
func (s *_AddInStatus) Name() string {
	return s.name
}

// InstanceFace 返回插件实例及其接口缓存。
func (s *_AddInStatus) InstanceFace() iface.FaceAny {
	return s.instanceFace
}

// Reflected 返回插件实例的反射值。
func (s *_AddInStatus) Reflected() reflect.Value {
	return s.reflected
}

// State 返回插件当前的生命周期状态。
func (s *_AddInStatus) State() extension.AddInState {
	return s.state
}

// String 实现 fmt.Stringer，返回包含 ID、名称和实例类型的 JSON 文本。
func (s *_AddInStatus) String() string {
	if s.stringer == "" {
		s.stringer = fmt.Sprintf(`{"id":%d,"name":%q,"instance":%q}`, s.id, s.name, s.reflected.Type())
	}
	return s.stringer
}

func (s *_AddInStatus) started() {
	s.setState(extension.AddInState_Running)
}

func (s *_AddInStatus) managedRuntimeRunningEventHandle(runtimeRunningEventHandle event.Handle) {
	if s.managedRuntimeHandles[0] != runtimeRunningEventHandle {
		s.managedRuntimeHandles[0].Unbind()
	}
	s.managedRuntimeHandles[0] = runtimeRunningEventHandle
}

func (s *_AddInStatus) managedUnbindRuntimeHandles() {
	event.UnbindHandles(s.managedRuntimeHandles[:])
}

func (s *_AddInStatus) uninstall() {
	s.reentrancyGuard.Call(addInStatusReentrancyGuardUninstall, func() {
		s.mgr.uninstallIfVersion(s.idx, s.ver)
	})
}

func (s *_AddInStatus) setState(state extension.AddInState) {
	slot := s.mgr.addInList.Get(s.idx)
	if slot.Version() != s.ver {
		return
	}
	if s.state >= state {
		return
	}

	s.state = state
	_EmitEventAddInStateChanged(s.mgr, s, state)
}
