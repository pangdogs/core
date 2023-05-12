package runtime

import (
	"fmt"
	"kit.golaxy.org/golaxy/ec"
)

// NewRet 创建调用结果
func NewRet(err error, val any) Ret {
	return Ret{
		Error: err,
		Value: val,
	}
}

// Ret 调用结果
type Ret struct {
	Error error // error
	Value any   // 返回值
}

// OK 是否成功
func (ret Ret) OK() bool {
	return ret.Error == nil
}

// String 字符串化
func (ret Ret) String() string {
	if ret.Error != nil {
		return ret.Error.Error()
	}
	return fmt.Sprintf("%v", ret.Value)
}

// AsyncRet 异步调用结果
type AsyncRet <-chan Ret

// Caller 异步调用发起者
type Caller interface {
	// AwaitCall 同步调用。在运行时中，将代码片段压入任务流水线，串行化的进行调用，会阻塞并等待返回值。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 当运行时的AwaitCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会转换为error返回。
	AwaitCall(segment func() Ret) Ret

	// AsyncCall 异步调用。在运行时中，将代码片段压入任务流水线，串行化的进行调用，不会阻塞，会返回AsyncRet。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 在代码片段中，如果向调用方所在的运行时发起同步调用，并且调用方也在阻塞AsyncRet等待返回值，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会转换为error返回。
	AsyncCall(segment func() Ret) AsyncRet

	// AwaitCallNoRet 同步调用，无返回值。在运行时中，将代码片段压入任务流水线，串行化的进行调用，会阻塞，没有返回值。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 当运行时的AwaitCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会抛弃。
	AwaitCallNoRet(segment func())

	// AsyncCallNoRet 异步调用，无返回值。在运行时中，将代码片段压入任务流水线，串行化的进行调用，不会阻塞，没有返回值。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//  - 调用过程中的panic信息，均会抛弃。
	AsyncCallNoRet(segment func())
}

// Callee 调用接收者
type Callee interface {
	// PushCall 将代码片段压入接收者的任务处理流水线，串行化的进行调用。
	PushCall(segment func())
}

func entityCaller(entity ec.Entity) Caller {
	return Get(entity)
}

func entityExist(entity ec.Entity) bool {
	_, ok := Get(entity).GetEntityMgr().GetEntity(entity.GetId())
	return ok
}

// AwaitCall 同步调用。在运行时中，将代码片段压入任务流水线，串行化的进行调用，会阻塞并等待返回值。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 当运行时的AwaitCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) AwaitCall(segment func() Ret) Ret {
	var ret Ret

	func() {
		defer func() {
			if info := recover(); info != nil {
				err, ok := info.(error)
				if !ok {
					err = fmt.Errorf("%v", info)
				}
				ret = NewRet(err, nil)
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

// AsyncCall 异步调用。在运行时中，将代码片段压入任务流水线，串行化的进行调用，不会阻塞，会返回AsyncRet。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 在代码片段中，如果向调用方所在的运行时发起同步调用，并且调用方也在阻塞AsyncRet等待返回值，那么会造成线程死锁。
//	- 调用过程中的panic信息，均会转换为error返回。
func (ctx *ContextBehavior) AsyncCall(segment func() Ret) AsyncRet {
	asyncRet := make(chan Ret)

	go func() {
		defer func() {
			if info := recover(); info != nil {
				err, ok := info.(error)
				if !ok {
					err = fmt.Errorf("%v", info)
				}
				asyncRet <- NewRet(err, nil)
				close(asyncRet)
			}
		}()

		if segment == nil {
			panic("nil segment")
		}

		ctx.callee.PushCall(func() {
			asyncRet <- segment()
			close(asyncRet)
		})
	}()

	return asyncRet
}

// AwaitCallNoRet 同步调用，无返回值。在运行时中，将代码片段压入任务流水线，串行化的进行调用，会阻塞，没有返回值。
//
//	注意：
//	- 代码片段中的线程安全问题。
//	- 当运行时的AwaitCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
//	- 调用过程中的panic信息，均会抛弃。
func (ctx *ContextBehavior) AwaitCallNoRet(segment func()) {
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
