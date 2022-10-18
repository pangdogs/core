package service

import (
	"github.com/pangdogs/galaxy/ec"
	_ "unsafe"
)

// EntityContext 从实体上获取服务上下文
func EntityContext(entity ec.Entity) Context {
	return entityServiceContext(entity)
}

// ComponentContext 从组件上获取服务上下文
func ComponentContext(comp ec.Component) Context {
	return componentServiceContext(comp)
}

//go:linkname entityServiceContext github.com/pangdogs/galaxy/runtime.entityServiceContext
func entityServiceContext(entity ec.Entity) Context

//go:linkname componentServiceContext github.com/pangdogs/galaxy/runtime.componentServiceContext
func componentServiceContext(comp ec.Component) Context
