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
	"git.golaxy.org/core/ec/ictx"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/reinterpret"
	"sync"
	"sync/atomic"
)

// NewService 创建服务
func NewService(svcCtx service.Context, settings ...option.Setting[ServiceOptions]) Service {
	return UnsafeNewService(svcCtx, option.Make(With.Service.Default(), settings...))
}

// Deprecated: UnsafeNewService 内部创建服务
func UnsafeNewService(svcCtx service.Context, options ServiceOptions) Service {
	if !options.InstanceFace.IsNil() {
		options.InstanceFace.Iface.init(svcCtx, options)
		return options.InstanceFace.Iface
	}

	service := &ServiceBehavior{}
	service.init(svcCtx, options)

	return service.options.InstanceFace.Iface
}

// Service 服务
type Service interface {
	iService
	iRunning
	reinterpret.InstanceProvider

	// GetContext 获取服务上下文
	GetContext() service.Context
}

type iService interface {
	init(svcCtx service.Context, options ServiceOptions)
	getOptions() *ServiceOptions
}

type _StatusChanges struct {
	status service.RunningStatus
	args   []any
}

type ServiceBehavior struct {
	ctx               service.Context
	options           ServiceOptions
	isRunning         atomic.Bool
	statusChangesCond *sync.Cond
	statusChanges     *_StatusChanges
}

// GetContext 获取服务上下文
func (svc *ServiceBehavior) GetContext() service.Context {
	return svc.ctx
}

// GetInstanceFaceCache 支持重新解释类型
func (svc *ServiceBehavior) GetInstanceFaceCache() iface.Cache {
	return svc.options.InstanceFace.Cache
}

func (svc *ServiceBehavior) init(svcCtx service.Context, options ServiceOptions) {
	if svcCtx == nil {
		exception.Panicf("%w: %w: svcCtx is nil", ErrService, ErrArgs)
	}

	if !ictx.UnsafeContext(svcCtx).SetPaired(true) {
		exception.Panicf("%w: context already paired", ErrService)
	}

	svc.ctx = svcCtx
	svc.options = options
	svc.statusChangesCond = sync.NewCond(&sync.Mutex{})

	if svc.options.InstanceFace.IsNil() {
		svc.options.InstanceFace = iface.MakeFaceT[Service](svc)
	}

	svc.changeRunningStatus(service.RunningStatus_Birth)
}

func (svc *ServiceBehavior) getOptions() *ServiceOptions {
	return &svc.options
}
