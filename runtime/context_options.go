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

	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/uid"
)

type (
	RunningEventCB = generic.ActionVar2[Context, RunningEvent, any] // 运行时运行事件回调。
)

// ContextOptions 定义创建运行时上下文时使用的选项。
type ContextOptions struct {
	InstanceFace   iface.Face[Context] // 自定义上下文实例及其接口缓存。
	Context        context.Context     // 父上下文；nil 时使用所属服务上下文。
	AutoRecover    bool                // 回调发生 panic 时是否自动恢复。
	ReportError    chan error          // 自动恢复后接收 panic 错误的通道。
	Name           string              // 运行时名称。
	PersistID      uid.ID              // 运行时持久化 ID；为 Nil 时自动生成。
	AddInManager   AddInManager        // 运行时插件管理器；nil 时创建默认管理器。
	RunningEventCB RunningEventCB      // 运行时运行事件回调。
}

// With 提供 Runtime 上下文选项构造器。
var With _ContextOption

type _ContextOption struct{}

// Default 返回运行时上下文选项的默认设置。
func (_ContextOption) Default() option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		With.InstanceFace(iface.Face[Context]{}).Apply(options)
		With.Context(nil).Apply(options)
		With.PanicHandling(false, nil).Apply(options)
		With.Name("").Apply(options)
		With.PersistID(uid.Nil).Apply(options)
		With.AddInManager(nil).Apply(options)
		With.RunningEventCB(nil).Apply(options)
	}
}

// InstanceFace 设置用于扩展运行时上下文能力的自定义实例。
func (_ContextOption) InstanceFace(face iface.Face[Context]) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.InstanceFace = face
	}
}

// Context 设置父上下文。
func (_ContextOption) Context(ctx context.Context) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.Context = ctx
	}
}

// PanicHandling 设置回调 panic 的自动恢复和错误上报方式。
func (_ContextOption) PanicHandling(autoRecover bool, reportError chan error) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.AutoRecover = autoRecover
		options.ReportError = reportError
	}
}

// Name 设置运行时名称。
func (_ContextOption) Name(name string) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.Name = name
	}
}

// PersistID 设置运行时持久化 ID。
func (_ContextOption) PersistID(id uid.ID) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.PersistID = id
	}
}

// AddInManager 设置运行时插件管理器。
func (_ContextOption) AddInManager(mgr AddInManager) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.AddInManager = mgr
	}
}

// RunningEventCB 设置运行时运行事件回调。
func (_ContextOption) RunningEventCB(cb RunningEventCB) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.RunningEventCB = cb
	}
}
