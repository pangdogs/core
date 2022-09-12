package core

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"sync"
	"sync/atomic"
)

// ServiceContext 服务上下文，线程安全
type ServiceContext interface {
	_Context
	_RunnableMark
	_ServiceContextEntityMgr

	init(opts *ServiceContextOptions)

	getOptions() *ServiceContextOptions

	// GenUID 生成运行时唯一ID，向负方向增长，非全局唯一，重启服务后也会生成相同ID，不能使用此ID持久化，性能较好，线程安全
	GenUID() int64

	// GenPersistID 生成持久化ID，向正方向增长，全局唯一，必须使用此ID持久化，使用snowflake算法，性能较差，默认情况下单个服务每毫秒仅能生成4096个，线程安全
	GenPersistID() int64

	// GetPrototype 获取服务原型
	GetPrototype() string
}

// ServiceContextGetOptions 获取服务上下文创建选项，线程安全
func ServiceContextGetOptions(servCtx ServiceContext) ServiceContextOptions {
	return *servCtx.getOptions()
}

// NewServiceContext 创建服务上下文，线程安全
func NewServiceContext(optSetterFuncs ...ServiceContextOptionSetterFunc) ServiceContext {
	opts := ServiceContextOptions{}
	ServiceContextOptionSetter.Default()(&opts)

	for i := range optSetterFuncs {
		optSetterFuncs[i](&opts)
	}

	return NewServiceContextWithOpts(opts)
}

// NewServiceContextWithOpts 创建服务上下文并传入参数，线程安全
func NewServiceContextWithOpts(opts ServiceContextOptions) ServiceContext {
	if !opts.Inheritor.IsNil() {
		opts.Inheritor.Iface.init(&opts)
		return opts.Inheritor.Iface
	}

	serv := &_ServiceContextBehavior{}
	serv.init(&opts)

	return serv.opts.Inheritor.Iface
}

type _ServiceContextBehavior struct {
	_ContextBehavior
	_RunnableMarkBehavior
	opts           ServiceContextOptions
	uidGenerator   int64
	snowflakeNode  *snowflake.Node
	entityMap      map[int64]Entity
	entityMapMutex sync.RWMutex
}

func (servCtx *_ServiceContextBehavior) init(opts *ServiceContextOptions) {
	if opts == nil {
		panic("nil opts")
	}

	servCtx.opts = *opts

	if servCtx.opts.Inheritor.IsNil() {
		servCtx.opts.Inheritor = NewFace[ServiceContext](servCtx)
	}

	if servCtx.opts.ParentContext == nil {
		servCtx.opts.ParentContext = context.Background()
	}

	snowflakeNode, err := snowflake.NewNode(servCtx.opts.NodeID)
	if err != nil {
		panic(err)
	}
	servCtx.snowflakeNode = snowflakeNode

	servCtx._ContextBehavior.init(servCtx.opts.ParentContext, servCtx.opts.ReportError)
}

func (servCtx *_ServiceContextBehavior) getOptions() *ServiceContextOptions {
	return &servCtx.opts
}

// GenUID 生成运行时唯一ID，向负方向增长，非全局唯一，重启服务后也会生成相同ID，不能使用此ID持久化，性能较好，线程安全
func (servCtx *_ServiceContextBehavior) GenUID() int64 {
	return atomic.AddInt64(&servCtx.uidGenerator, -1)
}

// GenPersistID 生成持久化ID，向正方向增长，全局唯一，必须使用此ID持久化，性能较差，单个服务每毫秒仅能生成4096个，线程安全
func (servCtx *_ServiceContextBehavior) GenPersistID() int64 {
	return int64(servCtx.snowflakeNode.Generate())
}

// GetPrototype 获取服务原型
func (servCtx *_ServiceContextBehavior) GetPrototype() string {
	return servCtx.opts.Prototype
}
