package golaxy

import (
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/service"
)

type LifecycleServicePluginInit interface {
	InitService(ctx service.Context)
}

type LifecycleServicePluginShut interface {
	ShutService(ctx service.Context)
}

type LifecycleRuntimePluginInit interface {
	InitRuntime(ctx runtime.Context)
}

type LifecycleRuntimePluginShut interface {
	ShutRuntime(ctx runtime.Context)
}
