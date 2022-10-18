package galaxy

import (
	"github.com/pangdogs/galaxy/comp"
	"github.com/pangdogs/galaxy/comp/helloworld"
	"github.com/pangdogs/galaxy/pt"
)

func init() {
	pt.RegisterComponent(comp.HelloWorld, "组件Start时，在控制台打印`Hello World`", helloworld.HelloWorld{})
}
