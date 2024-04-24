package runtime

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/option"
)

// FrameOptions 帧的所有选项
type FrameOptions struct {
	TargetFPS   float32 // 目标FPS
	TotalFrames uint64  // 运行帧数上限
}

type _FrameOption struct{}

// Default 默认值
func (_FrameOption) Default() option.Setting[FrameOptions] {
	return func(o *FrameOptions) {
		With.Frame.TargetFPS(30)(o)
		With.Frame.TotalFrames(0)(o)
	}
}

// TargetFPS 目标FPS
func (_FrameOption) TargetFPS(fps float32) option.Setting[FrameOptions] {
	return func(o *FrameOptions) {
		if fps <= 0 {
			panic(fmt.Errorf("%w: %w: TargetFPS less equal 0 is invalid", ErrFrame, exception.ErrArgs))
		}
		o.TargetFPS = fps
	}
}

// TotalFrames 运行帧数上限
func (_FrameOption) TotalFrames(v uint64) option.Setting[FrameOptions] {
	return func(o *FrameOptions) {
		if v < 0 {
			panic(fmt.Errorf("%w: %w: TotalFrames less 0 is invalid", ErrFrame, exception.ErrArgs))
		}
		o.TotalFrames = v
	}
}
