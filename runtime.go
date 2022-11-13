package galaxy

import (
	"github.com/galaxy-kit/galaxy-go/ec"
	"github.com/galaxy-kit/galaxy-go/internal"
	"github.com/galaxy-kit/galaxy-go/localevent"
	"github.com/galaxy-kit/galaxy-go/runtime"
	"github.com/galaxy-kit/galaxy-go/util"
)

// Runtime 运行时接口
type Runtime interface {
	internal.Running
	init(runtimeCtx runtime.Context, opts *RuntimeOptions)
	getOptions() *RuntimeOptions
	// GetRuntimeCtx 获取运行时上下文
	GetRuntimeCtx() runtime.Context
}

// NewRuntime 创建运行时
func NewRuntime(runtimeCtx runtime.Context, options ...WithRuntimeOption) Runtime {
	opts := RuntimeOptions{}
	RuntimeOption.Default()(&opts)

	for i := range options {
		options[i](&opts)
	}

	return UnsafeNewRuntime(runtimeCtx, opts)
}

func UnsafeNewRuntime(runtimeCtx runtime.Context, options RuntimeOptions) Runtime {
	if !options.Inheritor.IsNil() {
		options.Inheritor.Iface.init(runtimeCtx, &options)
		return options.Inheritor.Iface
	}

	runtime := &RuntimeBehavior{}
	runtime.init(runtimeCtx, &options)

	return runtime.opts.Inheritor.Iface
}

// RuntimeBehavior 运行时行为，在需要拓展运行时能力时，匿名嵌入至运行时结构体中
type RuntimeBehavior struct {
	opts            RuntimeOptions
	ctx             runtime.Context
	hooksMap        map[int64][3]localevent.Hook
	processQueue    chan func()
	eventUpdate     localevent.Event
	eventLateUpdate localevent.Event
}

func (_runtime *RuntimeBehavior) init(runtimeCtx runtime.Context, opts *RuntimeOptions) {
	if runtimeCtx == nil {
		panic("nil runtimeCtx")
	}

	if opts == nil {
		panic("nil opts")
	}

	_runtime.opts = *opts

	if _runtime.opts.Inheritor.IsNil() {
		_runtime.opts.Inheritor = util.NewFace[Runtime](_runtime)
	}

	_runtime.ctx = runtimeCtx
	_runtime.hooksMap = make(map[int64][3]localevent.Hook)

	_runtime.eventUpdate.Init(runtimeCtx.GetAutoRecover(), runtimeCtx.GetReportError(), localevent.EventRecursion_Disallow, runtimeCtx.GetHookCache(), _runtime)
	_runtime.eventLateUpdate.Init(runtimeCtx.GetAutoRecover(), runtimeCtx.GetReportError(), localevent.EventRecursion_Disallow, runtimeCtx.GetHookCache(), _runtime)

	if opts.EnableAutoRun {
		_runtime.opts.Inheritor.Iface.Run()
	}
}

func (_runtime *RuntimeBehavior) getOptions() *RuntimeOptions {
	return &_runtime.opts
}

// GetRuntimeCtx 获取运行时上下文
func (_runtime *RuntimeBehavior) GetRuntimeCtx() runtime.Context {
	return _runtime.ctx
}

// OnEntityMgrAddEntity 事件回调：实体管理器中添加实体
func (_runtime *RuntimeBehavior) OnEntityMgrAddEntity(entityMgr runtime.IEntityMgr, entity ec.Entity) {
	_runtime.initEntity(entity)
	_runtime.connectEntity(entity)
}

// OnEntityMgrRemoveEntity 事件回调：实体管理器中删除实体
func (_runtime *RuntimeBehavior) OnEntityMgrRemoveEntity(entityMgr runtime.IEntityMgr, entity ec.Entity) {
	_runtime.disconnectEntity(entity)
	_runtime.shutEntity(entity)
}

// OnEntityMgrEntityFirstAccessComponent 事件回调：实体管理器中的实体首次访问组件
func (_runtime *RuntimeBehavior) OnEntityMgrEntityFirstAccessComponent(entityMgr runtime.IEntityMgr, entity ec.Entity, component ec.Component) {
	comp := ec.UnsafeComponent(component)

	if comp.GetAwoke() {
		return
	}
	comp.SetAwoke(true)

	if compAwake, ok := component.(_ComponentAwake); ok {
		internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
			compAwake.Awake()
		})
	}
}

// OnEntityMgrEntityAddComponents 事件回调：实体管理器中的实体添加组件
func (_runtime *RuntimeBehavior) OnEntityMgrEntityAddComponents(entityMgr runtime.IEntityMgr, entity ec.Entity, components []ec.Component) {
	_runtime.addComponents(components)
}

// OnEntityMgrEntityRemoveComponent 事件回调：实体管理器中的实体删除组件
func (_runtime *RuntimeBehavior) OnEntityMgrEntityRemoveComponent(entityMgr runtime.IEntityMgr, entity ec.Entity, component ec.Component) {
	_runtime.removeComponent(component)
}

// OnEntityDestroySelf 事件回调：实体销毁自身
func (_runtime *RuntimeBehavior) OnEntityDestroySelf(entity ec.Entity) {
	_runtime.ctx.GetEntityMgr().RemoveEntity(entity.GetID())
}

// OnComponentDestroySelf 事件回调：组件销毁自身
func (_runtime *RuntimeBehavior) OnComponentDestroySelf(comp ec.Component) {
	comp.GetEntity().RemoveComponentByID(comp.GetID())
}

func (_runtime *RuntimeBehavior) initEntity(entity ec.Entity) {
	ec.UnsafeEntity(entity).SetInitialing(true)
	defer ec.UnsafeEntity(entity).SetInitialing(false)

	if entityInit, ok := entity.(_EntityInit); ok {
		internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
			entityInit.Init()
		})
	}

	entity.RangeComponents(func(comp ec.Component) bool {
		_comp := ec.UnsafeComponent(comp)

		if _comp.GetAwoke() {
			return true
		}
		_comp.SetAwoke(true)

		if compAwake, ok := comp.(_ComponentAwake); ok {
			internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
				compAwake.Awake()
			})
		}

		return true
	})

	entity.RangeComponents(func(comp ec.Component) bool {
		_comp := ec.UnsafeComponent(comp)

		if _comp.GetStarted() {
			return true
		}
		_comp.SetStarted(true)

		if compStart, ok := comp.(_ComponentStart); ok {
			internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
				compStart.Start()
			})
		}

		return true
	})

	if entityInitFin, ok := entity.(_EntityInitFin); ok {
		internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
			entityInitFin.InitFin()
		})
	}
}

func (_runtime *RuntimeBehavior) shutEntity(entity ec.Entity) {
	ec.UnsafeEntity(entity).SetShutting(true)
	defer ec.UnsafeEntity(entity).SetShutting(false)

	if entityShut, ok := entity.(_EntityShut); ok {
		internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
			entityShut.Shut()
		})
	}

	entity.RangeComponents(func(comp ec.Component) bool {
		if compShut, ok := comp.(_ComponentShut); ok {
			internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
				compShut.Shut()
			})
		}
		return true
	})

	if entityShutFin, ok := entity.(_EntityShutFin); ok {
		internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
			entityShutFin.ShutFin()
		})
	}
}

func (_runtime *RuntimeBehavior) connectEntity(entity ec.Entity) {
	var hooks [3]localevent.Hook

	if entityUpdate, ok := entity.(_EntityUpdate); ok {
		hooks[0] = localevent.BindEvent[_EntityUpdate](&_runtime.eventUpdate, entityUpdate)
	}

	if entityLateUpdate, ok := entity.(_EntityLateUpdate); ok {
		hooks[1] = localevent.BindEvent[_EntityLateUpdate](&_runtime.eventLateUpdate, entityLateUpdate)
	}

	hooks[2] = localevent.BindEvent[ec.EventEntityDestroySelf](ec.UnsafeEntity(entity).EventEntityDestroySelf(), _runtime)

	_runtime.hooksMap[entity.GetID()] = hooks

	entity.RangeComponents(func(comp ec.Component) bool {
		_runtime.connectComponent(comp)
		return true
	})
}

func (_runtime *RuntimeBehavior) disconnectEntity(entity ec.Entity) {
	entity.RangeComponents(func(comp ec.Component) bool {
		_runtime.disconnectComponent(comp)
		return true
	})

	hooks, ok := _runtime.hooksMap[entity.GetID()]
	if ok {
		delete(_runtime.hooksMap, entity.GetID())

		for i := range hooks {
			hooks[i].Unbind()
		}
	}
}

func (_runtime *RuntimeBehavior) addComponents(components []ec.Component) {
	for i := range components {
		_comp := ec.UnsafeComponent(components[i])

		if _comp.GetAwoke() {
			continue
		}
		_comp.SetAwoke(true)

		if compAwake, ok := components[i].(_ComponentAwake); ok {
			internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
				compAwake.Awake()
			})
		}
	}

	for i := range components {
		_comp := ec.UnsafeComponent(components[i])

		if _comp.GetStarted() {
			continue
		}
		_comp.SetStarted(true)

		if compStart, ok := components[i].(_ComponentStart); ok {
			internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
				compStart.Start()
			})
		}
	}

	for i := range components {
		_runtime.connectComponent(components[i])
	}
}

func (_runtime *RuntimeBehavior) removeComponent(component ec.Component) {
	_runtime.disconnectComponent(component)

	if compShut, ok := component.(_ComponentShut); ok {
		internal.CallOuterNoRet(_runtime.ctx.GetAutoRecover(), _runtime.ctx.GetReportError(), func() {
			compShut.Shut()
		})
	}
}

func (_runtime *RuntimeBehavior) connectComponent(comp ec.Component) {
	if _, ok := _runtime.hooksMap[comp.GetID()]; ok {
		return
	}

	var hooks [3]localevent.Hook

	if compUpdate, ok := comp.(_ComponentUpdate); ok {
		hooks[0] = localevent.BindEvent[_ComponentUpdate](&_runtime.eventUpdate, compUpdate)
	}

	if compLateUpdate, ok := comp.(_ComponentLateUpdate); ok {
		hooks[1] = localevent.BindEvent[_ComponentLateUpdate](&_runtime.eventLateUpdate, compLateUpdate)
	}

	hooks[2] = localevent.BindEvent[ec.EventComponentDestroySelf](ec.UnsafeComponent(comp).EventComponentDestroySelf(), _runtime)

	_runtime.hooksMap[comp.GetID()] = hooks
}

func (_runtime *RuntimeBehavior) disconnectComponent(comp ec.Component) {
	hooks, ok := _runtime.hooksMap[comp.GetID()]
	if ok {
		delete(_runtime.hooksMap, comp.GetID())

		for i := range hooks {
			hooks[i].Unbind()
		}
	}
}
