package service

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/pangdogs/galaxy/internal"
	"github.com/pangdogs/galaxy/util"
	"sync/atomic"
)

// Context 服务上下文
type Context interface {
	internal.Context
	internal.RunningMark
	_SafeCall

	init(opts *ContextOptions)

	getOptions() *ContextOptions

	// GetPrototype 获取服务原型
	GetPrototype() string

	// GenUID 生成运行时唯一ID，向负方向增长，非全局唯一，重启服务后也会生成相同ID，不能使用此ID持久化，性能较好，值小于0
	GenUID() int64

	// GenPersistID 生成持久化ID，向正方向增长，全局唯一，必须使用此ID持久化，使用snowflake算法，性能较差，默认情况下单个服务每毫秒仅能生成4096个，值大于0
	GenPersistID() int64

	// GetEntityMgr 获取实体管理器
	GetEntityMgr() IEntityMgr
}

// NewContext 创建服务上下文
func NewContext(optSetter ...ContextOptionSetter) Context {
	opts := ContextOptions{}
	ContextOption.Default()(&opts)

	for i := range optSetter {
		optSetter[i](&opts)
	}

	return UnsafeNewContext(opts)
}

func UnsafeNewContext(opts ContextOptions) Context {
	if !opts.Inheritor.IsNil() {
		opts.Inheritor.Iface.init(&opts)
		return opts.Inheritor.Iface
	}

	ctx := &ContextBehavior{}
	ctx.init(&opts)

	return ctx.opts.Inheritor.Iface
}

// ContextBehavior 服务上下文行为，在需要拓展服务上下文能力时，匿名嵌入至服务上下文结构体中
type ContextBehavior struct {
	internal.ContextBehavior
	internal.RunningMarkBehavior
	opts          ContextOptions
	uidGenerator  int64
	snowflakeNode *snowflake.Node
	entityMgr     _EntityMgr
}

func (ctx *ContextBehavior) init(opts *ContextOptions) {
	if opts == nil {
		panic("nil opts")
	}

	ctx.opts = *opts

	if ctx.opts.Inheritor.IsNil() {
		ctx.opts.Inheritor = util.NewFace[Context](ctx)
	}

	if ctx.opts.Context == nil {
		ctx.opts.Context = context.Background()
	}

	ctx.ContextBehavior.Init(ctx.opts.Context, ctx.opts.AutoRecover, ctx.opts.ReportError)

	snowflakeNode, err := snowflake.NewNode(ctx.opts.NodeID)
	if err != nil {
		panic(err)
	}
	ctx.snowflakeNode = snowflakeNode

	ctx.entityMgr.Init(ctx.opts.Inheritor.Iface)
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

// GetEntityMgr 获取实体管理器
func (ctx *ContextBehavior) GetEntityMgr() IEntityMgr {
	return &ctx.entityMgr
}
