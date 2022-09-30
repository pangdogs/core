package runtime

// Callee 安全调用接收者
type Callee interface {
	// PushCall 将代码片段压入接收者的任务处理流水线，串行化的进行调用。
	PushCall(segment func())
}
