package registry

import (
	"github.com/pangdogs/galaxy/define"
)

// Deregister 取消注册本插件
var Deregister = define.Plugin[Registry, any]().Deregister()

// Get 获取本插件
var Get = define.Plugin[Registry, any]().ServiceGet()
