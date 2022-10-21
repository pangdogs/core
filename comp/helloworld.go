package comp

import "github.com/pangdogs/galaxy/define"

// HelloWorld HelloWorld组件接口
type HelloWorld interface {
	HelloWorld()
}

// HelloWorldPt HelloWorld组件原型
var HelloWorldPt = define.DefineComponentPt[HelloWorld]().ComponentPt()
