package golaxy

import (
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/service"
)

type _ServicePluginInit interface {
	Init(ctx service.Context)
}

type _RuntimePluginInit interface {
	Init(ctx runtime.Context)
}

type _PluginShut interface {
	Shut()
}
