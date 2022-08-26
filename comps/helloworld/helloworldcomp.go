// Package helloworld 提供HelloWorld示例。
package helloworld

import (
	"fmt"
	"github.com/pangdogs/galaxy/core"
)

// HelloWorldComp HelloWorld示例组件，实体（Entity）创建时，在控制台打印`Hello World`。
type HelloWorldComp struct {
	core.ComponentBehavior
}

// Start 开始
func (comp *HelloWorldComp) Start() {
	fmt.Printf("Hello world, my id is %d", comp.GetEntity().GetID())
}
