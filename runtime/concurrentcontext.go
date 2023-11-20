package runtime

import (
	"kit.golaxy.org/golaxy/internal/concurrent"
	"kit.golaxy.org/golaxy/util/uid"
)

// ConcurrentContext 多线程安全的运行时上下文接口
type ConcurrentContext interface {
	concurrent.ConcurrentContextProvider
	concurrent.Context
	concurrent.Caller

	// GetName 获取名称
	GetName() string
	// GetId 获取运行时Id
	GetId() uid.Id
}
