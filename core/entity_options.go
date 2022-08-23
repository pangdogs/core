package core

import (
	"github.com/pangdogs/galaxy/core/container"
)

var NewEntityOption = &NewEntityOptions{}

type EntityOptions struct {
	Inheritor                  Face[Entity]
	FaceCache                  *container.Cache[FaceAny]
	HookCache                  *container.Cache[Hook]
	EnableFastGetComponent     bool
	EnableFastGetComponentByID bool
	Params                     EntityParams
}

type NewEntityOptionFunc func(o *EntityOptions)

type NewEntityOptions struct{}

func (*NewEntityOptions) Default() NewEntityOptionFunc {
	return func(o *EntityOptions) {
		o.Inheritor = Face[Entity]{}
		o.FaceCache = nil
		o.HookCache = nil
		o.EnableFastGetComponent = false
		o.EnableFastGetComponentByID = false
		o.Params = EntityParams{}
	}
}

func (*NewEntityOptions) Inheritor(v Face[Entity]) NewEntityOptionFunc {
	return func(o *EntityOptions) {
		o.Inheritor = v
	}
}

func (*NewEntityOptions) FaceCache(v *container.Cache[FaceAny]) NewEntityOptionFunc {
	return func(o *EntityOptions) {
		o.FaceCache = v
	}
}

func (*NewEntityOptions) HookCache(v *container.Cache[Hook]) NewEntityOptionFunc {
	return func(o *EntityOptions) {
		o.HookCache = v
	}
}

func (*NewEntityOptions) EnableFastGetComponent(v bool) NewEntityOptionFunc {
	return func(o *EntityOptions) {
		o.EnableFastGetComponent = v
	}
}

func (*NewEntityOptions) EnableFastGetComponentByID(v bool) NewEntityOptionFunc {
	return func(o *EntityOptions) {
		o.EnableFastGetComponentByID = v
	}
}

type EntityParams struct {
	PersistID  string
	Prototype  string
	RuntimeCtx RuntimeContext
	ParentID   uint64
}

func (*NewEntityOptions) Params(v EntityParams) NewEntityOptionFunc {
	return func(o *EntityOptions) {
		o.Params = v
	}
}
