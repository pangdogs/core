package golaxy

// Running 运行接口
type Running interface {
	// Run 运行，返回的channel用于线程同步，可以阻塞等待至运行结束
	Run() <-chan struct{}
	// Terminate 停止
	Terminate()
}
