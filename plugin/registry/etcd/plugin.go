package etcd

import (
	"github.com/pangdogs/galaxy/define"
	"github.com/pangdogs/galaxy/plugin/registry"
)

// Register 注册本插件
var Register = define.Plugin[registry.Registry, Option]().Register(newRegistry)
