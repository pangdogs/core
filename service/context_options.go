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
	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/utils/generic"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/uid"
)

type (
	RunningHandler = generic.ActionVar2[Context, RunningStatus, any] // 运行状态变化处理器
)

// ContextOptions 创建服务上下文的所有选项
type ContextOptions struct {
	InstanceFace   iface.Face[Context]    // 实例，用于扩展服务上下文能力
	Context        context.Context        // 父Context
	AutoRecover    bool                   // 是否开启panic时自动恢复
	ReportError    chan error             // panic时错误写入的error channel
	Name           string                 // 服务名称
	PersistId      uid.Id                 // 服务持久化Id
	EntityLib      pt.EntityLib           // 实体原型库
	AddInManager   extension.AddInManager // 插件管理器
	RunningHandler RunningHandler         // 运行状态变化处理器
}

var With _Option

type _Option struct{}

// Default 默认值
func (_Option) Default() option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		With.InstanceFace(iface.Face[Context]{})(o)
		With.Context(nil)(o)
		With.PanicHandling(false, nil)(o)
		With.Name("")(o)
		With.PersistId(uid.Nil)(o)
		With.EntityLib(pt.NewEntityLib(pt.DefaultComponentLib()))(o)
		With.AddInManager(extension.NewAddInManager())(o)
		With.RunningHandler(nil)(o)
	}
}

// InstanceFace 实例，用于扩展服务上下文能力
func (_Option) InstanceFace(face iface.Face[Context]) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.InstanceFace = face
	}
}

// Context 父Context
func (_Option) Context(ctx context.Context) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.Context = ctx
	}
}

// PanicHandling panic时的处理方式
func (_Option) PanicHandling(autoRecover bool, reportError chan error) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.AutoRecover = autoRecover
		o.ReportError = reportError
	}
}

// Name 服务名称
func (_Option) Name(name string) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.Name = name
	}
}

// PersistId 服务持久化Id
func (_Option) PersistId(id uid.Id) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.PersistId = id
	}
}

// EntityLib 实体原型库
func (_Option) EntityLib(lib pt.EntityLib) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.EntityLib = lib
	}
}

// AddInManager 插件管理器
func (_Option) AddInManager(mgr extension.AddInManager) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.AddInManager = mgr
	}
}

// RunningHandler 运行状态变化处理器
func (_Option) RunningHandler(handler RunningHandler) option.Setting[ContextOptions] {
	return func(o *ContextOptions) {
		o.RunningHandler = handler
	}
}
