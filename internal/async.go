package internal

import "fmt"

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
	// SyncCall 同步调用。在运行时中，将代码片段压入任务流水线，串行化的进行调用，会阻塞并等待返回值。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 当运行时的SyncCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会转换为error返回。
	SyncCall(segment func() Ret) Ret

	// AsyncCall 异步调用。在运行时中，将代码片段压入任务流水线，串行化的进行调用，不会阻塞，会返回AsyncRet。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 在代码片段中，如果向调用方所在的运行时发起同步调用，并且调用方也在阻塞AsyncRet等待返回值，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会转换为error返回。
	AsyncCall(segment func() Ret) AsyncRet

	// SyncCallNoRet 同步调用，无返回值。在运行时中，将代码片段压入任务流水线，串行化的进行调用，会阻塞，没有返回值。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 当运行时的SyncCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会抛弃。
	SyncCallNoRet(segment func())

	// AsyncCallNoRet 异步调用，无返回值。在运行时中，将代码片段压入任务流水线，串行化的进行调用，不会阻塞，没有返回值。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//  - 调用过程中的panic信息，均会抛弃。
	AsyncCallNoRet(segment func())
}
