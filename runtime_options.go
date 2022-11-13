package galaxy

import (
	"github.com/galaxy-kit/galaxy/runtime"
	"github.com/galaxy-kit/galaxy/util"
	"time"
)

// RuntimeOptions 创建运行时的所有选项
type RuntimeOptions struct {
	Inheritor            util.Face[Runtime] // 继承者，需要拓展运行时自身功能时需要使用
	EnableAutoRun        bool               // 是否开启自动运行
	ProcessQueueCapacity int                // 任务处理流水线大小
	ProcessQueueTimeout  time.Duration      // 任务插入流水线超时时长
	Frame                runtime.Frame      // 帧
	GCInterval           time.Duration      // GC间隔时长
}

// WithRuntimeOption 创建运行时的选项设置器
type WithRuntimeOption func(o *RuntimeOptions)

// RuntimeOption 创建运行时的选项
var RuntimeOption = _RuntimeOption{}

type _RuntimeOption struct{}

// Default 默认值
func (_RuntimeOption) Default() WithRuntimeOption {
	return func(o *RuntimeOptions) {
		o.Inheritor = util.Face[Runtime]{}
		o.EnableAutoRun = false
		o.ProcessQueueCapacity = 128
		o.ProcessQueueTimeout = 5 * time.Second
		o.Frame = nil
		o.GCInterval = 10 * time.Second
	}
}

// Inheritor 继承者，需要拓展运行时自身功能时需要使用
func (_RuntimeOption) Inheritor(v util.Face[Runtime]) WithRuntimeOption {
	return func(o *RuntimeOptions) {
		o.Inheritor = v
	}
}

// EnableAutoRun 是否开启自动运行
func (_RuntimeOption) EnableAutoRun(v bool) WithRuntimeOption {
	return func(o *RuntimeOptions) {
		o.EnableAutoRun = v
	}
}

// ProcessQueueCapacity 任务处理流水线大小
func (_RuntimeOption) ProcessQueueCapacity(v int) WithRuntimeOption {
	return func(o *RuntimeOptions) {
		if v <= 0 {
			panic("ProcessQueueCapacity less equal 0 invalid")
		}
		o.ProcessQueueCapacity = v
	}
}

// ProcessQueueTimeout 任务插入流水线超时时长
func (_RuntimeOption) ProcessQueueTimeout(v time.Duration) WithRuntimeOption {
	return func(o *RuntimeOptions) {
		if v <= 0 {
			panic("ProcessQueueTimeout less equal 0 invalid")
		}
		o.ProcessQueueTimeout = v
	}
}

// Frame 帧
func (_RuntimeOption) Frame(v runtime.Frame) WithRuntimeOption {
	return func(o *RuntimeOptions) {
		o.Frame = v
	}
}

// GCInterval GC间隔时长
func (_RuntimeOption) GCInterval(v time.Duration) WithRuntimeOption {
	return func(o *RuntimeOptions) {
		if v <= 0 {
			panic("GCInterval less equal 0 invalid")
		}
		o.GCInterval = v
	}
}
