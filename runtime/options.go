package runtime

// Option 所有选项设置器
type Option struct {
	Context _ContextOption // 运行时上下文的选项设置器
	Frame   _FrameOption   // 帧的选项设置器
}
