package core

var NewServiceOption = &NewServiceOptions{}

type ServiceOptions struct {
	Inheritor         Face[Service]
	EnableAutoRecover bool
}

type NewServiceOptionFunc func(o *ServiceOptions)

type NewServiceOptions struct{}

func (*NewServiceOptions) Default() NewServiceOptionFunc {
	return func(o *ServiceOptions) {
		o.Inheritor = Face[Service]{}
		o.EnableAutoRecover = false
	}
}

func (*NewServiceOptions) Inheritor(v Face[Service]) NewServiceOptionFunc {
	return func(o *ServiceOptions) {
		o.Inheritor = v
	}
}

func (*NewServiceOptions) EnableAutoRecover(v bool) NewServiceOptionFunc {
	return func(o *ServiceOptions) {
		o.EnableAutoRecover = v
	}
}
