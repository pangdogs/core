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

	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/reinterpret"
)

// NewService 创建服务
func NewService(svcCtx service.Context, settings ...option.Setting[ServiceOptions]) Service {
	return UnsafeNewService(svcCtx, option.New(With.Service.Default(), settings...))
}

// Deprecated: UnsafeNewService 内部创建服务
func UnsafeNewService(svcCtx service.Context, options ServiceOptions) Service {
	var svc Service

	if !options.InstanceFace.IsNil() {
		svc = options.InstanceFace.Iface
	} else {
		svc = &ServiceBehavior{}
	}
	svc.init(svcCtx, options)

	return svc
}

// Service 服务
type Service interface {
	iService
	iWorker
	iServiceStats
	reinterpret.InstanceProvider

	// Context 获取服务上下文
	Context() service.Context
}

type iService interface {
	init(svcCtx service.Context, options ServiceOptions)
	getOptions() *ServiceOptions
}

type ServiceBehavior struct {
	ctx       service.Context
	options   ServiceOptions
	isRunning atomic.Bool
}

// Context 获取服务上下文
func (svc *ServiceBehavior) Context() service.Context {
	return svc.ctx
}

// InstanceFaceCache 支持重新解释类型
func (svc *ServiceBehavior) InstanceFaceCache() iface.Cache {
	return svc.options.InstanceFace.Cache
}

func (svc *ServiceBehavior) init(svcCtx service.Context, options ServiceOptions) {
	if svcCtx == nil {
		exception.Panicf("%w: %w: svcCtx is nil", ErrService, ErrArgs)
	}

	if !service.UnsafeContext(svcCtx).Scoped().CompareAndSwap(false, true) {
		exception.Panicf("%w: %w: svcCtx is already bound to another service scope", ErrService, ErrArgs)
	}

	svc.ctx = svcCtx
	svc.options = options

	if svc.options.InstanceFace.IsNil() {
		svc.options.InstanceFace = iface.NewFaceT[Service](svc)
	}

	svc.emitEventRunningEvent(service.RunningEvent_Birth)
}

func (svc *ServiceBehavior) getOptions() *ServiceOptions {
	return &svc.options
}

func (svc *ServiceBehavior) getInstance() Service {
	return svc.options.InstanceFace.Iface
}
