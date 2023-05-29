package golaxy

import (
	"errors"
	"fmt"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/uid"
)

// NewEntityCreator 创建实体构建器
func NewEntityCreator(ctx runtime.Context, prototype string, options ...EntityOption) EntityCreator {
	if ctx == nil {
		panic("nil ctx")
	}

	opts := EntityOptions{}
	WithOption{}.EntityDefault()(&opts)

	for i := range options {
		options[i](&opts)
	}

	opts.Prototype = prototype

	if opts.FaceAnyAllocator == nil {
		opts.FaceAnyAllocator = ctx.GetFaceAnyAllocator()
	}
	if opts.HookAllocator == nil {
		opts.HookAllocator = ctx.GetHookAllocator()
	}

	return EntityCreator{
		ctx:     ctx,
		options: opts,
		inited:  true,
	}
}

type EntityCreator struct {
	ctx     runtime.Context
	options EntityOptions
	inited  bool
}

// Spawn 创建实体
func (creator EntityCreator) Spawn() (ec.Entity, error) {
	return creator.spawn(nil)
}

// SpawnWithId 使用指定Id创建实体
func (creator EntityCreator) SpawnWithId(id uid.Id) (ec.Entity, error) {
	return creator.spawn(func(options *EntityOptions) {
		options.PersistId = id
	})
}

func (creator EntityCreator) spawn(modifyOptions func(options *EntityOptions)) (ec.Entity, error) {
	if !creator.inited {
		return nil, errors.New("not inited")
	}

	runtimeCtx := creator.ctx
	serviceCtx := service.Get(runtimeCtx)

	entityPt, ok := pt.TryGetEntityPt(serviceCtx, creator.options.Prototype)
	if !ok {
		return nil, fmt.Errorf("entity prototype %q not registered", creator.options.Prototype)
	}

	if modifyOptions != nil {
		modifyOptions(&creator.options)
	}

	entity := entityPt.UnsafeConstruct(creator.options.EntityOptions)

	if err := runtimeCtx.GetEntityMgr().AddEntity(entity, creator.options.Scope); err != nil {
		return nil, fmt.Errorf("add entity to runtime context failed, %v", err)
	}

	if !creator.options.ParentID.IsNil() {
		err := runtimeCtx.GetECTree().AddChild(creator.options.ParentID, entity.GetId())
		if err != nil {
			runtimeCtx.GetEntityMgr().RemoveEntity(entity.GetId())
			return nil, fmt.Errorf("add entity to ec-tree failed, %v", err)
		}
	}

	return entity, nil
}
