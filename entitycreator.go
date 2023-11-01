package golaxy

import (
	"fmt"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/option"
)

var (
	ErrEntityCreator = fmt.Errorf("%w: entity-creator", internal.ErrGolaxy)
)

// EntityCreator 实体构建器
type EntityCreator struct {
	Context runtime.Context      // 运行时上下文
	Options EntityCreatorOptions // 实体构建器的所有选项
}

// Spawn 创建实体
func (creator EntityCreator) Spawn(settings ...option.Setting[EntityCreatorOptions]) (ec.Entity, error) {
	if creator.Context == nil {
		panic(fmt.Errorf("%w: setting Context is nil", ErrEntityCreator))
	}

	runtimeCtx := creator.Context
	serviceCtx := service.Current(runtimeCtx)

	opts := option.Append(creator.Options, settings...)

	entityPt, err := pt.Using(serviceCtx, opts.Prototype)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrEntityCreator, err)
	}

	entity := entityPt.UnsafeConstruct(opts.ConstructEntityOptions)

	if err := runtimeCtx.GetEntityMgr().AddEntity(entity, opts.Scope); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrEntityCreator, err)
	}

	if !opts.ParentID.IsNil() {
		if err := runtimeCtx.GetECTree().AddChild(opts.ParentID, entity.GetId()); err != nil {
			entity.DestroySelf()
			return nil, fmt.Errorf("%w, %w", ErrEntityCreator, err)
		}
	}

	return entity, nil
}
