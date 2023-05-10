package golaxy

import (
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/service"
)

type LifecycleServicePluginInit interface {
	InitSP(ctx service.Context)
}

type LifecycleServicePluginShut interface {
	ShutSP(ctx service.Context)
}

type LifecycleRuntimePluginInit interface {
	InitRP(ctx runtime.Context)
}

type LifecycleRuntimePluginShut interface {
	ShutRP(ctx runtime.Context)
}
