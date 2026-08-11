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
	"sync/atomic"
	"time"

	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/reinterpret"
)

// NewRuntime 创建运行时并将 rtCtx 绑定到该运行时。
// 同一个运行时上下文只能绑定一次；创建期间会同步触发 RunningEvent_Birth。
func NewRuntime(rtCtx runtime.Context, settings ...option.Setting[RuntimeOptions]) Runtime {
	return UnsafeNewRuntime(rtCtx, option.New(With.Runtime.Default(), settings...))
}

// Deprecated: UnsafeNewRuntime 仅供框架内部使用，请改用 NewRuntime。
func UnsafeNewRuntime(rtCtx runtime.Context, options RuntimeOptions) Runtime {
	var rt Runtime

	if !options.InstanceFace.IsNil() {
		rt = options.InstanceFace.Iface
	} else {
		rt = &RuntimeBehavior{}
	}
	rt.init(rtCtx, options)

	return rt
}

// Runtime 驱动单个 Actor 风格运行时的任务队列、实体生命周期和可选帧循环。
type Runtime interface {
	iRuntime
	iWorker
	iRuntimeStats
	corectx.CurrentContextProvider
	corectx.ConcurrentContextProvider
	reinterpret.InstanceProvider
	runtime.Callee
}

type iRuntime interface {
	init(rtCtx runtime.Context, options RuntimeOptions)
	getOptions() *RuntimeOptions
	getInstance() Runtime
}

// RuntimeBehavior 提供 Runtime 的默认实现。
// 扩展运行时类型时应匿名嵌入该类型，并通过 InstanceFace 传入扩展实例。
type RuntimeBehavior struct {
	ctx                                                  runtime.Context
	options                                              RuntimeOptions
	isRunning                                            atomic.Bool
	frame                                                *_Frame
	taskQueue                                            _TaskQueue
	handleEventEntityManagerAddEntity                    runtime.EventEntityManagerAddEntity
	handleEventEntityManagerRemoveEntity                 runtime.EventEntityManagerRemoveEntity
	handleEventEntityManagerEntityAddComponents          runtime.EventEntityManagerEntityAddComponents
	handleEventEntityManagerEntityRemoveComponent        runtime.EventEntityManagerEntityRemoveComponent
	handleEventEntityManagerEntityComponentEnableChanged runtime.EventEntityManagerEntityComponentEnableChanged
	handleEventEntityManagerEntityFirstTouchComponent    runtime.EventEntityManagerEntityFirstTouchComponent
	managedAddInManagerHandles                           [2]event.Handle
	lastProgressTime                                     atomic.Int64

	runtimeEventTab runtimeEventTab
}

// CurrentContextCache 返回仅供运行时 goroutine 使用的上下文接口缓存。
func (rt *RuntimeBehavior) CurrentContextCache() iface.Cache {
	return rt.ctx.CurrentContextCache()
}

// ConcurrentContextCache 返回可跨 goroutine 使用的上下文接口缓存。
func (rt *RuntimeBehavior) ConcurrentContextCache() iface.Cache {
	return rt.ctx.ConcurrentContextCache()
}

// InstanceFaceCache 返回运行时实例的接口缓存，用于 reinterpret.Cast。
func (rt *RuntimeBehavior) InstanceFaceCache() iface.Cache {
	return rt.options.InstanceFace.Cache
}

func (rt *RuntimeBehavior) init(rtCtx runtime.Context, options RuntimeOptions) {
	if rtCtx == nil {
		exception.Panicf("%w: %w: rtCtx is nil", ErrRuntime, ErrArgs)
	}

	if !runtime.UnsafeContext(rtCtx).Scoped().CompareAndSwap(false, true) {
		exception.Panicf("%w: %w: rtCtx is already bound to another runtime scope", ErrRuntime, ErrArgs)
	}

	rt.ctx = rtCtx
	rt.options = options
	rt.lastProgressTime.Store(time.Now().UnixNano())

	if rt.options.InstanceFace.IsNil() {
		rt.options.InstanceFace = iface.NewFaceT[Runtime](rt)
	}

	if rt.options.Frame.Enabled {
		rt.frame = &_Frame{}
		rt.frame.init(rt.options.Frame.TargetFPS, rt.options.Frame.TotalFrames)
		runtime.UnsafeContext(rtCtx).SetFrame(rt.frame)
	} else {
		runtime.UnsafeContext(rtCtx).SetFrame(nil)
	}

	rt.taskQueue.init(rt.options.TaskQueue.Unbounded, rt.options.TaskQueue.Capacity)
	runtime.UnsafeContext(rtCtx).SetCallee(rt.getInstance())

	rt.runtimeEventTab.SetPanicHandling(rtCtx.AutoRecover(), rtCtx.ReportError())

	rt.handleEventEntityManagerAddEntity = runtime.HandleEventEntityManagerAddEntity(rt.onEntityManagerAddEntity)
	rt.handleEventEntityManagerRemoveEntity = runtime.HandleEventEntityManagerRemoveEntity(rt.onEntityManagerRemoveEntity)
	rt.handleEventEntityManagerEntityAddComponents = runtime.HandleEventEntityManagerEntityAddComponents(rt.onEntityManagerEntityAddComponents)
	rt.handleEventEntityManagerEntityRemoveComponent = runtime.HandleEventEntityManagerEntityRemoveComponent(rt.onEntityManagerEntityRemoveComponent)
	rt.handleEventEntityManagerEntityComponentEnableChanged = runtime.HandleEventEntityManagerEntityComponentEnableChanged(rt.onEntityManagerEntityComponentEnableChanged)
	rt.handleEventEntityManagerEntityFirstTouchComponent = runtime.HandleEventEntityManagerEntityFirstTouchComponent(rt.onEntityManagerEntityFirstTouchComponent)

	runtime.BindEventContextRunningEvent(rtCtx, runtime.HandleEventContextRunningEvent(rt.onBeforeContextRunningEvent), -100)
	runtime.BindEventContextRunningEvent(rtCtx, runtime.HandleEventContextRunningEvent(rt.onAfterContextRunningEvent), 100)

	rt.emitEventRunningEvent(runtime.RunningEvent_Birth)
}

func (rt *RuntimeBehavior) getOptions() *RuntimeOptions {
	return &rt.options
}

func (rt *RuntimeBehavior) getInstance() Runtime {
	return rt.options.InstanceFace.Iface
}

// onEntityManagerAddEntity 在实体加入 Runtime 管理器后推进其激活流程。
func (rt *RuntimeBehavior) onEntityManagerAddEntity(entityManager runtime.EntityManager, entity ec.Entity) {
	if entity.State() != ec.EntityState_Entered {
		return
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Awakened)

	entity.EachComponents(func(comp ec.Component) {
		if comp.State() != ec.ComponentState_Attached {
			return
		}
		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Awakened)
	})

	{
		caller := newEntityLifecycleCaller(entity)

		if !caller.Call(func() {
			rt.emitEventRunningEvent(runtime.RunningEvent_EntityActivating, entity)
		}) {
			rt.emitEventRunningEvent(runtime.RunningEvent_EntityActivationAborted, entity)
			return
		}

		if !caller.Call(func() {
			if !caller.MarkProcessed() {
				return
			}
			if cb, ok := entity.(LifecycleEntityAwake); ok {
				rt.panicHandlingActivatingEntity(entity, generic.CastAction0(cb.Awake).Call(rt.ctx.AutoRecover(), rt.ctx.ReportError()))
			}
		}) {
			rt.emitEventRunningEvent(runtime.RunningEvent_EntityActivationAborted, entity)
			return
		}

		rt.observeEntity(entity)

		if !caller.Call(func() {
			entity.RangeComponents(func(comp ec.Component) bool {
				return caller.Call(func() {
					rt.panicHandlingActivatingEntity(entity, rt.awakeComponent(comp))
				})
			})
		}) {
			rt.emitEventRunningEvent(runtime.RunningEvent_EntityActivationAborted, entity)
			return
		}

		if !caller.Call(func() {
			entity.RangeComponents(func(comp ec.Component) bool {
				return caller.Call(func() {
					rt.panicHandlingActivatingEntity(entity, rt.enableAwakenedComponent(comp))
				})
			})
		}) {
			rt.emitEventRunningEvent(runtime.RunningEvent_EntityActivationAborted, entity)
			return
		}
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Starting)

	{
		caller := newEntityLifecycleCaller(entity)

		if !caller.Call(func() {
			entity.RangeComponents(func(comp ec.Component) bool {
				return caller.Call(func() {
					rt.panicHandlingActivatingEntity(entity, rt.startComponent(comp))
				})
			})
		}) {
			rt.emitEventRunningEvent(runtime.RunningEvent_EntityActivationAborted, entity)
			return
		}

		if !caller.Call(func() {
			if !caller.MarkProcessed() {
				return
			}
			if cb, ok := entity.(LifecycleEntityStart); ok {
				rt.panicHandlingActivatingEntity(entity, generic.CastAction0(cb.Start).Call(rt.ctx.AutoRecover(), rt.ctx.ReportError()))
			}
		}) {
			rt.emitEventRunningEvent(runtime.RunningEvent_EntityActivationAborted, entity)
			return
		}
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Alive)

	rt.emitEventRunningEvent(runtime.RunningEvent_EntityActivated, entity)
}

// onEntityManagerRemoveEntity 在实体离开 Runtime 管理器时推进其停用与销毁流程。
func (rt *RuntimeBehavior) onEntityManagerRemoveEntity(entityManager runtime.EntityManager, entity ec.Entity) {
	if entity.State() != ec.EntityState_Leaving {
		return
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Shutting)

	entity.EachComponents(func(comp ec.Component) {
		if comp.State() < ec.ComponentState_Awakened {
			return
		}
		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Shutting)
	})

	rt.emitEventRunningEvent(runtime.RunningEvent_EntityDeactivating, entity)

	{
		caller := newEntityLifecycleCaller(entity)

		if caller.IsProcessed(ec.EntityState_Starting) && caller.MarkProcessed() {
			if cb, ok := entity.(LifecycleEntityShut); ok {
				generic.CastAction0(cb.Shut).Call(rt.ctx.AutoRecover(), rt.ctx.ReportError())
			}
		}

		entity.ReversedEachComponents(func(comp ec.Component) {
			rt.shutComponent(comp)
		})
	}

	ec.UnsafeEntity(entity).SetState(ec.EntityState_Dead)

	{
		caller := newEntityLifecycleCaller(entity)

		entity.ReversedEachComponents(func(comp ec.Component) {
			rt.disableDeathComponent(comp)
		})

		entity.ReversedEachComponents(func(comp ec.Component) {
			rt.disposeComponent(comp, true)
		})

		if caller.IsProcessed(ec.EntityState_Awakened) && caller.MarkProcessed() {
			if cb, ok := entity.(LifecycleEntityDispose); ok {
				generic.CastAction0(cb.Dispose).Call(rt.ctx.AutoRecover(), rt.ctx.ReportError())
			}
		}
	}

	rt.emitEventRunningEvent(runtime.RunningEvent_EntityDeactivated, entity)
}

// onEntityManagerEntityAddComponents 为运行中实体推进新增组件的激活流程。
func (rt *RuntimeBehavior) onEntityManagerEntityAddComponents(entityManager runtime.EntityManager, entity ec.Entity, components []ec.Component) {
	if entity.State() < ec.EntityState_Awakened || entity.State() > ec.EntityState_Alive {
		return
	}

	for i := range components {
		comp := components[i]
		if comp.State() != ec.ComponentState_Attached {
			continue
		}
		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Awakened)
	}

	{
		caller := newEntityLifecycleCaller(entity)

		if !caller.Call(func() {
			rt.emitEventRunningEvent(runtime.RunningEvent_EntityComponentsActivating, entity, components)
		}) {
			rt.emitEventRunningEvent(runtime.RunningEvent_EntityComponentsActivationAborted, entity, components)
			return
		}

		if !caller.Call(func() {
			for i := range components {
				if !caller.Call(func() {
					rt.awakeComponent(components[i])
				}) {
					return
				}
			}
		}) {
			rt.emitEventRunningEvent(runtime.RunningEvent_EntityComponentsActivationAborted, entity, components)
			return
		}

		if !caller.Call(func() {
			for i := range components {
				if !caller.Call(func() {
					rt.enableAwakenedComponent(components[i])
				}) {
					return
				}
			}
		}) {
			rt.emitEventRunningEvent(runtime.RunningEvent_EntityComponentsActivationAborted, entity, components)
			return
		}

		if entity.State() >= ec.EntityState_Starting {
			if !caller.Call(func() {
				for i := range components {
					if !caller.Call(func() {
						rt.startComponent(components[i])
					}) {
						return
					}
				}
			}) {
				rt.emitEventRunningEvent(runtime.RunningEvent_EntityComponentsActivationAborted, entity, components)
				return
			}
		}
	}

	rt.emitEventRunningEvent(runtime.RunningEvent_EntityComponentsActivated, entity, components)
}

// onEntityManagerEntityRemoveComponent 为运行中实体推进组件移除流程。
func (rt *RuntimeBehavior) onEntityManagerEntityRemoveComponent(entityManager runtime.EntityManager, entity ec.Entity, component ec.Component) {
	if entity.State() < ec.EntityState_Awakened || entity.State() > ec.EntityState_Alive {
		return
	}

	if component.State() != ec.ComponentState_Detaching {
		return
	}

	{
		caller := newEntityLifecycleCaller(entity)

		if !caller.Call(func() {
			rt.emitEventRunningEvent(runtime.RunningEvent_EntityComponentDeactivating, entity, component)
		}) {
			rt.emitEventRunningEvent(runtime.RunningEvent_EntityComponentDeactivationAborted, entity, component)
			return
		}

		ec.UnsafeComponent(component).SetState(ec.ComponentState_Shutting)

		if !caller.Call(func() {
			rt.shutComponent(component)
		}) {
			rt.emitEventRunningEvent(runtime.RunningEvent_EntityComponentDeactivationAborted, entity, component)
			return
		}

		if !caller.Call(func() {
			rt.disableDeathComponent(component)
		}) {
			rt.emitEventRunningEvent(runtime.RunningEvent_EntityComponentDeactivationAborted, entity, component)
			return
		}

		if !caller.Call(func() {
			rt.disposeComponent(component, false)
		}) {
			rt.emitEventRunningEvent(runtime.RunningEvent_EntityComponentDeactivationAborted, entity, component)
			return
		}
	}

	rt.emitEventRunningEvent(runtime.RunningEvent_EntityComponentDeactivated, entity, component)
}

// onEntityManagerEntityComponentEnableChanged 推进组件启用状态变化对应的生命周期。
func (rt *RuntimeBehavior) onEntityManagerEntityComponentEnableChanged(entityManager runtime.EntityManager, entity ec.Entity, component ec.Component, enable bool) {
	if entity.State() < ec.EntityState_Awakened || entity.State() > ec.EntityState_Alive {
		return
	}

	{
		caller := newEntityLifecycleCaller(entity)

		if enable {
			if !caller.Call(func() {
				rt.enableComponent(component)
			}) {
				return
			}

			if !caller.Call(func() {
				rt.startComponent(component)
			}) {
				return
			}

		} else {
			if !caller.Call(func() {
				rt.disableComponent(component)
			}) {
				return
			}
		}
	}
}

// onEntityManagerEntityFirstTouchComponent 在首次访问时唤醒延迟激活的组件。
func (rt *RuntimeBehavior) onEntityManagerEntityFirstTouchComponent(entityManager runtime.EntityManager, entity ec.Entity, component ec.Component) {
	if entity.State() < ec.EntityState_Awakened || entity.State() > ec.EntityState_Alive {
		return
	}

	ec.UnsafeComponent(component).SetState(ec.ComponentState_Awakened)

	{
		caller := newEntityLifecycleCaller(entity)

		if !caller.Call(func() {
			rt.awakeComponent(component)
		}) {
			return
		}
	}
}

func (rt *RuntimeBehavior) observeEntity(entity ec.Entity) {
	if cb, ok := entity.(LifecycleEntityUpdate); ok {
		ec.UnsafeEntity(entity).ManagedRuntimeUpdateHandle(_BindEventUpdate(&rt.runtimeEventTab, cb))
	}
	if cb, ok := entity.(LifecycleEntityLateUpdate); ok {
		ec.UnsafeEntity(entity).ManagedRuntimeLateUpdateHandle(_BindEventLateUpdate(&rt.runtimeEventTab, cb))
	}
}

func (rt *RuntimeBehavior) observeComponent(comp ec.Component) {
	if cb, ok := comp.(LifecycleComponentUpdate); ok {
		ec.UnsafeComponent(comp).ManagedRuntimeUpdateHandle(_BindEventUpdate(&rt.runtimeEventTab, cb))
	}
	if cb, ok := comp.(LifecycleComponentLateUpdate); ok {
		ec.UnsafeComponent(comp).ManagedRuntimeLateUpdateHandle(_BindEventLateUpdate(&rt.runtimeEventTab, cb))
	}
}

func (rt *RuntimeBehavior) unobserveComponent(comp ec.Component) {
	ec.UnsafeComponent(comp).ManagedUnbindRuntimeHandles()
}

func (rt *RuntimeBehavior) awakeComponent(comp ec.Component) (err error) {
	if comp.State() != ec.ComponentState_Awakened {
		return
	}

	{
		caller := newComponentLifecycleCaller(comp)

		if !caller.Call(func() {
			if !caller.MarkProcessed() {
				return
			}
			if cb, ok := comp.(LifecycleComponentAwake); ok {
				err = generic.CastAction0(cb.Awake).Call(rt.ctx.AutoRecover(), rt.ctx.ReportError())
			}
		}) {
			return
		}
	}

	ec.UnsafeComponent(comp).SetState(ec.ComponentState_Enabling)

	return
}

func (rt *RuntimeBehavior) enableAwakenedComponent(comp ec.Component) (err error) {
	if comp.State() != ec.ComponentState_Enabling {
		return
	}

	if !comp.Enabled() {
		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Idle)
		return
	}

	{
		caller := newComponentLifecycleCaller(comp)

		if !caller.Call(func() {
			if !caller.MarkProcessed() {
				return
			}
			if cb, ok := comp.(LifecycleComponentOnEnable); ok {
				err = generic.CastAction0(cb.OnEnable).Call(rt.ctx.AutoRecover(), rt.ctx.ReportError())
			}
		}) {
			return
		}
	}

	if !comp.Enabled() {
		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Idle)
		return
	}

	rt.observeComponent(comp)

	ec.UnsafeComponent(comp).SetState(ec.ComponentState_Starting)

	return
}

func (rt *RuntimeBehavior) startComponent(comp ec.Component) (err error) {
	if comp.State() != ec.ComponentState_Starting {
		return
	}

	{
		caller := newComponentLifecycleCaller(comp)

		if !caller.Call(func() {
			if !caller.MarkProcessed() {
				return
			}
			if cb, ok := comp.(LifecycleComponentStart); ok {
				err = generic.CastAction0(cb.Start).Call(rt.ctx.AutoRecover(), rt.ctx.ReportError())
			}
		}) {
			return
		}
	}

	ec.UnsafeComponent(comp).SetState(ec.ComponentState_Alive)

	return
}

func (rt *RuntimeBehavior) shutComponent(comp ec.Component) {
	if comp.State() != ec.ComponentState_Shutting {
		return
	}

	{
		caller := newComponentLifecycleCaller(comp)

		if !caller.Call(func() {
			if caller.IsProcessed(ec.ComponentState_Starting) {
				if !caller.MarkProcessed() {
					return
				}
				if cb, ok := comp.(LifecycleComponentShut); ok {
					generic.CastAction0(cb.Shut).Call(rt.ctx.AutoRecover(), rt.ctx.ReportError())
				}
			}
		}) {
			return
		}
	}

	ec.UnsafeComponent(comp).SetState(ec.ComponentState_Disabling)
}

func (rt *RuntimeBehavior) disableDeathComponent(comp ec.Component) {
	if comp.State() != ec.ComponentState_Disabling {
		return
	}

	{
		caller := newComponentLifecycleCaller(comp)

		if !caller.Call(func() {
			if caller.IsProcessed(ec.ComponentState_Enabling) {
				if !caller.MarkProcessed() {
					return
				}
				if !comp.Enabled() {
					return
				}
				if cb, ok := comp.(LifecycleComponentOnDisable); ok {
					generic.CastAction0(cb.OnDisable).Call(rt.ctx.AutoRecover(), rt.ctx.ReportError())
				}
			}
		}) {
			return
		}
	}

	ec.UnsafeComponent(comp).SetState(ec.ComponentState_Dead)
}

func (rt *RuntimeBehavior) disposeComponent(comp ec.Component, stateDestroyed bool) {
	if comp.State() != ec.ComponentState_Dead {
		return
	}

	{
		caller := newComponentLifecycleCaller(comp)

		if !caller.Call(func() {
			if caller.IsProcessed(ec.ComponentState_Awakened) {
				if !caller.MarkProcessed() {
					return
				}
				if cb, ok := comp.(LifecycleComponentDispose); ok {
					generic.CastAction0(cb.Dispose).Call(rt.ctx.AutoRecover(), rt.ctx.ReportError())
				}
			}
		}) {
			return
		}
	}

	if stateDestroyed {
		ec.UnsafeComponent(comp).SetState(ec.ComponentState_Destroyed)
	}
}

func (rt *RuntimeBehavior) enableComponent(comp ec.Component) {
	if comp.State() < ec.ComponentState_Enabling || comp.State() >= ec.ComponentState_Disabling {
		return
	}

	{
		caller := newComponentLifecycleCaller(comp)

		if !caller.Call(func() {
			caller.SetProcessed(ec.ComponentState_Enabling)

			if cb, ok := comp.(LifecycleComponentOnEnable); ok {
				generic.CastAction0(cb.OnEnable).Call(rt.ctx.AutoRecover(), rt.ctx.ReportError())
			}
		}) {
			return
		}
	}

	rt.observeComponent(comp)

	ec.UnsafeComponent(comp).SetState(ec.ComponentState_Starting)
}

func (rt *RuntimeBehavior) disableComponent(comp ec.Component) {
	if comp.State() < ec.ComponentState_Enabling || comp.State() >= ec.ComponentState_Disabling {
		return
	}

	rt.unobserveComponent(comp)

	{
		caller := newComponentLifecycleCaller(comp)

		if !caller.Call(func() {
			if cb, ok := comp.(LifecycleComponentOnDisable); ok {
				generic.CastAction0(cb.OnDisable).Call(rt.ctx.AutoRecover(), rt.ctx.ReportError())
			}
		}) {
			return
		}
	}

	ec.UnsafeComponent(comp).SetState(ec.ComponentState_Idle)
}

func (rt *RuntimeBehavior) panicHandlingActivatingEntity(entity ec.Entity, err error) {
	if err != nil && !rt.options.ContinueOnActivatingEntityPanic {
		entity.Destroy()
	}
}
