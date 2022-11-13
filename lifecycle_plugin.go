package galaxy

import (
	"github.com/galaxy-kit/galaxy-go/runtime"
	"github.com/galaxy-kit/galaxy-go/service"
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
