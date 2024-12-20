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
	"git.golaxy.org/core/utils/iface"
	"reflect"
	"sync/atomic"
)

// AddInStatus 插件状态信息
type AddInStatus interface {
	iAddInStatus

	// Name 插件名称
	Name() string
	// InstanceFace 插件实例
	InstanceFace() iface.FaceAny
	// Reflected 插件反射值
	Reflected() reflect.Value
	// State 状态
	State() AddInState
}

type iAddInStatus interface {
	setState(state, must AddInState) bool
}

type _AddInStatus struct {
	name         string
	instanceFace iface.FaceAny
	reflected    reflect.Value
	state        atomic.Int32
}

// Name 插件名称
func (s *_AddInStatus) Name() string {
	return s.name
}

// InstanceFace 插件实例
func (s *_AddInStatus) InstanceFace() iface.FaceAny {
	return s.instanceFace
}

// Reflected 插件反射值
func (s *_AddInStatus) Reflected() reflect.Value {
	return s.reflected
}

// State 状态
func (s *_AddInStatus) State() AddInState {
	return AddInState(s.state.Load())
}

func (s *_AddInStatus) setState(state, must AddInState) bool {
	return s.state.CompareAndSwap(int32(must), int32(state))
}
