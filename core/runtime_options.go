package core

import "time"

var NewRuntimeOption = &NewRuntimeOptions{}

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

type NewRuntimeOptionFunc func(o *RuntimeOptions)

type NewRuntimeOptions struct{}

func (*NewRuntimeOptions) Default() NewRuntimeOptionFunc {
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

func (*NewRuntimeOptions) Inheritor(v Face[Runtime]) NewRuntimeOptionFunc {
	return func(o *RuntimeOptions) {
		o.Inheritor = v
	}
}

func (*NewRuntimeOptions) EnableAutoRun(v bool) NewRuntimeOptionFunc {
	return func(o *RuntimeOptions) {
		o.EnableAutoRun = v
	}
}

func (*NewRuntimeOptions) EnableAutoRecover(v bool) NewRuntimeOptionFunc {
	return func(o *RuntimeOptions) {
		o.EnableAutoRecover = v
	}
}

func (*NewRuntimeOptions) EnableSortCompStartOrder(v bool) NewRuntimeOptionFunc {
	return func(o *RuntimeOptions) {
		o.EnableSortCompStartOrder = v
	}
}

func (*NewRuntimeOptions) ProcessQueueCapacity(v int) NewRuntimeOptionFunc {
	return func(o *RuntimeOptions) {
		if v <= 0 {
			panic("ProcessQueueCapacity less equal 0 invalid")
		}
		o.ProcessQueueCapacity = v
	}
}

func (*NewRuntimeOptions) ProcessQueueTimeout(v time.Duration) NewRuntimeOptionFunc {
	return func(o *RuntimeOptions) {
		if v <= 0 {
			panic("ProcessQueueTimeout less equal 0 invalid")
		}
		o.ProcessQueueTimeout = v
	}
}

func (*NewRuntimeOptions) Frame(v Frame) NewRuntimeOptionFunc {
	return func(o *RuntimeOptions) {
		o.Frame = v
	}
}

func (*NewRuntimeOptions) GCInterval(v time.Duration) NewRuntimeOptionFunc {
	return func(o *RuntimeOptions) {
		if v <= 0 {
			panic("GCInterval less equal 0 invalid")
		}
		o.GCInterval = v
	}
}
