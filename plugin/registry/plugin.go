package registry

import (
	"github.com/pangdogs/galaxy/define"
)

// Deregister 取消注册本插件
var Deregister = define.DefinePlugin[Registry, any]().Deregister()

// Get 获取本插件
var Get = define.DefinePlugin[Registry, any]().ServiceGet()
