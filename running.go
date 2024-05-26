package core

// Running 运行接口
type Running interface {
	// Run 运行
	Run()
	// Terminate 停止
	Terminate() <-chan struct{}
}
