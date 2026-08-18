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

package service

import (
	"context"
	"fmt"
	"reflect"
	"sync/atomic"

	"git.golaxy.org/core/ec/pt"
	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/utils/corectx"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/reinterpret"
	"git.golaxy.org/core/utils/uid"
)

// NewContext 创建服务上下文。
// 未提供父上下文、持久化 ID、原型库或插件管理器时会自动创建默认值。
func NewContext(settings ...option.Setting[ContextOptions]) Context {
	return UnsafeNewContext(option.New(With.Default(), settings...))
}

// Deprecated: UnsafeNewContext 仅供框架内部使用，请改用 NewContext。
func UnsafeNewContext(options ContextOptions) Context {
	var ctx Context

	if !options.InstanceFace.IsNil() {
		ctx = options.InstanceFace.Iface
	} else {
		ctx = &ContextBehavior{}
	}
	ctx.init(options)

	return ctx
}

// Context 表示应用级共享作用域，并为多个 runtime 提供父上下文和全局资源。
type Context interface {
	iContext
	corectx.Context
	reinterpret.InstanceProvider
	extension.AddInProvider
	pt.EntityPTProvider
	Caller
	fmt.Stringer

	// Name 返回服务名称。
	Name() string
	// ID 返回服务的持久化 ID。
	ID() uid.ID
	// Reflected 返回实际服务上下文实例的反射值。
	Reflected() reflect.Value
	// EntityManager 返回并发安全的全局实体索引。
	EntityManager() EntityManager
}

type iContext interface {
	init(options ContextOptions)
	getOptions() *ContextOptions
	getInstance() Context
	emitEventRunningEvent(runningEvent RunningEvent, args ...any)
	getAddInManager() AddInManager
	getScoped() *atomic.Bool
}

// ContextBehavior 提供 Context 的默认实现。
// 扩展服务上下文时应匿名嵌入该类型，并通过 InstanceFace 传入扩展实例。
type ContextBehavior struct {
	corectx.ContextBehavior
	options       ContextOptions
	reflected     reflect.Value
	entityManager _EntityManager
	scoped        atomic.Bool
	stringerCache atomic.Pointer[string]
}

// Name 返回服务名称。
func (ctx *ContextBehavior) Name() string {
	return ctx.options.Name
}

// ID 返回服务的持久化 ID。
func (ctx *ContextBehavior) ID() uid.ID {
	return ctx.options.PersistID
}

// Reflected 返回实际服务上下文实例的反射值。
func (ctx *ContextBehavior) Reflected() reflect.Value {
	return ctx.reflected
}

// EntityManager 返回并发安全的全局实体索引。
func (ctx *ContextBehavior) EntityManager() EntityManager {
	return &ctx.entityManager
}

// InstanceFaceCache 返回上下文实例的接口缓存，用于 reinterpret.Cast。
func (ctx *ContextBehavior) InstanceFaceCache() iface.Cache {
	return ctx.options.InstanceFace.Cache
}

// String 实现 fmt.Stringer，返回包含服务 ID 和名称的 JSON 文本。
func (ctx *ContextBehavior) String() string {
	if cached := ctx.stringerCache.Load(); cached != nil {
		return *cached
	}

	value := fmt.Sprintf(`{"id":%q,"name":%q}`, ctx.ID(), ctx.Name())
	if ctx.stringerCache.CompareAndSwap(nil, &value) {
		return value
	}
	return *ctx.stringerCache.Load()
}

func (ctx *ContextBehavior) init(options ContextOptions) {
	ctx.options = options

	if ctx.options.InstanceFace.IsNil() {
		ctx.options.InstanceFace = iface.NewFaceT[Context](ctx)
	}

	if ctx.options.Context == nil {
		ctx.options.Context = context.Background()
	}

	if ctx.options.PersistID.IsNil() {
		ctx.options.PersistID = uid.New()
	}

	if ctx.options.EntityLib == nil {
		ctx.options.EntityLib = pt.NewEntityLib(pt.DefaultComponentLib())
	}

	if ctx.options.AddInManager == nil {
		ctx.options.AddInManager = NewAddInManager()
	}

	corectx.UnsafeContext(&ctx.ContextBehavior).Init(ctx.options.Context, ctx.options.AutoRecover, ctx.options.ReportError)
	ctx.reflected = reflect.ValueOf(ctx.getInstance())
	ctx.entityManager.init(ctx.getInstance())
}

func (ctx *ContextBehavior) getOptions() *ContextOptions {
	return &ctx.options
}

func (ctx *ContextBehavior) getInstance() Context {
	return ctx.options.InstanceFace.Iface
}

func (ctx *ContextBehavior) emitEventRunningEvent(runningEvent RunningEvent, args ...any) {
	ctx.options.RunningEventCB.Call(ctx.AutoRecover(), ctx.ReportError(), ctx.getInstance(), runningEvent, args...)
}

func (ctx *ContextBehavior) getScoped() *atomic.Bool {
	return &ctx.scoped
}
