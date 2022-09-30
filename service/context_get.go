package service

import (
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/runtime"
)

func EntityContext(entity ec.Entity) Context {
	return runtime.EntityContext(entity).GetServiceCtx()
}

func ComponentContext(comp ec.Component) Context {
	return runtime.ComponentContext(comp).GetServiceCtx()
}
