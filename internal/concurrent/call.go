package concurrent

import (
	"context"
	"fmt"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/generic"
)

var (
	ErrAsyncRetClosed = fmt.Errorf("%w: async result closed", exception.ErrCore)
)

var (
	VoidRet = MakeRet(nil, nil)
)

// MakeRet 创建调用结果
func MakeRet(val any, err error) Ret {
	return Ret{
		Value: val,
		Error: err,
	}
}

// Ret 调用结果
type Ret struct {
	Value any   // 返回值
	Error error // error
}

// OK 是否成功
func (ret Ret) OK() bool {
	return ret.Error == nil
}

// String implements fmt.Stringer
func (ret Ret) String() string {
	if ret.Error != nil {
		return ret.Error.Error()
	}
	return fmt.Sprintf("%v", ret.Value)
}

// MakeAsyncRet 创建异步调用结果
func MakeAsyncRet() chan Ret {
	return make(chan Ret, 1)
}

// AsyncRet 异步调用结果
type AsyncRet <-chan Ret

// Wait 等待异步调用结果
func (asyncRet AsyncRet) Wait(ctx context.Context) Ret {
	if ctx == nil {
		ctx = context.Background()
	}

	select {
	case ret, ok := <-asyncRet:
		if !ok {
			return MakeRet(nil, ErrAsyncRetClosed)
		}
		return ret
	case <-ctx.Done():
		return MakeRet(nil, context.Canceled)
	}
}

// Caller 异步调用发起者
type Caller interface {
	// Call 异步调用函数，有返回值。不会阻塞当前线程，会返回AsyncRet。
	//
	//	注意：
	//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
	//  - 调用过程中的panic信息，均会转换为error返回。
	Call(fun generic.FuncVar0[any, Ret], va ...any) AsyncRet

	// CallDelegate 异步调用委托，有返回值。不会阻塞当前线程，会返回AsyncRet。
	//
	//	注意：
	//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
	//  - 调用过程中的panic信息，均会转换为error返回。
	CallDelegate(fun generic.DelegateFuncVar0[any, Ret], va ...any) AsyncRet

	// CallVoid 异步调用函数，无返回值。在运行时中。不会阻塞当前线程，会返回AsyncRet。
	//
	//	注意：
	//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
	//  - 调用过程中的panic信息，均会转换为error返回。
	CallVoid(fun generic.ActionVar0[any], va ...any) AsyncRet

	// CallVoidDelegate 异步调用委托，无返回值。在运行时中。不会阻塞当前线程，会返回AsyncRet。
	//
	//	注意：
	//	- 代码片段中的线程安全问题，如临界区访问、线程死锁等。
	//  - 调用过程中的panic信息，均会转换为error返回。
	CallVoidDelegate(fun generic.DelegateActionVar0[any], va ...any) AsyncRet
}

// Callee 异步调用接受者
type Callee interface {
	// PushCall 将调用函数压入接受者的任务处理流水线，返回AsyncRet。
	PushCall(fun generic.FuncVar0[any, Ret], va ...any) AsyncRet
	// PushCallDelegate 将调用委托压入接受者的任务处理流水线，返回AsyncRet。
	PushCallDelegate(fun generic.DelegateFuncVar0[any, Ret], va ...any) AsyncRet
	// PushCallVoid 将调用函数压入接受者的任务处理流水线，返回AsyncRet。
	PushCallVoid(fun generic.ActionVar0[any], va ...any) AsyncRet
	// PushCallVoidDelegate 将调用委托压入接受者的任务处理流水线，返回AsyncRet。
	PushCallVoidDelegate(fun generic.DelegateActionVar0[any], va ...any) AsyncRet
}
