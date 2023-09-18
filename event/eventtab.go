package event

import "kit.golaxy.org/golaxy/util/container"

// IEventTab 本地事件表接口，我们可以把一些事件定义在同一个源码文件中，使用事件代码生成器的生成事件表功能，自动生成事件表
type IEventTab interface {
	// Init 初始化事件表
	Init(autoRecover bool, reportError chan error, hookAllocator container.Allocator[Hook], gcCollector container.GCCollector)
	// Get 获取事件
	Get(id int) IEvent
	// Open 打开事件表中所有事件
	Open()
	// Close 关闭事件表中所有事件
	Close()
	// Clean 事件表中的所有事件清除全部订阅者
	Clean()
}
