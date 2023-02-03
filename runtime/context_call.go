package runtime

import (
	"fmt"
	"github.com/golaxy-kit/golaxy/ec"
	"github.com/golaxy-kit/golaxy/internal"
)

// Ret 调用结果
type Ret = internal.Ret

// Callee 调用接收者
type Callee = internal.Callee

type _Call interface {
	internal.Call
	setCallee(callee Callee)
}

func entityCall(entity ec.Entity) internal.Call {
	return Get(entity)
}

func entityExist(entity ec.Entity) bool {
	_, ok := Get(entity).GetEntityMgr().GetEntity(entity.GetID())
	return ok
}

// SyncCall 同步调用。在运行时中，将代码片段压入任务流水线，串行化的进行调用，会阻塞并等待返回值。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 当运行时的SyncCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) SyncCall(segment func() Ret) Ret {
	var ret Ret

	func() {
		defer func() {
			if info := recover(); info != nil {
				err, ok := info.(error)
				if !ok {
					err = fmt.Errorf("%v", info)
				}
				ret = Ret{Err: err}
			}
		}()

		if segment == nil {
			panic("nil segment")
		}

		ctx.callee.PushCall(func() {
			ret = segment()
		})
	}()

	return ret
}

// AsyncCall 异步调用。在运行时中，将代码片段压入任务流水线，串行化的进行调用，不会阻塞，会返回result channel。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 在代码片段中，如果向调用方所在的运行时发起同步调用，并且调用方也在阻塞result channel等待返回值，那么会造成线程死锁。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) AsyncCall(segment func() Ret) <-chan Ret {
	ret := make(chan Ret)

	go func() {
		defer func() {
			if info := recover(); info != nil {
				err, ok := info.(error)
				if !ok {
					err = fmt.Errorf("%v", info)
				}
				ret <- Ret{Err: err}
			}
		}()

		if segment == nil {
			panic("nil segment")
		}

		ctx.callee.PushCall(func() {
			ret <- segment()
		})
	}()

	return ret
}

// SyncCallNoRet 同步调用，无返回值。在运行时中，将代码片段压入任务流水线，串行化的进行调用，会阻塞，没有返回值。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 当运行时的SyncCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
//	- 调用过程中的panic信息，均会抛弃。
func (ctx *ContextBehavior) SyncCallNoRet(segment func()) {
	func() {
		defer func() {
			recover()
		}()

		if segment == nil {
			panic("nil segment")
		}

		ctx.callee.PushCall(segment)
	}()
}

// AsyncCallNoRet 异步调用，无返回值。在运行时中，将代码片段压入任务流水线，串行化的进行调用，不会阻塞，没有返回值。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 调用过程中的panic信息，均会抛弃。
func (ctx *ContextBehavior) AsyncCallNoRet(segment func()) {
	go func() {
		defer func() {
			recover()
		}()

		if segment == nil {
			panic("nil segment")
		}

		ctx.callee.PushCall(segment)
	}()
}

func (ctx *ContextBehavior) setCallee(callee Callee) {
	ctx.callee = callee
}
