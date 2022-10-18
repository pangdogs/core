package main

import (
	"fmt"
	"github.com/pangdogs/galaxy"
	"github.com/pangdogs/galaxy/comps/helloworld"
	"github.com/pangdogs/galaxy/pt"
	"github.com/pangdogs/galaxy/runtime"
	"github.com/pangdogs/galaxy/service"
	"github.com/pangdogs/galaxy/util"
)

func main() {
	// 创建实体库，注册实体原型
	entityLib := pt.NewEntityLib()
	entityLib.Register("Demo", []string{
		util.TypeFullName[helloworld.HelloWorld](),
		util.TypeFullName[DemoComp](),
	})

	// 创建服务上下文
	serviceCtx := service.NewContext(
		service.ContextOption.EntityLib(entityLib),
	)

	// 创建运行时上下文与运行时
	runtime := galaxy.NewRuntime(runtime.NewContext(serviceCtx),
		galaxy.RuntimeOption.Frame(runtime.NewFrame(30, 0, false)),
		galaxy.RuntimeOption.EnableAutoRun(true),
	)

	// 在运行时线程环境中，创建实体
	runtime.GetRuntimeCtx().SafeCallNoRetNoWait(func() {
		entity, err := galaxy.EntityCreator().
			RuntimeCtx(runtime.GetRuntimeCtx()).
			Prototype("Demo").
			Build()
		if err != nil {
			panic(err)
		}

		fmt.Printf("create entity[%s:%d:%d] finish\n", entity.GetPrototype(), entity.GetID(), entity.GetSerialNo())
	})

	service := galaxy.NewService(serviceCtx)

	<-service.Run()
}