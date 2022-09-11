package core

// Runtime 运行时接口
type Runtime interface {
	_Runnable

	init(runtimeCtx RuntimeContext, opts *RuntimeOptions)

	getOptions() *RuntimeOptions

	// GetRuntimeCtx 获取运行时上下文（Runtime Context），线程安全
	GetRuntimeCtx() RuntimeContext
}

// RuntimeGetOptions 获取运行时创建选项，线程安全
func RuntimeGetOptions(runtime Runtime) RuntimeOptions {
	return *runtime.getOptions()
}

// NewRuntime 创建运行时，线程安全
func NewRuntime(runtimeCtx RuntimeContext, optSetterFuncs ...RuntimeOptionSetterFunc) Runtime {
	opts := RuntimeOptions{}
	RuntimeOptionSetter.Default()(&opts)

	for i := range optSetterFuncs {
		optSetterFuncs[i](&opts)
	}

	return NewRuntimeWithOpts(runtimeCtx, opts)
}

// NewRuntimeWithOpts 创建运行时并传入参数，线程安全
func NewRuntimeWithOpts(runtimeCtx RuntimeContext, opts RuntimeOptions) Runtime {
	if !opts.Inheritor.IsNil() {
		opts.Inheritor.Iface.init(runtimeCtx, &opts)
		return opts.Inheritor.Iface
	}

	runtime := &_RuntimeBehavior{}
	runtime.init(runtimeCtx, &opts)

	return runtime.opts.Inheritor.Iface
}

type _RuntimeBehavior struct {
	opts            RuntimeOptions
	ctx             RuntimeContext
	hooksMap        map[int64][3]Hook
	processQueue    chan func()
	eventUpdate     Event
	eventLateUpdate Event
}

func (runtime *_RuntimeBehavior) init(runtimeCtx RuntimeContext, opts *RuntimeOptions) {
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

	runtime.ctx = runtimeCtx
	runtime.hooksMap = make(map[int64][3]Hook)

	runtime.eventUpdate.Init(runtime.getOptions().EnableAutoRecover, runtimeCtx.GetReportError(), EventRecursion_Disallow, runtimeCtx.getOptions().HookCache, runtime)
	runtime.eventLateUpdate.Init(runtime.getOptions().EnableAutoRecover, runtimeCtx.GetReportError(), EventRecursion_Disallow, runtimeCtx.getOptions().HookCache, runtime)

	if opts.EnableAutoRun {
		runtime.opts.Inheritor.Iface.Run()
	}
}

func (runtime *_RuntimeBehavior) getOptions() *RuntimeOptions {
	return &runtime.opts
}

// GetRuntimeCtx 获取运行时上下文（Runtime Context），线程安全
func (runtime *_RuntimeBehavior) GetRuntimeCtx() RuntimeContext {
	return runtime.ctx
}

// OnEntityMgrAddEntity 事件回调：运行时上下文（Runtime Context）添加实体（Entity）
func (runtime *_RuntimeBehavior) OnEntityMgrAddEntity(runtimeCtx RuntimeContext, entity Entity) {
	runtime.initEntity(entity)
	runtime.connectEntity(entity)
}

// OnEntityMgrRemoveEntity 事件回调：运行时上下文（Runtime Context）删除实体（Entity）
func (runtime *_RuntimeBehavior) OnEntityMgrRemoveEntity(runtimeCtx RuntimeContext, entity Entity) {
	runtime.disconnectEntity(entity)
	runtime.shutEntity(entity)
}

// OnEntityMgrEntityAddComponents 事件回调：运行时上下文（Runtime Context）中的实体（Entity）添加组件（Component）
func (runtime *_RuntimeBehavior) OnEntityMgrEntityAddComponents(runtimeCtx RuntimeContext, entity Entity, components []Component) {
	runtime.addComponents(components)
}

// OnEntityMgrEntityRemoveComponent 事件回调：运行时上下文（Runtime Context）中的实体（Entity）删除组件（Component）
func (runtime *_RuntimeBehavior) OnEntityMgrEntityRemoveComponent(runtimeCtx RuntimeContext, entity Entity, component Component) {
	runtime.removeComponent(component)
}

func (runtime *_RuntimeBehavior) onEntityDestroySelf(entity Entity) {
	runtime.ctx.RemoveEntity(entity.GetID())
}

func (runtime *_RuntimeBehavior) onComponentDestroySelf(comp Component) {
	comp.GetEntity().RemoveComponentByID(comp.GetID())
}

func (runtime *_RuntimeBehavior) initEntity(entity Entity) {
	entity.setInitialing(true)
	defer entity.setInitialing(false)

	if entityInit, ok := entity.(_EntityInit); ok {
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

		if compAwake, ok := comp.(_ComponentAwake); ok {
			compAwake.Awake()
		}

		return true
	})

	entity.RangeComponents(func(comp Component) bool {
		if !comp.getPrimary() {
			return true
		}

		if compStart, ok := comp.(_ComponentStart); ok {
			compStart.Start()
		}

		return true
	})

	if entityInitFin, ok := entity.(_EntityInitFin); ok {
		entityInitFin.InitFin()
	}
}

func (runtime *_RuntimeBehavior) shutEntity(entity Entity) {
	if entityShut, ok := entity.(_EntityShut); ok {
		entityShut.Shut()
	}

	entity.RangeComponents(func(comp Component) bool {
		if compShut, ok := comp.(_ComponentShut); ok {
			compShut.Shut()
		}
		return true
	})

	if entityShutFin, ok := entity.(_EntityShutFin); ok {
		entityShutFin.ShutFin()
	}
}

func (runtime *_RuntimeBehavior) connectEntity(entity Entity) {
	var hooks [3]Hook

	if entityUpdate, ok := entity.(_EntityUpdate); ok {
		hooks[0] = BindEvent[_EntityUpdate](&runtime.eventUpdate, entityUpdate)
	}

	if entityLateUpdate, ok := entity.(_EntityLateUpdate); ok {
		hooks[1] = BindEvent[_EntityLateUpdate](&runtime.eventLateUpdate, entityLateUpdate)
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

func (runtime *_RuntimeBehavior) disconnectEntity(entity Entity) {
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

func (runtime *_RuntimeBehavior) addComponents(components []Component) {
	for i := range components {
		if compAwake, ok := components[i].(_ComponentAwake); ok {
			compAwake.Awake()
		}
	}

	for i := range components {
		if compStart, ok := components[i].(_ComponentStart); ok {
			compStart.Start()
		}
	}

	for i := range components {
		runtime.connectComponent(components[i])
	}
}

func (runtime *_RuntimeBehavior) removeComponent(component Component) {
	runtime.disconnectComponent(component)

	if compShut, ok := component.(_ComponentShut); ok {
		compShut.Shut()
	}
}

func (runtime *_RuntimeBehavior) connectComponent(comp Component) {
	var hooks [3]Hook

	if compUpdate, ok := comp.(_ComponentUpdate); ok {
		hooks[0] = BindEvent[_ComponentUpdate](&runtime.eventUpdate, compUpdate)
	}

	if compLateUpdate, ok := comp.(_ComponentLateUpdate); ok {
		hooks[1] = BindEvent[_ComponentLateUpdate](&runtime.eventLateUpdate, compLateUpdate)
	}

	hooks[2] = BindEvent[eventComponentDestroySelf](comp.eventComponentDestroySelf(), runtime)

	runtime.hooksMap[comp.GetID()] = hooks
}

func (runtime *_RuntimeBehavior) disconnectComponent(comp Component) {
	hooks, ok := runtime.hooksMap[comp.GetID()]
	if ok {
		delete(runtime.hooksMap, comp.GetID())

		for i := range hooks {
			hooks[i].Unbind()
		}
	}
}
