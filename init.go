package galaxy

import (
	"github.com/pangdogs/galaxy/api"
	"github.com/pangdogs/galaxy/componentlib"
	"github.com/pangdogs/galaxy/components/helloworld"
)

func init() {
	componentlib.RegisterPt(api.HelloWorldComp, "显示Hello World", helloworld.HelloWorldComp{})
}
