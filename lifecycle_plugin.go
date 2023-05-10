package golaxy

import (
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/service"
)

type LifecycleServicePluginInit interface {
	InitServicePlugin(ctx service.Context)
}

type LifecycleServicePluginShut interface {
	ShutServicePlugin(ctx service.Context)
}

type LifecycleRuntimePluginInit interface {
	InitRuntimePlugin(ctx runtime.Context)
}

type LifecycleRuntimePluginShut interface {
	ShutRuntimePlugin(ctx runtime.Context)
}
