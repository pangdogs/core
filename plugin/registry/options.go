package registry

import "time"

type RegisterOptions struct {
	TTL time.Duration
}

type WithRegisterOption func(options *RegisterOptions)

func TTL(ttl time.Duration) WithRegisterOption {
	return func(options *RegisterOptions) {
		options.TTL = ttl
	}
}
