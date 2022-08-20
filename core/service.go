package core

type Service interface {
	Runnable
	init(ctx ServiceContext, opts *ServiceOptions)
	getOptions() *ServiceOptions
	GetContext() ServiceContext
}

func ServiceGetOptions(serv Service) ServiceOptions {
	return *serv.getOptions()
}

func ServiceGetInheritor(serv Service) Face[Service] {
	return serv.getOptions().Inheritor
}

func ServiceGetInheritorIFace[T any](serv Service) T {
	return Cache2IFace[T](serv.getOptions().Inheritor.Cache)
}

func NewService(servCtx ServiceContext, optFuncs ...NewServiceOptionFunc) Service {
	opts := &ServiceOptions{}
	NewServiceOption.Default()(opts)

	for i := range optFuncs {
		optFuncs[i](opts)
	}

	if !opts.Inheritor.IsNil() {
		opts.Inheritor.IFace.init(servCtx, opts)
		return opts.Inheritor.IFace
	}

	serv := &ServiceBehavior{}
	serv.init(servCtx, opts)

	return serv.opts.Inheritor.IFace
}

type ServiceBehavior struct {
	opts ServiceOptions
	ctx  ServiceContext
}

func (serv *ServiceBehavior) init(servCtx ServiceContext, opts *ServiceOptions) {
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

func (serv *ServiceBehavior) getOptions() *ServiceOptions {
	return &serv.opts
}

func (serv *ServiceBehavior) GetContext() ServiceContext {
	return serv.ctx
}
