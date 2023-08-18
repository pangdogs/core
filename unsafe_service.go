package golaxy

import "kit.golaxy.org/golaxy/service"

// Deprecated: UnsafeService 访问服务内部方法
func UnsafeService(service Service) _UnsafeService {
	return _UnsafeService{
		Service: service,
	}
}

type _UnsafeService struct {
	Service
}

// Init 初始化
func (us _UnsafeService) Init(ctx service.Context, opts *ServiceOptions) {
	us.init(ctx, opts)
}

// GetOptions 获取服务所有选项
func (us _UnsafeService) GetOptions() *ServiceOptions {
	return us.getOptions()
}
