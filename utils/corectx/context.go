/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

package corectx

import (
	"context"

	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/generic"
)

// WaitGroup 等待组
type WaitGroup interface {
	// Join 添加任务
	Join(delta int64) bool
	// Done 任务完成
	Done()
	// Wait 等待所有任务完成
	Wait()
	// IsClosed 是否已关闭
	IsClosed() bool
	// Pending 等待任务数量
	Pending() int64
}

// Context 上下文
type Context interface {
	iContext
	context.Context

	// GetParentContext 获取父上下文
	GetParentContext() context.Context
	// GetAutoRecover panic时是否自动恢复
	GetAutoRecover() bool
	// GetReportError 在开启panic时自动恢复时，将会恢复并将错误写入此error channel
	GetReportError() chan error
	// GetWaitGroup 获取等待组
	GetWaitGroup() WaitGroup
	// Terminate 停止
	Terminate() async.AsyncRet
	// Terminated 已停止
	Terminated() async.AsyncRet
}

type iContext interface {
	init(parentCtx context.Context, autoRecover bool, reportError chan error)
	closeWaitGroup()
	returnTerminated()
}

// ContextBehavior 上下文行为
type ContextBehavior struct {
	context.Context
	parentCtx   context.Context
	autoRecover bool
	reportError chan error
	barrier     generic.Barrier
	terminate   context.CancelFunc
	terminated  chan async.Ret
}

// GetParentContext 获取父上下文
func (ctx *ContextBehavior) GetParentContext() context.Context {
	return ctx.parentCtx
}

// GetAutoRecover panic时是否自动恢复
func (ctx *ContextBehavior) GetAutoRecover() bool {
	return ctx.autoRecover
}

// GetReportError 在开启panic时自动恢复时，将会恢复并将错误写入此error channel
func (ctx *ContextBehavior) GetReportError() chan error {
	return ctx.reportError
}

// GetWaitGroup 获取等待组
func (ctx *ContextBehavior) GetWaitGroup() WaitGroup {
	return &ctx.barrier
}

// Terminate 停止
func (ctx *ContextBehavior) Terminate() async.AsyncRet {
	ctx.terminate()
	return ctx.terminated
}

// Terminated 已停止
func (ctx *ContextBehavior) Terminated() async.AsyncRet {
	return ctx.terminated
}

func (ctx *ContextBehavior) init(parentCtx context.Context, autoRecover bool, reportError chan error) {
	if parentCtx == nil {
		ctx.parentCtx = context.Background()
	} else {
		ctx.parentCtx = parentCtx
	}
	ctx.autoRecover = autoRecover
	ctx.reportError = reportError
	ctx.Context, ctx.terminate = context.WithCancel(ctx.parentCtx)
	ctx.terminated = async.NewAsyncRet()
}

func (ctx *ContextBehavior) closeWaitGroup() {
	ctx.barrier.Close()
}

func (ctx *ContextBehavior) returnTerminated() {
	async.Return(ctx.terminated, async.VoidRet)
}
