package golaxy

import (
	"fmt"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/service"
)

var (
	ErrEntityCreator = fmt.Errorf("%w: entity-creator", internal.ErrGolaxy)
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
		return nil, fmt.Errorf("%w: context is nil", ErrEntityCreator)
	}

	runtimeCtx := creator.Context
	serviceCtx := service.Current(runtimeCtx)

	creator.Options()

	opts := &creator.options
	if len(options) > 0 {
		copyOpts := creator.options
		for i := range options {
			options[i](&copyOpts)
		}
		opts = &copyOpts
	}

	entityPt, ok := pt.AccessEntityPt(serviceCtx, opts.Prototype)
	if !ok {
		return nil, fmt.Errorf("%w: entity %q not registered", ErrEntityCreator, opts.Prototype)
	}

	entity := entityPt.UnsafeConstruct(opts.ConstructEntityOptions)

	if err := runtimeCtx.GetEntityMgr().AddEntity(entity, opts.Scope); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrEntityCreator, err)
	}

	if !opts.ParentID.IsNil() {
		err := runtimeCtx.GetECTree().AddChild(opts.ParentID, entity.GetId())
		if err != nil {
			runtimeCtx.GetEntityMgr().RemoveEntity(entity.GetId())
			return nil, fmt.Errorf("%w, %w", ErrEntityCreator, err)
		}
	}

	return entity, nil
}
