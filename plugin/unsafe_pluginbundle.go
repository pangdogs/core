package plugin

// Deprecated: UnsafePluginBundle 访问插件包的内部方法
func UnsafePluginBundle(pluginBundle PluginBundle) _UnsafePluginBundle {
	return _UnsafePluginBundle{
		PluginBundle: pluginBundle,
	}
}

type _UnsafePluginBundle struct {
	PluginBundle
}

// Activate 设置插件活跃状态
func (up _UnsafePluginBundle) Activate(name string, b bool) {
	up.activate(name, b)
}
