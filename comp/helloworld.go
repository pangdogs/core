package comp

import "github.com/pangdogs/galaxy/define"

// IHelloWorld HelloWorld组件接口
type IHelloWorld interface {
	HelloWorld()
}

// HelloWorld HelloWorld组件名称
var HelloWorld = define.Component[IHelloWorld]().Name()
