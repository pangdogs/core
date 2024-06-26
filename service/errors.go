package service

import (
	"fmt"
	"git.golaxy.org/core/utils/exception"
)

var (
	ErrContext   = fmt.Errorf("%w: service-context", exception.ErrCore) // 服务上下文错误
	ErrEntityMgr = fmt.Errorf("%w: entity-mgr", ErrContext)             // 实体管理器错误
)
