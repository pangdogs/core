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
	return make(chan Result)
}

type FutureVoid chan Result

func (future FutureVoid) Out() Future {
	return chan Result(future)
}

func NewFutureStream(size ...int) FutureStream {
	return make(chan Result, max(1, pie.First(size)))
}

type FutureStream chan Result

func (future FutureStream) Out() Future {
	return chan Result(future)
}

type Future <-chan Result

func (future Future) Void() bool {
	return cap(future) <= 0
}

func (future Future) Wait(ctx context.Context) Result {
	if ctx == nil {
		ctx = context.Background()
	}

	select {
	case ret, ok := <-future:
		if !ok {
			if future.Void() {
				return Result{}
			}
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

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		for {
			select {
			case _, ok := <-ctx.Done():
				if !ok {
					return
				}
			case <-future:
				cancel()
				return
			}
		}
	}()

	return ctx
}
