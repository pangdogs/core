package service

import (
	"fmt"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/internal/concurrent"
	"git.golaxy.org/core/util/generic"
	"git.golaxy.org/core/util/uid"
	_ "unsafe"
)

type (
	Ret      = concurrent.Ret      // 调用结果
	AsyncRet = concurrent.AsyncRet // 异步调用结果
)

// Caller 异步调用发起者
type Caller interface {
	// Call 查找实体并异步调用函数，有返回值。不会阻塞当前线程，会返回AsyncRet。
	//
	//	注意：
	//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
	//  - 调用过程中的panic信息，均会转换为error返回。
	Call(entityId uid.Id, fun generic.FuncVar1[ec.Entity, any, Ret], va ...any) AsyncRet

	// CallDelegate 查找实体并异步调用委托，有返回值。不会阻塞当前线程，会返回AsyncRet。
	//
	//	注意：
	//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
	//  - 调用过程中的panic信息，均会转换为error返回。
	CallDelegate(entityId uid.Id, fun generic.DelegateFuncVar1[ec.Entity, any, Ret], va ...any) AsyncRet

	// CallVoid 查找实体并异步调用函数，无返回值。在运行时中。不会阻塞当前线程，会返回AsyncRet。
	//
	//	注意：
	//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
	//  - 调用过程中的panic信息，均会转换为error返回。
	CallVoid(entityId uid.Id, fun generic.ActionVar1[ec.Entity, any], va ...any) AsyncRet

	// CallVoidDelegate 查找实体并异步调用委托，无返回值。在运行时中。不会阻塞当前线程，会返回AsyncRet。
	//
	//	注意：
	//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
	//  - 调用过程中的panic信息，均会转换为error返回。
	CallVoidDelegate(entityId uid.Id, fun generic.DelegateActionVar1[ec.Entity, any], va ...any) AsyncRet
}

//go:linkname getCaller git.golaxy.org/core/runtime.getCaller
func getCaller(provider concurrent.ConcurrentContextProvider) concurrent.Caller

func makeAsyncErr(err error) AsyncRet {
	asyncRet := concurrent.MakeAsyncRet()
	asyncRet <- concurrent.MakeRet(nil, err)
	close(asyncRet)
	return asyncRet
}

func checkEntity(entity ec.Entity) error {
	if entity.GetState() >= ec.EntityState_Leave {
		return fmt.Errorf("%w: entity not alive", ErrContext)
	}
	return nil
}

// Call 查找实体并异步调用函数，有返回值。不会阻塞当前线程，会返回AsyncRet。
//
//		注意：
//		- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
//	 - 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) Call(entityId uid.Id, fun generic.FuncVar1[ec.Entity, any, Ret], va ...any) AsyncRet {
	entity, err := ctx.getEntity(entityId)
	if err != nil {
		return makeAsyncErr(err)
	}

	return getCaller(entity).Call(func(va ...any) concurrent.Ret {
		entity := va[0].(ec.Entity)
		fun := va[1].(generic.FuncVar1[ec.Entity, any, Ret])
		funVa := va[2].([]any)

		if err := checkEntity(entity); err != nil {
			return concurrent.MakeRet(nil, err)
		}

		return fun.Exec(entity, funVa...)
	}, entity, fun, va)
}

// CallDelegate 查找实体并异步调用委托，有返回值。不会阻塞当前线程，会返回AsyncRet。
//
//		注意：
//		- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
//	 - 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) CallDelegate(entityId uid.Id, fun generic.DelegateFuncVar1[ec.Entity, any, Ret], va ...any) AsyncRet {
	entity, err := ctx.getEntity(entityId)
	if err != nil {
		return makeAsyncErr(err)
	}

	return getCaller(entity).Call(func(va ...any) concurrent.Ret {
		entity := va[0].(ec.Entity)
		fun := va[1].(generic.DelegateFuncVar1[ec.Entity, any, Ret])
		funVa := va[2].([]any)

		if err := checkEntity(entity); err != nil {
			return concurrent.MakeRet(nil, err)
		}

		return fun.Exec(nil, entity, funVa...)
	}, entity, fun, va)
}

// CallVoid 查找实体并异步调用函数，无返回值。在运行时中。不会阻塞当前线程，会返回AsyncRet。
//
//		注意：
//		- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
//	 - 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) CallVoid(entityId uid.Id, fun generic.ActionVar1[ec.Entity, any], va ...any) AsyncRet {
	entity, err := ctx.getEntity(entityId)
	if err != nil {
		return makeAsyncErr(err)
	}

	return getCaller(entity).Call(func(va ...any) concurrent.Ret {
		entity := va[0].(ec.Entity)
		fun := va[1].(generic.ActionVar1[ec.Entity, any])
		funVa := va[2].([]any)

		if err := checkEntity(entity); err != nil {
			return concurrent.MakeRet(nil, err)
		}

		fun.Exec(entity, funVa...)

		return concurrent.VoidRet
	}, entity, fun, va)
}

// CallVoidDelegate 查找实体并异步调用委托，无返回值。在运行时中。不会阻塞当前线程，会返回AsyncRet。
//
//		注意：
//		- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
//	 - 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) CallVoidDelegate(entityId uid.Id, fun generic.DelegateActionVar1[ec.Entity, any], va ...any) AsyncRet {
	entity, err := ctx.getEntity(entityId)
	if err != nil {
		return makeAsyncErr(err)
	}

	return getCaller(entity).Call(func(va ...any) concurrent.Ret {
		entity := va[0].(ec.Entity)
		fun := va[1].(generic.DelegateActionVar1[ec.Entity, any])
		funVa := va[2].([]any)

		if err := checkEntity(entity); err != nil {
			return concurrent.MakeRet(nil, err)
		}

		fun.Exec(nil, entity, funVa...)

		return concurrent.VoidRet
	}, entity, fun, va)
}

func (ctx *ContextBehavior) getEntity(id uid.Id) (ec.ConcurrentEntity, error) {
	entity, ok := ctx.entityMgr.GetEntity(id)
	if !ok {
		return nil, fmt.Errorf("%w: entity not exist", ErrContext)
	}
	return entity, nil
}
