package registry

import (
	"context"
	"github.com/pangdogs/galaxy/service"
	"time"
)

func Register(serviceCtx service.Context, ctx context.Context, service Service, ttl time.Duration) error {
	return Plugin.Get(serviceCtx).Register(ctx, service, ttl)
}

func Deregister(serviceCtx service.Context, ctx context.Context, service Service) error {
	return Plugin.Get(serviceCtx).Deregister(ctx, service)
}

func GetService(serviceCtx service.Context, ctx context.Context, serviceName string) ([]Service, error) {
	return Plugin.Get(serviceCtx).GetService(ctx, serviceName)
}

func ListServices(serviceCtx service.Context, ctx context.Context) ([]Service, error) {
	return Plugin.Get(serviceCtx).ListServices(ctx)
}

func Watch(serviceCtx service.Context, ctx context.Context, serviceName string) (Watcher, error) {
	return Plugin.Get(serviceCtx).Watch(ctx, serviceName)
}
