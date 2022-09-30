// Package helloworld 提供HelloWorld示例。
package helloworld

import (
	"fmt"
	"github.com/pangdogs/galaxy/ec"
)

// HelloWorld HelloWorld示例组件，组件Start时，在控制台打印`Hello World`。
type HelloWorld struct {
	ec.ComponentBehavior
}

// Start 开始
func (comp *HelloWorld) Start() {
	comp.HelloWorld()
}

func (comp *HelloWorld) HelloWorld() {
	fmt.Printf("[%d] Hello World\n", comp.GetEntity().GetID())
}
