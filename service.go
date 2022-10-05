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

	if !opts.Inheritor.IsNil() {
		opts.Inheritor.Iface.init(serviceCtx, &opts)
		return opts.Inheritor.Iface
	}

	service := &_ServiceBehavior{}
	service.init(serviceCtx, &opts)

	return service.opts.Inheritor.Iface
}

type _ServiceBehavior struct {
	opts ServiceOptions
	ctx  service.Context
}

func (serv *_ServiceBehavior) init(serviceCtx service.Context, opts *ServiceOptions) {
	if serviceCtx == nil {
		panic("nil serviceCtx")
	}

	if opts == nil {
		panic("nil opts")
	}

	serv.opts = *opts

	if serv.opts.Inheritor.IsNil() {
		serv.opts.Inheritor = util.NewFace[Service](serv)
	}

	serv.ctx = serviceCtx
}

func (serv *_ServiceBehavior) getOptions() *ServiceOptions {
	return &serv.opts
}

// GetContext 获取服务上下文
func (serv *_ServiceBehavior) GetContext() service.Context {
	return serv.ctx
}
