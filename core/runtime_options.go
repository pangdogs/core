package core

import "time"

// NewRuntimeOption ...
var NewRuntimeOption = &_NewRuntimeOptions{}

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

type _NewRuntimeOptionFunc func(o *RuntimeOptions)

type _NewRuntimeOptions struct{}

// Default ...
func (*_NewRuntimeOptions) Default() _NewRuntimeOptionFunc {
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
func (*_NewRuntimeOptions) Inheritor(v Face[Runtime]) _NewRuntimeOptionFunc {
	return func(o *RuntimeOptions) {
		o.Inheritor = v
	}
}

// EnableAutoRun ...
func (*_NewRuntimeOptions) EnableAutoRun(v bool) _NewRuntimeOptionFunc {
	return func(o *RuntimeOptions) {
		o.EnableAutoRun = v
	}
}

// EnableAutoRecover ...
func (*_NewRuntimeOptions) EnableAutoRecover(v bool) _NewRuntimeOptionFunc {
	return func(o *RuntimeOptions) {
		o.EnableAutoRecover = v
	}
}

// EnableSortCompStartOrder ...
func (*_NewRuntimeOptions) EnableSortCompStartOrder(v bool) _NewRuntimeOptionFunc {
	return func(o *RuntimeOptions) {
		o.EnableSortCompStartOrder = v
	}
}

// ProcessQueueCapacity ...
func (*_NewRuntimeOptions) ProcessQueueCapacity(v int) _NewRuntimeOptionFunc {
	return func(o *RuntimeOptions) {
		if v <= 0 {
			panic("ProcessQueueCapacity less equal 0 invalid")
		}
		o.ProcessQueueCapacity = v
	}
}

// ProcessQueueTimeout ...
func (*_NewRuntimeOptions) ProcessQueueTimeout(v time.Duration) _NewRuntimeOptionFunc {
	return func(o *RuntimeOptions) {
		if v <= 0 {
			panic("ProcessQueueTimeout less equal 0 invalid")
		}
		o.ProcessQueueTimeout = v
	}
}

// Frame ...
func (*_NewRuntimeOptions) Frame(v Frame) _NewRuntimeOptionFunc {
	return func(o *RuntimeOptions) {
		o.Frame = v
	}
}

// GCInterval ...
func (*_NewRuntimeOptions) GCInterval(v time.Duration) _NewRuntimeOptionFunc {
	return func(o *RuntimeOptions) {
		if v <= 0 {
			panic("GCInterval less equal 0 invalid")
		}
		o.GCInterval = v
	}
}
