package galaxy

import (
	"github.com/galaxy-kit/galaxy/runtime"
	"github.com/galaxy-kit/galaxy/service"
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
