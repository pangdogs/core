package golaxy

import "kit.golaxy.org/golaxy/service"

func UnsafeService(service Service) _UnsafeService {
	return _UnsafeService{
		Service: service,
	}
}

type _UnsafeService struct {
	Service
}

func (us _UnsafeService) Init(ctx service.Context, opts *ServiceOptions) {
	us.init(ctx, opts)
}

func (us _UnsafeService) GetOptions() *ServiceOptions {
	return us.getOptions()
}
