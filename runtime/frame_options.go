package runtime

// FrameOptions 帧的所有选项
type FrameOptions struct {
	TargetFPS   float32 // 目标FPS
	TotalFrames uint64  // 运行帧数上限
	Blink       bool    // 是否是瞬时运行
}

// FrameOption 实体构建器的选项设置器
type FrameOption func(o *FrameOptions)

// DefaultFrame 默认值
func (Option) DefaultFrame() FrameOption {
	return func(o *FrameOptions) {
		Option{}.TargetFPS(30)(o)
		Option{}.TotalFrames(0)(o)
		Option{}.Blink(false)(o)
	}
}

// TargetFPS 目标FPS
func (Option) TargetFPS(fps float32) FrameOption {
	return func(o *FrameOptions) {
		if fps <= 0 {
			panic("TargetFPS less equal 0 is invalid")
		}
		o.TargetFPS = fps
	}
}

// TotalFrames 运行帧数上限
func (Option) TotalFrames(v uint64) FrameOption {
	return func(o *FrameOptions) {
		if v < 0 {
			panic("TotalFrames less 0 is invalid")
		}
		o.TotalFrames = v
	}
}

// Blink 是否是瞬时运行
func (Option) Blink(blink bool) FrameOption {
	return func(o *FrameOptions) {
		o.Blink = blink
	}
}
