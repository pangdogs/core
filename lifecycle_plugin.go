package golaxy

import (
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/service"
)

// LifecycleServicePluginInit 服务上的插件初始化回调，插件实现此接口即可使用
type LifecycleServicePluginInit interface {
	InitSP(ctx service.Context)
}

// LifecycleServicePluginShut 服务上的插件结束回调，插件实现此接口即可使用
type LifecycleServicePluginShut interface {
	ShutSP(ctx service.Context)
}

// LifecycleRuntimePluginInit 运行时上的插件初始化回调，插件实现此接口即可使用
type LifecycleRuntimePluginInit interface {
	InitRP(ctx runtime.Context)
}

// LifecycleRuntimePluginShut 运行时上的插件结束回调，插件实现此接口即可使用
type LifecycleRuntimePluginShut interface {
	ShutRP(ctx runtime.Context)
}
