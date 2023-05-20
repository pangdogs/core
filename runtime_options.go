package golaxy

import (
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/util"
	"time"
)

// RuntimeOptions 创建运行时的所有选项
type RuntimeOptions struct {
	CompositeFace        util.Face[Runtime] // 扩展者，需要扩展运行时自身功能时需要使用
	EnableAutoRun        bool               // 是否开启自动运行
	ProcessQueueCapacity int                // 任务处理流水线大小
	ProcessQueueTimeout  time.Duration      // 当任务处理流水线满时，向其插入代码片段的超时时间，为0表示不等待直接报错
	SyncCallTimeout      time.Duration      // 同步调用超时时间，为0表示不处理超时，此时两个运行时互相同步调用会死锁
	Frame                runtime.Frame      // 帧，设置为nil表示不使用帧更新特性
	GCInterval           time.Duration      // GC间隔时长
	CustomGC             func(rt Runtime)   // 自定义GC
}

// RuntimeOption 创建运行时的选项设置器
type RuntimeOption func(o *RuntimeOptions)

// WithRuntimeOption 创建运行时的所有选项设置器
type WithRuntimeOption struct{}

// Default 默认值
func (WithRuntimeOption) Default() RuntimeOption {
	return func(o *RuntimeOptions) {
		WithRuntimeOption{}.CompositeFace(util.Face[Runtime]{})(o)
		WithRuntimeOption{}.EnableAutoRun(false)(o)
		WithRuntimeOption{}.ProcessQueueCapacity(128)(o)
		WithRuntimeOption{}.ProcessQueueTimeout(0)(o)
		WithRuntimeOption{}.SyncCallTimeout(3 * time.Second)(o)
		WithRuntimeOption{}.Frame(nil)(o)
		WithRuntimeOption{}.GCInterval(10 * time.Second)(o)
		WithRuntimeOption{}.CustomGC(nil)(o)
	}
}

// CompositeFace 扩展者，需要扩展运行时自身功能时需要使用
func (WithRuntimeOption) CompositeFace(face util.Face[Runtime]) RuntimeOption {
	return func(o *RuntimeOptions) {
		o.CompositeFace = face
	}
}

// EnableAutoRun 是否开启自动运行
func (WithRuntimeOption) EnableAutoRun(b bool) RuntimeOption {
	return func(o *RuntimeOptions) {
		o.EnableAutoRun = b
	}
}

// ProcessQueueCapacity 任务处理流水线大小
func (WithRuntimeOption) ProcessQueueCapacity(cap int) RuntimeOption {
	return func(o *RuntimeOptions) {
		if cap <= 0 {
			panic("ProcessQueueCapacity less equal 0 is invalid")
		}
		o.ProcessQueueCapacity = cap
	}
}

// ProcessQueueTimeout 当任务处理流水线满时，向其插入代码片段的超时时间，为0表示不等待直接报错
func (WithRuntimeOption) ProcessQueueTimeout(dur time.Duration) RuntimeOption {
	return func(o *RuntimeOptions) {
		o.ProcessQueueTimeout = dur
	}
}

// SyncCallTimeout 同步调用超时时间，为0表示不处理超时，此时两个运行时互相同步调用会死锁
func (WithRuntimeOption) SyncCallTimeout(dur time.Duration) RuntimeOption {
	return func(o *RuntimeOptions) {
		o.SyncCallTimeout = dur
	}
}

// Frame 帧，设置为nil表示不使用帧更新特性
func (WithRuntimeOption) Frame(frame runtime.Frame) RuntimeOption {
	return func(o *RuntimeOptions) {
		o.Frame = frame
	}
}

// GCInterval GC间隔时长
func (WithRuntimeOption) GCInterval(dur time.Duration) RuntimeOption {
	return func(o *RuntimeOptions) {
		if dur <= 0 {
			panic("GCInterval less equal 0 is invalid")
		}
		o.GCInterval = dur
	}
}

// CustomGC 自定义GC
func (WithRuntimeOption) CustomGC(fn func(rt Runtime)) RuntimeOption {
	return func(o *RuntimeOptions) {
		o.CustomGC = fn
	}
}
