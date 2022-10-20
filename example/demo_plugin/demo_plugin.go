package main

import (
	"fmt"
	"github.com/pangdogs/galaxy/define"
	"github.com/pangdogs/galaxy/service"
)

type DemoPlugin interface {
	Test()
}

var DemoPluginName = define.Plugin[DemoPlugin, any]().Name()

var RegisterDemoPlugin = define.Plugin[DemoPlugin, any]().Register(func(options ...any) DemoPlugin {
	return &_DemoPlugin{
		options: options,
	}
})

var DeregisterDemoPlugin = define.Plugin[DemoPlugin, any]().Deregister()

var GetDemoPlugin = define.Plugin[DemoPlugin, any]().ServiceGet()

type _DemoPlugin struct {
	options []any
}

func (d *_DemoPlugin) Init(ctx service.Context) {
	fmt.Printf("%s Init.\n", DemoPluginName)
}

func (d *_DemoPlugin) Shut() {
	fmt.Printf("%s Shut.\n", DemoPluginName)
}

func (d *_DemoPlugin) Test() {
	fmt.Printf("%s Test.\n", DemoPluginName)
}
