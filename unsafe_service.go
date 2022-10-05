package galaxy

import "github.com/pangdogs/galaxy/service"

func UnsafeService(service Service) _UnsafeService {
	return _UnsafeService{
		Service: service,
	}
}

type _UnsafeService struct {
	Service
}

func (us _UnsafeService) Init(serviceCtx service.Context, opts *ServiceOptions) {
	us.init(serviceCtx, opts)
}

func (us _UnsafeService) GetOptions() *ServiceOptions {
	return us.getOptions()
}
