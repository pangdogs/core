package service

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/pangdogs/galaxy/ec"
	"github.com/pangdogs/galaxy/internal"
	"github.com/pangdogs/galaxy/util"
	"sync"
	"sync/atomic"
)

// Context 服务上下文
type Context interface {
	internal.Context
	internal.RunningMark
	_ServiceContextEntityMgr

	init(opts *ContextOptions)

	getOptions() *ContextOptions

	// GetPrototype 获取服务原型
	GetPrototype() string

	// GenUID 生成运行时唯一ID，向负方向增长，非全局唯一，重启服务后也会生成相同ID，不能使用此ID持久化，性能较好，值小于0
	GenUID() int64

	// GenPersistID 生成持久化ID，向正方向增长，全局唯一，必须使用此ID持久化，使用snowflake算法，性能较差，默认情况下单个服务每毫秒仅能生成4096个，值大于0
	GenPersistID() int64
}

// NewContext 创建服务上下文
func NewContext(optSetter ...ContextOptionSetter) Context {
	opts := ContextOptions{}
	ContextOption.Default()(&opts)

	for i := range optSetter {
		optSetter[i](&opts)
	}

	if !opts.Inheritor.IsNil() {
		opts.Inheritor.Iface.init(&opts)
		return opts.Inheritor.Iface
	}

	serv := &ContextBehavior{}
	serv.init(&opts)

	return serv.opts.Inheritor.Iface
}

type ContextBehavior struct {
	internal.ContextBehavior
	internal.RunningMarkBehavior
	opts           ContextOptions
	uidGenerator   int64
	snowflakeNode  *snowflake.Node
	entityMap      map[int64]ec.Entity
	entityMapMutex sync.RWMutex
}

func (ctx *ContextBehavior) init(opts *ContextOptions) {
	if opts == nil {
		panic("nil opts")
	}

	ctx.opts = *opts

	if ctx.opts.Inheritor.IsNil() {
		ctx.opts.Inheritor = util.NewFace[Context](ctx)
	}

	if ctx.opts.ParentContext == nil {
		ctx.opts.ParentContext = context.Background()
	}

	ctx.ContextBehavior.Init(ctx.opts.ParentContext, ctx.opts.ctx.opts.ReportError)

	snowflakeNode, err := snowflake.NewNode(ctx.opts.NodeID)
	if err != nil {
		panic(err)
	}
	ctx.snowflakeNode = snowflakeNode
}

func (ctx *ContextBehavior) getOptions() *ContextOptions {
	return &ctx.opts
}

// GetPrototype 获取服务原型
func (ctx *ContextBehavior) GetPrototype() string {
	return ctx.opts.Prototype
}

// GenUID 生成运行时唯一ID，向负方向增长，非全局唯一，重启服务后也会生成相同ID，不能使用此ID持久化，性能较好，值小于0
func (ctx *ContextBehavior) GenUID() int64 {
	return atomic.AddInt64(&ctx.uidGenerator, -1)
}

// GenPersistID 生成持久化ID，向正方向增长，全局唯一，必须使用此ID持久化，性能较差，单个服务每毫秒仅能生成4096个，值大于0
func (ctx *ContextBehavior) GenPersistID() int64 {
	return int64(ctx.snowflakeNode.Generate())
}
