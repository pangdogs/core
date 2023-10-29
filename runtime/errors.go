package runtime

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
)

var (
	ErrContext   = fmt.Errorf("%w: runtime-context", internal.ErrGolaxy) // 运行时上下文错误
	ErrECTree    = fmt.Errorf("%w: ec-tree", ErrContext)                 // EC树错误
	ErrEntityMgr = fmt.Errorf("%w: entity-mgr", ErrContext)              // 实体管理器错误
	ErrFrame     = fmt.Errorf("%w: frame", ErrContext)                   // 帧错误
)
