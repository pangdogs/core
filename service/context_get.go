package service

import (
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/runtime"
)

// EntityContext 从实体上获取服务上下文
func EntityContext(entity ec.Entity) Context {
	return runtime.EntityContext(entity).GetServiceCtx()
}

// ComponentContext 从组件上获取服务上下文
func ComponentContext(comp ec.Component) Context {
	return runtime.ComponentContext(comp).GetServiceCtx()
}
