package core

import (
	"time"
)

// RuntimeOptions ...
type RuntimeOptions struct {
	Inheritor                Face[Runtime]
	EnableAutoRun            bool
	EnableAutoRecover        bool
	EnableSortCompStartOrder bool
	ProcessQueueCapacity     int
	ProcessQueueTimeout      time.Duration
	Frame                    Frame
	GCInterval               time.Duration
}

// RuntimeOptionSetter ...
var RuntimeOptionSetter = &_RuntimeOptionSetter{}

type _RuntimeOptionSetterFunc func(o *RuntimeOptions)

type _RuntimeOptionSetter struct{}

// Default ...
func (*_RuntimeOptionSetter) Default() _RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		o.Inheritor = Face[Runtime]{}
		o.EnableAutoRun = false
		o.EnableAutoRecover = false
		o.EnableSortCompStartOrder = false
		o.ProcessQueueCapacity = 128
		o.ProcessQueueTimeout = 5 * time.Second
		o.Frame = nil
		o.GCInterval = 10 * time.Second
	}
}

// Inheritor ...
func (*_RuntimeOptionSetter) Inheritor(v Face[Runtime]) _RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		o.Inheritor = v
	}
}

// EnableAutoRun ...
func (*_RuntimeOptionSetter) EnableAutoRun(v bool) _RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		o.EnableAutoRun = v
	}
}

// EnableAutoRecover ...
func (*_RuntimeOptionSetter) EnableAutoRecover(v bool) _RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		o.EnableAutoRecover = v
	}
}

// EnableSortCompStartOrder ...
func (*_RuntimeOptionSetter) EnableSortCompStartOrder(v bool) _RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		o.EnableSortCompStartOrder = v
	}
}

// ProcessQueueCapacity ...
func (*_RuntimeOptionSetter) ProcessQueueCapacity(v int) _RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		if v <= 0 {
			panic("ProcessQueueCapacity less equal 0 invalid")
		}
		o.ProcessQueueCapacity = v
	}
}

// ProcessQueueTimeout ...
func (*_RuntimeOptionSetter) ProcessQueueTimeout(v time.Duration) _RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		if v <= 0 {
			panic("ProcessQueueTimeout less equal 0 invalid")
		}
		o.ProcessQueueTimeout = v
	}
}

// Frame ...
func (*_RuntimeOptionSetter) Frame(v Frame) _RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		o.Frame = v
	}
}

// GCInterval ...
func (*_RuntimeOptionSetter) GCInterval(v time.Duration) _RuntimeOptionSetterFunc {
	return func(o *RuntimeOptions) {
		if v <= 0 {
			panic("GCInterval less equal 0 invalid")
		}
		o.GCInterval = v
	}
}
