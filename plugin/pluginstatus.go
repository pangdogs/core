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
	setState(state, old PluginState) bool
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

func (s *_PluginStatus) setState(state, old PluginState) bool {
	return s.state.CompareAndSwap(int32(old), int32(state))
}
