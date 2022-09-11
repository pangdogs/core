package api

import (
	"github.com/pangdogs/galaxy/core"
)

// HelloWorld HelloWorld组件接口名称
var HelloWorld = core.TypeFullName[IHelloWorld]()

// IHelloWorld HelloWorld组件接口定义
type IHelloWorld interface{}
