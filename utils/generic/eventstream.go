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
	"context"
	"sync"
)

// NewEventStream 创建空事件流；EventStream 的零值同样可用。
func NewEventStream[T any]() *EventStream[T] {
	return &EventStream[T]{}
}

// EventStream 将每次 Publish 广播给所有当前订阅者。
//
// EventStream 可并发使用且首次使用后不可复制。每个订阅者使用无界队列，慢消费者
// 不会丢失事件，但可能持续占用内存。
type EventStream[T any] struct {
	_           noCopy
	mutex       sync.RWMutex
	subscribers map[*UnboundedChannel[T]]chan struct{}
}

// Subscribe 创建订阅，并先按参数顺序发送 catchUp，再接收后续发布。
//
// nil ctx 按 context.Background 处理。ctx 结束或 Clear 被调用时，返回频道会在已排队
// 事件处理完毕后关闭。
func (es *EventStream[T]) Subscribe(ctx context.Context, catchUp ...T) <-chan T {
	if ctx == nil {
		ctx = context.Background()
	}

	es.mutex.Lock()
	defer es.mutex.Unlock()

	subscriber := NewUnboundedChannel[T]()

	for _, e := range catchUp {
		subscriber.In() <- e
	}

	if es.subscribers == nil {
		es.subscribers = map[*UnboundedChannel[T]]chan struct{}{}
	}

	closed := make(chan struct{})
	es.subscribers[subscriber] = closed

	go func() {
		select {
		case <-ctx.Done():
		case <-closed:
		}
		es.mutex.Lock()
		defer es.mutex.Unlock()
		if es.subscribers == nil {
			return
		}
		if _, ok := es.subscribers[subscriber]; ok {
			subscriber.Close()
			close(closed)
			delete(es.subscribers, subscriber)
		}
	}()

	return subscriber.Out()
}

// Publish 向调用时存在的全部订阅者广播 event。
func (es *EventStream[T]) Publish(event T) {
	es.mutex.RLock()
	defer es.mutex.RUnlock()

	for subscriber := range es.subscribers {
		subscriber.In() <- event
	}
}

// Clear 关闭并移除全部当前订阅；清空后的 EventStream 可以继续订阅和发布。
func (es *EventStream[T]) Clear() {
	es.mutex.Lock()
	defer es.mutex.Unlock()

	for subscriber, closed := range es.subscribers {
		subscriber.Close()
		close(closed)
	}
	es.subscribers = nil
}
