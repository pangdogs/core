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

}

// Update 组件更新
func (comp *_DemoComp) Update() {
	// 服务上下文
	serviceCtx := service.Get(comp)

	// 注册服务
	registry.Register(
		serviceCtx,
		context.Background(),
		registry.Service{
			Name:    "demo",
			Version: "1.0.0",
			Nodes: []registry.Node{
				{
					Id:       "1",
					Address:  "",
					Metadata: nil,
				},
			},
		},
		registry.RegisterOption.TTL(10*time.Second))
}
