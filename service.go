package golaxy

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/concurrent"
	"kit.golaxy.org/golaxy/service"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/option"
	"sync/atomic"
)

// NewService 创建服务
func NewService(ctx service.Context, settings ...option.Setting[ServiceOptions]) Service {
	return UnsafeNewService(ctx, option.Make(_ServiceOption{}.Default(), settings...))
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
	_Service
	Running

	// GetContext 获取服务上下文
	GetContext() service.Context
}

type _Service interface {
	init(ctx service.Context, opts ServiceOptions)
	getOptions() *ServiceOptions
}

type ServiceBehavior struct {
	ctx     service.Context
	opts    ServiceOptions
	started atomic.Bool
}

// GetContext 获取服务上下文
func (serv *ServiceBehavior) GetContext() service.Context {
	return serv.ctx
}

func (serv *ServiceBehavior) init(ctx service.Context, opts ServiceOptions) {
	if ctx == nil {
		panic(fmt.Errorf("%w: %w: ctx is nil", ErrService, ErrArgs))
	}

	if !concurrent.UnsafeContext(ctx).SetPaired(true) {
		panic(fmt.Errorf("%w: context already paired", ErrService))
	}

	serv.ctx = ctx
	serv.opts = opts

	if serv.opts.CompositeFace.IsNil() {
		serv.opts.CompositeFace = iface.MakeFace[Service](serv)
	}

	serv.changeRunningState(service.RunningState_Birth)
}

func (serv *ServiceBehavior) getOptions() *ServiceOptions {
	return &serv.opts
}
