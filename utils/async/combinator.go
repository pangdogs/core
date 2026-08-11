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
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"git.golaxy.org/core/utils/exception"
)

// Pair 是 Zip2 成功时保存在 Result.Value 中的二元结果。
type Pair struct {
	First  any
	Second any
}

// Race 返回第一个完成的 Future 结果。候选为空时返回 ErrNoCandidates。
//
// Race 只解除失败者订阅，不取消可能被其他调用者共享的生产任务。
func Race(futures ...Future) Future {
	futures = compactFutures(futures)
	if len(futures) <= 0 {
		return Rejected(ErrNoCandidates)
	}

	promise, result := NewPromise(commonCompletionExecutorID(futures))
	group := &subscriptionGroup{}
	for _, future := range futures {
		future := future
		group.add(future.OnComplete(func(ret Result) {
			if promise.Resolve(ret) {
				group.cancel()
			}
		}))
	}
	return result
}

// FirstSuccess 返回第一个成功结果；所有候选均失败时返回包装 ErrNoFutureSucceeded 的错误。
func FirstSuccess(futures ...Future) Future {
	futures = compactFutures(futures)
	if len(futures) <= 0 {
		return Rejected(ErrNoCandidates)
	}

	promise, result := NewPromise(commonCompletionExecutorID(futures))
	group := &subscriptionGroup{}
	errs := make([]error, len(futures))
	var remaining atomic.Int64
	remaining.Store(int64(len(futures)))

	for i, future := range futures {
		i, future := i, future
		group.add(future.OnComplete(func(ret Result) {
			if ret.OK() {
				if promise.Resolve(ret) {
					group.cancel()
				}
				return
			}

			errs[i] = ret.Error
			if remaining.Add(-1) != 0 {
				return
			}
			joined := errors.Join(errs...)
			promise.Resolve(NewResult(nil, fmt.Errorf("%w: %w", ErrNoFutureSucceeded, joined)))
			group.cancel()
		}))
	}
	return result
}

// All 按输入顺序收集所有值；任一 Future 失败时立即失败。
// 空输入成功返回空 []any。
func All(futures ...Future) Future {
	futures = compactFutures(futures)
	if len(futures) <= 0 {
		return Resolved(NewResult([]any{}, nil))
	}

	promise, result := NewPromise(commonCompletionExecutorID(futures))
	group := &subscriptionGroup{}
	values := make([]any, len(futures))
	var remaining atomic.Int64
	remaining.Store(int64(len(futures)))

	for i, future := range futures {
		i, future := i, future
		group.add(future.OnComplete(func(ret Result) {
			if !ret.OK() {
				if promise.Resolve(NewResult(nil, ret.Error)) {
					group.cancel()
				}
				return
			}
			values[i] = ret.Value
			if remaining.Add(-1) == 0 && promise.Resolve(NewResult(values, nil)) {
				group.cancel()
			}
		}))
	}
	return result
}

// AllSettled 按输入顺序返回全部 Result，无论各项成功或失败。
// 空输入成功返回空 []Result。
func AllSettled(futures ...Future) Future {
	futures = compactFutures(futures)
	if len(futures) <= 0 {
		return Resolved(NewResult([]Result{}, nil))
	}

	promise, result := NewPromise(commonCompletionExecutorID(futures))
	results := make([]Result, len(futures))
	var remaining atomic.Int64
	remaining.Store(int64(len(futures)))

	for i, future := range futures {
		i, future := i, future
		future.OnComplete(func(ret Result) {
			results[i] = ret
			if remaining.Add(-1) == 0 {
				promise.Resolve(NewResult(results, nil))
			}
		})
	}
	return result
}

// Zip2 等待两个 Future 成功，并按参数顺序返回 Pair；任一失败时立即失败。
func Zip2(first, second Future) Future {
	if first.IsNil() || second.IsNil() {
		return Rejected(ErrNoCandidates)
	}
	return Map(All(first, second), func(ret Result) Result {
		if !ret.OK() {
			return ret
		}
		values, ok := ret.Value.([]any)
		if !ok || len(values) != 2 {
			return NewResult(nil, fmt.Errorf("%w: invalid Zip2 result", ErrAsync))
		}
		return NewResult(Pair{First: values[0], Second: values[1]}, nil)
	})
}

// Map 在源 Future 完成的 goroutine 中转换结果。fn 必须快速返回。
func Map(future Future, fn func(Result) Result) Future {
	if future.IsNil() {
		return Rejected(ErrNoCandidates)
	}
	if fn == nil {
		exception.Panicf("%w: %w: future map function is nil", ErrAsync, exception.ErrArgs)
	}
	promise, mapped := NewPromise(future.CompletionExecutorID())
	future.OnComplete(func(ret Result) {
		promise.Resolve(safeResultCall(fn, ret))
	})
	return mapped
}

// FlatMap 使用源结果选择下一个 Future，并将其结果展平到返回 Future。
func FlatMap(future Future, fn func(Result) Future) Future {
	if future.IsNil() {
		return Rejected(ErrNoCandidates)
	}
	if fn == nil {
		exception.Panicf("%w: %w: future flat-map function is nil", ErrAsync, exception.ErrArgs)
	}
	promise, flattened := NewPromise()
	future.OnComplete(func(ret Result) {
		next, panicErr := safeFutureCall(fn, ret)
		if panicErr != nil {
			promise.Resolve(NewResult(nil, panicErr))
			return
		}
		if next.IsNil() {
			promise.Resolve(NewResult(nil, ErrNoCandidates))
			return
		}
		next.OnComplete(func(ret Result) { promise.Resolve(ret) })
	})
	return flattened
}

// Timeout 返回在源 Future、ctx 取消或 duration 到期三者中最先完成的结果。
// 它只解除源 Future 的订阅，不取消源任务。
func Timeout(ctx context.Context, future Future, duration time.Duration) Future {
	if future.IsNil() {
		return Rejected(ErrNoCandidates)
	}
	if ret, ok := future.TryGet(); ok {
		return Resolved(ret)
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if duration <= 0 {
		return Rejected(ErrFutureTimeout)
	}

	promise, result := NewPromise()
	subscription := future.OnComplete(func(ret Result) { promise.Resolve(ret) })
	timer := time.AfterFunc(duration, func() {
		promise.Resolve(NewResult(nil, ErrFutureTimeout))
	})
	stopContext := context.AfterFunc(ctx, func() {
		promise.Resolve(NewResult(nil, ctx.Err()))
	})
	result.OnComplete(func(Result) {
		subscription.Cancel()
		timer.Stop()
		stopContext()
	})
	return result
}

func compactFutures(futures []Future) []Future {
	compacted := make([]Future, 0, len(futures))
	for _, future := range futures {
		if !future.IsNil() {
			compacted = append(compacted, future)
		}
	}
	return compacted
}

func commonCompletionExecutorID(futures []Future) ExecutorID {
	if len(futures) <= 0 {
		return 0
	}
	executorID := futures[0].CompletionExecutorID()
	if executorID == 0 {
		return 0
	}
	for _, future := range futures[1:] {
		if future.CompletionExecutorID() != executorID {
			return 0
		}
	}
	return executorID
}

type subscriptionGroup struct {
	mu            sync.Mutex
	closed        bool
	subscriptions []Subscription
}

func (group *subscriptionGroup) add(subscription Subscription) {
	group.mu.Lock()
	if group.closed {
		group.mu.Unlock()
		subscription.Cancel()
		return
	}
	group.subscriptions = append(group.subscriptions, subscription)
	group.mu.Unlock()
}

func (group *subscriptionGroup) cancel() {
	group.mu.Lock()
	if group.closed {
		group.mu.Unlock()
		return
	}
	group.closed = true
	subscriptions := group.subscriptions
	group.subscriptions = nil
	group.mu.Unlock()

	for _, subscription := range subscriptions {
		subscription.Cancel()
	}
}

func safeResultCall(fn func(Result) Result, input Result) (ret Result) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			ret = NewResult(nil, panicToError(panicValue))
		}
	}()
	return fn(input)
}

func safeFutureCall(fn func(Result) Future, input Result) (future Future, panicErr error) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			panicErr = panicToError(panicValue)
		}
	}()
	return fn(input), nil
}

func panicToError(panicValue any) error {
	if err, ok := panicValue.(error); ok {
		return exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, err))
	}
	return exception.TraceStack(fmt.Errorf("%w: %v", exception.ErrPanicked, panicValue))
}
