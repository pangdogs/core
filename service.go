package golaxy

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/iface"
)

// NewService 创建服务
func NewService(ctx service.Context, options ...ServiceOption) Service {
	opts := ServiceOptions{}
	Option{}.Service.Default()(&opts)

	for i := range options {
		options[i](&opts)
	}

	return UnsafeNewService(ctx, opts)
}

// Deprecated: UnsafeNewService 内部创建服务
func UnsafeNewService(ctx service.Context, options ServiceOptions) Service {
	if !options.CompositeFace.IsNil() {
		options.CompositeFace.Iface.init(ctx, &options)
		return options.CompositeFace.Iface
	}

	service := &ServiceBehavior{}
	service.init(ctx, &options)

	return service.opts.CompositeFace.Iface
}

// Service 服务
type Service interface {
	_Service
	internal.Running

	// GetContext 获取服务上下文
	GetContext() service.Context
}

type _Service interface {
	init(ctx service.Context, opts *ServiceOptions)
	getOptions() *ServiceOptions
}

type ServiceBehavior struct {
	opts ServiceOptions
	ctx  service.Context
}

// GetContext 获取服务上下文
func (_service *ServiceBehavior) GetContext() service.Context {
	return _service.ctx
}

func (_service *ServiceBehavior) init(ctx service.Context, opts *ServiceOptions) {
	if ctx == nil {
		panic(fmt.Errorf("%w: %w: ctx is nil", ErrService, ErrArgs))
	}

	if opts == nil {
		panic(fmt.Errorf("%w: %w: opts is nil", ErrService, ErrArgs))
	}

	if !service.UnsafeContext(ctx).MarkPaired(true) {
		panic(fmt.Errorf("%w: ctx already paired", ErrService))
	}

	_service.opts = *opts

	if _service.opts.CompositeFace.IsNil() {
		_service.opts.CompositeFace = iface.NewFace[Service](_service)
	}

	_service.ctx = ctx

	_service.changeRunningState(service.RunningState_Birth)
}

func (_service *ServiceBehavior) getOptions() *ServiceOptions {
	return &_service.opts
}
