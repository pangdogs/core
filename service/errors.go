package service

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/errors"
)

var (
	ErrContext   = fmt.Errorf("%w: service-context", errors.ErrGolaxy) // 服务上下文错误
	ErrEntityMgr = fmt.Errorf("%w: entity-mgr", ErrContext)            // 实体管理器错误
)
