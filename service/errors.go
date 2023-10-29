package service

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
)

var (
	ErrContext   = fmt.Errorf("%w: service-context", internal.ErrGolaxy) // 服务上下文错误
	ErrEntityMgr = fmt.Errorf("%w: entity-mgr", ErrContext)              // 实体管理器错误
)
