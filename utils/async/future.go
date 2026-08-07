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
	"fmt"

	"git.golaxy.org/core/utils/exception"
	"github.com/elliotchance/pie/v2"
)

var (
	// ErrFutureClosed 表示结果 Future 在未取得下一项时已经结束。
	ErrFutureClosed = fmt.Errorf("%w: future closed", exception.ErrCore)
)

// NewFutureVoid 创建只表示完成信号的 FutureVoid。
func NewFutureVoid() FutureVoid {
	return make(chan struct{})
}

// FutureVoid 是可由生产者关闭的无结果 Future。
//
// 其零值不可用；生产者应且只能调用一次 ReturnVoid。
type FutureVoid chan struct{}

// Out 返回供消费者使用的只读 Future 视图。
func (future FutureVoid) Out() Future {
	return Future{
		done: future,
		void: true,
	}
}

// NewFutureChan 创建可传递 Result 的 FutureChan。
//
// size 指定结果缓冲区容量，省略或小于 1 时使用 1；仅使用第一个值。
func NewFutureChan(size ...int) FutureChan {
	return FutureChan{
		ch:   make(chan Result, max(1, pie.First(size))),
		done: make(chan struct{}),
	}
}

// FutureChan 是 Future 的生产者端。
//
// 其零值不可用。生产者可选择一次 Return，或多次 YieldReturn 后调用一次 YieldBreak。
type FutureChan struct {
	ch   chan Result
	done chan struct{}
}

// IsNil 报告生产者端是否为零值。
func (future FutureChan) IsNil() bool {
	return future.ch == nil && future.done == nil
}

// Out 返回供消费者使用的只读 Future 视图。
func (future FutureChan) Out() Future {
	return Future{
		ch:   future.ch,
		done: future.done,
	}
}

// Future 是异步结果或完成信号的只读消费者视图。
//
// 结果频道采用竞争消费语义，不会向多个消费者广播。
type Future struct {
	ch   <-chan Result
	done <-chan struct{}
	void bool
}

// IsNil 报告 Future 是否为零值。
func (future Future) IsNil() bool {
	return future.ch == nil && future.done == nil
}

// Void 报告 Future 是否只表示完成信号；零值 Future 会导致 panic。
func (future Future) Void() bool {
	if future.IsNil() {
		exception.Panic("future is nil, cannot check void")
	}
	return future.void
}

// Chan 返回结果频道；零值 Future 会导致 panic。
//
// 对无结果 Future，返回的频道不会产生值，并在 Future 完成时关闭；每次调用可能创建
// 一个新的桥接频道。对结果 Future，返回其共享的消费频道。
func (future Future) Chan() <-chan Result {
	if future.IsNil() {
		exception.Panic("future is nil, cannot get channel")
	}
	if future.Void() {
		ch := make(chan Result)
		select {
		case <-future.done:
			close(ch)
		default:
			go func(done <-chan struct{}) {
				<-done
				close(ch)
			}(future.done)
		}
		return ch
	}
	return future.ch
}

// Done 返回 Future 完成时关闭的频道；零值 Future 会导致 panic。
func (future Future) Done() <-chan struct{} {
	if future.IsNil() {
		exception.Panic("future is nil, cannot get done channel")
	}
	return future.done
}

// Wait 等待下一项结果、无结果 Future 完成或 ctx 取消。
//
// nil ctx 按 context.Background 处理。结果 Future 已结束且没有下一项时返回
// ErrFutureClosed；无结果 Future 完成时返回成功的空结果。零值 Future 会导致 panic。
func (future Future) Wait(ctx context.Context) Result {
	if ctx == nil {
		ctx = context.Background()
	}
	if future.IsNil() {
		exception.Panic("future is nil, cannot wait")
	}

	if future.Void() {
		select {
		case <-future.done:
			return NewResult(nil, nil)
		case <-ctx.Done():
			return NewResult(nil, ctx.Err())
		}
	}

	select {
	case ret, ok := <-future.ch:
		if !ok {
			return NewResult(nil, ErrFutureClosed)
		}
		return ret
	case <-ctx.Done():
		return NewResult(nil, ctx.Err())
	}
}

// Context 返回由 ctx 派生、并在 Future 完成时取消的上下文。
// nil ctx 按 context.Background 处理；零值 Future 会导致 panic。
func (future Future) Context(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if future.IsNil() {
		exception.Panic("future is nil, cannot convert to context")
	}

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		select {
		case <-ctx.Done():
		case <-future.done:
			cancel()
		}
	}()

	return ctx
}
