package galaxy

import (
	"github.com/golaxy-kit/golaxy/runtime"
	"github.com/golaxy-kit/golaxy/service"
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
