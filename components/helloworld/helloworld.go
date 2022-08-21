package helloworld

import (
	"fmt"
	"github.com/pangdogs/galaxy/core"
)

type HelloWorldComp struct {
	core.ComponentBehavior
}

func (comp *HelloWorldComp) Start() {
	fmt.Println("Hello World")
}
