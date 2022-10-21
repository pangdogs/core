package main

import (
	"fmt"
	"github.com/pangdogs/galaxy/define"
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/pt"
)

func init() {
	pt.RegisterComponent(DemoComp, "demo组件", _DemoComp{})
}

type IDemoComp interface{}

var DemoComp = define.DefineComponentPt[IDemoComp]().Name()

type _DemoComp struct {
	ec.ComponentBehavior
	count int
}

func (comp *_DemoComp) Awake() {
	fmt.Printf("I'm entity[%s:%d:%d], %s Awake.\n", comp.GetEntity().GetPrototype(), comp.GetEntity().GetID(), comp.GetEntity().GetSerialNo(), comp.GetName())
}

func (comp *_DemoComp) Start() {
	fmt.Printf("I'm entity[%s:%d:%d], %s Start.\n", comp.GetEntity().GetPrototype(), comp.GetEntity().GetID(), comp.GetEntity().GetSerialNo(), comp.GetName())
}

func (comp *_DemoComp) Update() {
	if comp.count%30 == 0 {
		fmt.Printf("I'm entity[%s:%d:%d], %s Update(%d).\n",
			comp.GetEntity().GetPrototype(), comp.GetEntity().GetID(), comp.GetEntity().GetSerialNo(), comp.GetName(), comp.count)
	}
}

func (comp *_DemoComp) LateUpdate() {
	if comp.count%30 == 0 {
		fmt.Printf("I'm entity[%s:%d:%d], %s LateUpdate(%d).\n",
			comp.GetEntity().GetPrototype(), comp.GetEntity().GetID(), comp.GetEntity().GetSerialNo(), comp.GetName(), comp.count)
	}
	comp.count++
}

func (comp *_DemoComp) Shut() {
	fmt.Printf("I'm entity[%s:%d:%d], %s Shut.\n", comp.GetEntity().GetPrototype(), comp.GetEntity().GetID(), comp.GetEntity().GetSerialNo(), comp.GetName())
}
