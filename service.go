package galaxy

import (
	"github.com/pangdogs/galaxy/internal"
	"github.com/pangdogs/galaxy/service"
	"github.com/pangdogs/galaxy/util"
)

// Service 服务
type Service interface {
	internal.Running

	init(ctx service.Context, opts *ServiceOptions)

	getOptions() *ServiceOptions

	// GetContext 获取服务上下文
	GetContext() service.Context
}

// NewService 创建服务
func NewService(serviceCtx service.Context, optSetter ...ServiceOptionSetterFunc) Service {
	opts := ServiceOptions{}
	ServiceOption.Default()(&opts)

	for i := range optSetter {
		optSetter[i](&opts)
	}

	return UnsafeNewService(serviceCtx, opts)
}

func UnsafeNewService(serviceCtx service.Context, opts ServiceOptions) Service {
	if !opts.Inheritor.IsNil() {
		opts.Inheritor.Iface.init(serviceCtx, &opts)
		return opts.Inheritor.Iface
	}

	service := &ServiceBehavior{}
	service.init(serviceCtx, &opts)

	return service.opts.Inheritor.Iface
}

type ServiceBehavior struct {
	opts ServiceOptions
	ctx  service.Context
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

// GetContext 获取服务上下文
func (_service *ServiceBehavior) GetContext() service.Context {
	return _service.ctx
}
