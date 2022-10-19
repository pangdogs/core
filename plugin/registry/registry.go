package registry

import (
	"context"
	"time"
)

// Registry 分布式服务注册器
type Registry interface {
	Register(ctx context.Context, service Service, ttl time.Duration) error
	Deregister(ctx context.Context, service Service) error
	GetService(ctx context.Context, serviceName string) ([]Service, error)
	ListServices(ctx context.Context) ([]Service, error)
	Watch(ctx context.Context, serviceName string) (Watcher, error)
}

type Service struct {
	Name      string            `json:"name"`
	Version   string            `json:"version"`
	Metadata  map[string]string `json:"metadata"`
	Endpoints []*Endpoint       `json:"endpoints"`
	Nodes     []*Node           `json:"nodes"`
}

type Node struct {
	Id       string            `json:"id"`
	Address  string            `json:"address"`
	Metadata map[string]string `json:"metadata"`
}

type Endpoint struct {
	Name     string            `json:"name"`
	Request  *Value            `json:"request"`
	Response *Value            `json:"response"`
	Metadata map[string]string `json:"metadata"`
}

type Value struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Values []*Value `json:"values"`
}
