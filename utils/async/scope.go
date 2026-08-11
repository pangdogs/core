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

package async

import (
	"context"
	"sync"

	"git.golaxy.org/core/utils/exception"
)

// ScopeStats 是异步作用域的瞬时统计快照。
type ScopeStats struct {
	Spawned   int64 // 成功注册过的任务总数。
	Active    int64 // 当前仍未退出的任务数。
	Completed int64 // 正常返回的任务总数。
	Canceled  int64 // 在 Context 已取消状态下退出的任务总数。
	Rejected  int64 // Scope 关闭后拒绝的任务总数。
	Closed    bool  // Scope 是否已经关闭。
}

// Scope 将一组后台任务绑定到同一生命周期。
//
// Scope 组合 Context 取消、关闭后拒绝注册、活动任务计数和完成信号。Close 不会
// 强制终止 goroutine；任务必须观察传入的 Context 才能及时退出。
type Scope struct {
	ctx       context.Context
	cancel    context.CancelFunc
	mu        sync.Mutex
	closed    bool
	active    int64
	stats     ScopeStats
	completer Completer
	done      Signal
	stopWatch func() bool
}

// NewScope 创建由 parent 控制生命周期的异步作用域。nil parent 使用 Background。
func NewScope(parent context.Context) *Scope {
	if parent == nil {
		parent = context.Background()
	}
	ctx, cancel := context.WithCancel(parent)
	completer, done := NewSignal()
	scope := &Scope{
		ctx:       ctx,
		cancel:    cancel,
		completer: completer,
		done:      done,
	}
	if ctx.Err() != nil {
		scope.Close()
		return scope
	}

	// AfterFunc 无需常驻 watcher goroutine 即可跟随上层取消。主动 Close 会先停止回调再
	// 取消 ctx，避免为同一次关闭额外启动重复回调。
	stopWatch := context.AfterFunc(ctx, func() { scope.Close() })
	scope.mu.Lock()
	if scope.closed {
		scope.mu.Unlock()
		stopWatch()
	} else {
		scope.stopWatch = stopWatch
		scope.mu.Unlock()
	}
	return scope
}

// Context 返回传递给所属异步任务的取消上下文。
func (scope *Scope) Context() context.Context {
	if scope == nil {
		return context.Background()
	}
	return scope.ctx
}

// Err 返回 Scope 的取消原因；尚未关闭时返回 nil。
func (scope *Scope) Err() error {
	if scope == nil {
		return ErrScopeClosed
	}
	return scope.ctx.Err()
}

// Done 返回 Scope 关闭且所有已注册任务退出后完成的 Signal。
func (scope *Scope) Done() Signal {
	if scope == nil {
		return CompletedSignal()
	}
	return scope.done
}

// Close 幂等关闭 Scope、取消 Context 并禁止注册新任务。
// 返回值表示本次调用是否首次关闭 Scope。
func (scope *Scope) Close() bool {
	if scope == nil {
		return false
	}

	scope.mu.Lock()
	if scope.closed {
		scope.mu.Unlock()
		return false
	}
	scope.closed = true
	scope.stats.Closed = true
	complete := scope.active == 0
	stopWatch := scope.stopWatch
	scope.stopWatch = nil
	scope.mu.Unlock()

	if stopWatch != nil {
		stopWatch()
	}
	scope.cancel()
	if complete {
		scope.completer.Complete()
	}
	return true
}

// Stats 返回 Scope 的并发安全统计快照。
func (scope *Scope) Stats() ScopeStats {
	if scope == nil {
		return ScopeStats{Closed: true}
	}
	scope.mu.Lock()
	defer scope.mu.Unlock()
	stats := scope.stats
	stats.Active = scope.active
	return stats
}

// Spawn 在 Scope 中启动后台任务并返回其一次性结果 Future。
// panic 会转换为带堆栈的 Result.Error。
func Spawn(scope *Scope, task func(context.Context) Result) Future {
	if task == nil {
		exception.Panicf("%w: %w: async task is nil", ErrAsync, exception.ErrArgs)
	}
	if scope == nil || !scope.begin() {
		return Rejected(ErrScopeClosed)
	}

	promise, future := NewPromise()
	go func() {
		defer scope.end()
		promise.Resolve(safeTaskCall(task, scope.Context()))
	}()
	return future
}

// SpawnVoid 在 Scope 中启动无业务返回值的后台任务，并返回可等待错误的 Future。
func SpawnVoid(scope *Scope, task func(context.Context)) Future {
	if task == nil {
		exception.Panicf("%w: %w: async void task is nil", ErrAsync, exception.ErrArgs)
	}
	return Spawn(scope, func(ctx context.Context) Result {
		task(ctx)
		return NewResult(nil, nil)
	})
}

func (scope *Scope) begin() bool {
	if scope == nil {
		return false
	}
	scope.mu.Lock()
	defer scope.mu.Unlock()
	if scope.closed || scope.ctx.Err() != nil {
		scope.stats.Rejected++
		return false
	}
	scope.active++
	scope.stats.Spawned++
	return true
}

func (scope *Scope) end() {
	scope.mu.Lock()
	scope.active--
	if scope.ctx.Err() != nil {
		scope.stats.Canceled++
	} else {
		scope.stats.Completed++
	}
	complete := scope.closed && scope.active == 0
	scope.mu.Unlock()
	if complete {
		scope.completer.Complete()
	}
}

func safeTaskCall(task func(context.Context) Result, ctx context.Context) (ret Result) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			ret = NewResult(nil, panicToError(panicValue))
		}
	}()
	return task(ctx)
}
