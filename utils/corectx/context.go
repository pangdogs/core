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

// WaitGroup 是用于协调宿主关闭的并发安全屏障。
type WaitGroup interface {
	// Join 调整任务计数；屏障已关闭时拒绝正增量并返回 false，delta 为 0 时 panic。
	Join(delta int64) bool
	// Done 将任务计数减一。
	Done()
	// Wait 阻塞到屏障关闭且全部已加入任务完成。
	Wait()
	// Closed 报告屏障是否已关闭并停止接收新任务。
	Closed() bool
	// Count 返回内部屏障计数；关闭前包含一个宿主持有的基准计数。
	Count() int64
}

// Context 定义 Service 与 Runtime 共享的生命周期上下文能力。
type Context interface {
	iContext
	context.Context
	AsyncScopeProvider

	// ParentContext 返回创建当前上下文时使用的父上下文。
	ParentContext() context.Context
	// AutoRecover 报告框架回调是否自动恢复 panic。
	AutoRecover() bool
	// ReportError 返回自动恢复 panic 后用于非阻塞上报错误的频道。
	ReportError() chan error
	// WaitGroup 返回用于协调关闭的任务屏障。
	WaitGroup() WaitGroup
	// Terminate 发出取消信号，并返回宿主完成清理时兑现的 Signal。
	Terminate() async.Signal
	// Terminated 返回宿主完成清理时兑现的 Signal。
	Terminated() async.Signal
}

type iContext interface {
	init(parentCtx context.Context, autoRecover bool, reportError chan error)
	closeWaitGroup()
	returnTerminated()
}

// ContextBehavior 提供 Context 的通用实现，供 Service 与 Runtime 上下文嵌入。
type ContextBehavior struct {
	context.Context
	parentCtx   context.Context
	autoRecover bool
	reportError chan error
	barrier     generic.Barrier
	asyncScope  *async.Scope
	terminated  async.Completer
}

// ParentContext 返回创建当前上下文时使用的父上下文。
func (ctx *ContextBehavior) ParentContext() context.Context {
	return ctx.parentCtx
}

// AutoRecover 报告框架回调是否自动恢复 panic。
func (ctx *ContextBehavior) AutoRecover() bool {
	return ctx.autoRecover
}

// ReportError 返回自动恢复 panic 后用于非阻塞上报错误的频道。
func (ctx *ContextBehavior) ReportError() chan error {
	return ctx.reportError
}

// WaitGroup 返回用于协调关闭的任务屏障。
func (ctx *ContextBehavior) WaitGroup() WaitGroup {
	return &ctx.barrier
}

// AsyncScope 返回绑定宿主生命周期的后台任务作用域。
func (ctx *ContextBehavior) AsyncScope() *async.Scope {
	return ctx.asyncScope
}

// Terminate 发出取消信号，并返回宿主完成清理时兑现的 Signal。
func (ctx *ContextBehavior) Terminate() async.Signal {
	ctx.asyncScope.Close()
	return ctx.terminated.Signal()
}

// Terminated 返回宿主完成清理时兑现的 Signal。
func (ctx *ContextBehavior) Terminated() async.Signal {
	return ctx.terminated.Signal()
}

func (ctx *ContextBehavior) init(parentCtx context.Context, autoRecover bool, reportError chan error) {
	if parentCtx == nil {
		ctx.parentCtx = context.Background()
	} else {
		ctx.parentCtx = parentCtx
	}
	ctx.autoRecover = autoRecover
	ctx.reportError = reportError
	ctx.asyncScope = async.NewScope(ctx.parentCtx)
	ctx.Context = ctx.asyncScope.Context()
	ctx.terminated, _ = async.NewSignal()
}

func (ctx *ContextBehavior) closeWaitGroup() {
	ctx.barrier.Close()
}

func (ctx *ContextBehavior) returnTerminated() {
	ctx.terminated.Complete()
}
