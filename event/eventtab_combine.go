package event

// CombineEventTab 联合事件表，可以将多个事件表联合在一起，方便管理多个事件表
type CombineEventTab []IEventTab

// Init 初始化事件
func (c *CombineEventTab) Init(autoRecover bool, reportError chan error, recursion EventRecursion) {
	for _, tab := range *c {
		tab.Init(autoRecover, reportError, recursion)
	}
}

// Open 打开事件
func (c *CombineEventTab) Open() {
	for _, tab := range *c {
		tab.Open()
	}
}

// Close 关闭事件
func (c *CombineEventTab) Close() {
	for _, tab := range *c {
		tab.Close()
	}
}

// Clean 清除全部订阅者
func (c *CombineEventTab) Clean() {
	for _, tab := range *c {
		tab.Clean()
	}
}

// Get 获取事件
func (c *CombineEventTab) Get(id uint64) IEvent {
	for _, tab := range *c {
		event := tab.Get(id)
		if event != nil {
			return event
		}
	}
	return nil
}
