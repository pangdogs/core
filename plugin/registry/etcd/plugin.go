package etcd

import (
	"github.com/pangdogs/galaxy/define"
	"github.com/pangdogs/galaxy/plugin/registry"
)

// Plugin 定义本插件
var Plugin = define.DefinePlugin[registry.Registry, WithOption]().ServicePlugin(newRegistry)
