package galaxy

import (
	"errors"
	"fmt"
	"github.com/golaxy-kit/golaxy/ec"
	"github.com/golaxy-kit/golaxy/pt"
	"github.com/golaxy-kit/golaxy/runtime"
	"github.com/golaxy-kit/golaxy/service"
)

// EntityCreator 创建实体构建器
func EntityCreator(ctx runtime.Context, options ...pt.EntityOption) _EntityCreator {
	opts := pt.EntityOptions{}
	pt.WithEntityOption{}.Default()(&opts)

	for i := range options {
		options[i](&opts)
	}

	if opts.FaceCache == nil {
		opts.FaceCache = ctx.GetFaceCache()
	}
	if opts.HookCache == nil {
		opts.HookCache = ctx.GetHookCache()
	}

	return _EntityCreator{
		runtimeCtx: ctx,
		options:    opts,
	}
}

type _EntityCreator struct {
	runtimeCtx runtime.Context
	options    pt.EntityOptions
}

// Spawn 创建实体
func (creator _EntityCreator) Spawn() (ec.Entity, error) {
	return creator.spawn(nil)
}

// SpawnWithID 使用指定ID创建实体
func (creator _EntityCreator) SpawnWithID(id ec.ID) (ec.Entity, error) {
	return creator.spawn(func(options *pt.EntityOptions) {
		options.PersistID = id
	})
}

func (creator _EntityCreator) spawn(modifyOptions func(options *pt.EntityOptions)) (ec.Entity, error) {
	if creator.runtimeCtx == nil {
		return nil, errors.New("nil runtimeCtx")
	}

	runtimeCtx := creator.runtimeCtx
	serviceCtx := runtimeCtx.GetServiceCtx()

	entityLib := service.UnsafeContext(serviceCtx).GetOptions().EntityLib
	if entityLib == nil {
		return nil, errors.New("nil entityLib")
	}

	options := creator.options

	entityPt, ok := entityLib.Get(options.Prototype)
	if !ok {
		return nil, fmt.Errorf("entity '%s' not registered", options.Prototype)
	}

	if modifyOptions != nil {
		modifyOptions(&options)
	}

	entity := entityPt.UnsafeConstruct(options)

	if err := runtimeCtx.GetEntityMgr().AddEntity(entity, options.Accessibility); err != nil {
		return nil, fmt.Errorf("add entity to runtime context failed, %v", err)
	}

	return entity, nil
}
