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
	ctx                             runtime.Context
	opts                            RuntimeOptions
	hooksMap                        map[uid.Id][3]event.Hook
	processQueue                    chan _Task
	eventUpdate                     event.Event
	eventLateUpdate                 event.Event
	eventRuntimeRunningStateChanged event.Event
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
	rtCtx.ActivateEvent(&rt.eventRuntimeRunningStateChanged, event.EventRecursion_Allow)

	rt.changeRunningState(runtime.RunningState_Birth)

	if rt.opts.AutoRun {
		rt.opts.InstanceFace.Iface.Run()
	}
}

func (rt *RuntimeBehavior) getOptions() *RuntimeOptions {
	return &rt.opts
}

// OnEntityManagerAddEntity 事件处理器：实体管理器添加实体
func (rt *RuntimeBehavior) OnEntityManagerAddEntity(entityManager runtime.EntityManager, entity ec.Entity) {
	rt.activateEntity(entity)
	rt.initEntity(entity)
}

// OnEntityManagerRemoveEntity 事件处理器：实体管理器删除实体
func (rt *RuntimeBehavior) OnEntityManagerRemoveEntity(entityManager runtime.EntityManager, entity ec.Entity) {
	rt.deactivateEntity(entity)
	rt.shutEntity(entity)
}

// OnEntityManagerEntityFirstTouchComponent 事件处理器：实体管理器中的实体首次访问组件
func (rt *RuntimeBehavior) OnEntityManagerEntityFirstTouchComponent(entityManager runtime.EntityManager, entity ec.Entity, component ec.Component) {
	if component.GetState() != ec.ComponentState_Attach {
		return
	}

	ec.UnsafeComponent(component).SetState(ec.ComponentState_Awake)

	if compAwake, ok := component.(LifecycleComponentAwake); ok {
		generic.CastAction0(compAwake.Awake).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}

	ec.UnsafeComponent(component).SetState(ec.ComponentState_Start)
}

// OnEntityManagerEntityAddComponents 事件处理器：实体管理器中的实体添加组件
func (rt *RuntimeBehavior) OnEntityManagerEntityAddComponents(entityManager runtime.EntityManager, entity ec.Entity, components []ec.Component) {
	rt.addComponents(entity, components)
}

// OnEntityManagerEntityRemoveComponent 事件处理器：实体管理器中的实体删除组件
func (rt *RuntimeBehavior) OnEntityManagerEntityRemoveComponent(entityManager runtime.EntityManager, entity ec.Entity, component ec.Component) {
	rt.deactivateComponent(component)
	rt.removeComponent(component)
}

// OnEntityDestroySelf 事件处理器：实体销毁自身
func (rt *RuntimeBehavior) OnEntityDestroySelf(entity ec.Entity) {
	rt.ctx.GetEntityManager().RemoveEntity(entity.GetId())
}

// OnComponentDestroySelf 事件处理器：组件销毁自身
func (rt *RuntimeBehavior) OnComponentDestroySelf(comp ec.Component) {
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

		if compAwake, ok := comp.(LifecycleComponentAwake); ok {
			generic.CastAction0(compAwake.Awake).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
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

		if compStart, ok := comp.(LifecycleComponentStart); ok {
			generic.CastAction0(compStart.Start).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
		}

		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Alive)
	}
}

func (rt *RuntimeBehavior) removeComponent(component ec.Component) {
	if component.GetState() != ec.ComponentState_Shut {
		return
	}

	if compShut, ok := component.(LifecycleComponentShut); ok {
		generic.CastAction0(compShut.Shut).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}

	ec.UnsafeComponent(component).SetState(ec.ComponentState_Death)

	if compDispose, ok := component.(LifecycleComponentDispose); ok {
		generic.CastAction0(compDispose.Dispose).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}

	ec.UnsafeComponent(component).CleanManagedHooks()
}

func (rt *RuntimeBehavior) activateEntity(entity ec.Entity) {
	if entity.GetState() != ec.EntityState_Enter {
		return
	}

	var hooks [3]event.Hook

	if entityUpdate, ok := entity.(LifecycleEntityUpdate); ok {
		hooks[0] = event.Bind[LifecycleEntityUpdate](&rt.eventUpdate, entityUpdate)
	}
	if entityLateUpdate, ok := entity.(LifecycleEntityLateUpdate); ok {
		hooks[1] = event.Bind[LifecycleEntityLateUpdate](&rt.eventLateUpdate, entityLateUpdate)
	}
	hooks[2] = ec.BindEventEntityDestroySelf(entity, rt)

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

	if compUpdate, ok := comp.(LifecycleComponentUpdate); ok {
		hooks[0] = event.Bind[LifecycleComponentUpdate](&rt.eventUpdate, compUpdate)
		bound = true
	}
	if compLateUpdate, ok := comp.(LifecycleComponentLateUpdate); ok {
		hooks[1] = event.Bind[LifecycleComponentLateUpdate](&rt.eventLateUpdate, compLateUpdate)
		bound = true
	}
	if !comp.GetNonRemovable() {
		hooks[2] = ec.BindEventComponentDestroySelf(comp, rt)
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

	if entityAwake, ok := entity.(LifecycleEntityAwake); ok {
		generic.CastAction0(entityAwake.Awake).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}

	if entity.GetState() != ec.EntityState_Awake {
		return
	}

	entity.RangeComponents(func(comp ec.Component) bool {
		if comp.GetState() != ec.ComponentState_Awake {
			return true
		}

		if compAwake, ok := comp.(LifecycleComponentAwake); ok {
			generic.CastAction0(compAwake.Awake).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
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

		if compStart, ok := comp.(LifecycleComponentStart); ok {
			generic.CastAction0(compStart.Start).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
		}

		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Alive)

		return entity.GetState() == ec.EntityState_Awake
	})

	if entity.GetState() != ec.EntityState_Awake {
		return
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Start)

	if entityStart, ok := entity.(LifecycleEntityStart); ok {
		generic.CastAction0(entityStart.Start).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
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

	if entityShut, ok := entity.(LifecycleEntityShut); ok {
		generic.CastAction0(entityShut.Shut).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}

	entity.RangeComponents(func(comp ec.Component) bool {
		if comp.GetState() != ec.ComponentState_Shut {
			return true
		}

		if compShut, ok := comp.(LifecycleComponentShut); ok {
			generic.CastAction0(compShut.Shut).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
		}

		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Death)

		if compDispose, ok := comp.(LifecycleComponentDispose); ok {
			generic.CastAction0(compDispose.Dispose).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
		}

		ec.UnsafeComponent(comp).CleanManagedHooks()

		return true
	})

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Death)

	if entityDispose, ok := entity.(LifecycleEntityDispose); ok {
		generic.CastAction0(entityDispose.Dispose).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}

	ec.UnsafeEntity(entity).CleanManagedHooks()
}
