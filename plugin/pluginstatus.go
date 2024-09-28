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

package plugin

import (
	"git.golaxy.org/core/utils/iface"
	"reflect"
	"sync/atomic"
)

// PluginStatus 插件状态信息
type PluginStatus interface {
	iPluginStatus

	// Name 插件名称
	Name() string
	// InstanceFace 插件实例
	InstanceFace() iface.FaceAny
	// Reflected 插件反射值
	Reflected() reflect.Value
	// State 状态
	State() PluginState
}

type iPluginStatus interface {
	setState(state, must PluginState) bool
}

type _PluginStatus struct {
	name         string
	instanceFace iface.FaceAny
	reflected    reflect.Value
	state        atomic.Int32
}

// Name 插件名称
func (s *_PluginStatus) Name() string {
	return s.name
}

// InstanceFace 插件实例
func (s *_PluginStatus) InstanceFace() iface.FaceAny {
	return s.instanceFace
}

// Reflected 插件反射值
func (s *_PluginStatus) Reflected() reflect.Value {
	return s.reflected
}

// State 状态
func (s *_PluginStatus) State() PluginState {
	return PluginState(s.state.Load())
}

func (s *_PluginStatus) setState(state, must PluginState) bool {
	return s.state.CompareAndSwap(int32(must), int32(state))
}
