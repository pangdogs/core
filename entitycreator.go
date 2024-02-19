package core

import (
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/pt"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/util/option"
)

// CreateEntity 创建实体
func CreateEntity(ctxProvider runtime.CurrentContextProvider, settings ...option.Setting[EntityCreatorOptions]) EntityCreator {
	return EntityCreator{
		rtCtx:   runtime.Current(ctxProvider),
		options: option.Make(With.EntityCreator.Default(), settings...),
	}
}

// EntityCreator 实体构建器
type EntityCreator struct {
	rtCtx   runtime.Context      // 运行时上下文
	options EntityCreatorOptions // 实体构建器的所有选项
}

// Spawn 创建实体
func (creator EntityCreator) Spawn(settings ...option.Setting[EntityCreatorOptions]) (ec.Entity, error) {
	if creator.rtCtx == nil {
		panic(fmt.Errorf("%w: setting rtCtx is nil", ErrGolaxy))
	}

	ctx := creator.rtCtx
	opts := option.Append(creator.options, settings...)

	if !opts.ParentID.IsNil() {
		_, err := runtime.UnsafeECTree(ctx.GetECTree()).FetchEntity(opts.ParentID)
		if err != nil {
			return nil, err
		}
	}

	entity := pt.Using(service.Current(ctx), opts.Prototype).UnsafeConstruct(opts.ConstructEntityOptions)

	if err := ctx.GetEntityMgr().AddEntity(entity); err != nil {
		return nil, err
	}

	if !opts.ParentID.IsNil() {
		if err := ctx.GetECTree().AddChild(opts.ParentID, entity.GetId()); err != nil {
			entity.DestroySelf()
			return nil, err
		}
	}

	return entity, nil
}
