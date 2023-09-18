package golaxy

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/util/iface"
	"time"
)

type _RuntimeOption struct{}

type (
	CustomGC = func(rt Runtime) // 自定义GC函数
)

// RuntimeOptions 创建运行时的所有选项
type RuntimeOptions struct {
	CompositeFace        iface.Face[Runtime] // 扩展者，需要扩展运行时自身功能时需要使用
	AutoRun              bool                // 是否开启自动运行
	ProcessQueueCapacity int                 // 任务处理流水线大小
	ProcessQueueTimeout  time.Duration       // 当任务处理流水线满时，向其插入代码片段的超时时间，为0表示不等待直接报错
	SyncCallTimeout      time.Duration       // 同步调用超时时间，为0表示不处理超时，此时两个运行时互相同步调用会死锁
	Frame                runtime.Frame       // 帧，设置为nil表示不使用帧更新特性
	GCInterval           time.Duration       // GC间隔时长
	CustomGC             CustomGC            // 自定义GC
}

// RuntimeOption 创建运行时的选项设置器
type RuntimeOption func(o *RuntimeOptions)

// Default 运行时的默认值
func (_RuntimeOption) Default() RuntimeOption {
	return func(o *RuntimeOptions) {
		_RuntimeOption{}.CompositeFace(iface.Face[Runtime]{})(o)
		_RuntimeOption{}.AutoRun(false)(o)
		_RuntimeOption{}.ProcessQueueCapacity(128)(o)
		_RuntimeOption{}.ProcessQueueTimeout(0)(o)
		_RuntimeOption{}.SyncCallTimeout(3 * time.Second)(o)
		_RuntimeOption{}.Frame(nil)(o)
		_RuntimeOption{}.GCInterval(10 * time.Second)(o)
		_RuntimeOption{}.CustomGC(nil)(o)
	}
}

// CompositeFace 运行时的扩展者，需要扩展运行时自身功能时需要使用
func (_RuntimeOption) CompositeFace(face iface.Face[Runtime]) RuntimeOption {
	return func(o *RuntimeOptions) {
		o.CompositeFace = face
	}
}

// AutoRun 运行时是否开启自动运行
func (_RuntimeOption) AutoRun(b bool) RuntimeOption {
	return func(o *RuntimeOptions) {
		o.AutoRun = b
	}
}

// ProcessQueueCapacity 运行时的任务处理流水线大小
func (_RuntimeOption) ProcessQueueCapacity(cap int) RuntimeOption {
	return func(o *RuntimeOptions) {
		if cap <= 0 {
			panic(fmt.Errorf("%w: %w: ProcessQueueCapacity less equal 0 is invalid", ErrRuntime, internal.ErrArgs))
		}
		o.ProcessQueueCapacity = cap
	}
}

// ProcessQueueTimeout 运行时的当任务处理流水线满时，向其插入代码片段的超时时间，为0表示不等待直接报错
func (_RuntimeOption) ProcessQueueTimeout(dur time.Duration) RuntimeOption {
	return func(o *RuntimeOptions) {
		o.ProcessQueueTimeout = dur
	}
}

// SyncCallTimeout 运行时的同步调用超时时间，为0表示不处理超时，此时两个运行时互相同步调用会死锁
func (_RuntimeOption) SyncCallTimeout(dur time.Duration) RuntimeOption {
	return func(o *RuntimeOptions) {
		o.SyncCallTimeout = dur
	}
}

// Frame 运行时的帧，设置为nil表示不使用帧更新特性
func (_RuntimeOption) Frame(frame runtime.Frame) RuntimeOption {
	return func(o *RuntimeOptions) {
		o.Frame = frame
	}
}

// GCInterval 运行时的GC间隔时长
func (_RuntimeOption) GCInterval(dur time.Duration) RuntimeOption {
	return func(o *RuntimeOptions) {
		if dur <= 0 {
			panic(fmt.Errorf("%w: %w: GCInterval less equal 0 is invalid", ErrRuntime, internal.ErrArgs))
		}
		o.GCInterval = dur
	}
}

// CustomGC 运行时的自定义GC
func (_RuntimeOption) CustomGC(fn CustomGC) RuntimeOption {
	return func(o *RuntimeOptions) {
		o.CustomGC = fn
	}
}
