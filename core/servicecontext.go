package core

import (
	"context"
	"sync"
	"sync/atomic"
)

type ServiceContext interface {
	Context
	_RunnableMark
	EntityMgr
	init(ctx context.Context, opts *ServiceContextOptions)
	getOptions() *ServiceContextOptions
	genUID() uint64
}

func ServiceContextGetOptions(servCtx ServiceContext) ServiceContextOptions {
	return *servCtx.getOptions()
}

func ServiceContextGetInheritor(servCtx ServiceContext) Face[ServiceContext] {
	return servCtx.getOptions().Inheritor
}

func ServiceContextGetInheritorIFace[T any](servCtx ServiceContext) T {
	return Cache2IFace[T](servCtx.getOptions().Inheritor.Cache)
}

func NewServiceContext(ctx context.Context, optFuncs ...NewServiceContextOptionFunc) ServiceContext {
	opts := &ServiceContextOptions{}
	NewServiceContextOption.Default()(opts)

	for i := range optFuncs {
		optFuncs[i](opts)
	}

	if !opts.Inheritor.IsNil() {
		opts.Inheritor.IFace.init(ctx, opts)
		return opts.Inheritor.IFace
	}

	serv := &ServiceContextBehavior{}
	serv.init(ctx, opts)

	return serv.opts.Inheritor.IFace
}

type ServiceContextBehavior struct {
	_ContextBehavior
	_RunnableMarkBehavior
	opts      ServiceContextOptions
	uidGen    uint64
	entityMap sync.Map
}

func (servCtx *ServiceContextBehavior) init(ctx context.Context, opts *ServiceContextOptions) {
	if ctx == nil {
		panic("nil ctx")
	}

	if opts == nil {
		panic("nil opts")
	}

	servCtx.opts = *opts

	if servCtx.opts.Inheritor.IsNil() {
		servCtx.opts.Inheritor = NewFace[ServiceContext](servCtx)
	}

	servCtx._ContextBehavior.init(ctx, servCtx.opts.ReportError)
}

func (servCtx *ServiceContextBehavior) getOptions() *ServiceContextOptions {
	return &servCtx.opts
}

func (servCtx *ServiceContextBehavior) genUID() uint64 {
	return atomic.AddUint64(&servCtx.uidGen, 1)
}
