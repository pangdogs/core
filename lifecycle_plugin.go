package golaxy

import (
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/service"
)

type _ServicePluginInit interface {
	InitService(ctx service.Context)
}

type _ServicePluginShut interface {
	ShutService(ctx service.Context)
}

type _RuntimePluginInit interface {
	InitRuntime(ctx runtime.Context)
}

type _RuntimePluginShut interface {
	ShutRuntime(ctx runtime.Context)
}
