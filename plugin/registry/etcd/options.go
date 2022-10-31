package etcd

import (
	"crypto/tls"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"net"
	"time"
)

type Options struct {
	Username   string
	Password   string
	Endpoints  []string
	Timeout    time.Duration
	Secure     bool
	TLSConfig  *tls.Config
	ZapLogger  *zap.Logger
	ZapConfig  *zap.Config
	EtcdConfig *clientv3.Config
}

type WithOption func(options *Options)

func Default() WithOption {
	return func(options *Options) {
		Auth("", "")(options)
		Endpoints("127.0.0.1:2379")(options)
		Timeout(5 * time.Second)(options)
		Secure(false)(options)
		TLSConfig(nil)(options)
		ZapLogger(nil)(options)
		ZapConfig(nil)(options)
		EtcdConfig(nil)(options)
	}
}

func Auth(username, password string) WithOption {
	return func(options *Options) {
		options.Username = username
		options.Password = password
	}
}

func Endpoints(endpoints ...string) WithOption {
	return func(options *Options) {
		for _, endpoint := range endpoints {
			if _, _, err := net.SplitHostPort(endpoint); err != nil {
				panic(err)
			}
		}
		options.Endpoints = endpoints
	}
}

func Timeout(dur time.Duration) WithOption {
	return func(options *Options) {
		options.Timeout = dur
	}
}

func Secure(secure bool) WithOption {
	return func(o *Options) {
		o.Secure = secure
	}
}

func TLSConfig(config *tls.Config) WithOption {
	return func(o *Options) {
		o.TLSConfig = config
	}
}

func ZapLogger(zapLogger *zap.Logger) WithOption {
	return func(o *Options) {
		o.ZapLogger = zapLogger
	}
}

func ZapConfig(config *zap.Config) WithOption {
	return func(o *Options) {
		o.ZapConfig = config
	}
}

func EtcdConfig(config *clientv3.Config) WithOption {
	return func(o *Options) {
		o.EtcdConfig = config
	}
}
