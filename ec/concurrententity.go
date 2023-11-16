package ec

import (
	"kit.golaxy.org/golaxy/internal/concurrent"
	"kit.golaxy.org/golaxy/util/uid"
)

// ConcurrentEntity 多线程安全的实体接口
type ConcurrentEntity interface {
	concurrent.ConcurrentContextProvider

	// GetId 获取实体Id
	GetId() uid.Id
	// GetPrototype 获取实体原型
	GetPrototype() string
}
