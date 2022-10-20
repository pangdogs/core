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
	ZapConfig  *zap.Config
	EtcdConfig *clientv3.Config
}

type Option func(options *Options)

func Default() Option {
	return func(options *Options) {
		Auth("", "")
		Endpoints("127.0.0.1:2379")
		Timeout(5 * time.Second)
		Secure(false)
		TLSConfig(nil)
		ZapConfig(nil)
		EtcdConfig(nil)
	}
}

func Auth(username, password string) Option {
	return func(options *Options) {
		options.Username = username
		options.Password = password
	}
}

func Endpoints(endpoints ...string) Option {
	return func(options *Options) {
		for _, endpoint := range endpoints {
			if _, _, err := net.SplitHostPort(endpoint); err != nil {
				panic(err)
			}
		}
		options.Endpoints = endpoints
	}
}

func Timeout(dur time.Duration) Option {
	return func(options *Options) {
		options.Timeout = dur
	}
}

func Secure(secure bool) Option {
	return func(o *Options) {
		o.Secure = secure
	}
}

func TLSConfig(config *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = config
	}
}

func ZapConfig(config *zap.Config) Option {
	return func(o *Options) {
		o.ZapConfig = config
	}
}

func EtcdConfig(config *clientv3.Config) Option {
	return func(o *Options) {
		o.EtcdConfig = config
	}
}
