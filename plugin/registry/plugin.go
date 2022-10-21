package registry

import "github.com/pangdogs/galaxy/define"

var Plugin = define.DefinePluginInterface[Registry]().ServicePluginInterface()
