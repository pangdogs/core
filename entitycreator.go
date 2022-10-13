package galaxy

import (
	"errors"
	"fmt"
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/runtime"
	"github.com/pangdogs/galaxy/service"
)

func EntityCreator() _EntityCreator {
	return _EntityCreator{}
}

type _EntityCreator struct {
	runtimeCtx runtime.Context
	prototype  string
	optSetters []ec.EntityOptionSetter
}

func (creator _EntityCreator) RuntimeCtx(ctx runtime.Context) _EntityCreator {
	creator.runtimeCtx = ctx
	return creator
}

func (creator _EntityCreator) Prototype(prototype string) _EntityCreator {
	creator.prototype = prototype
	return creator
}

func (creator _EntityCreator) Options(optSetter ...ec.EntityOptionSetter) _EntityCreator {
	creator.optSetters = optSetter
	return creator
}

func (creator _EntityCreator) Build() (ec.Entity, error) {
	if creator.runtimeCtx == nil {
		return nil, errors.New("RuntimeCtx has not been setup")
	}

	runtimeCtx := creator.runtimeCtx
	serviceCtx := runtimeCtx.GetServiceCtx()

	entityLib := service.UnsafeContext(serviceCtx).GetOptions().EntityLib
	if entityLib == nil {
		return nil, errors.New("ServiceCtx option EntityLib has not been setup")
	}

	entityPt, ok := entityLib.Get(creator.prototype)
	if !ok {
		return nil, fmt.Errorf("entity '%s' not registered", creator.prototype)
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
	entityPt.AddComponents(entity)

	if err := runtimeCtx.GetEntityMgr().AddEntity(entity); err != nil {
		panic(err)
	}

	_, loaded, err := serviceCtx.GetEntityMgr().GetOrAddEntity(entity)
	if loaded {

	}

	return entity
}
