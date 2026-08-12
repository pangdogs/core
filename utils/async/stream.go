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

// NewStream 创建单生产者、单消费语义的 Emitter/Stream。
// buffer 省略或小于 1 时使用 1；仅使用第一个值。
func NewStream(buffer ...int) (Emitter, Stream) {
	size := 1
	if len(buffer) > 0 && buffer[0] > 0 {
		size = buffer[0]
	}
	state := &streamState{
		ch:   make(chan Result, size),
		done: make(chan struct{}),
	}
	return Emitter{state: state}, Stream{state: state}
}

// Emitter 是 Stream 的生产者端。通常只应由一个 goroutine 使用。
type Emitter struct {
	state *streamState
}

// IsNil 报告 Emitter 是否为零值。
func (emitter Emitter) IsNil() bool {
	return emitter.state == nil
}

// Stream 返回共享状态的只读流。
func (emitter Emitter) Stream() Stream {
	if emitter.IsNil() {
		exception.Panicf("%w: %w: emitter is nil, cannot get stream", ErrAsync, exception.ErrArgs)
	}
	return Stream{state: emitter.state}
}

// Emit 等待写入一项结果；流关闭或 ctx 取消时返回 false。
func (emitter Emitter) Emit(ctx context.Context, ret Result) bool {
	if emitter.IsNil() {
		exception.Panicf("%w: %w: emitter is nil, cannot emit", ErrAsync, exception.ErrArgs)
	}
	if ctx == nil {
		ctx = context.Background()
	}

	if !emitter.beginSend() {
		return false
	}
	defer emitter.endSend()
	select {
	case emitter.state.ch <- ret:
		return true
	case <-ctx.Done():
		return false
	case <-emitter.state.done:
		return false
	}
}

// TryEmit 无阻塞地写入一项结果。
func (emitter Emitter) TryEmit(ret Result) bool {
	if emitter.IsNil() {
		exception.Panicf("%w: %w: emitter is nil, cannot emit", ErrAsync, exception.ErrArgs)
	}
	if !emitter.beginSend() {
		return false
	}
	defer emitter.endSend()
	select {
	case emitter.state.ch <- ret:
		return true
	case <-emitter.state.done:
		return false
	default:
		return false
	}
}

// Close 幂等关闭流。首次关闭返回 true。
func (emitter Emitter) Close() bool {
	if emitter.IsNil() {
		exception.Panicf("%w: %w: emitter is nil, cannot close", ErrAsync, exception.ErrArgs)
	}
	emitter.state.mu.Lock()
	defer emitter.state.mu.Unlock()
	if emitter.state.closed {
		return false
	}
	emitter.state.closed = true
	close(emitter.state.done)
	emitter.state.closeDataIfIdle()
	return true
}

func (emitter Emitter) beginSend() bool {
	emitter.state.mu.Lock()
	defer emitter.state.mu.Unlock()
	if emitter.state.closed {
		return false
	}
	emitter.state.activeSenders++
	return true
}

func (emitter Emitter) endSend() {
	emitter.state.mu.Lock()
	emitter.state.activeSenders--
	emitter.state.closeDataIfIdle()
	emitter.state.mu.Unlock()
}

// Stream 是连续 Result 的只读消费者视图。
//
// Stream 采用单消费语义；多个消费者读取 Chan 会竞争元素。需要广播时应在更高层
// 使用事件或消息总线，而不是共享 Stream。
type Stream struct {
	state *streamState
}

// IsNil 报告 Stream 是否为零值。
func (stream Stream) IsNil() bool {
	return stream.state == nil
}

// Chan 返回结果频道。
func (stream Stream) Chan() <-chan Result {
	if stream.IsNil() {
		exception.Panicf("%w: %w: stream is nil, cannot get channel", ErrAsync, exception.ErrArgs)
	}
	return stream.state.ch
}

// Done 返回 Stream 关闭时关闭的频道。
func (stream Stream) Done() <-chan struct{} {
	if stream.IsNil() {
		exception.Panicf("%w: %w: stream is nil, cannot get done channel", ErrAsync, exception.ErrArgs)
	}
	return stream.state.done
}

// Next 等待下一项结果。流关闭时 ok 为 false；ctx 取消时返回包含 ctx.Err 的结果。
func (stream Stream) Next(ctx context.Context) (ret Result, ok bool) {
	if stream.IsNil() {
		exception.Panicf("%w: %w: stream is nil, cannot read", ErrAsync, exception.ErrArgs)
	}
	if ctx == nil {
		ctx = context.Background()
	}
	select {
	case ret, ok = <-stream.state.ch:
		return ret, ok
	case <-ctx.Done():
		return NewResult(nil, ctx.Err()), true
	}
}

type streamState struct {
	mu            sync.Mutex
	closed        bool
	activeSenders int
	dataClosed    bool
	ch            chan Result
	done          chan struct{}
}

func (state *streamState) closeDataIfIdle() {
	if state.closed && state.activeSenders == 0 && !state.dataClosed {
		state.dataClosed = true
		close(state.ch)
	}
}
