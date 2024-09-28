package plugin

// Deprecated: UnsafePluginStatus 访问插件状态信息的内部方法
func UnsafePluginStatus(status PluginStatus) _UnsafePluginStatus {
	return _UnsafePluginStatus{
		PluginStatus: status,
	}
}

type _UnsafePluginStatus struct {
	PluginStatus
}

// SetState 修改状态
func (up _UnsafePluginStatus) SetState(state, must PluginState) bool {
	return up.setState(state, must)
}
