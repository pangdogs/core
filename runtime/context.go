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

package runtime

import (
	"reflect"
	"sync/atomic"

	"git.golaxy.org/core/event"
	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/reinterpret"
	"git.golaxy.org/core/utils/uid"
)

// NewContext 创建隶属于 svcCtx 的运行时上下文。
// 未提供父上下文、持久化 ID 或插件管理器时会自动创建默认值。
func NewContext(svcCtx service.Context, settings ...option.Setting[ContextOptions]) Context {
	return UnsafeNewContext(svcCtx, option.New(With.Default(), settings...))
}

// Deprecated: UnsafeNewContext 仅供框架内部使用，请改用 NewContext。
func UnsafeNewContext(svcCtx service.Context, options ContextOptions) Context {
	var ctx Context

	if !options.InstanceFace.IsNil() {
		ctx = options.InstanceFace.Iface
	} else {
		ctx = &ContextBehavior{}
	}
	ctx.init(svcCtx, options)

	return ctx
}

// Context 表示单个 Actor 风格运行时的当前执行作用域。
// 实体、实体树、帧和插件的直接操作应在所属运行时 goroutine 中进行。
type Context interface {
	iContext
	ConcurrentContext
	corectx.CurrentContextProvider
	async.WaitGuard
	reinterpret.InstanceProvider
	extension.AddInProvider
	GCCollector

	// Name 返回运行时名称。
	Name() string
	// Id 返回运行时的持久化 ID。
	Id() uid.Id
	// Reflected 返回实际运行时上下文实例的反射值。
	Reflected() reflect.Value
	// Frame 返回帧统计接口；未启用帧循环时返回 nil。
	Frame() Frame
	// EntityManager 返回当前运行时的本地实体管理器。
	EntityManager() EntityManager
	// EntityTree 返回当前运行时的实体树。
	EntityTree() EntityTree
	// Managed 返回随运行时上下文统一解绑的事件句柄集合。
	Managed() *event.ManagedHandles

	IContextRunningEventTab
}

type iContext interface {
	init(svcCtx service.Context, options ContextOptions)
	getOptions() *ContextOptions
	emitEventRunningEvent(runningEvent RunningEvent, args ...any)
	setFrame(frame Frame)
	setCaller(caller Caller)
	getServiceContext() service.Context
	getAddInManager() AddInManager
	getScoped() *atomic.Bool
	gc()
}

// ContextBehavior 提供 Context 的默认实现。
// 扩展运行时上下文时应匿名嵌入该类型，并通过 InstanceFace 传入扩展实例。
type ContextBehavior struct {
	corectx.ContextBehavior
	svcCtx         service.Context
	options        ContextOptions
	reflected      reflect.Value
	frame          Frame
	entityManager  _EntityManager
	caller         Caller
	scoped         atomic.Bool
	gcList         []GC
	managed        event.ManagedHandles
	executorID     async.ExecutorID
	blockedFuture  atomic.Uint64
	lastWaitReject atomic.Uint64
	stringerCache  atomic.Pointer[string]

	contextRunningEventTab contextRunningEventTab
}

// Name 返回运行时名称。
func (ctx *ContextBehavior) Name() string {
	return ctx.options.Name
}

// Id 返回运行时的持久化 ID。
func (ctx *ContextBehavior) Id() uid.Id {
	return ctx.options.PersistId
}

// Reflected 返回实际运行时上下文实例的反射值。
func (ctx *ContextBehavior) Reflected() reflect.Value {
	return ctx.reflected
}

// Frame 返回帧统计接口；未启用帧循环时返回 nil。
func (ctx *ContextBehavior) Frame() Frame {
	return ctx.frame
}

// EntityManager 返回当前运行时的本地实体管理器。
func (ctx *ContextBehavior) EntityManager() EntityManager {
	return &ctx.entityManager
}

// EntityTree 返回当前运行时的实体树。
func (ctx *ContextBehavior) EntityTree() EntityTree {
	return &ctx.entityManager
}

// Managed 返回随运行时上下文统一解绑的事件句柄集合。
func (ctx *ContextBehavior) Managed() *event.ManagedHandles {
	return &ctx.managed
}

// EventContextRunningEvent 返回运行时运行事件。
func (ctx *ContextBehavior) EventContextRunningEvent() event.IEvent {
	return ctx.contextRunningEventTab.EventContextRunningEvent()
}

// CurrentContextCache 返回仅供运行时 goroutine 使用的上下文接口缓存。
func (ctx *ContextBehavior) CurrentContextCache() iface.Cache {
	return iface.Iface2Cache[Context](ctx.options.InstanceFace.Iface)
}

// BeforeFutureWait 实现 async.WaitGuard。Runtime Context 只允许读取已经完成的
// Future；对 pending Future 的等待会破坏 Actor 串行执行语义。
func (ctx *ContextBehavior) BeforeFutureWait(futureID async.FutureID, completionExecutorID async.ExecutorID) error {
	ctx.blockedFuture.Store(uint64(futureID))
	ctx.lastWaitReject.Store(uint64(futureID))
	if completionExecutorID != 0 && completionExecutorID == ctx.executorID {
		return ErrRuntimeSelfWait
	}
	return ErrBlockingWaitInRuntime
}

// AfterFutureWait 清除诊断中的等待 Future ID。
func (ctx *ContextBehavior) AfterFutureWait(futureID async.FutureID) {
	ctx.blockedFuture.CompareAndSwap(uint64(futureID), 0)
}

// InstanceFaceCache 返回上下文实例的接口缓存，用于 reinterpret.Cast。
func (ctx *ContextBehavior) InstanceFaceCache() iface.Cache {
	return ctx.options.InstanceFace.Cache
}

// CollectGC 将需要清理的对象加入本轮运行时 GC 队列。
func (ctx *ContextBehavior) CollectGC(gc GC) {
	if gc == nil || !gc.NeedGC() {
		return
	}

	ctx.gcList = append(ctx.gcList, gc)
}

func (ctx *ContextBehavior) init(svcCtx service.Context, options ContextOptions) {
	if svcCtx == nil {
		exception.Panicf("%w: %w: svcCtx is nil", ErrContext, exception.ErrArgs)
	}

	ctx.options = options
	ctx.executorID = async.GenExecutorID()

	if ctx.options.InstanceFace.IsNil() {
		ctx.options.InstanceFace = iface.NewFaceT[Context](ctx)
	}

	if ctx.options.Context == nil {
		ctx.options.Context = svcCtx
	}

	if ctx.options.PersistId.IsNil() {
		ctx.options.PersistId = uid.New()
	}

	if ctx.options.AddInManager == nil {
		ctx.options.AddInManager = NewAddInManager()
	}

	corectx.UnsafeContext(&ctx.ContextBehavior).Init(ctx.options.Context, ctx.options.AutoRecover, ctx.options.ReportError)

	ctx.svcCtx = svcCtx
	ctx.reflected = reflect.ValueOf(ctx.getInstance())
	ctx.contextRunningEventTab.SetPanicHandling(ctx.AutoRecover(), ctx.ReportError())

	ctx.entityManager.init(ctx.getInstance())

	event.UnsafeEvent(ctx.getAddInManager().EventInstallAddIn()).Ctrl().SetPanicHandling(ctx.AutoRecover(), ctx.ReportError())
	event.UnsafeEvent(ctx.getAddInManager().EventUninstallAddIn()).Ctrl().SetPanicHandling(ctx.AutoRecover(), ctx.ReportError())
	event.UnsafeEvent(ctx.getAddInManager().EventAddInStateChanged()).Ctrl().SetPanicHandling(ctx.AutoRecover(), ctx.ReportError())

	if ctx.options.RunningEventCB != nil {
		BindEventContextRunningEvent(ctx, HandleEventContextRunningEvent(ctx.options.RunningEventCB))
	}
	BindEventContextRunningEvent(ctx, HandleEventContextRunningEvent(ctx.entityManager.onContextRunningEvent))
}

func (ctx *ContextBehavior) getOptions() *ContextOptions {
	return &ctx.options
}

func (ctx *ContextBehavior) emitEventRunningEvent(runningEvent RunningEvent, args ...any) {
	_EmitEventContextRunningEvent(ctx, ctx.getInstance(), runningEvent, args...)

	switch runningEvent {
	case RunningEvent_Terminated:
		ctx.contextRunningEventTab.SetEnabled(false)
		ctx.managed.UnbindAllEventHandles()
	}
}

func (ctx *ContextBehavior) setFrame(frame Frame) {
	ctx.frame = frame
}

func (ctx *ContextBehavior) setCaller(caller Caller) {
	ctx.caller = caller
}

func (ctx *ContextBehavior) getServiceContext() service.Context {
	return ctx.svcCtx
}

func (ctx *ContextBehavior) getScoped() *atomic.Bool {
	return &ctx.scoped
}
