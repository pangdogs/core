package plugin

import "git.golaxy.org/core/util/generic"

// Deprecated: UnsafePluginBundle 访问插件包的内部方法
func UnsafePluginBundle(pluginBundle PluginBundle) _UnsafePluginBundle {
	return _UnsafePluginBundle{
		PluginBundle: pluginBundle,
	}
}

type _UnsafePluginBundle struct {
	PluginBundle
}

// SetActive 设置插件活跃状态
func (up _UnsafePluginBundle) SetActive(name string, b bool) {
	up.setActive(name, b)
}

// SetInstallCB 设置安装插件回调
func (up _UnsafePluginBundle) SetInstallCB(cb generic.Action1[PluginInfo]) {
	up.setInstallCB(cb)
}

// SetUninstallCB 设置卸载插件回调
func (up _UnsafePluginBundle) SetUninstallCB(cb generic.Action1[PluginInfo]) {
	up.setUninstallCB(cb)
}
