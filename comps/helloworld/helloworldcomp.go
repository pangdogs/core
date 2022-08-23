package helloworld

import (
	"fmt"
	"github.com/pangdogs/galaxy/core"
)

// HelloWorldComp HelloWorld组件
type HelloWorldComp struct {
	core.ComponentBehavior
}

// Start 开始
func (comp *HelloWorldComp) Start() {
	fmt.Println("Hello World")
}
