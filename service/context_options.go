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

	"git.golaxy.org/core/ec/pt"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/uid"
)

type (
	RunningEventCB = generic.ActionVar2[Context, RunningEvent, any] // 服务运行事件回调。
)

// ContextOptions 定义创建服务上下文时使用的选项。
type ContextOptions struct {
	InstanceFace   iface.Face[Context] // 自定义上下文实例及其接口缓存。
	Context        context.Context     // 父上下文；nil 时使用 context.Background。
	AutoRecover    bool                // 回调发生 panic 时是否自动恢复。
	ReportError    chan error          // 自动恢复后接收 panic 错误的通道。
	Name           string              // 服务名称。
	PersistID      uid.ID              // 服务持久化 ID；为 Nil 时自动生成。
	EntityLib      pt.EntityLib        // 实体原型库；nil 时创建独立实体库。
	AddInManager   AddInManager        // 服务插件管理器；nil 时创建默认管理器。
	RunningEventCB RunningEventCB      // 服务运行事件回调。
}

// With 提供 Service 上下文选项构造器。
var With _ContextOption

type _ContextOption struct{}

// Default 返回服务上下文选项的默认设置。
func (_ContextOption) Default() option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		With.InstanceFace(iface.Face[Context]{}).Apply(options)
		With.Context(nil).Apply(options)
		With.PanicHandling(false, nil).Apply(options)
		With.Name("").Apply(options)
		With.PersistID(uid.Nil).Apply(options)
		With.EntityLib(nil).Apply(options)
		With.AddInManager(nil).Apply(options)
		With.RunningEventCB(nil).Apply(options)
	}
}

// InstanceFace 设置用于扩展服务上下文能力的自定义实例。
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

// Name 设置服务名称。
func (_ContextOption) Name(name string) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.Name = name
	}
}

// PersistID 设置服务持久化 ID。
func (_ContextOption) PersistID(id uid.ID) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.PersistID = id
	}
}

// EntityLib 设置实体原型库。
func (_ContextOption) EntityLib(lib pt.EntityLib) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.EntityLib = lib
	}
}

// AddInManager 设置服务插件管理器。
func (_ContextOption) AddInManager(mgr AddInManager) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.AddInManager = mgr
	}
}

// RunningEventCB 设置服务运行事件回调。
func (_ContextOption) RunningEventCB(cb RunningEventCB) option.Setting[ContextOptions] {
	return func(options *ContextOptions) {
		options.RunningEventCB = cb
	}
}
