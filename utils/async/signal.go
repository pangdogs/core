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
	"sync/atomic"

	"git.golaxy.org/core/utils/exception"
)

// NewSignal 创建一对完成端 Completer 和等待端 Signal。
func NewSignal() (Completer, Signal) {
	state := &signalState{done: make(chan struct{})}
	return Completer{state: state}, Signal{state: state}
}

// CompletedSignal 创建已经完成的 Signal。
func CompletedSignal() Signal {
	completer, signal := NewSignal()
	completer.Complete()
	return signal
}

// Completer 是无结果完成信号的生产者端。
type Completer struct {
	state *signalState
}

// IsNil 报告 Completer 是否为零值。
func (completer Completer) IsNil() bool {
	return completer.state == nil
}

// Signal 返回共享状态的只读完成信号。
func (completer Completer) Signal() Signal {
	if completer.IsNil() {
		exception.Panicf("%w: %w: completer is nil, cannot get signal", ErrAsync, exception.ErrArgs)
	}
	return Signal{state: completer.state}
}

// Complete 完成 Signal。首次完成返回 true，后续调用返回 false。
func (completer Completer) Complete() bool {
	if completer.IsNil() {
		exception.Panicf("%w: %w: completer is nil, cannot complete", ErrAsync, exception.ErrArgs)
	}
	if !completer.state.completed.CompareAndSwap(false, true) {
		return false
	}
	close(completer.state.done)
	return true
}

// Signal 是可重放、无结果的完成通知。
type Signal struct {
	state *signalState
}

// IsNil 报告 Signal 是否为零值。
func (signal Signal) IsNil() bool {
	return signal.state == nil
}

// Done 返回 Signal 完成时关闭的频道。
func (signal Signal) Done() <-chan struct{} {
	if signal.IsNil() {
		exception.Panicf("%w: %w: signal is nil, cannot get done channel", ErrAsync, exception.ErrArgs)
	}
	return signal.state.done
}

// Completed 报告 Signal 是否已经完成。
func (signal Signal) Completed() bool {
	return !signal.IsNil() && signal.state.completed.Load()
}

// Wait 等待 Signal 完成或 ctx 取消，并返回取消错误。
func (signal Signal) Wait(ctx context.Context) error {
	if signal.IsNil() {
		exception.Panicf("%w: %w: signal is nil, cannot wait", ErrAsync, exception.ErrArgs)
	}
	if ctx == nil {
		ctx = context.Background()
	}
	select {
	case <-signal.state.done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

type signalState struct {
	completed atomic.Bool
	done      chan struct{}
}
