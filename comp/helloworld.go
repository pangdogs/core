package comp

import (
	"github.com/pangdogs/galaxy/util"
)

// HelloWorld HelloWorld组件接口名称
var HelloWorld = util.TypeFullName[IHelloWorld]()

// IHelloWorld HelloWorld组件接口定义
type IHelloWorld interface {
	HelloWorld()
}
