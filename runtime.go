/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

package core

import (
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/internal/ictx"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/reinterpret"
	"git.golaxy.org/core/utils/uid"
)

// NewRuntime 创建运行时
func NewRuntime(rtCtx runtime.Context, settings ...option.Setting[RuntimeOptions]) Runtime {
	return UnsafeNewRuntime(rtCtx, option.Make(With.Runtime.Default(), settings...))
}

// Deprecated: UnsafeNewRuntime 内部创建运行时
func UnsafeNewRuntime(rtCtx runtime.Context, options RuntimeOptions) Runtime {
	if !options.InstanceFace.IsNil() {
		options.InstanceFace.Iface.init(rtCtx, options)
		return options.InstanceFace.Iface
	}

	runtime := &RuntimeBehavior{}
	runtime.init(rtCtx, options)

	return runtime.opts.InstanceFace.Iface
}

// Runtime 运行时接口
type Runtime interface {
	iRuntime
	iRunning
	ictx.CurrentContextProvider
	ictx.ConcurrentContextProvider
	reinterpret.InstanceProvider
	async.Callee
}

type iRuntime interface {
	init(rtCtx runtime.Context, opts RuntimeOptions)
	getOptions() *RuntimeOptions
}

// RuntimeBehavior 运行时行为，在扩展运行时能力时，匿名嵌入至运行时结构体中
type RuntimeBehavior struct {
	ctx                                          runtime.Context
	opts                                         RuntimeOptions
	hooksMap                                     map[uid.Id][3]event.Hook
	processQueue                                 chan _Task
	eventUpdate                                  event.Event
	eventLateUpdate                              event.Event
	eventRuntimeRunningStatusChanged             event.Event
	handleEntityManagerAddEntity                 runtime.EventEntityManagerAddEntity
	handleEntityManagerRemoveEntity              runtime.EventEntityManagerRemoveEntity
	handleEntityManagerEntityFirstTouchComponent runtime.EventEntityManagerEntityFirstTouchComponent
	handleEntityManagerEntityAddComponents       runtime.EventEntityManagerEntityAddComponents
	handleEntityManagerEntityRemoveComponent     runtime.EventEntityManagerEntityRemoveComponent
	handleEventEntityDestroySelf                 ec.EventEntityDestroySelfHandler
	handleEventComponentDestroySelf              ec.EventComponentDestroySelf
}

// GetCurrentContext 获取当前上下文
func (rt *RuntimeBehavior) GetCurrentContext() iface.Cache {
	return rt.ctx.GetCurrentContext()
}

// GetConcurrentContext 获取多线程安全的上下文
func (rt *RuntimeBehavior) GetConcurrentContext() iface.Cache {
	return rt.ctx.GetConcurrentContext()
}

// GetInstanceFaceCache 支持重新解释类型
func (rt *RuntimeBehavior) GetInstanceFaceCache() iface.Cache {
	return rt.opts.InstanceFace.Cache
}

func (rt *RuntimeBehavior) init(rtCtx runtime.Context, opts RuntimeOptions) {
	if rtCtx == nil {
		exception.Panicf("%w: %w: rtCtx is nil", ErrRuntime, ErrArgs)
	}

	if !ictx.UnsafeContext(rtCtx).SetPaired(true) {
		exception.Panicf("%w: context already paired", ErrRuntime)
	}

	rt.ctx = rtCtx
	rt.opts = opts

	if rt.opts.InstanceFace.IsNil() {
		rt.opts.InstanceFace = iface.MakeFaceT[Runtime](rt)
	}

	rt.hooksMap = make(map[uid.Id][3]event.Hook)
	rt.processQueue = make(chan _Task, rt.opts.ProcessQueueCapacity)

	runtime.UnsafeContext(rtCtx).SetFrame(rt.opts.Frame)
	runtime.UnsafeContext(rtCtx).SetCallee(rt.opts.InstanceFace.Iface)

	rtCtx.ActivateEvent(&rt.eventUpdate, event.EventRecursion_Disallow)
	rtCtx.ActivateEvent(&rt.eventLateUpdate, event.EventRecursion_Disallow)
	rtCtx.ActivateEvent(&rt.eventRuntimeRunningStatusChanged, event.EventRecursion_Allow)

	rt.handleEntityManagerAddEntity = runtime.HandleEventEntityManagerAddEntity(rt.onEntityManagerAddEntity)
	rt.handleEntityManagerRemoveEntity = runtime.HandleEventEntityManagerRemoveEntity(rt.onEntityManagerRemoveEntity)
	rt.handleEntityManagerEntityFirstTouchComponent = runtime.HandleEventEntityManagerEntityFirstTouchComponent(rt.onEntityManagerEntityFirstTouchComponent)
	rt.handleEntityManagerEntityAddComponents = runtime.HandleEventEntityManagerEntityAddComponents(rt.onEntityManagerEntityAddComponents)
	rt.handleEntityManagerEntityRemoveComponent = runtime.HandleEventEntityManagerEntityRemoveComponent(rt.onEntityManagerEntityRemoveComponent)
	rt.handleEventEntityDestroySelf = ec.HandleEventEntityDestroySelf(rt.onEntityDestroySelf)
	rt.handleEventComponentDestroySelf = ec.HandleEventComponentDestroySelf(rt.onComponentDestroySelf)

	rt.changeRunningStatus(runtime.RunningStatus_Birth)

	if rt.opts.AutoRun {
		rt.opts.InstanceFace.Iface.Run()
	}
}

func (rt *RuntimeBehavior) getOptions() *RuntimeOptions {
	return &rt.opts
}

// onEntityManagerAddEntity 事件处理器：实体管理器添加实体
func (rt *RuntimeBehavior) onEntityManagerAddEntity(entityManager runtime.EntityManager, entity ec.Entity) {
	rt.activateEntity(entity)
	rt.initEntity(entity)
}

// onEntityManagerRemoveEntity 事件处理器：实体管理器删除实体
func (rt *RuntimeBehavior) onEntityManagerRemoveEntity(entityManager runtime.EntityManager, entity ec.Entity) {
	rt.deactivateEntity(entity)
	rt.shutEntity(entity)
}

// onEntityManagerEntityFirstTouchComponent 事件处理器：实体管理器中的实体首次访问组件
func (rt *RuntimeBehavior) onEntityManagerEntityFirstTouchComponent(entityManager runtime.EntityManager, entity ec.Entity, component ec.Component) {
	if component.GetState() != ec.ComponentState_Attach {
		return
	}

	ec.UnsafeComponent(component).SetState(ec.ComponentState_Awake)

	if cb, ok := component.(LifecycleComponentAwake); ok {
		generic.CastAction0(cb.Awake).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}

	ec.UnsafeComponent(component).SetState(ec.ComponentState_Start)
}

// onEntityManagerEntityAddComponents 事件处理器：实体管理器中的实体添加组件
func (rt *RuntimeBehavior) onEntityManagerEntityAddComponents(entityManager runtime.EntityManager, entity ec.Entity, components []ec.Component) {
	rt.addComponents(entity, components)
}

// onEntityManagerEntityRemoveComponent 事件处理器：实体管理器中的实体删除组件
func (rt *RuntimeBehavior) onEntityManagerEntityRemoveComponent(entityManager runtime.EntityManager, entity ec.Entity, component ec.Component) {
	rt.deactivateComponent(component)
	rt.removeComponent(component)
}

// onEntityDestroySelf 事件处理器：实体销毁自身
func (rt *RuntimeBehavior) onEntityDestroySelf(entity ec.Entity) {
	rt.ctx.GetEntityManager().RemoveEntity(entity.GetId())
}

// onComponentDestroySelf 事件处理器：组件销毁自身
func (rt *RuntimeBehavior) onComponentDestroySelf(comp ec.Component) {
	comp.GetEntity().RemoveComponentById(comp.GetId())
}

func (rt *RuntimeBehavior) addComponents(entity ec.Entity, components []ec.Component) {
	switch entity.GetState() {
	case ec.EntityState_Awake, ec.EntityState_Start, ec.EntityState_Alive:
	default:
		return
	}

	for i := range components {
		rt.activateComponent(components[i])
	}

	for i := range components {
		comp := components[i]

		if comp.GetState() != ec.ComponentState_Awake {
			continue
		}

		if cb, ok := comp.(LifecycleComponentAwake); ok {
			generic.CastAction0(cb.Awake).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
		}

		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Start)
	}

	switch entity.GetState() {
	case ec.EntityState_Awake, ec.EntityState_Start, ec.EntityState_Alive:
	default:
		return
	}

	for i := range components {
		comp := components[i]

		if comp.GetState() != ec.ComponentState_Start {
			continue
		}

		if cb, ok := comp.(LifecycleComponentStart); ok {
			generic.CastAction0(cb.Start).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
		}

		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Alive)
	}
}

func (rt *RuntimeBehavior) removeComponent(component ec.Component) {
	if component.GetState() != ec.ComponentState_Shut {
		return
	}

	if cb, ok := component.(LifecycleComponentShut); ok {
		generic.CastAction0(cb.Shut).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}

	ec.UnsafeComponent(component).SetState(ec.ComponentState_Death)

	if cb, ok := component.(LifecycleComponentDispose); ok {
		generic.CastAction0(cb.Dispose).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}

	ec.UnsafeComponent(component).CleanManagedHooks()
}

func (rt *RuntimeBehavior) activateEntity(entity ec.Entity) {
	if entity.GetState() != ec.EntityState_Enter {
		return
	}

	var hooks [3]event.Hook

	if cb, ok := entity.(LifecycleEntityUpdate); ok {
		hooks[0] = event.Bind[LifecycleEntityUpdate](&rt.eventUpdate, cb)
	}
	if cb, ok := entity.(LifecycleEntityLateUpdate); ok {
		hooks[1] = event.Bind[LifecycleEntityLateUpdate](&rt.eventLateUpdate, cb)
	}
	hooks[2] = ec.BindEventEntityDestroySelf(entity, rt.handleEventEntityDestroySelf)

	rt.hooksMap[entity.GetId()] = hooks

	entity.RangeComponents(func(comp ec.Component) bool {
		rt.activateComponent(comp)
		return true
	})

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Awake)
}

func (rt *RuntimeBehavior) deactivateEntity(entity ec.Entity) {
	entityId := entity.GetId()

	hooks, ok := rt.hooksMap[entityId]
	if ok {
		delete(rt.hooksMap, entityId)
		event.Clean(hooks[:])
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Shut)

	entity.RangeComponents(func(comp ec.Component) bool {
		rt.deactivateComponent(comp)
		return true
	})
}

func (rt *RuntimeBehavior) activateComponent(comp ec.Component) {
	if comp.GetState() != ec.ComponentState_Attach {
		return
	}

	var hooks [3]event.Hook
	bound := false

	if cb, ok := comp.(LifecycleComponentUpdate); ok {
		hooks[0] = event.Bind[LifecycleComponentUpdate](&rt.eventUpdate, cb)
		bound = true
	}
	if cb, ok := comp.(LifecycleComponentLateUpdate); ok {
		hooks[1] = event.Bind[LifecycleComponentLateUpdate](&rt.eventLateUpdate, cb)
		bound = true
	}
	if !comp.GetNonRemovable() {
		hooks[2] = ec.BindEventComponentDestroySelf(comp, rt.handleEventComponentDestroySelf)
		bound = true
	}

	if bound {
		rt.hooksMap[comp.GetId()] = hooks
	}

	ec.UnsafeComponent(comp).SetState(ec.ComponentState_Awake)
}

func (rt *RuntimeBehavior) deactivateComponent(comp ec.Component) {
	compId := comp.GetId()

	hooks, ok := rt.hooksMap[compId]
	if ok {
		delete(rt.hooksMap, compId)
		event.Clean(hooks[:])
	}

	ec.UnsafeComponent(comp).SetState(ec.ComponentState_Shut)
}

func (rt *RuntimeBehavior) initEntity(entity ec.Entity) {
	if entity.GetState() != ec.EntityState_Awake {
		return
	}

	if cb, ok := entity.(LifecycleEntityAwake); ok {
		generic.CastAction0(cb.Awake).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}

	if entity.GetState() != ec.EntityState_Awake {
		return
	}

	entity.RangeComponents(func(comp ec.Component) bool {
		if comp.GetState() != ec.ComponentState_Awake {
			return true
		}

		if cb, ok := comp.(LifecycleComponentAwake); ok {
			generic.CastAction0(cb.Awake).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
		}

		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Start)

		return entity.GetState() == ec.EntityState_Awake
	})

	if entity.GetState() != ec.EntityState_Awake {
		return
	}

	entity.RangeComponents(func(comp ec.Component) bool {
		if comp.GetState() != ec.ComponentState_Start {
			return true
		}

		if cb, ok := comp.(LifecycleComponentStart); ok {
			generic.CastAction0(cb.Start).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
		}

		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Alive)

		return entity.GetState() == ec.EntityState_Awake
	})

	if entity.GetState() != ec.EntityState_Awake {
		return
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Start)

	if cb, ok := entity.(LifecycleEntityStart); ok {
		generic.CastAction0(cb.Start).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}

	if entity.GetState() != ec.EntityState_Start {
		return
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Alive)
}

func (rt *RuntimeBehavior) shutEntity(entity ec.Entity) {
	if entity.GetState() != ec.EntityState_Shut {
		return
	}

	if cb, ok := entity.(LifecycleEntityShut); ok {
		generic.CastAction0(cb.Shut).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}

	entity.RangeComponents(func(comp ec.Component) bool {
		if comp.GetState() != ec.ComponentState_Shut {
			return true
		}

		if cb, ok := comp.(LifecycleComponentShut); ok {
			generic.CastAction0(cb.Shut).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
		}

		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Death)

		return true
	})

	entity.RangeComponents(func(comp ec.Component) bool {
		if comp.GetState() != ec.ComponentState_Death {
			return true
		}

		if cb, ok := comp.(LifecycleComponentDispose); ok {
			generic.CastAction0(cb.Dispose).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
		}

		ec.UnsafeComponent(comp).CleanManagedHooks()

		return true
	})

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Death)

	if cb, ok := entity.(LifecycleEntityDispose); ok {
		generic.CastAction0(cb.Dispose).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}

	ec.UnsafeEntity(entity).CleanManagedHooks()
}
