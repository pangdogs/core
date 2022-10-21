package galaxy

import (
	"github.com/pangdogs/galaxy/comp"
	"github.com/pangdogs/galaxy/comp/helloworld"
)

func init() {
	comp.HelloWorldPt.Register("组件Start时，在控制台打印`Hello World`", helloworld.HelloWorld{})
}
