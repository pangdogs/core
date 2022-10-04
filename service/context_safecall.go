package service

import (
	"errors"
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/runtime"
)

type _SafeCall interface {
	// SafeCall 在运行时中，将代码片段压入任务流水线，串行化的进行调用，返回result channel，可以选择阻塞并等待返回结果。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 如果任务流水线一直满时，会阻塞。
	//	- 如果在代码片段中，又向调用方所在的运行时发起安全调用，并且调用方阻塞并等待返回结果，那么这次调用会阻塞等待后超时。
	SafeCall(entityID int64, segment func(entity ec.Entity) runtime.SafeRet) <-chan runtime.SafeRet

	// SafeCallNoWait 在运行时中，将代码片段压入任务流水线，串行化的进行调用，返回result channel，可以选择阻塞并等待返回结果。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 如果任务流水线一直满时，不会阻塞。
	//	- 如果在代码片段中，又向调用方所在的运行时发起安全调用，并且调用方阻塞并等待返回结果，那么这次调用会阻塞等待后超时。
	SafeCallNoWait(entityID int64, segment func(entity ec.Entity) runtime.SafeRet) <-chan runtime.SafeRet

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

// SafeCall 在运行时中，将代码片段压入任务流水线，串行化的进行调用，返回result channel，可以选择阻塞并等待返回结果。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 如果任务流水线一直满时，会阻塞。
//	- 如果在代码片段中，又向调用方所在的运行时发起安全调用，并且调用方阻塞并等待返回结果，那么这次调用会阻塞等待后超时。
func (ctx *ContextBehavior) SafeCall(entityID int64, segment func(entity ec.Entity) runtime.SafeRet) <-chan runtime.SafeRet {
	if entityID == 0 {
		ret := make(chan runtime.SafeRet, 1)
		ret <- runtime.SafeRet{
			Err: errors.New("entity id invalid"),
		}
		return ret
	}

	if segment == nil {
		ret := make(chan runtime.SafeRet, 1)
		ret <- runtime.SafeRet{
			Err: errors.New("nil segment"),
		}
		return ret
	}

	entity, ok := ctx.entityMgr.GetEntity(entityID)
	if ok {
		ret := make(chan runtime.SafeRet, 1)
		ret <- runtime.SafeRet{
			Err: errors.New("entity not exist"),
		}
		return ret
	}

	return runtime.EntityContext(entity).SafeCall(func() runtime.SafeRet {
		return segment(entity)
	})
}

// SafeCallNoWait 在运行时中，将代码片段压入任务流水线，串行化的进行调用，返回result channel，可以选择阻塞并等待返回结果。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 如果任务流水线一直满时，不会阻塞。
//	- 如果在代码片段中，又向调用方所在的运行时发起安全调用，并且调用方阻塞并等待返回结果，那么这次调用会阻塞等待后超时。
func (ctx *ContextBehavior) SafeCallNoWait(entityID int64, segment func(entity ec.Entity) runtime.SafeRet) <-chan runtime.SafeRet {
	if entityID == 0 {
		ret := make(chan runtime.SafeRet, 1)
		ret <- runtime.SafeRet{
			Err: errors.New("entity id invalid"),
		}
		return ret
	}

	if segment == nil {
		ret := make(chan runtime.SafeRet, 1)
		ret <- runtime.SafeRet{
			Err: errors.New("nil segment"),
		}
		return ret
	}

	entity, ok := ctx.entityMgr.GetEntity(entityID)
	if ok {
		ret := make(chan runtime.SafeRet, 1)
		ret <- runtime.SafeRet{
			Err: errors.New("entity not exist"),
		}
		return ret
	}

	return runtime.EntityContext(entity).SafeCallNoWait(func() runtime.SafeRet {
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

	runtime.EntityContext(entity).SafeCallNoRet(func() {
		segment(entity)
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

	runtime.EntityContext(entity).SafeCallNoRetNoWait(func() {
		segment(entity)
	})
}
