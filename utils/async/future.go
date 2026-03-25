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
	ErrFutureClosed = fmt.Errorf("%w: future closed", exception.ErrCore)
)

func NewFutureVoid() FutureVoid {
	return make(chan struct{})
}

type FutureVoid chan struct{}

func (future FutureVoid) Out() Future {
	return Future{
		done: future,
		void: true,
	}
}

func NewFutureChan(size ...int) FutureChan {
	return FutureChan{
		ch:   make(chan Result, max(1, pie.First(size))),
		done: make(chan struct{}),
	}
}

type FutureChan struct {
	ch   chan Result
	done chan struct{}
}

func (future FutureChan) IsNil() bool {
	return future.ch == nil && future.done == nil
}

func (future FutureChan) Out() Future {
	return Future{
		ch:   future.ch,
		done: future.done,
	}
}

type Future struct {
	ch   <-chan Result
	done <-chan struct{}
	void bool
}

func (future Future) IsNil() bool {
	return future.ch == nil && future.done == nil
}

func (future Future) Void() bool {
	if future.IsNil() {
		exception.Panic("future is nil, cannot check void")
	}
	return future.void
}

func (future Future) Chan() <-chan Result {
	if future.IsNil() {
		exception.Panic("future is nil, cannot get channel")
	}
	return future.ch
}

func (future Future) Done() <-chan struct{} {
	if future.IsNil() {
		exception.Panic("future is nil, cannot get done channel")
	}
	return future.done
}

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
