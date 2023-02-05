package service

import (
	"errors"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/util"
	_ "unsafe"
)

// Ret 调用结果
type Ret = internal.Ret

type _Call interface {
	// SyncCall 同步调用。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，会阻塞并等待返回值。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 当运行时的SyncCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会转换为error返回。
	SyncCall(entityID ec.ID, segment func(entity ec.Entity) Ret) Ret

	// SyncCallWithSerialNo 与SyncCall相同，只是同时使用id与serialNo可以在多线程环境中准确定位实体
	SyncCallWithSerialNo(entityID ec.ID, entitySerialNo int64, segment func(entity ec.Entity) Ret) Ret

	// AsyncCall 异步调用。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，不会阻塞，会返回result channel。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 在代码片段中，如果向调用方所在的运行时发起同步调用，并且调用方也在阻塞result channel等待返回值，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会转换为error返回。
	AsyncCall(entityID ec.ID, segment func(entity ec.Entity) Ret) <-chan Ret

	// AsyncCallWithSerialNo 与AsyncCall相同，只是同时使用id与serialNo可以在多线程环境中准确定位实体
	AsyncCallWithSerialNo(entityID ec.ID, entitySerialNo int64, segment func(entity ec.Entity) Ret) <-chan Ret

	// SyncCallNoRet 同步调用，无返回值。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，会阻塞，没有返回值。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 当运行时的SyncCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会抛弃。
	SyncCallNoRet(entityID ec.ID, segment func(entity ec.Entity))

	// SyncCallNoRetWithSerialNo 与SyncCallNoRet相同，只是同时使用id与serialNo可以在多线程环境中准确定位实体
	SyncCallNoRetWithSerialNo(entityID ec.ID, entitySerialNo int64, segment func(entity ec.Entity))

	// AsyncCallNoRet 异步调用，无返回值。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，不会阻塞，没有返回值。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//  - 调用过程中的panic信息，均会抛弃。
	AsyncCallNoRet(entityID ec.ID, segment func(entity ec.Entity))

	// AsyncCallNoRetWithSerialNo 与AsyncCallNoRet相同，只是同时使用id与serialNo可以在多线程环境中准确定位实体
	AsyncCallNoRetWithSerialNo(entityID ec.ID, entitySerialNo int64, segment func(entity ec.Entity))
}

//go:linkname entityCall kit.golaxy.org/golaxy/runtime.entityCall
func entityCall(entity ec.Entity) internal.Call

//go:linkname entityExist kit.golaxy.org/golaxy/runtime.entityExist
func entityExist(entity ec.Entity) bool

func checkEntityID(entityID ec.ID) error {
	if entityID == util.Zero[ec.ID]() {
		return errors.New("entity id is zero invalid")
	}
	return nil
}

func checkEntitySerialNo(entitySerialNo int64) error {
	if entitySerialNo <= 0 {
		return errors.New("entity serial no less equal 0 invalid")
	}
	return nil
}

func checkSegment(segment func(entity ec.Entity) Ret) error {
	if segment == nil {
		return errors.New("nil segment")
	}
	return nil
}

func getEntity(entityMgr _EntityMgr, id ec.ID) (ec.Entity, error) {
	entity, ok := entityMgr.GetEntity(id)
	if !ok {
		return nil, errors.New("entity not exist in service context")
	}
	return entity, nil
}

func getEntityWithSerialNo(entityMgr _EntityMgr, id ec.ID, serialNo int64) (ec.Entity, error) {
	entity, ok := entityMgr.GetEntityWithSerialNo(id, serialNo)
	if !ok {
		return nil, errors.New("entity not exist in service context")
	}
	return entity, nil
}

func syncCall(entity ec.Entity, segment func(entity ec.Entity) Ret) Ret {
	return entityCall(entity).SyncCall(func() Ret {
		if !entityExist(entity) {
			return Ret{
				Err: errors.New("entity not exist in runtime context"),
			}
		}
		return segment(entity)
	})
}

func asyncCall(entity ec.Entity, segment func(entity ec.Entity) Ret) <-chan Ret {
	return entityCall(entity).AsyncCall(func() Ret {
		if !entityExist(entity) {
			return Ret{
				Err: errors.New("entity not exist in runtime context"),
			}
		}
		return segment(entity)
	})
}

// SyncCall 同步调用。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，会阻塞并等待返回值。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 当运行时的SyncCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) SyncCall(entityID ec.ID, segment func(entity ec.Entity) Ret) Ret {
	if err := checkEntityID(entityID); err != nil {
		return Ret{Err: err}
	}

	if err := checkSegment(segment); err != nil {
		return Ret{Err: err}
	}

	entity, err := getEntity(ctx.entityMgr, entityID)
	if err != nil {
		return Ret{Err: err}
	}

	return syncCall(entity, segment)
}

// SyncCallWithSerialNo 与SyncCall相同，只是同时使用id与serialNo可以在多线程环境中准确定位实体
func (ctx *ContextBehavior) SyncCallWithSerialNo(entityID ec.ID, entitySerialNo int64, segment func(entity ec.Entity) Ret) Ret {
	if err := checkEntityID(entityID); err != nil {
		return Ret{Err: err}
	}

	if err := checkEntitySerialNo(entitySerialNo); err != nil {
		return Ret{Err: err}
	}

	if err := checkSegment(segment); err != nil {
		return Ret{Err: err}
	}

	entity, err := getEntityWithSerialNo(ctx.entityMgr, entityID, entitySerialNo)
	if err != nil {
		return Ret{Err: err}
	}

	return syncCall(entity, segment)
}

// AsyncCall 异步调用。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，不会阻塞，会返回result channel。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 在代码片段中，如果向调用方所在的运行时发起同步调用，并且调用方也在阻塞result channel等待返回值，那么会造成线程死锁。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) AsyncCall(entityID ec.ID, segment func(entity ec.Entity) Ret) <-chan Ret {
	if err := checkEntityID(entityID); err != nil {
		retChan := make(chan Ret)
		retChan <- Ret{Err: err}
		return retChan
	}

	if err := checkSegment(segment); err != nil {
		retChan := make(chan Ret)
		retChan <- Ret{Err: err}
		return retChan
	}

	entity, err := getEntity(ctx.entityMgr, entityID)
	if err != nil {
		retChan := make(chan Ret)
		retChan <- Ret{Err: err}
		return retChan
	}

	return asyncCall(entity, segment)
}

// AsyncCallWithSerialNo 与AsyncCall相同，只是同时使用id与serialNo可以在多线程环境中准确定位实体
func (ctx *ContextBehavior) AsyncCallWithSerialNo(entityID ec.ID, entitySerialNo int64, segment func(entity ec.Entity) Ret) <-chan Ret {
	if err := checkEntityID(entityID); err != nil {
		retChan := make(chan Ret)
		retChan <- Ret{Err: err}
		return retChan
	}

	if err := checkEntitySerialNo(entitySerialNo); err != nil {
		retChan := make(chan Ret)
		retChan <- Ret{Err: err}
		return retChan
	}

	if err := checkSegment(segment); err != nil {
		retChan := make(chan Ret)
		retChan <- Ret{Err: err}
		return retChan
	}

	entity, err := getEntityWithSerialNo(ctx.entityMgr, entityID, entitySerialNo)
	if err != nil {
		retChan := make(chan Ret)
		retChan <- Ret{Err: err}
		return retChan
	}

	return asyncCall(entity, segment)
}

// SyncCallNoRet 同步调用，无返回值。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，会阻塞，没有返回值。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 当运行时的SyncCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
//	- 调用过程中的panic信息，均会抛弃。
func (ctx *ContextBehavior) SyncCallNoRet(entityID ec.ID, segment func(entity ec.Entity)) {
	if entityID == util.Zero[ec.ID]() || segment == nil {
		return
	}

	entity, ok := ctx.entityMgr.GetEntity(entityID)
	if !ok {
		return
	}

	entityCall(entity).SyncCallNoRet(func() {
		if entityExist(entity) {
			segment(entity)
		}
	})
}

// SyncCallNoRetWithSerialNo 与SyncCallNoRet相同，只是同时使用id与serialNo可以在多线程环境中准确定位实体
func (ctx *ContextBehavior) SyncCallNoRetWithSerialNo(entityID ec.ID, entitySerialNo int64, segment func(entity ec.Entity)) {
	if entityID == util.Zero[ec.ID]() || entitySerialNo <= 0 || segment == nil {
		return
	}

	entity, ok := ctx.entityMgr.GetEntityWithSerialNo(entityID, entitySerialNo)
	if !ok {
		return
	}

	entityCall(entity).SyncCallNoRet(func() {
		if entityExist(entity) {
			segment(entity)
		}
	})
}

// AsyncCallNoRet 异步调用，无返回值。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，不会阻塞，没有返回值。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 调用过程中的panic信息，均会抛弃。
func (ctx *ContextBehavior) AsyncCallNoRet(entityID ec.ID, segment func(entity ec.Entity)) {
	if entityID == util.Zero[ec.ID]() || segment == nil {
		return
	}

	entity, ok := ctx.entityMgr.GetEntity(entityID)
	if !ok {
		return
	}

	entityCall(entity).AsyncCallNoRet(func() {
		if entityExist(entity) {
			segment(entity)
		}
	})
}

// AsyncCallNoRetWithSerialNo 与AsyncCallNoRet相同，只是同时使用id与serialNo可以在多线程环境中准确定位实体
func (ctx *ContextBehavior) AsyncCallNoRetWithSerialNo(entityID ec.ID, entitySerialNo int64, segment func(entity ec.Entity)) {
	if entityID == util.Zero[ec.ID]() || entitySerialNo <= 0 || segment == nil {
		return
	}

	entity, ok := ctx.entityMgr.GetEntityWithSerialNo(entityID, entitySerialNo)
	if !ok {
		return
	}

	entityCall(entity).AsyncCallNoRet(func() {
		if entityExist(entity) {
			segment(entity)
		}
	})
}
