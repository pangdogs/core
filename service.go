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
	"fmt"
	"git.golaxy.org/core/internal/gctx"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/iface"
	"git.golaxy.org/core/utils/option"
	"git.golaxy.org/core/utils/reinterpret"
)

// NewService 创建服务
func NewService(ctx service.Context, settings ...option.Setting[ServiceOptions]) Service {
	return UnsafeNewService(ctx, option.Make(With.Service.Default(), settings...))
}

// Deprecated: UnsafeNewService 内部创建服务
func UnsafeNewService(ctx service.Context, options ServiceOptions) Service {
	if !options.CompositeFace.IsNil() {
		options.CompositeFace.Iface.init(ctx, options)
		return options.CompositeFace.Iface
	}

	service := &ServiceBehavior{}
	service.init(ctx, options)

	return service.opts.CompositeFace.Iface
}

// Service 服务
type Service interface {
	iService
	reinterpret.CompositeProvider
	Running

	// GetContext 获取服务上下文
	GetContext() service.Context
}

type iService interface {
	init(ctx service.Context, opts ServiceOptions)
	getOptions() *ServiceOptions
}

type ServiceBehavior struct {
	ctx  service.Context
	opts ServiceOptions
}

// GetContext 获取服务上下文
func (serv *ServiceBehavior) GetContext() service.Context {
	return serv.ctx
}

// GetCompositeFaceCache 支持重新解释类型
func (serv *ServiceBehavior) GetCompositeFaceCache() iface.Cache {
	return serv.opts.CompositeFace.Cache
}

func (serv *ServiceBehavior) init(ctx service.Context, opts ServiceOptions) {
	if ctx == nil {
		panic(fmt.Errorf("%w: %w: ctx is nil", ErrService, ErrArgs))
	}

	if !gctx.UnsafeContext(ctx).SetPaired(true) {
		panic(fmt.Errorf("%w: context already paired", ErrService))
	}

	serv.ctx = ctx
	serv.opts = opts

	if serv.opts.CompositeFace.IsNil() {
		serv.opts.CompositeFace = iface.MakeFaceT[Service](serv)
	}

	serv.changeRunningState(service.RunningState_Birth)
}

func (serv *ServiceBehavior) getOptions() *ServiceOptions {
	return &serv.opts
}
