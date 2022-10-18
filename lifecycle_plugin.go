package galaxy

import (
	"github.com/pangdogs/galaxy/runtime"
	"github.com/pangdogs/galaxy/service"
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
