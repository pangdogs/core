package golaxy

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/exception"
	"kit.golaxy.org/golaxy/runtime"
	"kit.golaxy.org/golaxy/util/generic"
	"kit.golaxy.org/golaxy/util/iface"
	"kit.golaxy.org/golaxy/util/option"
	"time"
)

type _RuntimeOption struct{}

type (
	CustomGC = generic.DelegateAction1[Runtime] // 自定义GC函数
)

// RuntimeOptions 创建运行时的所有选项
type RuntimeOptions struct {
	CompositeFace        iface.Face[Runtime] // 扩展者，需要扩展运行时自身功能时需要使用
	AutoRun              bool                // 是否开启自动运行
	ProcessQueueCapacity int                 // 任务处理流水线大小
	ProcessQueueTimeout  time.Duration       // 任务处理流水线满时，向其插入任务的超时时间，为0表示不等待直接报错
	Frame                runtime.Frame       // 帧，设置为nil表示不使用帧更新特性
	GCInterval           time.Duration       // GC间隔时长
	CustomGC             CustomGC            // 自定义GC
}

// Default 运行时的默认值
func (_RuntimeOption) Default() option.Setting[RuntimeOptions] {
	return func(o *RuntimeOptions) {
		_RuntimeOption{}.CompositeFace(iface.Face[Runtime]{})(o)
		_RuntimeOption{}.AutoRun(false)(o)
		_RuntimeOption{}.ProcessQueueCapacity(128)(o)
		_RuntimeOption{}.ProcessQueueTimeout(0)(o)
		_RuntimeOption{}.Frame(nil)(o)
		_RuntimeOption{}.GCInterval(10 * time.Second)(o)
		_RuntimeOption{}.CustomGC(nil)(o)
	}
}

// CompositeFace 运行时的扩展者，需要扩展运行时自身功能时需要使用
func (_RuntimeOption) CompositeFace(face iface.Face[Runtime]) option.Setting[RuntimeOptions] {
	return func(o *RuntimeOptions) {
		o.CompositeFace = face
	}
}

// AutoRun 运行时是否开启自动运行
func (_RuntimeOption) AutoRun(b bool) option.Setting[RuntimeOptions] {
	return func(o *RuntimeOptions) {
		o.AutoRun = b
	}
}

// ProcessQueueCapacity 任务处理流水线大小
func (_RuntimeOption) ProcessQueueCapacity(cap int) option.Setting[RuntimeOptions] {
	return func(o *RuntimeOptions) {
		if cap <= 0 {
			panic(fmt.Errorf("%w: %w: ProcessQueueCapacity less equal 0 is invalid", ErrRuntime, exception.ErrArgs))
		}
		o.ProcessQueueCapacity = cap
	}
}

// ProcessQueueTimeout 任务处理流水线满时，向其插入任务的超时时间，为0表示不等待直接报错
func (_RuntimeOption) ProcessQueueTimeout(dur time.Duration) option.Setting[RuntimeOptions] {
	return func(o *RuntimeOptions) {
		o.ProcessQueueTimeout = dur
	}
}

// Frame 运行时的帧，设置为nil表示不使用帧更新特性
func (_RuntimeOption) Frame(frame runtime.Frame) option.Setting[RuntimeOptions] {
	return func(o *RuntimeOptions) {
		o.Frame = frame
	}
}

// GCInterval 运行时的GC间隔时长
func (_RuntimeOption) GCInterval(dur time.Duration) option.Setting[RuntimeOptions] {
	return func(o *RuntimeOptions) {
		if dur <= 0 {
			panic(fmt.Errorf("%w: %w: GCInterval less equal 0 is invalid", ErrRuntime, exception.ErrArgs))
		}
		o.GCInterval = dur
	}
}

// CustomGC 运行时的自定义GC
func (_RuntimeOption) CustomGC(fn CustomGC) option.Setting[RuntimeOptions] {
	return func(o *RuntimeOptions) {
		o.CustomGC = fn
	}
}
