package golaxy

import (
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util"
)

// NewService 创建服务
func NewService(serviceCtx service.Context, options ...ServiceOption) Service {
	opts := ServiceOptions{}
	WithServiceOption{}.Default()(&opts)

	for i := range options {
		options[i](&opts)
	}

	return UnsafeNewService(serviceCtx, opts)
}

func UnsafeNewService(serviceCtx service.Context, options ServiceOptions) Service {
	if !options.CompositeFace.IsNil() {
		options.CompositeFace.Iface.init(serviceCtx, &options)
		return options.CompositeFace.Iface
	}

	service := &ServiceBehavior{}
	service.init(serviceCtx, &options)

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

func (_service *ServiceBehavior) init(serviceCtx service.Context, opts *ServiceOptions) {
	if serviceCtx == nil {
		panic("nil serviceCtx")
	}

	if opts == nil {
		panic("nil opts")
	}

	if !internal.UnsafeContext(serviceCtx).Paired() {
		panic("service context already paired")
	}

	_service.opts = *opts

	if _service.opts.CompositeFace.IsNil() {
		_service.opts.CompositeFace = util.NewFace[Service](_service)
	}

	_service.ctx = serviceCtx
}

func (_service *ServiceBehavior) getOptions() *ServiceOptions {
	return &_service.opts
}
