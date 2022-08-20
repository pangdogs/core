package core

import (
	"github.com/pangdogs/galaxy/core/container"
)

type Runtime interface {
	container.GC
	Runnable
	init(runtimeCtx RuntimeContext, opts *RuntimeOptions)
	getOptions() *RuntimeOptions
	GetID() uint64
	GetRuntimeCtx() RuntimeContext
}

func RuntimeGetOptions(runtime Runtime) RuntimeOptions {
	return *runtime.getOptions()
}

func RuntimeGetInheritor(runtime Runtime) Face[Runtime] {
	return runtime.getOptions().Inheritor
}

func RuntimeGetInheritorIFace[T any](runtime Runtime) T {
	return Cache2IFace[T](runtime.getOptions().Inheritor.Cache)
}

func NewRuntime(runtimeCtx RuntimeContext, optFuncs ...NewRuntimeOptionFunc) Runtime {
	opts := &RuntimeOptions{}
	NewRuntimeOption.Default()(opts)

	for i := range optFuncs {
		optFuncs[i](opts)
	}

	if !opts.Inheritor.IsNil() {
		opts.Inheritor.IFace.init(runtimeCtx, opts)
		return opts.Inheritor.IFace
	}

	runtime := &RuntimeBehavior{}
	runtime.init(runtimeCtx, opts)

	return runtime.opts.Inheritor.IFace
}

type RuntimeBehavior struct {
	id              uint64
	opts            RuntimeOptions
	ctx             RuntimeContext
	hooksMap        map[uint64][3]Hook
	processQueue    chan func()
	eventUpdate     Event
	eventLateUpdate Event
}

func (runtime *RuntimeBehavior) GC() {
	runtime.ctx.GC()
	runtime.eventUpdate.GC()
	runtime.eventLateUpdate.GC()
}

func (runtime *RuntimeBehavior) NeedGC() bool {
	return true
}

func (runtime *RuntimeBehavior) CollectGC(gc container.GC) {
}

func (runtime *RuntimeBehavior) init(runtimeCtx RuntimeContext, opts *RuntimeOptions) {
	if runtimeCtx == nil {
		panic("nil runtimeCtx")
	}

	if opts == nil {
		panic("nil opts")
	}

	runtime.opts = *opts

	if runtime.opts.Inheritor.IsNil() {
		runtime.opts.Inheritor = NewFace[Runtime](runtime)
	}

	runtime.id = runtimeCtx.GetServiceCtx().genUID()
	runtime.ctx = runtimeCtx
	runtime.hooksMap = make(map[uint64][3]Hook)

	runtime.eventUpdate.Init(runtime.getOptions().EnableAutoRecover, runtimeCtx.GetReportError(), EventRecursion_Disallow, runtimeCtx.getOptions().HookCache, runtime)
	runtime.eventLateUpdate.Init(runtime.getOptions().EnableAutoRecover, runtimeCtx.GetReportError(), EventRecursion_Disallow, runtimeCtx.getOptions().HookCache, runtime)

	if opts.EnableAutoRun {
		runtime.opts.Inheritor.IFace.Run()
	}
}

func (runtime *RuntimeBehavior) getOptions() *RuntimeOptions {
	return &runtime.opts
}

func (runtime *RuntimeBehavior) GetID() uint64 {
	return runtime.id
}

func (runtime *RuntimeBehavior) GetRuntimeCtx() RuntimeContext {
	return runtime.ctx
}

func (runtime *RuntimeBehavior) OnEntityMgrAddEntity(runtimeCtx RuntimeContext, entity Entity) {
	runtime.initEntity(entity)
	runtime.connectEntity(entity)
}

func (runtime *RuntimeBehavior) OnEntityMgrRemoveEntity(runtimeCtx RuntimeContext, entity Entity) {
	runtime.disconnectEntity(entity)
	runtime.shutEntity(entity)
}

func (runtime *RuntimeBehavior) OnEntityMgrEntityAddComponents(runtimeCtx RuntimeContext, entity Entity, components []Component) {
	runtime.addComponents(components)
}

func (runtime *RuntimeBehavior) OnEntityMgrEntityRemoveComponent(runtimeCtx RuntimeContext, entity Entity, component Component) {
	runtime.removeComponent(component)
}

func (runtime *RuntimeBehavior) onEntityDestroySelf(entity Entity) {
	runtime.ctx.RemoveEntity(entity.GetID())
}

func (runtime *RuntimeBehavior) onComponentDestroySelf(comp Component) {
	comp.GetEntity().RemoveComponentByID(comp.GetID())
}

func (runtime *RuntimeBehavior) initEntity(entity Entity) {
	entity.setInitialing(true)
	defer entity.setInitialing(false)

	if entityInit, ok := entity.(EntityInit); ok {
		entityInit.Init()
	}

	entity.RangeComponents(func(comp Component) bool {
		comp.setPrimary(true)
		return true
	})

	entity.RangeComponents(func(comp Component) bool {
		if !comp.getPrimary() {
			return true
		}

		if compAwake, ok := comp.(ComponentAwake); ok {
			compAwake.Awake()
		}

		return true
	})

	entity.RangeComponents(func(comp Component) bool {
		if !comp.getPrimary() {
			return true
		}

		if compStart, ok := comp.(ComponentStart); ok {
			compStart.Start()
		}

		return true
	})

	if entityInitFin, ok := entity.(EntityInitFin); ok {
		entityInitFin.InitFin()
	}
}

func (runtime *RuntimeBehavior) shutEntity(entity Entity) {
	if entityShut, ok := entity.(EntityShut); ok {
		entityShut.Shut()
	}

	entity.RangeComponents(func(comp Component) bool {
		if compShut, ok := comp.(ComponentShut); ok {
			compShut.Shut()
		}
		return true
	})

	if entityShutFin, ok := entity.(EntityShutFin); ok {
		entityShutFin.ShutFin()
	}
}

func (runtime *RuntimeBehavior) connectEntity(entity Entity) {
	var hooks [3]Hook

	if entityUpdate, ok := entity.(EntityUpdate); ok {
		hooks[0] = BindEvent[EntityUpdate](&runtime.eventUpdate, entityUpdate)
	}

	if entityLateUpdate, ok := entity.(EntityLateUpdate); ok {
		hooks[1] = BindEvent[EntityLateUpdate](&runtime.eventLateUpdate, entityLateUpdate)
	}

	hooks[2] = BindEvent[eventEntityDestroySelf](entity.eventEntityDestroySelf(), runtime)

	runtime.hooksMap[entity.GetID()] = hooks

	entity.RangeComponents(func(comp Component) bool {
		if comp.getPrimary() {
			runtime.connectComponent(comp)
		}
		return true
	})
}

func (runtime *RuntimeBehavior) disconnectEntity(entity Entity) {
	entity.RangeComponents(func(comp Component) bool {
		runtime.disconnectComponent(comp)
		return true
	})

	hooks, ok := runtime.hooksMap[entity.GetID()]
	if ok {
		delete(runtime.hooksMap, entity.GetID())

		for i := range hooks {
			hooks[i].Unbind()
		}
	}
}

func (runtime *RuntimeBehavior) addComponents(components []Component) {
	for i := range components {
		if compAwake, ok := components[i].(ComponentAwake); ok {
			compAwake.Awake()
		}
	}

	for i := range components {
		if compStart, ok := components[i].(ComponentStart); ok {
			compStart.Start()
		}
	}

	for i := range components {
		runtime.connectComponent(components[i])
	}
}

func (runtime *RuntimeBehavior) removeComponent(component Component) {
	runtime.disconnectComponent(component)

	if compShut, ok := component.(ComponentShut); ok {
		compShut.Shut()
	}
}

func (runtime *RuntimeBehavior) connectComponent(comp Component) {
	var hooks [3]Hook

	if compUpdate, ok := comp.(ComponentUpdate); ok {
		hooks[0] = BindEvent[ComponentUpdate](&runtime.eventUpdate, compUpdate)
	}

	if compLateUpdate, ok := comp.(ComponentLateUpdate); ok {
		hooks[1] = BindEvent[ComponentLateUpdate](&runtime.eventLateUpdate, compLateUpdate)
	}

	hooks[2] = BindEvent[eventComponentDestroySelf](comp.eventComponentDestroySelf(), runtime)

	runtime.hooksMap[comp.GetID()] = hooks
}

func (runtime *RuntimeBehavior) disconnectComponent(comp Component) {
	hooks, ok := runtime.hooksMap[comp.GetID()]
	if ok {
		delete(runtime.hooksMap, comp.GetID())

		for i := range hooks {
			hooks[i].Unbind()
		}
	}
}
