package runtime

import (
	"fmt"
	"git.golaxy.org/core/internal/gctx"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/uid"
)

// ConcurrentContext 多线程安全的运行时上下文接口
type ConcurrentContext interface {
	gctx.ConcurrentContextProvider
	gctx.Context
	async.Caller
	fmt.Stringer

	// GetName 获取名称
	GetName() string
	// GetId 获取运行时Id
	GetId() uid.Id
}
