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

package generic

import (
	"sync"
	"sync/atomic"
)

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type Barrier struct {
	_                   noCopy
	initOnce, closeOnce sync.Once
	n                   atomic.Int64
	closed              atomic.Bool
	done                chan struct{}
}

func (b *Barrier) Join(delta int64) bool {
	b.initOnce.Do(b.init)
	if delta == 0 {
		return false
	}
	if delta > 0 && b.closed.Load() {
		return false
	}
	n := b.n.Add(delta)
	if delta > 0 {
		if b.closed.Load() {
			if b.n.Add(-delta) == 0 {
				b.closeOnce.Do(b.close)
			}
			return false
		}
	} else {
		if n == 0 {
			b.closed.Store(true)
			b.closeOnce.Do(b.close)
		}
	}
	return true
}

func (b *Barrier) Done() {
	b.Join(-1)
}

func (b *Barrier) Wait() {
	b.initOnce.Do(b.init)
	<-b.done
}

func (b *Barrier) Close() {
	b.initOnce.Do(b.init)
	if b.closed.CompareAndSwap(false, true) {
		if b.n.Add(-1) == 0 {
			b.closeOnce.Do(b.close)
		}
	}
}

func (b *Barrier) IsClosed() bool {
	b.initOnce.Do(b.init)
	return b.closed.Load()
}

func (b *Barrier) Pending() int64 {
	b.initOnce.Do(b.init)
	return b.n.Load()
}

func (b *Barrier) init() {
	b.n.Store(1)
	b.done = make(chan struct{})
}

func (b *Barrier) close() {
	close(b.done)
}
