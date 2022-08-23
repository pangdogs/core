package galaxy

import (
	"github.com/pangdogs/galaxy/api"
	"github.com/pangdogs/galaxy/complib"
	"github.com/pangdogs/galaxy/comps/helloworld"
)

func init() {
	complib.RegisterPt(api.HelloWorld, "实体创建时，在控制台打印Hello World。", helloworld.HelloWorldComp{})
}
