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
	RunningEventCB = generic.ActionVar2[Context, RunningEvent, any] // 运行事件回调
)

// ContextOptions 创建运行时上下文的所有选项
type ContextOptions struct {
	InstanceFace   iface.Face[Context]           // 实例，用于扩展运行时上下文能力
	Context        context.Context               // 父Context
	AutoRecover    bool                          // 是否开启panic时自动恢复
	ReportError    chan error                    // panic时错误写入的error channel
	Name           string                        // 运行时名称
	PersistId      uid.Id                        // 运行时持久化Id
	AddInManager   extension.RuntimeAddInManager // 插件管理器
	RunningEventCB RunningEventCB                // 运行事件回调
}

type _ContextOption struct{}

// Default 默认值
func (_ContextOption) Default() option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		With.Context.InstanceFace(iface.Face[Context]{}).Apply(options)
		With.Context.Context(nil).Apply(options)
		With.Context.PanicHandling(false, nil).Apply(options)
		With.Context.Name("").Apply(options)
		With.Context.PersistId(uid.Nil).Apply(options)
		With.Context.AddInManager(nil).Apply(options)
		With.Context.RunningEventCB(nil).Apply(options)
	}
}

// InstanceFace 实例，用于扩展运行时上下文能力
func (_ContextOption) InstanceFace(face iface.Face[Context]) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.InstanceFace = face
	}
}

// Context 父Context
func (_ContextOption) Context(ctx context.Context) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.Context = ctx
	}
}

// PanicHandling panic时的处理方式
func (_ContextOption) PanicHandling(autoRecover bool, reportError chan error) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.AutoRecover = autoRecover
		options.ReportError = reportError
	}
}

// Name 运行时名称
func (_ContextOption) Name(name string) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.Name = name
	}
}

// PersistId 运行时持久化Id
func (_ContextOption) PersistId(id uid.Id) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.PersistId = id
	}
}

// AddInManager 插件管理器
func (_ContextOption) AddInManager(mgr extension.RuntimeAddInManager) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.AddInManager = mgr
	}
}

// RunningEventCB 运行事件回调
func (_ContextOption) RunningEventCB(cb RunningEventCB) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.RunningEventCB = cb
	}
}
