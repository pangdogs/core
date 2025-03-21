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
	"fmt"
	"git.golaxy.org/core/ec/ectx"
	"git.golaxy.org/core/event"
	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/reinterpret"
	"git.golaxy.org/core/utils/uid"
	"reflect"
)

// NewContext 创建运行时上下文
func NewContext(svcCtx service.Context, settings ...option.Setting[ContextOptions]) Context {
	return UnsafeNewContext(svcCtx, option.Make(With.Context.Default(), settings...))
}

// Deprecated: UnsafeNewContext 内部创建运行时上下文
func UnsafeNewContext(svcCtx service.Context, options ContextOptions) Context {
	if !options.InstanceFace.IsNil() {
		options.InstanceFace.Iface.init(svcCtx, options)
		return options.InstanceFace.Iface
	}

	ctx := &ContextBehavior{}
	ctx.init(svcCtx, options)

	return ctx.opts.InstanceFace.Iface
}

// Context 运行时上下文接口
type Context interface {
	iContext
	iConcurrentContext
	ectx.Context
	ectx.CurrentContextProvider
	reinterpret.InstanceProvider
	extension.AddInProvider
	async.Caller
	GCCollector
	fmt.Stringer

	// GetName 获取名称
	GetName() string
	// GetId 获取运行时Id
	GetId() uid.Id
	// GetReflected 获取反射值
	GetReflected() reflect.Value
	// GetFrame 获取帧
	GetFrame() Frame
	// GetEntityManager 获取实体管理器
	GetEntityManager() EntityManager
	// GetEntityTree 获取实体树
	GetEntityTree() EntityTree
	// ActivateEvent 启用事件
	ActivateEvent(event event.IEventCtrl, recursion event.EventRecursion)
	// ManagedAddHooks 托管事件钩子（event.Hook），在运行时停止时自动解绑定
	ManagedAddHooks(hooks ...event.Hook)
	// ManagedAddTagHooks 根据标签托管事件钩子（event.Hook），在运行时停止时自动解绑定
	ManagedAddTagHooks(tag string, hooks ...event.Hook)
	// ManagedGetTagHooks 根据标签获取托管事件钩子（event.Hook）
	ManagedGetTagHooks(tag string) []event.Hook
	// ManagedCleanTagHooks 清理根据标签托管的事件钩子（event.Hook）
	ManagedCleanTagHooks(tag string)
}

type iContext interface {
	init(svcCtx service.Context, opts ContextOptions)
	getOptions() *ContextOptions
	setFrame(frame Frame)
	setCallee(callee async.Callee)
	getServiceCtx() service.Context
	changeRunningStatus(status RunningStatus, args ...any)
	gc()
}

// ContextBehavior 运行时上下文行为，在扩展运行时上下文能力时，匿名嵌入至运行时上下文结构体中
type ContextBehavior struct {
	ectx.ContextBehavior
	svcCtx          service.Context
	opts            ContextOptions
	reflected       reflect.Value
	frame           Frame
	entityManager   _EntityManagerBehavior
	callee          async.Callee
	managedHooks    []event.Hook
	managedTagHooks generic.SliceMap[string, []event.Hook]
	gcList          []GC
}

// GetName 获取名称
func (ctx *ContextBehavior) GetName() string {
	return ctx.opts.Name
}

// GetId 获取运行时Id
func (ctx *ContextBehavior) GetId() uid.Id {
	return ctx.opts.PersistId
}

// GetReflected 获取反射值
func (ctx *ContextBehavior) GetReflected() reflect.Value {
	return ctx.reflected
}

// GetFrame 获取帧
func (ctx *ContextBehavior) GetFrame() Frame {
	return ctx.frame
}

// GetEntityManager 获取实体管理器
func (ctx *ContextBehavior) GetEntityManager() EntityManager {
	return &ctx.entityManager
}

// GetEntityTree 获取主实体树
func (ctx *ContextBehavior) GetEntityTree() EntityTree {
	return &ctx.entityManager
}

// ActivateEvent 启用事件
func (ctx *ContextBehavior) ActivateEvent(event event.IEventCtrl, recursion event.EventRecursion) {
	if event == nil {
		exception.Panicf("%w: %w: event is nil", ErrContext, exception.ErrArgs)
	}
	event.Init(ctx.GetAutoRecover(), ctx.GetReportError(), recursion)
}

// GetCurrentContext 获取当前上下文
func (ctx *ContextBehavior) GetCurrentContext() iface.Cache {
	return iface.Iface2Cache[Context](ctx.opts.InstanceFace.Iface)
}

// GetConcurrentContext 获取多线程安全的上下文
func (ctx *ContextBehavior) GetConcurrentContext() iface.Cache {
	return iface.Iface2Cache[Context](ctx.opts.InstanceFace.Iface)
}

// GetInstanceFaceCache 支持重新解释类型
func (ctx *ContextBehavior) GetInstanceFaceCache() iface.Cache {
	return ctx.opts.InstanceFace.Cache
}

// CollectGC 收集GC
func (ctx *ContextBehavior) CollectGC(gc GC) {
	if gc == nil || !gc.NeedGC() {
		return
	}

	ctx.gcList = append(ctx.gcList, gc)
}

// String implements fmt.Stringer
func (ctx *ContextBehavior) String() string {
	return fmt.Sprintf(`{"id":%q, "name":%q}`, ctx.GetId(), ctx.GetName())
}

func (ctx *ContextBehavior) init(svcCtx service.Context, opts ContextOptions) {
	if svcCtx == nil {
		exception.Panicf("%w: %w: svcCtx is nil", ErrContext, exception.ErrArgs)
	}

	ctx.opts = opts

	if ctx.opts.InstanceFace.IsNil() {
		ctx.opts.InstanceFace = iface.MakeFaceT[Context](ctx)
	}

	if ctx.opts.Context == nil {
		ctx.opts.Context = svcCtx
	}

	if ctx.opts.PersistId.IsNil() {
		ctx.opts.PersistId = uid.New()
	}

	ectx.UnsafeContext(&ctx.ContextBehavior).Init(ctx.opts.Context, ctx.opts.AutoRecover, ctx.opts.ReportError)
	ctx.svcCtx = svcCtx
	ctx.reflected = reflect.ValueOf(ctx.opts.InstanceFace.Iface)
	ctx.entityManager.init(ctx.opts.InstanceFace.Iface)
}

func (ctx *ContextBehavior) getOptions() *ContextOptions {
	return &ctx.opts
}

func (ctx *ContextBehavior) setFrame(frame Frame) {
	ctx.frame = frame
}

func (ctx *ContextBehavior) setCallee(callee async.Callee) {
	ctx.callee = callee
}

func (ctx *ContextBehavior) getServiceCtx() service.Context {
	return ctx.svcCtx
}

func (ctx *ContextBehavior) changeRunningStatus(status RunningStatus, args ...any) {
	ctx.entityManager.changeRunningStatus(status, args...)

	ctx.opts.RunningStatusChangedCB.Call(ctx.GetAutoRecover(), ctx.GetReportError(), ctx.opts.InstanceFace.Iface, status, args...)

	switch status {
	case RunningStatus_Terminated:
		ctx.managedCleanAllHooks()
	}
}
