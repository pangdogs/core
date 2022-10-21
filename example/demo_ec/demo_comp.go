package main

import (
	"fmt"
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
	count int
}

// Awake 组件唤醒
func (comp *_DemoComp) Awake() {
	fmt.Printf("I'm entity[%s:%d:%d], %s Awake.\n", comp.GetEntity().GetPrototype(), comp.GetEntity().GetID(), comp.GetEntity().GetSerialNo(), comp.GetName())
}

// Start 组件开始
func (comp *_DemoComp) Start() {
	fmt.Printf("I'm entity[%s:%d:%d], %s Start.\n", comp.GetEntity().GetPrototype(), comp.GetEntity().GetID(), comp.GetEntity().GetSerialNo(), comp.GetName())
}

// Update 组件更新
func (comp *_DemoComp) Update() {
	if comp.count%30 == 0 {
		fmt.Printf("I'm entity[%s:%d:%d], %s Update(%d).\n",
			comp.GetEntity().GetPrototype(), comp.GetEntity().GetID(), comp.GetEntity().GetSerialNo(), comp.GetName(), comp.count)
	}
}

// LateUpdate 组件滞后更新
func (comp *_DemoComp) LateUpdate() {
	if comp.count%30 == 0 {
		fmt.Printf("I'm entity[%s:%d:%d], %s LateUpdate(%d).\n",
			comp.GetEntity().GetPrototype(), comp.GetEntity().GetID(), comp.GetEntity().GetSerialNo(), comp.GetName(), comp.count)
	}
	comp.count++
}

// Shut 组件停止
func (comp *_DemoComp) Shut() {
	fmt.Printf("I'm entity[%s:%d:%d], %s Shut.\n", comp.GetEntity().GetPrototype(), comp.GetEntity().GetID(), comp.GetEntity().GetSerialNo(), comp.GetName())
}
