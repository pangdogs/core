package galaxy

import (
	"errors"
	"fmt"
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/runtime"
	"github.com/pangdogs/galaxy/service"
)

// EntityCreator 实体构建器
func EntityCreator() _EntityCreator {
	return _EntityCreator{}
}

// Accessibility 可访问性
type Accessibility int32

const (
	Local     Accessibility = iota // 本地可以访问
	Global                         // 全局可以访问
	TryGlobal                      // 尝试全局可以访问
)

type _EntityCreator struct {
	runtimeCtx    runtime.Context
	prototype     string
	optSetters    []ec.EntityOptionSetter
	accessibility Accessibility
}

// RuntimeCtx 设置运行时上下文
func (creator _EntityCreator) RuntimeCtx(ctx runtime.Context) _EntityCreator {
	creator.runtimeCtx = ctx
	return creator
}

// Prototype 设置实体原型
func (creator _EntityCreator) Prototype(prototype string) _EntityCreator {
	creator.prototype = prototype
	return creator
}

// Options 设置创建实体的选项
func (creator _EntityCreator) Options(optSetter ...ec.EntityOptionSetter) _EntityCreator {
	creator.optSetters = optSetter
	return creator
}

// Accessibility 设置实体的可访问性
func (creator _EntityCreator) Accessibility(accessibility Accessibility) _EntityCreator {
	creator.accessibility = accessibility
	return creator
}

// Build 构建实体
func (creator _EntityCreator) Build() (ec.Entity, error) {
	if creator.runtimeCtx == nil {
		return nil, errors.New("nil runtimeCtx")
	}

	runtimeCtx := creator.runtimeCtx
	serviceCtx := runtimeCtx.GetServiceCtx()

	entityLib := service.UnsafeContext(serviceCtx).GetOptions().EntityLib
	if entityLib == nil {
		return nil, errors.New("nil entityLib")
	}

	entityPt, ok := entityLib.Get(creator.prototype)
	if !ok {
		return nil, fmt.Errorf("entity '%s' not registered", creator.prototype)
	}

	var addEntity func(entity ec.Entity) error

	switch creator.accessibility {
	case Local:
		addEntity = runtimeCtx.GetEntityMgr().AddEntity
	case Global:
		addEntity = runtimeCtx.GetEntityMgr().AddGlobalEntity
	case TryGlobal:
		addEntity = runtimeCtx.GetEntityMgr().TryAddGlobalEntity
	default:
		return nil, errors.New("accessibility invalid")
	}

	opts := ec.EntityOptions{}
	ec.EntityOption.Default()(&opts)

	for i := range creator.optSetters {
		creator.optSetters[i](&opts)
	}

	opts.Prototype = creator.prototype
	if opts.FaceCache == nil {
		opts.FaceCache = runtimeCtx.GetFaceCache()
	}
	if opts.HookCache == nil {
		opts.HookCache = runtimeCtx.GetHookCache()
	}

	entity := ec.UnsafeNewEntity(opts)
	entityPt.InstallComponents(entity)

	if err := addEntity(entity); err != nil {
		return nil, fmt.Errorf("runtime context add entity '%s:%d:%d' failed, %v", entity.GetPrototype(), entity.GetID(), entity.GetSerialNo(), err)
	}

	return entity, nil
}
