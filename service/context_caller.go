package service

import (
	"errors"
	"kit.golaxy.org/golaxy/ec"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/uid"
	_ "unsafe"
)

// Ret 调用结果
type Ret = internal.Ret

// AsyncRet 异步调用结果
type AsyncRet = internal.AsyncRet

// Caller 异步调用发起者
type Caller interface {
	// SyncCall 同步调用。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，会阻塞并等待返回值。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 当运行时的SyncCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会转换为error返回。
	SyncCall(entityId uid.Id, segment func(entity ec.Entity) Ret) Ret

	// AsyncCall 异步调用。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，不会阻塞，会返回AsyncRet。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 在代码片段中，如果向调用方所在的运行时发起同步调用，并且调用方也在阻塞AsyncRet等待返回值，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会转换为error返回。
	AsyncCall(entityId uid.Id, segment func(entity ec.Entity) Ret) AsyncRet

	// SyncCallVoid 同步调用，无返回值。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，会阻塞，没有返回值。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 当运行时的SyncCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会抛弃。
	SyncCallVoid(entityId uid.Id, segment func(entity ec.Entity))

	// AsyncCallVoid 异步调用，无返回值。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，不会阻塞，没有返回值。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//  - 调用过程中的panic信息，均会抛弃。
	AsyncCallVoid(entityId uid.Id, segment func(entity ec.Entity))
}

//go:linkname entityCaller kit.golaxy.org/golaxy/runtime.entityCaller
func entityCaller(entity ec.Entity) internal.Caller

//go:linkname entityExist kit.golaxy.org/golaxy/runtime.entityExist
func entityExist(entity ec.Entity) bool

func checkEntityId(entityId uid.Id) error {
	if entityId.IsNil() {
		return errors.New("entity id equal zero is invalid")
	}
	return nil
}

func checkSegment(segment func(entity ec.Entity) Ret) error {
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

func syncCall(entity ec.Entity, segment func(entity ec.Entity) Ret) Ret {
	return entityCaller(entity).SyncCall(func() Ret {
		if !entityExist(entity) {
			return Ret{
				Error: errors.New("entity not exist in runtime context"),
			}
		}
		return segment(entity)
	})
}

func asyncCall(entity ec.Entity, segment func(entity ec.Entity) Ret) AsyncRet {
	return entityCaller(entity).AsyncCall(func() Ret {
		if !entityExist(entity) {
			return Ret{
				Error: errors.New("entity not exist in runtime context"),
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
func (ctx *ContextBehavior) SyncCall(entityId uid.Id, segment func(entity ec.Entity) Ret) Ret {
	if err := checkEntityId(entityId); err != nil {
		return internal.NewRet(err, nil)
	}

	if err := checkSegment(segment); err != nil {
		return internal.NewRet(err, nil)
	}

	entity, err := getEntity(ctx.entityMgr, entityId)
	if err != nil {
		return internal.NewRet(err, nil)
	}

	return syncCall(entity, segment)
}

// AsyncCall 异步调用。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，不会阻塞，会返回AsyncRet。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 在代码片段中，如果向调用方所在的运行时发起同步调用，并且调用方也在阻塞AsyncRet等待返回值，那么会造成线程死锁。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) AsyncCall(entityId uid.Id, segment func(entity ec.Entity) Ret) AsyncRet {
	returnAsyncRet := func(err error, val any) AsyncRet {
		asyncRet := make(chan Ret, 1)
		asyncRet <- internal.NewRet(err, val)
		close(asyncRet)
		return asyncRet
	}

	if err := checkEntityId(entityId); err != nil {
		return returnAsyncRet(err, nil)
	}

	if err := checkSegment(segment); err != nil {
		return returnAsyncRet(err, nil)
	}

	entity, err := getEntity(ctx.entityMgr, entityId)
	if err != nil {
		return returnAsyncRet(err, nil)
	}

	return asyncCall(entity, segment)
}

// SyncCallVoid 同步调用，无返回值。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，会阻塞，没有返回值。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 当运行时的SyncCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
//	- 调用过程中的panic信息，均会抛弃。
func (ctx *ContextBehavior) SyncCallVoid(entityId uid.Id, segment func(entity ec.Entity)) {
	if entityId.IsNil() || segment == nil {
		return
	}

	entity, ok := ctx.entityMgr.GetEntity(entityId)
	if !ok {
		return
	}

	entityCaller(entity).SyncCallVoid(func() {
		if entityExist(entity) {
			segment(entity)
		}
	})
}

// AsyncCallVoid 异步调用，无返回值。查找实体，并获取实体的运行时，将代码片段压入运行时的任务流水线，串行化的进行调用，不会阻塞，没有返回值。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 调用过程中的panic信息，均会抛弃。
func (ctx *ContextBehavior) AsyncCallVoid(entityId uid.Id, segment func(entity ec.Entity)) {
	if entityId.IsNil() || segment == nil {
		return
	}

	entity, ok := ctx.entityMgr.GetEntity(entityId)
	if !ok {
		return
	}

	entityCaller(entity).AsyncCallVoid(func() {
		if entityExist(entity) {
			segment(entity)
		}
	})
}
