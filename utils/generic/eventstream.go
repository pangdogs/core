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

func NewEventStream[T any]() *EventStream[T] {
	return &EventStream[T]{}
}

type EventStream[T any] struct {
	_           noCopy
	mutex       sync.RWMutex
	subscribers map[*UnboundedChannel[T]]chan struct{}
}

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

func (es *EventStream[T]) Publish(event T) {
	es.mutex.RLock()
	defer es.mutex.RUnlock()

	for subscriber := range es.subscribers {
		subscriber.In() <- event
	}
}

func (es *EventStream[T]) Clear() {
	es.mutex.Lock()
	defer es.mutex.Unlock()

	for subscriber, closed := range es.subscribers {
		subscriber.Close()
		close(closed)
	}
	es.subscribers = nil
}
