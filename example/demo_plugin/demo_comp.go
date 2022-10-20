package main

import (
	"github.com/pangdogs/galaxy/define"
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/pt"
	"github.com/pangdogs/galaxy/service"
)

func init() {
	pt.RegisterComponent(DemoComp, "demo组件", _DemoComp{})
}

type IDemoComp interface{}

var DemoComp = define.Component[IDemoComp]().Name()

type _DemoComp struct {
	ec.ComponentBehavior
	count int
}

func (comp *_DemoComp) Start() {
	GetDemoPlugin(service.ComponentContext(comp)).Test()
}
