package ec

import (
	"kit.golaxy.org/golaxy/util"
)

// ContextResolver 用于从实体或组件上获取上下文
type ContextResolver interface {
	getContext() util.IfaceCache
}
