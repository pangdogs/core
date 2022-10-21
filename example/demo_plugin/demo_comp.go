package main

import (
	"github.com/pangdogs/galaxy/define"
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/service"
)

func init() {
	DemoCompPt.Register("demo组件", _DemoComp{})
}

var DemoCompPt = define.DefineComponentPt[IDemoComp]().ComponentPt()

type IDemoComp interface{}

type _DemoComp struct {
	ec.ComponentBehavior
	count int
}

func (comp *_DemoComp) Start() {
	DemoPlugin.Get(service.ComponentContext(comp)).Test()
}
