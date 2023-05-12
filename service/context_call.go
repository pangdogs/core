package service

import (
	"errors"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/uid"
	_ "unsafe"
)

type _Call interface {
	// AwaitCall 同步调用。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，会阻塞并等待返回值。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 当运行时的AwaitCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会转换为error返回。
	AwaitCall(entityId uid.Id, segment func(entity ec.Entity) runtime.Ret) runtime.Ret

	// AwaitCallWithSerialNo 与AwaitCall()相同，只是同时使用id与serialNo可以在多线程环境中准确定位实体
	AwaitCallWithSerialNo(entityId uid.Id, entitySerialNo int64, segment func(entity ec.Entity) runtime.Ret) runtime.Ret

	// AsyncCall 异步调用。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，不会阻塞，会返回AsyncRet。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 在代码片段中，如果向调用方所在的运行时发起同步调用，并且调用方也在阻塞AsyncRet等待返回值，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会转换为error返回。
	AsyncCall(entityId uid.Id, segment func(entity ec.Entity) runtime.Ret) runtime.AsyncRet

	// AsyncCallWithSerialNo 与AsyncCall()相同，只是同时使用id与serialNo可以在多线程环境中准确定位实体
	AsyncCallWithSerialNo(entityId uid.Id, entitySerialNo int64, segment func(entity ec.Entity) runtime.Ret) runtime.AsyncRet

	// AwaitCallNoRet 同步调用，无返回值。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，会阻塞，没有返回值。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 当运行时的AwaitCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会抛弃。
	AwaitCallNoRet(entityId uid.Id, segment func(entity ec.Entity))

	// AwaitCallNoRetWithSerialNo 与AwaitCallNoRet()相同，只是同时使用id与serialNo可以在多线程环境中准确定位实体
	AwaitCallNoRetWithSerialNo(entityId uid.Id, entitySerialNo int64, segment func(entity ec.Entity))

	// AsyncCallNoRet 异步调用，无返回值。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，不会阻塞，没有返回值。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//  - 调用过程中的panic信息，均会抛弃。
	AsyncCallNoRet(entityId uid.Id, segment func(entity ec.Entity))

	// AsyncCallNoRetWithSerialNo 与AsyncCallNoRet()相同，只是同时使用id与serialNo可以在多线程环境中准确定位实体
	AsyncCallNoRetWithSerialNo(entityId uid.Id, entitySerialNo int64, segment func(entity ec.Entity))
}

//go:linkname entityCaller kit.golaxy.org/golaxy/runtime.entityCaller
func entityCaller(entity ec.Entity) runtime.Caller

//go:linkname entityExist kit.golaxy.org/golaxy/runtime.entityExist
func entityExist(entity ec.Entity) bool

func checkEntityId(entityId uid.Id) error {
	if entityId.IsNil() {
		return errors.New("entity id equal zero is invalid")
	}
	return nil
}

func checkEntitySerialNo(entitySerialNo int64) error {
	if entitySerialNo <= 0 {
		return errors.New("entity serial no less equal 0 is invalid")
	}
	return nil
}

func checkSegment(segment func(entity ec.Entity) runtime.Ret) error {
	if segment == nil {
		return errors.New("nil segment")
	}
	return nil
}

func getEntity(entityMgr _EntityMgr, id uid.Id) (ec.Entity, error) {
	entity, ok := entityMgr.GetEntity(id)
	if !ok {
		return nil, errors.New("entity not exist in service context")
	}
	return entity, nil
}

func getEntityWithSerialNo(entityMgr _EntityMgr, id uid.Id, serialNo int64) (ec.Entity, error) {
	entity, ok := entityMgr.GetEntityWithSerialNo(id, serialNo)
	if !ok {
		return nil, errors.New("entity not exist in service context")
	}
	return entity, nil
}

func awaitCall(entity ec.Entity, segment func(entity ec.Entity) runtime.Ret) runtime.Ret {
	return entityCaller(entity).AwaitCall(func() runtime.Ret {
		if !entityExist(entity) {
			return runtime.Ret{
				Err: errors.New("entity not exist in runtime context"),
			}
		}
		return segment(entity)
	})
}

func asyncCall(entity ec.Entity, segment func(entity ec.Entity) runtime.Ret) runtime.AsyncRet {
	return entityCaller(entity).AsyncCall(func() runtime.Ret {
		if !entityExist(entity) {
			return runtime.Ret{
				Err: errors.New("entity not exist in runtime context"),
			}
		}
		return segment(entity)
	})
}

// AwaitCall 同步调用。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，会阻塞并等待返回值。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 当运行时的AwaitCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) AwaitCall(entityId uid.Id, segment func(entity ec.Entity) runtime.Ret) runtime.Ret {
	if err := checkEntityId(entityId); err != nil {
		return runtime.Ret{Err: err}
	}

	if err := checkSegment(segment); err != nil {
		return runtime.Ret{Err: err}
	}

	entity, err := getEntity(ctx.entityMgr, entityId)
	if err != nil {
		return runtime.Ret{Err: err}
	}

	return awaitCall(entity, segment)
}

// AwaitCallWithSerialNo 与AwaitCall()相同，只是同时使用id与serialNo可以在多线程环境中准确定位实体
func (ctx *ContextBehavior) AwaitCallWithSerialNo(entityId uid.Id, entitySerialNo int64, segment func(entity ec.Entity) runtime.Ret) runtime.Ret {
	if err := checkEntityId(entityId); err != nil {
		return runtime.Ret{Err: err}
	}

	if err := checkEntitySerialNo(entitySerialNo); err != nil {
		return runtime.Ret{Err: err}
	}

	if err := checkSegment(segment); err != nil {
		return runtime.Ret{Err: err}
	}

	entity, err := getEntityWithSerialNo(ctx.entityMgr, entityId, entitySerialNo)
	if err != nil {
		return runtime.Ret{Err: err}
	}

	return awaitCall(entity, segment)
}

// AsyncCall 异步调用。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，不会阻塞，会返回AsyncRet。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 在代码片段中，如果向调用方所在的运行时发起同步调用，并且调用方也在阻塞AsyncRet等待返回值，那么会造成线程死锁。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) AsyncCall(entityId uid.Id, segment func(entity ec.Entity) runtime.Ret) runtime.AsyncRet {
	if err := checkEntityId(entityId); err != nil {
		retChan := make(chan runtime.Ret)
		retChan <- runtime.Ret{Err: err}
		return retChan
	}

	if err := checkSegment(segment); err != nil {
		retChan := make(chan runtime.Ret)
		retChan <- runtime.Ret{Err: err}
		return retChan
	}

	entity, err := getEntity(ctx.entityMgr, entityId)
	if err != nil {
		retChan := make(chan runtime.Ret)
		retChan <- runtime.Ret{Err: err}
		return retChan
	}

	return asyncCall(entity, segment)
}

// AsyncCallWithSerialNo 与AsyncCall()相同，只是同时使用id与serialNo可以在多线程环境中准确定位实体
func (ctx *ContextBehavior) AsyncCallWithSerialNo(entityId uid.Id, entitySerialNo int64, segment func(entity ec.Entity) runtime.Ret) runtime.AsyncRet {
	if err := checkEntityId(entityId); err != nil {
		retChan := make(chan runtime.Ret)
		retChan <- runtime.Ret{Err: err}
		return retChan
	}

	if err := checkEntitySerialNo(entitySerialNo); err != nil {
		retChan := make(chan runtime.Ret)
		retChan <- runtime.Ret{Err: err}
		return retChan
	}

	if err := checkSegment(segment); err != nil {
		retChan := make(chan runtime.Ret)
		retChan <- runtime.Ret{Err: err}
		return retChan
	}

	entity, err := getEntityWithSerialNo(ctx.entityMgr, entityId, entitySerialNo)
	if err != nil {
		retChan := make(chan runtime.Ret)
		retChan <- runtime.Ret{Err: err}
		return retChan
	}

	return asyncCall(entity, segment)
}

// AwaitCallNoRet 同步调用，无返回值。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，会阻塞，没有返回值。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 当运行时的AwaitCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
//	- 调用过程中的panic信息，均会抛弃。
func (ctx *ContextBehavior) AwaitCallNoRet(entityId uid.Id, segment func(entity ec.Entity)) {
	if entityId.IsNil() || segment == nil {
		return
	}

	entity, ok := ctx.entityMgr.GetEntity(entityId)
	if !ok {
		return
	}

	entityCaller(entity).AwaitCallNoRet(func() {
		if entityExist(entity) {
			segment(entity)
		}
	})
}

// AwaitCallNoRetWithSerialNo 与AwaitCallNoRet()相同，只是同时使用id与serialNo可以在多线程环境中准确定位实体
func (ctx *ContextBehavior) AwaitCallNoRetWithSerialNo(entityId uid.Id, entitySerialNo int64, segment func(entity ec.Entity)) {
	if entityId.IsNil() || entitySerialNo <= 0 || segment == nil {
		return
	}

	entity, ok := ctx.entityMgr.GetEntityWithSerialNo(entityId, entitySerialNo)
	if !ok {
		return
	}

	entityCaller(entity).AwaitCallNoRet(func() {
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
func (ctx *ContextBehavior) AsyncCallNoRet(entityId uid.Id, segment func(entity ec.Entity)) {
	if entityId.IsNil() || segment == nil {
		return
	}

	entity, ok := ctx.entityMgr.GetEntity(entityId)
	if !ok {
		return
	}

	entityCaller(entity).AsyncCallNoRet(func() {
		if entityExist(entity) {
			segment(entity)
		}
	})
}

// AsyncCallNoRetWithSerialNo 与AsyncCallNoRet()相同，只是同时使用id与serialNo可以在多线程环境中准确定位实体
func (ctx *ContextBehavior) AsyncCallNoRetWithSerialNo(entityId uid.Id, entitySerialNo int64, segment func(entity ec.Entity)) {
	if entityId.IsNil() || entitySerialNo <= 0 || segment == nil {
		return
	}

	entity, ok := ctx.entityMgr.GetEntityWithSerialNo(entityId, entitySerialNo)
	if !ok {
		return
	}

	entityCaller(entity).AsyncCallNoRet(func() {
		if entityExist(entity) {
			segment(entity)
		}
	})
}
