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

// Clone 克隆
func (creator EntityCreator) Clone() *EntityCreator {
	return &creator
}

// Options 创建实体的选项
func (creator *EntityCreator) Options(options ...EntityCreatorOption) *EntityCreator {
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
func (creator *EntityCreator) Spawn(options ...EntityCreatorOption) (ec.Entity, error) {
	if creator.Context == nil {
		return nil, errors.New("nil context")
	}

	runtimeCtx := creator.Context
	serviceCtx := service.Get(runtimeCtx)

	creator.Options()

	opts := &creator.options
	if len(options) > 0 {
		copyOpts := creator.options
		for i := range options {
			options[i](&copyOpts)
		}
		opts = &copyOpts
	}

	entityPt, ok := pt.TryGetEntityPt(serviceCtx, opts.Prototype)
	if !ok {
		return nil, fmt.Errorf("entity prototype %q not registered", opts.Prototype)
	}

	entity := entityPt.UnsafeConstruct(opts.ConstructEntityOptions)

	if err := runtimeCtx.GetEntityMgr().AddEntity(entity, opts.Scope); err != nil {
		return nil, fmt.Errorf("add entity to runtime context failed, %v", err)
	}

	if !opts.ParentID.IsNil() {
		err := runtimeCtx.GetECTree().AddChild(opts.ParentID, entity.GetId())
		if err != nil {
			runtimeCtx.GetEntityMgr().RemoveEntity(entity.GetId())
			return nil, fmt.Errorf("add entity to ec-tree failed, %v", err)
		}
	}

	return entity, nil
}
