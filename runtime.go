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
	"git.golaxy.org/core/ec/ictx"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/reinterpret"
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

const (
	tagForRuntimeObserveComponentEnableChanged = "runtime_observe_component_enable_changed"
	tagForRuntimeObserveComponentUpdate        = "runtime_observe_component_update"
)

func makeEntityLifecycleCaller(entity ec.Entity) _EntityLifecycleCaller {
	return _EntityLifecycleCaller{entity: entity, state: entity.GetState()}
}

type _EntityLifecycleCaller struct {
	entity ec.Entity
	state  ec.EntityState
}

func (c _EntityLifecycleCaller) Call(fun func(state ec.EntityState)) bool {
	if c.entity.GetState() != c.state {
		return false
	}

	fun(c.state)

	return c.entity.GetState() == c.state
}

func makeComponentLifecycleCaller(comp ec.Component) _ComponentLifecycleCaller {
	return _ComponentLifecycleCaller{component: comp, state: comp.GetState()}
}

type _ComponentLifecycleCaller struct {
	component ec.Component
	state     ec.ComponentState
}

func (c _ComponentLifecycleCaller) Call(fun func(state ec.ComponentState)) bool {
	state := c.component.GetState()

	if state != c.state {
		return false
	}

	bits := ec.UnsafeComponent(c.component).GetCallingStateBits()
	if bits.Is(int8(state)) {
		return true
	}

	bits.Set(int8(state), true)
	defer bits.Set(int8(state), false)

	fun(c.state)

	return c.component.GetState() == c.state
}

// RuntimeBehavior 运行时行为，在扩展运行时能力时，匿名嵌入至运行时结构体中
type RuntimeBehavior struct {
	ctx                                               runtime.Context
	opts                                              RuntimeOptions
	processQueue                                      chan _Task
	eventUpdate                                       event.Event
	eventLateUpdate                                   event.Event
	eventRuntimeRunningStatusChanged                  event.Event
	handleEventEntityManagerAddEntity                 runtime.EventEntityManagerAddEntity
	handleEventEntityManagerRemoveEntity              runtime.EventEntityManagerRemoveEntity
	handleEventEntityManagerEntityFirstTouchComponent runtime.EventEntityManagerEntityFirstTouchComponent
	handleEventEntityManagerEntityAddComponents       runtime.EventEntityManagerEntityAddComponents
	handleEventEntityManagerEntityRemoveComponent     runtime.EventEntityManagerEntityRemoveComponent
	handleEventEntityDestroySelf                      ec.EventEntityDestroySelf
	handleEventComponentEnableChanged                 ec.EventComponentEnableChanged
	handleEventComponentDestroySelf                   ec.EventComponentDestroySelf
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

	rt.processQueue = make(chan _Task, rt.opts.ProcessQueueCapacity)

	runtime.UnsafeContext(rtCtx).SetFrame(rt.opts.Frame)
	runtime.UnsafeContext(rtCtx).SetCallee(rt.opts.InstanceFace.Iface)

	rtCtx.ActivateEvent(&rt.eventUpdate, event.EventRecursion_Disallow)
	rtCtx.ActivateEvent(&rt.eventLateUpdate, event.EventRecursion_Disallow)
	rtCtx.ActivateEvent(&rt.eventRuntimeRunningStatusChanged, event.EventRecursion_Allow)

	rt.handleEventEntityManagerAddEntity = runtime.HandleEventEntityManagerAddEntity(rt.onEntityManagerAddEntity)
	rt.handleEventEntityManagerRemoveEntity = runtime.HandleEventEntityManagerRemoveEntity(rt.onEntityManagerRemoveEntity)
	rt.handleEventEntityManagerEntityFirstTouchComponent = runtime.HandleEventEntityManagerEntityFirstTouchComponent(rt.onEntityManagerEntityFirstTouchComponent)
	rt.handleEventEntityManagerEntityAddComponents = runtime.HandleEventEntityManagerEntityAddComponents(rt.onEntityManagerEntityAddComponents)
	rt.handleEventEntityManagerEntityRemoveComponent = runtime.HandleEventEntityManagerEntityRemoveComponent(rt.onEntityManagerEntityRemoveComponent)
	rt.handleEventEntityDestroySelf = ec.HandleEventEntityDestroySelf(rt.onEntityDestroySelf)
	rt.handleEventComponentEnableChanged = ec.HandleEventComponentEnableChanged(rt.onComponentEnableChanged)
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
	if entity.GetState() != ec.EntityState_Enter {
		return
	}

	rt.observeEntity(entity)
	rt.activateEntity(entity)
}

// onEntityManagerRemoveEntity 事件处理器：实体管理器删除实体
func (rt *RuntimeBehavior) onEntityManagerRemoveEntity(entityManager runtime.EntityManager, entity ec.Entity) {
	if entity.GetState() != ec.EntityState_Leave {
		return
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Shut)

	rt.deactivateEntity(entity)
}

// onEntityManagerEntityFirstTouchComponent 事件处理器：实体管理器中的实体首次访问组件
func (rt *RuntimeBehavior) onEntityManagerEntityFirstTouchComponent(entityManager runtime.EntityManager, entity ec.Entity, component ec.Component) {
	if entity.GetState() < ec.EntityState_Awake || entity.GetState() > ec.EntityState_Alive {
		return
	}

	rt.observeComponentDestroySelf(component)
	rt.awakeComponent(component)
}

// onEntityManagerEntityAddComponents 事件处理器：实体管理器中的实体添加组件
func (rt *RuntimeBehavior) onEntityManagerEntityAddComponents(entityManager runtime.EntityManager, entity ec.Entity, components []ec.Component) {
	if entity.GetState() < ec.EntityState_Awake || entity.GetState() > ec.EntityState_Alive {
		return
	}

	for i := range components {
		rt.observeComponentDestroySelf(components[i])
	}

	rt.changeRunningStatus(runtime.RunningStatus_EntityAddComponents, entity, components)

	{
		caller := makeEntityLifecycleCaller(entity)

		if !caller.Call(func(state ec.EntityState) {
			for i := range components {
				if entity.GetState() != state {
					return
				}
				rt.awakeComponent(components[i])
			}
		}) {
			return
		}

		if !caller.Call(func(state ec.EntityState) {
			for i := range components {
				if entity.GetState() != state {
					return
				}
				rt.enableAwokeComponent(components[i])
			}
		}) {
			return
		}

		if !caller.Call(func(state ec.EntityState) {
			for i := range components {
				if entity.GetState() != state {
					return
				}
				rt.startComponent(components[i])
			}
		}) {
			return
		}
	}
}

// onEntityManagerEntityRemoveComponent 事件处理器：实体管理器中的实体删除组件
func (rt *RuntimeBehavior) onEntityManagerEntityRemoveComponent(entityManager runtime.EntityManager, entity ec.Entity, component ec.Component) {
	if entity.GetState() < ec.EntityState_Awake || entity.GetState() > ec.EntityState_Alive {
		return
	}

	if component.GetState() != ec.ComponentState_Detach {
		return
	}

	if !ec.UnsafeComponent(component).GetProcessedStateBits().Is(int8(ec.ComponentState_Awake)) {
		ec.UnsafeComponent(component).SetState(ec.ComponentState_Destroyed)
		return
	}

	if ec.UnsafeComponent(component).GetProcessedStateBits().Is(int8(ec.ComponentState_Alive)) {
		ec.UnsafeComponent(component).SetState(ec.ComponentState_Shut)
	} else {
		ec.UnsafeComponent(component).SetState(ec.ComponentState_Disable)
	}

	{
		caller := makeEntityLifecycleCaller(entity)

		if !caller.Call(func(state ec.EntityState) {
			rt.shutComponent(component)
		}) {
			return
		}

		if !caller.Call(func(state ec.EntityState) {
			rt.disableDeathComponent(component)
		}) {
			return
		}

		if !caller.Call(func(state ec.EntityState) {
			rt.disposeComponent(component)
		}) {
			return
		}
	}
}

// onEntityDestroySelf 事件处理器：实体销毁自身
func (rt *RuntimeBehavior) onEntityDestroySelf(entity ec.Entity) {
	rt.ctx.GetEntityManager().RemoveEntity(entity.GetId())
}

// onComponentEnableChanged 事件处理器：组件启用状态改变
func (rt *RuntimeBehavior) onComponentEnableChanged(comp ec.Component, enable bool) {
	if comp.GetEnable() != enable {
		return
	}

	caller := makeEntityLifecycleCaller(comp.GetEntity())

	if enable {
		if !caller.Call(func(ec.EntityState) {
			rt.enableComponent(comp)
		}) {
			return
		}

		if !caller.Call(func(ec.EntityState) {
			rt.startComponent(comp)
		}) {
			return
		}

	} else {
		if !caller.Call(func(ec.EntityState) {
			rt.disableComponent(comp)
		}) {
			return
		}
	}
}

// onComponentDestroySelf 事件处理器：组件销毁自身
func (rt *RuntimeBehavior) onComponentDestroySelf(comp ec.Component) {
	ec.UnsafeEntity(comp.GetEntity()).RemoveComponentByRef(comp)
}

func (rt *RuntimeBehavior) observeEntity(entity ec.Entity) {
	if entity.GetState() != ec.EntityState_Enter {
		return
	}

	if cb, ok := entity.(LifecycleEntityUpdate); ok {
		event.Bind[LifecycleEntityUpdate](&rt.eventUpdate, cb)
	}

	if cb, ok := entity.(LifecycleEntityLateUpdate); ok {
		event.Bind[LifecycleEntityLateUpdate](&rt.eventLateUpdate, cb)
	}

	ec.BindEventEntityDestroySelf(entity, rt.handleEventEntityDestroySelf)

	entity.RangeComponents(func(comp ec.Component) bool {
		rt.observeComponentDestroySelf(comp)
		return true
	})

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Awake)
}

func (rt *RuntimeBehavior) observeComponentDestroySelf(comp ec.Component) {
	if comp.GetState() != ec.ComponentState_Attach {
		return
	}

	if comp.GetRemovable() {
		ec.BindEventComponentDestroySelf(comp, rt.handleEventComponentDestroySelf)
	}

	ec.UnsafeComponent(comp).SetState(ec.ComponentState_Awake)
}

func (rt *RuntimeBehavior) observeComponentEnableChanged(comp ec.Component) {
	comp.ManagedAddTagHooks(tagForRuntimeObserveComponentEnableChanged, ec.BindEventComponentEnableChanged(comp, rt.handleEventComponentEnableChanged))
}

func (rt *RuntimeBehavior) unobserveComponentEnableChanged(comp ec.Component) {
	comp.ManagedCleanTagHooks(tagForRuntimeObserveComponentEnableChanged)
}

func (rt *RuntimeBehavior) observeComponentUpdate(comp ec.Component) {
	var hooks []event.Hook

	if cb, ok := comp.(LifecycleComponentUpdate); ok {
		hooks = append(hooks, event.Bind[LifecycleComponentUpdate](&rt.eventUpdate, cb))
	}

	if cb, ok := comp.(LifecycleComponentLateUpdate); ok {
		hooks = append(hooks, event.Bind[LifecycleComponentLateUpdate](&rt.eventLateUpdate, cb))
	}

	comp.ManagedAddTagHooks(tagForRuntimeObserveComponentUpdate, hooks...)
}

func (rt *RuntimeBehavior) unobserveComponentUpdate(comp ec.Component) {
	comp.ManagedCleanTagHooks(tagForRuntimeObserveComponentUpdate)
}

func (rt *RuntimeBehavior) activateEntity(entity ec.Entity) {
	if entity.GetState() != ec.EntityState_Awake {
		return
	}

	rt.changeRunningStatus(runtime.RunningStatus_ActivatingEntity, entity)

	{
		caller := makeEntityLifecycleCaller(entity)

		if !caller.Call(func(ec.EntityState) {
			if cb, ok := entity.(LifecycleEntityAwake); ok {
				generic.CastAction0(cb.Awake).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
			}
		}) {
			return
		}

		rt.changeRunningStatus(runtime.RunningStatus_EntityInitComponents, entity)

		if !caller.Call(func(state ec.EntityState) {
			entity.RangeComponents(func(comp ec.Component) bool {
				rt.awakeComponent(comp)
				return entity.GetState() == state
			})
		}) {
			return
		}

		if !caller.Call(func(state ec.EntityState) {
			entity.RangeComponents(func(comp ec.Component) bool {
				rt.enableAwokeComponent(comp)
				return entity.GetState() == state
			})
		}) {
			return
		}
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Start)

	{
		caller := makeEntityLifecycleCaller(entity)

		if !caller.Call(func(state ec.EntityState) {
			entity.RangeComponents(func(comp ec.Component) bool {
				rt.startComponent(comp)
				return entity.GetState() == state
			})
		}) {
			return
		}

		if !caller.Call(func(ec.EntityState) {
			if cb, ok := entity.(LifecycleEntityStart); ok {
				generic.CastAction0(cb.Start).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
			}
		}) {
			return
		}
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Alive)

	rt.changeRunningStatus(runtime.RunningStatus_EntityActivated, entity)
}

func (rt *RuntimeBehavior) deactivateEntity(entity ec.Entity) {
	if entity.GetState() != ec.EntityState_Shut {
		return
	}

	rt.changeRunningStatus(runtime.RunningStatus_DeactivatingEntity, entity)

	{
		if cb, ok := entity.(LifecycleEntityShut); ok {
			generic.CastAction0(cb.Shut).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
		}

		entity.RangeComponents(func(comp ec.Component) bool {
			if comp.GetState() > ec.ComponentState_Attach && comp.GetState() < ec.ComponentState_Detach {
				if ec.UnsafeComponent(comp).GetProcessedStateBits().Is(int8(ec.ComponentState_Alive)) {
					ec.UnsafeComponent(comp).SetState(ec.ComponentState_Shut)
				} else {
					ec.UnsafeComponent(comp).SetState(ec.ComponentState_Disable)
				}
			}
			return true
		})

		entity.RangeComponents(func(comp ec.Component) bool {
			rt.shutComponent(comp)
			return true
		})
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Death)

	{
		entity.RangeComponents(func(comp ec.Component) bool {
			rt.disableDeathComponent(comp)
			return true
		})

		entity.RangeComponents(func(comp ec.Component) bool {
			rt.disposeComponent(comp)
			return true
		})

		if cb, ok := entity.(LifecycleEntityDispose); ok {
			generic.CastAction0(cb.Dispose).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
		}
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Destroyed)

	rt.changeRunningStatus(runtime.RunningStatus_EntityDeactivated, entity)
}

func (rt *RuntimeBehavior) awakeComponent(comp ec.Component) {
	if comp.GetState() != ec.ComponentState_Awake {
		return
	}

	{
		caller := makeComponentLifecycleCaller(comp)

		if !caller.Call(func(ec.ComponentState) {
			if cb, ok := comp.(LifecycleComponentAwake); ok {
				generic.CastAction0(cb.Awake).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
			}
		}) {
			return
		}
	}

	ec.UnsafeComponent(comp).SetState(ec.ComponentState_Enable)
}

func (rt *RuntimeBehavior) enableAwokeComponent(comp ec.Component) {
	if comp.GetState() != ec.ComponentState_Enable {
		return
	}

	if !comp.GetEnable() {
		rt.observeComponentEnableChanged(comp)
		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Idle)
		return
	}

	{
		caller := makeComponentLifecycleCaller(comp)

		if !caller.Call(func(ec.ComponentState) {
			if cb, ok := comp.(LifecycleComponentOnEnable); ok {
				generic.CastAction0(cb.OnEnable).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
			}
		}) {
			return
		}
	}

	if !comp.GetEnable() {
		rt.observeComponentEnableChanged(comp)
		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Idle)
		return
	}

	rt.observeComponentEnableChanged(comp)
	rt.observeComponentUpdate(comp)

	if ec.UnsafeComponent(comp).GetProcessedStateBits().Is(int8(ec.ComponentState_Alive)) {
		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Alive)
	} else {
		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Start)
	}
}

func (rt *RuntimeBehavior) startComponent(comp ec.Component) {
	if comp.GetState() != ec.ComponentState_Start {
		return
	}

	{
		caller := makeComponentLifecycleCaller(comp)

		if !caller.Call(func(ec.ComponentState) {
			if cb, ok := comp.(LifecycleComponentStart); ok {
				generic.CastAction0(cb.Start).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
			}
		}) {
			return
		}
	}

	ec.UnsafeComponent(comp).SetState(ec.ComponentState_Alive)
}

func (rt *RuntimeBehavior) shutComponent(comp ec.Component) {
	if comp.GetState() != ec.ComponentState_Shut {
		return
	}

	{
		caller := makeComponentLifecycleCaller(comp)

		if !caller.Call(func(ec.ComponentState) {
			if cb, ok := comp.(LifecycleComponentShut); ok {
				generic.CastAction0(cb.Shut).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
			}
		}) {
			return
		}
	}

	ec.UnsafeComponent(comp).SetState(ec.ComponentState_Disable)
}

func (rt *RuntimeBehavior) disableDeathComponent(comp ec.Component) {
	if comp.GetState() != ec.ComponentState_Disable {
		return
	}

	rt.unobserveComponentEnableChanged(comp)
	rt.unobserveComponentUpdate(comp)

	if comp.GetEnable() && ec.UnsafeComponent(comp).GetProcessedStateBits().Is(int8(ec.ComponentState_Start)) {
		caller := makeComponentLifecycleCaller(comp)

		if !caller.Call(func(ec.ComponentState) {
			if cb, ok := comp.(LifecycleComponentOnDisable); ok {
				generic.CastAction0(cb.OnDisable).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
			}
		}) {
			return
		}
	}

	ec.UnsafeComponent(comp).SetState(ec.ComponentState_Death)
}

func (rt *RuntimeBehavior) disposeComponent(comp ec.Component) {
	if comp.GetState() != ec.ComponentState_Death {
		return
	}

	{
		caller := makeComponentLifecycleCaller(comp)

		if !caller.Call(func(ec.ComponentState) {
			if cb, ok := comp.(LifecycleComponentDispose); ok {
				generic.CastAction0(cb.Dispose).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
			}
		}) {
			return
		}
	}

	ec.UnsafeComponent(comp).SetState(ec.ComponentState_Destroyed)
}

func (rt *RuntimeBehavior) enableComponent(comp ec.Component) {
	if !comp.GetEnable() {
		return
	}

	{
		caller := makeComponentLifecycleCaller(comp)

		if !caller.Call(func(ec.ComponentState) {
			if cb, ok := comp.(LifecycleComponentOnEnable); ok {
				generic.CastAction0(cb.OnEnable).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
			}
		}) {
			return
		}
	}

	if !comp.GetEnable() {
		return
	}

	if comp.GetState() > ec.ComponentState_Alive {
		return
	}

	rt.observeComponentUpdate(comp)

	if ec.UnsafeComponent(comp).GetProcessedStateBits().Is(int8(ec.ComponentState_Alive)) {
		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Alive)
	} else {
		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Start)
	}
}

func (rt *RuntimeBehavior) disableComponent(comp ec.Component) {
	if comp.GetEnable() {
		return
	}

	rt.unobserveComponentUpdate(comp)

	if cb, ok := comp.(LifecycleComponentOnDisable); ok {
		generic.CastAction0(cb.OnDisable).Call(rt.ctx.GetAutoRecover(), rt.ctx.GetReportError())
	}

	if comp.GetEnable() {
		return
	}

	if comp.GetState() > ec.ComponentState_Alive {
		return
	}

	ec.UnsafeComponent(comp).SetState(ec.ComponentState_Idle)
}
