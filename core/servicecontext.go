package core

import (
	"context"
	"github.com/rs/xid"
	"sync"
	"sync/atomic"
)

type ServiceContext interface {
	Context
	_RunnableMark
	EntityMgr
	EntityFactory
	init(ctx context.Context, opts *ServiceContextOptions)
	getOptions() *ServiceContextOptions
	genUID() uint64
	GetPersistID() string
	GetPrototype() string
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

	serv := &_ServiceContextBehavior{}
	serv.init(ctx, opts)

	return serv.opts.Inheritor.IFace
}

type _ServiceContextBehavior struct {
	_ContextBehavior
	_RunnableMarkBehavior
	opts                ServiceContextOptions
	uidGen              uint64
	entityMap           sync.Map
	persistentEntityMap sync.Map
	singletonEntityMap  map[string]Entity
}

func (servCtx *_ServiceContextBehavior) init(ctx context.Context, opts *ServiceContextOptions) {
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

	if servCtx.opts.Params.PersistID == "" {
		servCtx.opts.Params.PersistID = xid.New().String()
	}

	servCtx._ContextBehavior.init(ctx, servCtx.opts.ReportError)
}

func (servCtx *_ServiceContextBehavior) getOptions() *ServiceContextOptions {
	return &servCtx.opts
}

func (servCtx *_ServiceContextBehavior) genUID() uint64 {
	return atomic.AddUint64(&servCtx.uidGen, 1)
}

func (servCtx *_ServiceContextBehavior) GetPersistID() string {
	return servCtx.opts.Params.PersistID
}

func (servCtx *_ServiceContextBehavior) GetPrototype() string {
	return servCtx.opts.Params.Prototype
}
