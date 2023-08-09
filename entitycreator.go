package golaxy

import (
	"errors"
	"fmt"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/service"
)

// EntityCreator 实体构建器
type EntityCreator struct {
	Context runtime.Context      // 运行时上下文
	options EntityCreatorOptions // 实体构建器的所有选项
	mutable bool                 // 是否已改变选项
}

// Options 创建实体的选项
func (creator EntityCreator) Options(options ...EntityCreatorOption) EntityCreator {
	if !creator.mutable {
		Option{}.EntityCreator.Default()(&creator.options)
		if creator.Context != nil {
			creator.options.FaceAnyAllocator = creator.Context.GetFaceAnyAllocator()
			creator.options.HookAllocator = creator.Context.GetHookAllocator()
		}
		creator.mutable = true
	}
	for i := range options {
		options[i](&creator.options)
	}
	return creator
}

// Spawn 创建实体
func (creator EntityCreator) Spawn() (ec.Entity, error) {
	if creator.Context == nil {
		return nil, errors.New("nil context")
	}

	runtimeCtx := creator.Context
	serviceCtx := service.Get(runtimeCtx)

	if !creator.mutable {
		Option{}.EntityCreator.Default()(&creator.options)
		creator.options.FaceAnyAllocator = runtimeCtx.GetFaceAnyAllocator()
		creator.options.HookAllocator = runtimeCtx.GetHookAllocator()
	}

	entityPt, ok := pt.TryGetEntityPt(serviceCtx, creator.options.Prototype)
	if !ok {
		return nil, fmt.Errorf("entity prototype %q not registered", creator.options.Prototype)
	}

	entity := entityPt.UnsafeConstruct(creator.options.ConstructEntityOptions)

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
