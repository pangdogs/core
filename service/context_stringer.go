package service

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
)

func (ctx *ContextBehavior) String() string {
	return ContextStringer.String(ctx.opts.CompositeFace.Iface)
}

var ContextStringer internal.Stringer[Context] = DefaultContextStringer{}

type DefaultContextStringer struct{}

func (DefaultContextStringer) String(ctx Context) string {
	return fmt.Sprintf("{Id:%s Name:%s}", ctx.GetId(), ctx.GetName())
}
