package internal

// Running 运行接口
type Running interface {
	// Run 运行，返回的channel用于线程同步，可以阻塞等待至运行结束
	Run() <-chan struct{}
	// Stop 停止
	Stop()
}
