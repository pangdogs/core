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
	"git.golaxy.org/core/ec/ectx"
	"git.golaxy.org/core/ec/pt"
	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/reinterpret"
	"git.golaxy.org/core/utils/uid"
	"reflect"
)

// NewContext 创建服务上下文
func NewContext(settings ...option.Setting[ContextOptions]) Context {
	return UnsafeNewContext(option.Make(With.Default(), settings...))
}

// Deprecated: UnsafeNewContext 内部创建服务上下文
func UnsafeNewContext(options ContextOptions) Context {
	if !options.InstanceFace.IsNil() {
		options.InstanceFace.Iface.init(options)
		return options.InstanceFace.Iface
	}

	ctx := &ContextBehavior{}
	ctx.init(options)

	return ctx.opts.InstanceFace.Iface
}

// Context 服务上下文
type Context interface {
	iContext
	ectx.Context
	reinterpret.InstanceProvider
	extension.AddInProvider
	pt.EntityPTProvider
	Caller
	fmt.Stringer

	// GetName 获取名称
	GetName() string
	// GetId 获取服务Id
	GetId() uid.Id
	// GetReflected 获取反射值
	GetReflected() reflect.Value
	// GetEntityManager 获取实体管理器
	GetEntityManager() EntityManager
}

type iContext interface {
	init(opts ContextOptions)
	getOptions() *ContextOptions
	changeRunningStatus(status RunningStatus, args ...any)
}

// ContextBehavior 服务上下文行为，在扩展服务上下文能力时，匿名嵌入至服务上下文结构体中
type ContextBehavior struct {
	ectx.ContextBehavior
	opts          ContextOptions
	reflected     reflect.Value
	entityManager _EntityManagerBehavior
}

// GetName 获取名称
func (ctx *ContextBehavior) GetName() string {
	return ctx.opts.Name
}

// GetId 获取服务Id
func (ctx *ContextBehavior) GetId() uid.Id {
	return ctx.opts.PersistId
}

// GetReflected 获取反射值
func (ctx *ContextBehavior) GetReflected() reflect.Value {
	return ctx.reflected
}

// GetEntityManager 获取实体管理器
func (ctx *ContextBehavior) GetEntityManager() EntityManager {
	return &ctx.entityManager
}

// GetInstanceFaceCache 支持重新解释类型
func (ctx *ContextBehavior) GetInstanceFaceCache() iface.Cache {
	return ctx.opts.InstanceFace.Cache
}

// String implements fmt.Stringer
func (ctx *ContextBehavior) String() string {
	return fmt.Sprintf(`{"id":%q, "name":%q}`, ctx.GetId(), ctx.GetName())
}

func (ctx *ContextBehavior) init(opts ContextOptions) {
	ctx.opts = opts

	if ctx.opts.InstanceFace.IsNil() {
		ctx.opts.InstanceFace = iface.MakeFaceT[Context](ctx)
	}

	if ctx.opts.Context == nil {
		ctx.opts.Context = context.Background()
	}

	if ctx.opts.PersistId.IsNil() {
		ctx.opts.PersistId = uid.New()
	}

	ectx.UnsafeContext(&ctx.ContextBehavior).Init(ctx.opts.Context, ctx.opts.AutoRecover, ctx.opts.ReportError)
	ctx.reflected = reflect.ValueOf(ctx.opts.InstanceFace.Iface)
	ctx.entityManager.init(ctx.opts.InstanceFace.Iface)
}

func (ctx *ContextBehavior) getOptions() *ContextOptions {
	return &ctx.opts
}

func (ctx *ContextBehavior) changeRunningStatus(status RunningStatus, args ...any) {
	ctx.opts.RunningHandler.Call(ctx.GetAutoRecover(), ctx.GetReportError(), nil, ctx.opts.InstanceFace.Iface, status, args...)
}
