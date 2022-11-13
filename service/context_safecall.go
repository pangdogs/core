package service

import (
	"errors"
	"github.com/galaxy-kit/galaxy-go/ec"
	"github.com/galaxy-kit/galaxy-go/internal"
	_ "unsafe"
)

// SafeRet 安全调用结果
type SafeRet = internal.SafeRet

// _SafeCall 安全调用
type _SafeCall interface {
	// SafeCall 在运行时中，将代码片段压入任务流水线，串行化的进行调用，返回result channel，可以选择阻塞并等待返回结果。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 如果任务流水线一直满时，会阻塞。
	//	- 如果在代码片段中，又向调用方所在的运行时发起安全调用，并且调用方阻塞并等待返回结果，那么这次调用会阻塞等待后超时。
	SafeCall(entityID int64, segment func(entity ec.Entity) SafeRet) <-chan SafeRet

	// SafeCallNoWait 在运行时中，将代码片段压入任务流水线，串行化的进行调用，返回result channel，可以选择阻塞并等待返回结果。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 如果任务流水线一直满时，不会阻塞。
	//	- 如果在代码片段中，又向调用方所在的运行时发起安全调用，并且调用方阻塞并等待返回结果，那么这次调用会阻塞等待后超时。
	SafeCallNoWait(entityID int64, segment func(entity ec.Entity) SafeRet) <-chan SafeRet

	// SafeCallNoRet 在运行时中，将代码片段压入任务流水线，串行化的进行调用，没有返回值，任务流水线满时会阻塞。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 如果任务流水线一直满时，会阻塞。
	SafeCallNoRet(entityID int64, segment func(entity ec.Entity))

	// SafeCallNoRetNoWait 在运行时中，将代码片段压入任务流水线，串行化的进行调用，没有返回值，任务流水线满时不会阻塞。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 如果任务流水线一直满时，不会阻塞。
	SafeCallNoRetNoWait(entityID int64, segment func(entity ec.Entity))
}

//go:linkname entitySafeCall github.com/galaxy-kit/galaxy-go/runtime.entitySafeCall
func entitySafeCall(entity ec.Entity) internal.SafeCall

//go:linkname entityExist github.com/galaxy-kit/galaxy-go/runtime.entityExist
func entityExist(entity ec.Entity) bool

// SafeCall 在运行时中，将代码片段压入任务流水线，串行化的进行调用，返回result channel，可以选择阻塞并等待返回结果。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 如果任务流水线一直满时，会阻塞。
//	- 如果在代码片段中，又向调用方所在的运行时发起安全调用，并且调用方阻塞并等待返回结果，那么这次调用会阻塞等待后超时。
func (ctx *ContextBehavior) SafeCall(entityID int64, segment func(entity ec.Entity) SafeRet) <-chan SafeRet {
	if entityID == 0 {
		ret := make(chan SafeRet, 1)
		ret <- SafeRet{
			Err: errors.New("entity id invalid"),
		}
		return ret
	}

	if segment == nil {
		ret := make(chan SafeRet, 1)
		ret <- SafeRet{
			Err: errors.New("nil segment"),
		}
		return ret
	}

	entity, ok := ctx.entityMgr.GetEntity(entityID)
	if ok {
		ret := make(chan SafeRet, 1)
		ret <- SafeRet{
			Err: errors.New("entity not exist in service context"),
		}
		return ret
	}

	return entitySafeCall(entity).SafeCall(func() SafeRet {
		if !entityExist(entity) {
			return SafeRet{
				Err: errors.New("entity not exist in runtime context"),
			}
		}
		return segment(entity)
	})
}

// SafeCallNoWait 在运行时中，将代码片段压入任务流水线，串行化的进行调用，返回result channel，可以选择阻塞并等待返回结果。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 如果任务流水线一直满时，不会阻塞。
//	- 如果在代码片段中，又向调用方所在的运行时发起安全调用，并且调用方阻塞并等待返回结果，那么这次调用会阻塞等待后超时。
func (ctx *ContextBehavior) SafeCallNoWait(entityID int64, segment func(entity ec.Entity) SafeRet) <-chan SafeRet {
	if entityID == 0 {
		ret := make(chan SafeRet, 1)
		ret <- SafeRet{
			Err: errors.New("entity id invalid"),
		}
		return ret
	}

	if segment == nil {
		ret := make(chan SafeRet, 1)
		ret <- SafeRet{
			Err: errors.New("nil segment"),
		}
		return ret
	}

	entity, ok := ctx.entityMgr.GetEntity(entityID)
	if ok {
		ret := make(chan SafeRet, 1)
		ret <- SafeRet{
			Err: errors.New("entity not exist in service context"),
		}
		return ret
	}

	return entitySafeCall(entity).SafeCallNoWait(func() SafeRet {
		if !entityExist(entity) {
			return SafeRet{
				Err: errors.New("entity not exist in runtime context"),
			}
		}
		return segment(entity)
	})
}

// SafeCallNoRet 在运行时中，将代码片段压入任务流水线，串行化的进行调用，没有返回值，任务流水线满时会阻塞。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 如果任务流水线一直满时，会阻塞。
func (ctx *ContextBehavior) SafeCallNoRet(entityID int64, segment func(entity ec.Entity)) {
	if entityID == 0 {
		return
	}

	if segment == nil {
		return
	}

	entity, ok := ctx.entityMgr.GetEntity(entityID)
	if ok {
		return
	}

	entitySafeCall(entity).SafeCallNoRet(func() {
		if entityExist(entity) {
			segment(entity)
		}
	})
}

// SafeCallNoRetNoWait 在运行时中，将代码片段压入任务流水线，串行化的进行调用，没有返回值，任务流水线满时不会阻塞。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 如果任务流水线一直满时，不会阻塞。
func (ctx *ContextBehavior) SafeCallNoRetNoWait(entityID int64, segment func(entity ec.Entity)) {
	if entityID == 0 {
		return
	}

	if segment == nil {
		return
	}

	entity, ok := ctx.entityMgr.GetEntity(entityID)
	if ok {
		return
	}

	entitySafeCall(entity).SafeCallNoRetNoWait(func() {
		if entityExist(entity) {
			segment(entity)
		}
	})
}
