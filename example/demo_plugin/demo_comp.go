package main

import (
	"github.com/pangdogs/galaxy/define"
	"github.com/pangdogs/galaxy/ec"
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
	DemoPlugin.ECGet(comp).Test()
}
