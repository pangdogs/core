package internal

import (
	"kit.golaxy.org/golaxy/util/iface"
)

// ContextResolver 上下文获取器
type ContextResolver interface {
	// ResolveContext 解析上下文
	ResolveContext() iface.Cache
}
