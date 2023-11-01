package runtime

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/util/option"
)

type _FrameOption struct{}

// FrameOptions 帧的所有选项
type FrameOptions struct {
	TargetFPS   float32 // 目标FPS
	TotalFrames uint64  // 运行帧数上限
	Blink       bool    // 是否是瞬时运行
}

// Default 默认值
func (_FrameOption) Default() option.Setting[FrameOptions] {
	return func(o *FrameOptions) {
		_FrameOption{}.TargetFPS(30)(o)
		_FrameOption{}.TotalFrames(0)(o)
		_FrameOption{}.Blink(false)(o)
	}
}

// TargetFPS 目标FPS
func (_FrameOption) TargetFPS(fps float32) option.Setting[FrameOptions] {
	return func(o *FrameOptions) {
		if fps <= 0 {
			panic(fmt.Errorf("%w: %w: TargetFPS less equal 0 is invalid", ErrFrame, internal.ErrArgs))
		}
		o.TargetFPS = fps
	}
}

// TotalFrames 运行帧数上限
func (_FrameOption) TotalFrames(v uint64) option.Setting[FrameOptions] {
	return func(o *FrameOptions) {
		if v < 0 {
			panic(fmt.Errorf("%w: %w: TotalFrames less 0 is invalid", ErrFrame, internal.ErrArgs))
		}
		o.TotalFrames = v
	}
}

// Blink 是否是瞬时运行
func (_FrameOption) Blink(blink bool) option.Setting[FrameOptions] {
	return func(o *FrameOptions) {
		o.Blink = blink
	}
}
