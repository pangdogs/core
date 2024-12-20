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
	"context"
	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/uid"
)

type (
	RunningHandler = generic.DelegateActionVar2[Context, RunningState, any] // 运行状态变化处理器
)

// ContextOptions 创建运行时上下文的所有选项
type ContextOptions struct {
	InstanceFace   iface.Face[Context]    // 实例，用于扩展运行时上下文能力
	Context        context.Context        // 父Context
	AutoRecover    bool                   // 是否开启panic时自动恢复
	ReportError    chan error             // panic时错误写入的error channel
	Name           string                 // 运行时名称
	PersistId      uid.Id                 // 运行时持久化Id
	AddInManager   extension.AddInManager // 插件管理器
	RunningHandler RunningHandler         // 运行状态变化处理器
}

type _ContextOption struct{}

// Default 默认值
func (_ContextOption) Default() option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		With.Context.InstanceFace(iface.Face[Context]{})(o)
		With.Context.Context(nil)(o)
		With.Context.PanicHandling(false, nil)(o)
		With.Context.Name("")(o)
		With.Context.PersistId(uid.Nil)(o)
		With.Context.AddInManager(extension.NewAddInManager())(o)
		With.Context.RunningHandler(nil)(o)
	}
}

// InstanceFace 实例，用于扩展运行时上下文能力
func (_ContextOption) InstanceFace(face iface.Face[Context]) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.InstanceFace = face
	}
}

// Context 父Context
func (_ContextOption) Context(ctx context.Context) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.Context = ctx
	}
}

// PanicHandling panic时的处理方式
func (_ContextOption) PanicHandling(autoRecover bool, reportError chan error) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.AutoRecover = autoRecover
		o.ReportError = reportError
	}
}

// Name 运行时名称
func (_ContextOption) Name(name string) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.Name = name
	}
}

// PersistId 运行时持久化Id
func (_ContextOption) PersistId(id uid.Id) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.PersistId = id
	}
}

// AddInManager 插件管理器
func (_ContextOption) AddInManager(bundle extension.AddInManager) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.AddInManager = bundle
	}
}

// RunningHandler 运行状态变化处理器
func (_ContextOption) RunningHandler(handler RunningHandler) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.RunningHandler = handler
	}
}
