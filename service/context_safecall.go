package service

import (
	"errors"
	"github.com/golaxy-kit/golaxy/ec"
	"github.com/golaxy-kit/golaxy/internal"
	"github.com/golaxy-kit/golaxy/util"
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
	SafeCall(entityID ec.ID, segment func(entity ec.Entity) SafeRet) <-chan SafeRet

	// SafeCallWithSerialNo 同SafeCall，同时使用id与serialNo可以在多线程环境中准确定位实体
	SafeCallWithSerialNo(entityID ec.ID, entitySerialNo int64, segment func(entity ec.Entity) SafeRet) <-chan SafeRet

	// SafeCallNoWait 在运行时中，将代码片段压入任务流水线，串行化的进行调用，返回result channel，可以选择阻塞并等待返回结果。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 如果任务流水线一直满时，不会阻塞。
	//	- 如果在代码片段中，又向调用方所在的运行时发起安全调用，并且调用方阻塞并等待返回结果，那么这次调用会阻塞等待后超时。
	SafeCallNoWait(entityID ec.ID, segment func(entity ec.Entity) SafeRet) <-chan SafeRet

	// SafeCallNoWaitWithSerialNo 同SafeCallNoWait，同时使用id与serialNo可以在多线程环境中准确定位实体
	SafeCallNoWaitWithSerialNo(entityID ec.ID, entitySerialNo int64, segment func(entity ec.Entity) SafeRet) <-chan SafeRet

	// SafeCallNoRet 在运行时中，将代码片段压入任务流水线，串行化的进行调用，没有返回值，任务流水线满时会阻塞。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 如果任务流水线一直满时，会阻塞。
	SafeCallNoRet(entityID ec.ID, segment func(entity ec.Entity))

	// SafeCallNoRetWithSerialNo 同SafeCallNoRet，同时使用id与serialNo可以在多线程环境中准确定位实体
	SafeCallNoRetWithSerialNo(entityID ec.ID, entitySerialNo int64, segment func(entity ec.Entity))

	// SafeCallNoRetNoWait 在运行时中，将代码片段压入任务流水线，串行化的进行调用，没有返回值，任务流水线满时不会阻塞。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 如果任务流水线一直满时，不会阻塞。
	SafeCallNoRetNoWait(entityID ec.ID, segment func(entity ec.Entity))

	// SafeCallNoRetNoWaitWithSerialNo 同SafeCallNoRetNoWait，同时使用id与serialNo可以在多线程环境中准确定位实体
	SafeCallNoRetNoWaitWithSerialNo(entityID ec.ID, entitySerialNo int64, segment func(entity ec.Entity))
}

//go:linkname entitySafeCall github.com/golaxy-kit/golaxy/runtime.entitySafeCall
func entitySafeCall(entity ec.Entity) internal.SafeCall

//go:linkname entityExist github.com/golaxy-kit/golaxy/runtime.entityExist
func entityExist(entity ec.Entity) bool

func checkEntityID(entityID ec.ID) <-chan SafeRet {
	if entityID == util.Zero[ec.ID]() {
		ret := make(chan SafeRet, 1)
		ret <- SafeRet{
			Err: errors.New("entity id is zero invalid"),
		}
		return ret
	}
	return nil
}

func checkEntitySerialNo(entitySerialNo int64) <-chan SafeRet {
	if entitySerialNo <= 0 {
		ret := make(chan SafeRet, 1)
		ret <- SafeRet{
			Err: errors.New("entity serial no less equal 0 invalid"),
		}
		return ret
	}
	return nil
}

func checkSegment(segment func(entity ec.Entity) SafeRet) <-chan SafeRet {
	if segment == nil {
		ret := make(chan SafeRet, 1)
		ret <- SafeRet{
			Err: errors.New("nil segment"),
		}
		return ret
	}
	return nil
}

func getEntity(entityMgr _EntityMgr, id ec.ID) (ec.Entity, <-chan SafeRet) {
	entity, ok := entityMgr.GetEntity(id)
	if !ok {
		ret := make(chan SafeRet, 1)
		ret <- SafeRet{
			Err: errors.New("entity not exist in service context"),
		}
		return nil, ret
	}
	return entity, nil
}

func getEntityWithSerialNo(entityMgr _EntityMgr, id ec.ID, serialNo int64) (ec.Entity, <-chan SafeRet) {
	entity, ok := entityMgr.GetEntityWithSerialNo(id, serialNo)
	if !ok {
		ret := make(chan SafeRet, 1)
		ret <- SafeRet{
			Err: errors.New("entity not exist in service context"),
		}
		return nil, ret
	}
	return entity, nil
}

func safeCall(entity ec.Entity, segment func(entity ec.Entity) SafeRet) <-chan SafeRet {
	return entitySafeCall(entity).SafeCall(func() SafeRet {
		if !entityExist(entity) {
			return SafeRet{
				Err: errors.New("entity not exist in runtime context"),
			}
		}
		return segment(entity)
	})
}

func safeCallNoWait(entity ec.Entity, segment func(entity ec.Entity) SafeRet) <-chan SafeRet {
	return entitySafeCall(entity).SafeCallNoWait(func() SafeRet {
		if !entityExist(entity) {
			return SafeRet{
				Err: errors.New("entity not exist in runtime context"),
			}
		}
		return segment(entity)
	})
}

// SafeCall 在运行时中，将代码片段压入任务流水线，串行化的进行调用，返回result channel，可以选择阻塞并等待返回结果。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 如果任务流水线一直满时，会阻塞。
//	- 如果在代码片段中，又向调用方所在的运行时发起安全调用，并且调用方阻塞并等待返回结果，那么这次调用会阻塞等待后超时。
func (ctx *ContextBehavior) SafeCall(entityID ec.ID, segment func(entity ec.Entity) SafeRet) <-chan SafeRet {
	if ret := checkEntityID(entityID); ret != nil {
		return ret
	}

	if ret := checkSegment(segment); ret != nil {
		return ret
	}

	entity, ret := getEntity(ctx.entityMgr, entityID)
	if ret != nil {
		return ret
	}

	return safeCall(entity, segment)
}

// SafeCallWithSerialNo 同SafeCall，同时使用id与serialNo可以在多线程环境中准确定位实体
func (ctx *ContextBehavior) SafeCallWithSerialNo(entityID ec.ID, entitySerialNo int64, segment func(entity ec.Entity) SafeRet) <-chan SafeRet {
	if ret := checkEntityID(entityID); ret != nil {
		return ret
	}

	if ret := checkEntitySerialNo(entitySerialNo); ret != nil {
		return ret
	}

	if ret := checkSegment(segment); ret != nil {
		return ret
	}

	entity, ret := getEntityWithSerialNo(ctx.entityMgr, entityID, entitySerialNo)
	if ret != nil {
		return ret
	}

	return safeCall(entity, segment)
}

// SafeCallNoWait 在运行时中，将代码片段压入任务流水线，串行化的进行调用，返回result channel，可以选择阻塞并等待返回结果。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 如果任务流水线一直满时，不会阻塞。
//	- 如果在代码片段中，又向调用方所在的运行时发起安全调用，并且调用方阻塞并等待返回结果，那么这次调用会阻塞等待后超时。
func (ctx *ContextBehavior) SafeCallNoWait(entityID ec.ID, segment func(entity ec.Entity) SafeRet) <-chan SafeRet {
	if ret := checkEntityID(entityID); ret != nil {
		return ret
	}

	if ret := checkSegment(segment); ret != nil {
		return ret
	}

	entity, ret := getEntity(ctx.entityMgr, entityID)
	if ret != nil {
		return ret
	}

	return safeCallNoWait(entity, segment)
}

// SafeCallNoWaitWithSerialNo 同SafeCallNoWait，同时使用id与serialNo可以在多线程环境中准确定位实体
func (ctx *ContextBehavior) SafeCallNoWaitWithSerialNo(entityID ec.ID, entitySerialNo int64, segment func(entity ec.Entity) SafeRet) <-chan SafeRet {
	if ret := checkEntityID(entityID); ret != nil {
		return ret
	}

	if ret := checkEntitySerialNo(entitySerialNo); ret != nil {
		return ret
	}

	if ret := checkSegment(segment); ret != nil {
		return ret
	}

	entity, ret := getEntityWithSerialNo(ctx.entityMgr, entityID, entitySerialNo)
	if ret != nil {
		return ret
	}

	return safeCallNoWait(entity, segment)
}

// SafeCallNoRet 在运行时中，将代码片段压入任务流水线，串行化的进行调用，没有返回值，任务流水线满时会阻塞。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 如果任务流水线一直满时，会阻塞。
func (ctx *ContextBehavior) SafeCallNoRet(entityID ec.ID, segment func(entity ec.Entity)) {
	if entityID == util.Zero[ec.ID]() || segment == nil {
		return
	}

	entity, ok := ctx.entityMgr.GetEntity(entityID)
	if !ok {
		return
	}

	entitySafeCall(entity).SafeCallNoRet(func() {
		if entityExist(entity) {
			segment(entity)
		}
	})
}

// SafeCallNoRetWithSerialNo 同SafeCallNoRet，同时使用id与serialNo可以在多线程环境中准确定位实体
func (ctx *ContextBehavior) SafeCallNoRetWithSerialNo(entityID ec.ID, entitySerialNo int64, segment func(entity ec.Entity)) {
	if entityID == util.Zero[ec.ID]() || entitySerialNo <= 0 || segment == nil {
		return
	}

	entity, ok := ctx.entityMgr.GetEntityWithSerialNo(entityID, entitySerialNo)
	if !ok {
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
func (ctx *ContextBehavior) SafeCallNoRetNoWait(entityID ec.ID, segment func(entity ec.Entity)) {
	if entityID == util.Zero[ec.ID]() || segment == nil {
		return
	}

	entity, ok := ctx.entityMgr.GetEntity(entityID)
	if !ok {
		return
	}

	entitySafeCall(entity).SafeCallNoRetNoWait(func() {
		if entityExist(entity) {
			segment(entity)
		}
	})
}

// SafeCallNoRetNoWaitWithSerialNo 同SafeCallNoRetNoWait，同时使用id与serialNo可以在多线程环境中准确定位实体
func (ctx *ContextBehavior) SafeCallNoRetNoWaitWithSerialNo(entityID ec.ID, entitySerialNo int64, segment func(entity ec.Entity)) {
	if entityID == util.Zero[ec.ID]() || entitySerialNo <= 0 || segment == nil {
		return
	}

	entity, ok := ctx.entityMgr.GetEntityWithSerialNo(entityID, entitySerialNo)
	if !ok {
		return
	}

	entitySafeCall(entity).SafeCallNoRetNoWait(func() {
		if entityExist(entity) {
			segment(entity)
		}
	})
}
