package galaxy

import (
	"github.com/pangdogs/galaxy/compi"
	"github.com/pangdogs/galaxy/comps/helloworld"
	"github.com/pangdogs/galaxy/pt"
)

func init() {
	pt.RegisterComponent(compi.HelloWorld, "组件Start时，在控制台打印`Hello World`", helloworld.HelloWorld{})
}
