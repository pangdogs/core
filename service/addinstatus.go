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
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"

	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/utils/iface"
)

// AddInStatus 描述与服务生命周期绑定的插件状态。
type AddInStatus interface {
	iAddInStatus
	extension.AddInStatus
}

type iAddInStatus interface {
	started()
	stopped()
}

type _AddInStatus struct {
	mgr          *_AddInManager
	id           uint64
	name         string
	instanceFace iface.FaceAny
	reflected    reflect.Value
	state        atomic.Int32
	stringerOnce sync.Once
	stringer     string
}

// Id 返回由插件名称生成的 ID。
func (s *_AddInStatus) Id() uint64 {
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
	return extension.AddInState(s.state.Load())
}

// String 实现 fmt.Stringer，返回包含 ID、名称和实例类型的 JSON 文本。
func (s *_AddInStatus) String() string {
	s.stringerOnce.Do(func() {
		s.stringer = fmt.Sprintf(`{"id":%d,"name":%q,"instance":%q}`, s.id, s.name, s.reflected.Type())
	})
	return s.stringer
}

func (s *_AddInStatus) started() {
	s.state.CompareAndSwap(int32(extension.AddInState_Loaded), int32(extension.AddInState_Running))
}

func (s *_AddInStatus) stopped() {
	s.mgr.stop(s)
}
