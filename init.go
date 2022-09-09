package galaxy

import (
	"github.com/pangdogs/galaxy/api"
	"github.com/pangdogs/galaxy/comps/helloworld"
	"github.com/pangdogs/galaxy/pt"
)

func init() {
	pt.RegisterComp(api.HelloWorld, "组件（Component）Start时，在控制台打印`Hello World`", helloworld.HelloWorldComp{})
}
