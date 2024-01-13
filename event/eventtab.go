package event

// IEventTab 本地事件表接口，方便管理多个事件
/*
使用方式：
	1.在定义事件的源码文件（.go）头部添加以下注释，在编译前自动化生成代码：
	//go:generate go run git.golaxy.org/core/event/eventcode --decl_file=$GOFILE gen_eventtab --package=$GOPACKAGE --name={事件表名称}

定义事件的选项（添加到定义事件的注释里）：
	1.事件表初始化时，该事件使用的递归处理方式，不填表示使用事件表初始化参数值
		[EventRecursion_Allow]
		[EventRecursion_Disallow]
		[EventRecursion_NotEmit]
		[EventRecursion_Discard]
		[EventRecursion_Deepest]
*/
type IEventTab interface {
	IEventCtrl
	// Get 获取事件
	Get(id int) IEvent
}
