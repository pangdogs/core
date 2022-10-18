package main

import (
	"fmt"
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/pt"
	"github.com/pangdogs/galaxy/service"
)

func init() {
	pt.RegisterComponent("", "demo组件", DemoComp{})
}

type DemoComp struct {
	ec.ComponentBehavior
	count int
}

func (comp *DemoComp) Awake() {
	fmt.Printf("I'm entity[%s:%d:%d], DemoComp Awake.\n", comp.GetEntity().GetPrototype(), comp.GetEntity().GetID(), comp.GetEntity().GetSerialNo())
}

func (comp *DemoComp) Start() {
	fmt.Printf("I'm entity[%s:%d:%d], DemoComp Start.\n", comp.GetEntity().GetPrototype(), comp.GetEntity().GetID(), comp.GetEntity().GetSerialNo())
}

func (comp *DemoComp) Update() {
	if comp.count%30 == 0 {
		fmt.Printf("I'm entity[%s:%d:%d], DemoComp Update.\n", comp.GetEntity().GetPrototype(), comp.GetEntity().GetID(), comp.GetEntity().GetSerialNo())
	}

	if comp.count >= 300 {
		service.EntityContext(comp.GetEntity()).GetCancelFunc()()
	}
	comp.count++
}

func (comp *DemoComp) LateUpdate() {
	if comp.count%30 == 0 {
		fmt.Printf("I'm entity[%s:%d:%d], DemoComp LateUpdate.\n", comp.GetEntity().GetPrototype(), comp.GetEntity().GetID(), comp.GetEntity().GetSerialNo())
	}
}

func (comp *DemoComp) Shut() {
	fmt.Printf("I'm entity[%s:%d:%d], DemoComp Shut.\n", comp.GetEntity().GetPrototype(), comp.GetEntity().GetID(), comp.GetEntity().GetSerialNo())
}
