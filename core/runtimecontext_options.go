package core

import (
	"github.com/pangdogs/galaxy/core/container"
)

// NewRuntimeContextOption ...
var NewRuntimeContextOption = &NewRuntimeContextOptions{}

// RuntimeContextOptions ...
type RuntimeContextOptions struct {
	Inheritor   Face[RuntimeContext]
	ReportError chan error
	StartedCallback,
	StoppedCallback func(runtime Runtime)
	FaceCache *container.Cache[FaceAny]
	HookCache *container.Cache[Hook]
}

// NewRuntimeContextOptionFunc ...
type NewRuntimeContextOptionFunc func(o *RuntimeContextOptions)

// NewRuntimeContextOptions ...
type NewRuntimeContextOptions struct{}

// Default ...
func (*NewRuntimeContextOptions) Default() NewRuntimeContextOptionFunc {
	return func(o *RuntimeContextOptions) {
		o.Inheritor = Face[RuntimeContext]{}
		o.StartedCallback = nil
		o.StoppedCallback = nil
		o.FaceCache = nil
		o.HookCache = nil
	}
}

// Inheritor ...
func (*NewRuntimeContextOptions) Inheritor(v Face[RuntimeContext]) NewRuntimeContextOptionFunc {
	return func(o *RuntimeContextOptions) {
		o.Inheritor = v
	}
}

// ReportError ...
func (*NewRuntimeContextOptions) ReportError(v chan error) NewRuntimeContextOptionFunc {
	return func(o *RuntimeContextOptions) {
		o.ReportError = v
	}
}

// StartFunc ...
func (*NewRuntimeContextOptions) StartFunc(v func(rt Runtime)) NewRuntimeContextOptionFunc {
	return func(o *RuntimeContextOptions) {
		o.StartedCallback = v
	}
}

// StopFunc ...
func (*NewRuntimeContextOptions) StopFunc(v func(rt Runtime)) NewRuntimeContextOptionFunc {
	return func(o *RuntimeContextOptions) {
		o.StoppedCallback = v
	}
}

// FaceCache ...
func (*NewRuntimeContextOptions) FaceCache(v *container.Cache[FaceAny]) NewRuntimeContextOptionFunc {
	return func(o *RuntimeContextOptions) {
		o.FaceCache = v
	}
}

// HookCache ...
func (*NewRuntimeContextOptions) HookCache(v *container.Cache[Hook]) NewRuntimeContextOptionFunc {
	return func(o *RuntimeContextOptions) {
		o.HookCache = v
	}
}
