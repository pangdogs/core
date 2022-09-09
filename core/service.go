package core

// Service 服务
type Service interface {
	_Runnable

	init(ctx ServiceContext, opts *ServiceOptions)

	getOptions() *ServiceOptions

	// GetContext 获取服务上下文（Service Context）
	GetContext() ServiceContext
}

// ServiceGetOptions 获取服务创建选项，线程安全
func ServiceGetOptions(serv Service) ServiceOptions {
	return *serv.getOptions()
}

// NewService 创建服务，线程安全
func NewService(servCtx ServiceContext, optSetterFuncs ...ServiceOptionSetterFunc) Service {
	opts := ServiceOptions{}
	ServiceOptionSetter.Default()(&opts)

	for i := range optSetterFuncs {
		optSetterFuncs[i](&opts)
	}

	return NewServiceWithOpts(servCtx, opts)
}

// NewServiceWithOpts 创建服务并传入参数，线程安全
func NewServiceWithOpts(servCtx ServiceContext, opts ServiceOptions) Service {
	if !opts.Inheritor.IsNil() {
		opts.Inheritor.IFace.init(servCtx, &opts)
		return opts.Inheritor.IFace
	}

	e := &_ServiceBehavior{}
	e.init(servCtx, &opts)

	return e.opts.Inheritor.IFace
}

type _ServiceBehavior struct {
	opts ServiceOptions
	ctx  ServiceContext
}

func (serv *_ServiceBehavior) init(servCtx ServiceContext, opts *ServiceOptions) {
	if servCtx == nil {
		panic("nil servCtx")
	}

	if opts == nil {
		panic("nil opts")
	}

	serv.opts = *opts

	if serv.opts.Inheritor.IsNil() {
		serv.opts.Inheritor = NewFace[Service](serv)
	}

	serv.ctx = servCtx
}

func (serv *_ServiceBehavior) getOptions() *ServiceOptions {
	return &serv.opts
}

// GetContext 获取服务上下文（Service Context）
func (serv *_ServiceBehavior) GetContext() ServiceContext {
	return serv.ctx
}
