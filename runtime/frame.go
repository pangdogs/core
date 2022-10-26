package runtime

import (
	"time"
)

// Frame 帧，在运行时初始化时可以设置帧，用于设置运行时帧更新方式，在逻辑运行过程中可以在运行时上下文中获取帧信息
type Frame interface {
	// GetTargetFPS 获取目标FPS
	GetTargetFPS() float32
	// GetCurFPS 获取当前FPS
	GetCurFPS() float32
	// GetTotalFrames 获取运行帧数上限
	GetTotalFrames() uint64
	// GetCurFrames 获取当前帧数
	GetCurFrames() uint64
	setCurFrames(v uint64)
	// Blink 是否是瞬时运行
	Blink() bool
	// GetRunningBeginTime 获取运行开始时间
	GetRunningBeginTime() time.Time
	// GetRunningElapseTime 获取运行持续时间
	GetRunningElapseTime() time.Duration
	// GetFrameBeginTime 获取当前帧开始时间
	GetFrameBeginTime() time.Time
	// GetLastFrameElapseTime 获取上一帧耗时
	GetLastFrameElapseTime() time.Duration
	// GetUpdateBeginTime 获取当前帧更新开始时间
	GetUpdateBeginTime() time.Time
	// GetLastUpdateElapseTime 获取上一次帧更新耗时
	GetLastUpdateElapseTime() time.Duration
	runningBegin()
	runningEnd()
	frameBegin()
	frameEnd()
	updateBegin()
	updateEnd()
}

// NewFrame 创建帧，在运行时初始化时可以设置帧，用于设置运行时帧更新方式，在逻辑运行过程中可以在运行时上下文中获取帧信息。
//
//	@param targetFPS 目标FPS。
//	@param totalFrames 运行帧数上限。
//	@param blink 是否是瞬时运行。
//	@return 帧。
func NewFrame(targetFPS float32, totalFrames uint64, blink bool) Frame {
	frame := &_Frame{}
	frame.init(targetFPS, totalFrames, blink)
	return frame
}

type _Frame struct {
	targetFPS, curFPS      float32
	totalFrames, curFrames uint64
	blink                  bool
	blinkFrameTime         time.Duration
	runningBeginTime       time.Time
	runningElapseTime      time.Duration
	frameBeginTime         time.Time
	lastFrameElapseTime    time.Duration
	updateBeginTime        time.Time
	lastUpdateElapseTime   time.Duration
	statFPSBeginTime       time.Time
	statFPSFrames          uint64
}

func (frame *_Frame) init(targetFPS float32, totalFrames uint64, blink bool) {
	if targetFPS <= 0 {
		panic("targetFPS less equal 0 invalid")
	}

	if totalFrames < 0 {
		panic("totalFrames less 0 invalid")
	}

	frame.targetFPS = targetFPS
	frame.totalFrames = totalFrames
	frame.blink = blink

	if blink {
		frame.blinkFrameTime = time.Duration(float64(time.Second) / float64(targetFPS))
	}
}

func (frame *_Frame) GetTargetFPS() float32 {
	return frame.targetFPS
}

func (frame *_Frame) GetCurFPS() float32 {
	return frame.curFPS
}

func (frame *_Frame) GetTotalFrames() uint64 {
	return frame.totalFrames
}

func (frame *_Frame) GetCurFrames() uint64 {
	return frame.curFrames
}

func (frame *_Frame) setCurFrames(v uint64) {
	frame.curFrames = v
}

func (frame *_Frame) Blink() bool {
	return frame.blink
}

func (frame *_Frame) GetRunningBeginTime() time.Time {
	return frame.runningBeginTime
}

func (frame *_Frame) GetRunningElapseTime() time.Duration {
	return frame.runningElapseTime
}

func (frame *_Frame) GetFrameBeginTime() time.Time {
	return frame.frameBeginTime
}

func (frame *_Frame) GetLastFrameElapseTime() time.Duration {
	return frame.lastFrameElapseTime
}

func (frame *_Frame) GetUpdateBeginTime() time.Time {
	return frame.updateBeginTime
}

func (frame *_Frame) GetLastUpdateElapseTime() time.Duration {
	return frame.lastUpdateElapseTime
}

func (frame *_Frame) runningBegin() {
	now := time.Now()

	frame.curFPS = 0
	frame.curFrames = 0

	frame.statFPSBeginTime = now
	frame.statFPSFrames = 0

	frame.runningBeginTime = now
	frame.runningElapseTime = 0

	frame.frameBeginTime = now
	frame.lastFrameElapseTime = 0

	frame.updateBeginTime = now
	frame.lastUpdateElapseTime = 0
}

func (frame *_Frame) runningEnd() {
	if frame.blink {
		frame.curFPS = float32(float64(frame.curFrames) / time.Now().Sub(frame.runningBeginTime).Seconds())
	}
}

func (frame *_Frame) frameBegin() {
	now := time.Now()

	frame.frameBeginTime = now

	if !frame.blink {
		statInterval := now.Sub(frame.statFPSBeginTime).Seconds()
		if statInterval >= 1 {
			frame.curFPS = float32(float64(frame.statFPSFrames) / statInterval)
			frame.statFPSBeginTime = now
			frame.statFPSFrames = 0
		}
	}
}

func (frame *_Frame) frameEnd() {
	if frame.blink {
		frame.runningElapseTime += frame.blinkFrameTime
	} else {
		frame.lastFrameElapseTime = time.Now().Sub(frame.frameBeginTime)
		frame.runningElapseTime += frame.lastFrameElapseTime
		frame.statFPSFrames++
	}
}

func (frame *_Frame) updateBegin() {
	frame.updateBeginTime = time.Now()
}

func (frame *_Frame) updateEnd() {
	frame.lastUpdateElapseTime = time.Now().Sub(frame.updateBeginTime)
}
