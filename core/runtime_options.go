package core

import (
	"time"
)

// RuntimeOptions 创建运行时（Runtime）的所有选项
type RuntimeOptions struct {
	Inheritor            Face[Runtime] // 继承者，需要拓展运行时（Runtime）自身功能时需要使用
	EnableAutoRun        bool          // 是否开启自动运行
	EnableAutoRecover    bool          // 是否开启panic时自动恢复
	ProcessQueueCapacity int           // 任务处理流水线大小
	ProcessQueueTimeout  time.Duration // 任务插入流水线超时时长
	Frame                Frame         // 帧
	GCInterval           time.Duration // GC间隔时长
}

// RuntimeOptionSetter 运行时（Runtime）选项设置器
var RuntimeOptionSetter = &_RuntimeOptionSetter{}

// RuntimeOptionSetterFunc 运行时（Runtime）选项设置器函数
type RuntimeOptionSetterFunc func(o *RuntimeOptions)

type _RuntimeOptionSetter struct{}

// Default 默认值
func (*_RuntimeOptionSetter) Default() RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		o.Inheritor = Face[Runtime]{}
		o.EnableAutoRun = false
		o.EnableAutoRecover = false
		o.ProcessQueueCapacity = 128
		o.ProcessQueueTimeout = 5 * time.Second
		o.Frame = nil
		o.GCInterval = 10 * time.Second
	}
}

// Inheritor 继承者，需要拓展运行时（Runtime）自身功能时需要使用
func (*_RuntimeOptionSetter) Inheritor(v Face[Runtime]) RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		o.Inheritor = v
	}
}

// EnableAutoRun 是否开启自动运行
func (*_RuntimeOptionSetter) EnableAutoRun(v bool) RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		o.EnableAutoRun = v
	}
}

// EnableAutoRecover 是否开启panic时自动恢复
func (*_RuntimeOptionSetter) EnableAutoRecover(v bool) RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		o.EnableAutoRecover = v
	}
}

// ProcessQueueCapacity 任务处理流水线大小
func (*_RuntimeOptionSetter) ProcessQueueCapacity(v int) RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		if v <= 0 {
			panic("ProcessQueueCapacity less equal 0 invalid")
		}
		o.ProcessQueueCapacity = v
	}
}

// ProcessQueueTimeout 任务插入流水线超时时长
func (*_RuntimeOptionSetter) ProcessQueueTimeout(v time.Duration) RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		if v <= 0 {
			panic("ProcessQueueTimeout less equal 0 invalid")
		}
		o.ProcessQueueTimeout = v
	}
}

// Frame 帧
func (*_RuntimeOptionSetter) Frame(v Frame) RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		o.Frame = v
	}
}

// GCInterval GC间隔时长
func (*_RuntimeOptionSetter) GCInterval(v time.Duration) RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		if v <= 0 {
			panic("GCInterval less equal 0 invalid")
		}
		o.GCInterval = v
	}
}
