package core

var NewServiceContextOption = &NewServiceContextOptions{}

type ServiceContextOptions struct {
	Inheritor   Face[ServiceContext]
	ReportError chan error
	StartedCallback,
	StoppingCallback,
	StoppedCallback func(serv Service)
	Params ServiceParams
}

type NewServiceContextOptionFunc func(o *ServiceContextOptions)

type NewServiceContextOptions struct{}

func (*NewServiceContextOptions) Default() NewServiceContextOptionFunc {
	return func(o *ServiceContextOptions) {
		o.Inheritor = Face[ServiceContext]{}
		o.ReportError = nil
		o.StartedCallback = nil
		o.StoppingCallback = nil
		o.StoppedCallback = nil
	}
}

func (*NewServiceContextOptions) Inheritor(v Face[ServiceContext]) NewServiceContextOptionFunc {
	return func(o *ServiceContextOptions) {
		o.Inheritor = v
	}
}

func (*NewServiceContextOptions) ReportError(v chan error) NewServiceContextOptionFunc {
	return func(o *ServiceContextOptions) {
		o.ReportError = v
	}
}

func (*NewServiceContextOptions) StartedCallback(v func(serv Service)) NewServiceContextOptionFunc {
	return func(o *ServiceContextOptions) {
		o.StartedCallback = v
	}
}

func (*NewServiceContextOptions) StoppingCallback(v func(serv Service)) NewServiceContextOptionFunc {
	return func(o *ServiceContextOptions) {
		o.StoppingCallback = v
	}
}

func (*NewServiceContextOptions) StoppedCallback(v func(serv Service)) NewServiceContextOptionFunc {
	return func(o *ServiceContextOptions) {
		o.StoppedCallback = v
	}
}

type ServiceParams struct {
	PersistID string
	Prototype string
}

func (*NewServiceContextOptions) Params(v ServiceParams) NewServiceContextOptionFunc {
	return func(o *ServiceContextOptions) {
		o.Params = v
	}
}
