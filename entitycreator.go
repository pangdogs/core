package golaxy

import (
	"fmt"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/pt"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/option"
)

// CreateEntity 创建实体
func CreateEntity(ctxProvider runtime.CurrentContextProvider, settings ...option.Setting[EntityCreatorOptions]) EntityCreator {
	return EntityCreator{
		Context: runtime.Current(ctxProvider),
		Options: option.Make(_EntityCreatorOption{}.Default(), settings...),
	}
}

// EntityCreator 实体构建器
type EntityCreator struct {
	Context runtime.Context      // 运行时上下文
	Options EntityCreatorOptions // 实体构建器的所有选项
}

// Spawn 创建实体
func (creator EntityCreator) Spawn(settings ...option.Setting[EntityCreatorOptions]) (ec.Entity, error) {
	if creator.Context == nil {
		panic(fmt.Errorf("%w: setting context is nil", ErrGolaxy))
	}

	rtCtx := creator.Context
	servCtx := service.Current(rtCtx)

	opts := option.Append(creator.Options, settings...)

	entityPt, err := pt.Using(servCtx, opts.Prototype)
	if err != nil {
		return nil, err
	}

	if !opts.ParentID.IsNil() {
		_, err := runtime.UnsafeECTree(rtCtx.GetECTree()).FetchEntity(opts.ParentID)
		if err != nil {
			return nil, err
		}
	}

	entity := entityPt.UnsafeConstruct(opts.ConstructEntityOptions)

	if err := rtCtx.GetEntityMgr().AddEntity(entity, opts.Scope); err != nil {
		return nil, err
	}

	if !opts.ParentID.IsNil() {
		if err := rtCtx.GetECTree().AddChild(opts.ParentID, entity.GetId()); err != nil {
			entity.DestroySelf()
			return nil, err
		}
	}

	return entity, nil
}
