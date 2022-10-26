package main

import (
	"context"
	"github.com/pangdogs/galaxy/define"
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/plugin/registry"
	"github.com/pangdogs/galaxy/service"
	"time"
)

func init() {
	// 注册Demo组件
	DemoCompPt.Register("demo组件", _DemoComp{})
}

// DemoCompPt 定义Demo组件原型
var DemoCompPt = define.DefineComponentPt[DemoComp]().ComponentPt()

// DemoComp Demo组件接口
type DemoComp interface{}

// _DemoComp Demo组件实现类
type _DemoComp struct {
	ec.ComponentBehavior
}

// Start 组件开始
func (comp *_DemoComp) Start() {
	DemoPlugin.Get(service.Get(comp)).Test()
	registry.Register(
		service.Get(comp),
		context.Background(),
		registry.Service{
			Name:      "demo",
			Version:   "1.0.0",
			Metadata:  nil,
			Endpoints: nil,
			Nodes:     nil,
		}, 5*time.Second)
}
