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
	"sync/atomic"

	"git.golaxy.org/core/utils/exception"
)

// ExecutorID 标识进程内的异步结果完成执行器。零值表示结果由外部执行器完成或归属未知。
//
// ExecutorID 只用于运行时阻塞检测，不是持久化 ID，也不应跨进程传输。
type ExecutorID uint64

// GenExecutorID 生成进程内唯一的非零异步执行器 ID。
func GenExecutorID() ExecutorID {
	return ExecutorID(executorIDGen.Add(1))
}

// WaitGuard 允许执行上下文在 Future 进入阻塞等待前实施调度约束。
// Runtime Context 使用该接口阻止 Actor 执行协程等待自身队列产生的结果。
type WaitGuard interface {
	BeforeFutureWait(futureID uint64, completionExecutorID ExecutorID) error
	AfterFutureWait(futureID uint64)
}

// NewPromise 创建一对生产者 Promise 和消费者 Future。
//
// completionExecutorID 可选；仅使用第一个值。它应表示完成该 Future 必须运行的
// Runtime 执行器，外部 I/O、后台 goroutine 或未知来源使用零值。
func NewPromise(completionExecutorID ...ExecutorID) (Promise, Future) {
	state := &futureState{
		id:          futureIDGen.Add(1),
		done:        make(chan struct{}),
		subscribers: make(map[uint64]func(Result)),
	}
	if len(completionExecutorID) > 0 {
		state.completionExecutorID = completionExecutorID[0]
	}
	return Promise{state: state}, Future{state: state}
}

// Resolved 创建已经以 ret 完成的 Future。
func Resolved(ret Result) Future {
	promise, future := NewPromise()
	promise.Resolve(ret)
	return future
}

// Rejected 创建已经以 err 失败的 Future。
func Rejected(err error) Future {
	return Resolved(NewResult(nil, err))
}

// Promise 是一次性异步结果的生产者端。
//
// Promise 可安全地由多个 goroutine 竞争完成；只有第一次 Resolve 成功。
type Promise struct {
	state *futureState
}

// IsNil 报告 Promise 是否为零值。
func (promise Promise) IsNil() bool {
	return promise.state == nil
}

// Future 返回与 Promise 共享状态的只读 Future。
func (promise Promise) Future() Future {
	if promise.IsNil() {
		exception.Panicf("%w: %w: promise is nil, cannot get future", ErrAsync, exception.ErrArgs)
	}
	return Future{state: promise.state}
}

// Resolve 以 ret 完成 Future。首次完成返回 true，后续调用返回 false。
//
// 回调在完成者 goroutine 中、状态锁之外执行，因此回调必须快速返回；需要修改
// Runtime 状态时应通过 core.ContinueOn 投递续体。
func (promise Promise) Resolve(ret Result) bool {
	if promise.IsNil() {
		exception.Panicf("%w: %w: promise is nil, cannot resolve", ErrAsync, exception.ErrArgs)
	}

	state := promise.state
	state.mu.Lock()
	if state.completed {
		state.mu.Unlock()
		return false
	}

	state.completed = true
	state.result = ret
	callbacks := make([]func(Result), 0, len(state.subscribers))
	for _, id := range state.subscriberOrder {
		if callback, ok := state.subscribers[id]; ok {
			callbacks = append(callbacks, callback)
		}
	}
	state.subscribers = nil
	state.subscriberOrder = nil
	close(state.done)
	state.mu.Unlock()

	for _, callback := range callbacks {
		callback(ret)
	}
	return true
}

// Future 是一次性异步结果的只读消费者视图。
//
// 完成结果会保存在共享状态中，Wait、TryGet 和 OnComplete 均具有重放语义；多个
// 消费者不会竞争同一个结果。
type Future struct {
	state *futureState
}

// IsNil 报告 Future 是否为零值。
func (future Future) IsNil() bool {
	return future.state == nil
}

// ID 返回 Future 的进程内诊断 ID；零值 Future 返回 0。
func (future Future) ID() uint64 {
	if future.IsNil() {
		return 0
	}
	return future.state.id
}

// CompletionExecutorID 返回完成 Future 所依赖的执行器 ID。
func (future Future) CompletionExecutorID() ExecutorID {
	if future.IsNil() {
		return 0
	}
	return future.state.completionExecutorID
}

// Done 返回 Future 完成时关闭的共享频道；零值 Future 会导致 panic。
func (future Future) Done() <-chan struct{} {
	if future.IsNil() {
		exception.Panicf("%w: %w: future is nil, cannot get done channel", ErrAsync, exception.ErrArgs)
	}
	return future.state.done
}

// TryGet 无阻塞地读取完成结果。Future 尚未完成时 ok 为 false。
func (future Future) TryGet() (ret Result, ok bool) {
	if future.IsNil() {
		exception.Panicf("%w: %w: future is nil, cannot get result", ErrAsync, exception.ErrArgs)
	}

	future.state.mu.Lock()
	defer future.state.mu.Unlock()
	if !future.state.completed {
		return Result{}, false
	}
	return future.state.result, true
}

// Wait 等待 Future 完成或 ctx 取消。
//
// nil ctx 按 context.Background 处理。结果具有重放语义。若 ctx 实现 WaitGuard，
// Future 会在真正阻塞前调用它，以阻止 Runtime 自等待等非法调度。
func (future Future) Wait(ctx context.Context) Result {
	if future.IsNil() {
		exception.Panicf("%w: %w: future is nil, cannot wait", ErrAsync, exception.ErrArgs)
	}
	if ret, ok := future.TryGet(); ok {
		return ret
	}
	if ctx == nil {
		ctx = context.Background()
	}

	var guard WaitGuard
	if candidate, ok := ctx.(WaitGuard); ok {
		guard = candidate
		if ret, ready := future.TryGet(); ready {
			return ret
		}
		if err := guard.BeforeFutureWait(future.ID(), future.CompletionExecutorID()); err != nil {
			guard.AfterFutureWait(future.ID())
			return NewResult(nil, err)
		}
		defer guard.AfterFutureWait(future.ID())
	}

	select {
	case <-future.state.done:
		ret, _ := future.TryGet()
		return ret
	case <-ctx.Done():
		return NewResult(nil, ctx.Err())
	}
}

// OnComplete 订阅 Future 完成。Future 已完成时 callback 会在调用者 goroutine 中立即执行。
// callback 为 nil 时会导致 panic。
func (future Future) OnComplete(callback func(Result)) Subscription {
	if future.IsNil() {
		exception.Panicf("%w: %w: future is nil, cannot subscribe", ErrAsync, exception.ErrArgs)
	}
	if callback == nil {
		exception.Panicf("%w: %w: future completion callback is nil", ErrAsync, exception.ErrArgs)
	}

	state := future.state
	state.mu.Lock()
	if state.completed {
		ret := state.result
		state.mu.Unlock()
		callback(ret)
		return Subscription{}
	}

	state.nextSubscriberID++
	id := state.nextSubscriberID
	state.subscribers[id] = callback
	state.subscriberOrder = append(state.subscriberOrder, id)
	state.mu.Unlock()
	return Subscription{state: state, id: id}
}

// Context 返回由 ctx 派生、并在 Future 完成时取消的上下文。
func (future Future) Context(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	derived, cancel := context.WithCancel(ctx)
	subscription := future.OnComplete(func(Result) { cancel() })
	context.AfterFunc(derived, func() { subscription.Cancel() })
	return derived
}

// Subscription 表示一个可取消的 Future 完成订阅。
type Subscription struct {
	state *futureState
	id    uint64
}

// Cancel 取消尚未执行的订阅。成功移除返回 true；已完成、已取消或零值返回 false。
func (subscription Subscription) Cancel() bool {
	if subscription.state == nil || subscription.id == 0 {
		return false
	}

	subscription.state.mu.Lock()
	defer subscription.state.mu.Unlock()
	if subscription.state.completed || subscription.state.subscribers == nil {
		return false
	}
	if _, ok := subscription.state.subscribers[subscription.id]; !ok {
		return false
	}
	delete(subscription.state.subscribers, subscription.id)
	if len(subscription.state.subscriberOrder) > 64 && len(subscription.state.subscribers)*2 < len(subscription.state.subscriberOrder) {
		order := make([]uint64, 0, len(subscription.state.subscribers))
		for _, id := range subscription.state.subscriberOrder {
			if _, ok := subscription.state.subscribers[id]; ok {
				order = append(order, id)
			}
		}
		subscription.state.subscriberOrder = order
	}
	return true
}

var (
	executorIDGen atomic.Uint64
	futureIDGen   atomic.Uint64
)

type futureState struct {
	mu                   sync.Mutex
	id                   uint64
	completionExecutorID ExecutorID
	completed            bool
	result               Result
	done                 chan struct{}
	nextSubscriberID     uint64
	subscribers          map[uint64]func(Result)
	subscriberOrder      []uint64
}
