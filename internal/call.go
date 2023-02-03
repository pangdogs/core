package internal

// Ret 调用结果
type Ret struct {
	Err error // error
	Ret any   // 返回值
}

// Call 调用
type Call interface {
	// SyncCall 同步调用。在运行时中，将代码片段压入任务流水线，串行化的进行调用，会阻塞并等待返回值。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 当运行时的SyncCallTimeout选项设置为0时，在代码片段中，如果向调用方所在的运行时发起同步调用，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会转换为error返回。
	SyncCall(segment func() Ret) Ret

	// AsyncCall 异步调用。在运行时中，将代码片段压入任务流水线，串行化的进行调用，不会阻塞，会返回result channel。
	//
	//	注意：
	//	- 代码片段中的线程安全问题。
	//	- 在代码片段中，如果向调用方所在的运行时发起同步调用，并且调用方也在阻塞result channel等待返回值，那么会造成线程死锁。
	//  - 调用过程中的panic信息，均会转换为error返回。
	AsyncCall(segment func() Ret) <-chan Ret

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

// Callee 调用接收者
type Callee interface {
	// PushCall 将代码片段压入接收者的任务处理流水线，串行化的进行调用。
	PushCall(segment func())
}
