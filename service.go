package galaxy

import (
	"github.com/golaxy-kit/golaxy/internal"
	"github.com/golaxy-kit/golaxy/service"
	"github.com/golaxy-kit/golaxy/util"
)

// NewService 创建服务
func NewService(serviceCtx service.Context, options ...ServiceOption) Service {
	opts := ServiceOptions{}
	WithServiceOption.Default()(&opts)

	for i := range options {
		options[i](&opts)
	}

	return UnsafeNewService(serviceCtx, opts)
}

func UnsafeNewService(serviceCtx service.Context, options ServiceOptions) Service {
	if !options.Inheritor.IsNil() {
		options.Inheritor.Iface.init(serviceCtx, &options)
		return options.Inheritor.Iface
	}

	service := &ServiceBehavior{}
	service.init(serviceCtx, &options)

	return service.opts.Inheritor.Iface
}

// Service 服务
type Service interface {
	internal.Running

	// GetContext 获取服务上下文
	GetContext() service.Context

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

	_service.opts = *opts

	if _service.opts.Inheritor.IsNil() {
		_service.opts.Inheritor = util.NewFace[Service](_service)
	}

	_service.ctx = serviceCtx
}

func (_service *ServiceBehavior) getOptions() *ServiceOptions {
	return &_service.opts
}
