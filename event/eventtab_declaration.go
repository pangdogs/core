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

package event

import (
	"fmt"
	"hash/fnv"
	"math"
	"reflect"
	"sync"

	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/types"
)

// GenEventTabId 生成事件表Id
func GenEventTabId(eventTab any) uint64 {
	if eventTab == nil {
		exception.Panicf("%w: %w: eventTab is nil", ErrEvent, exception.ErrArgs)
	}

	eventTabRT, ok := eventTab.(reflect.Type)
	if !ok {
		eventTabRT = reflect.ValueOf(eventTab).Type()
	}

	for eventTabRT.Kind() == reflect.Pointer {
		eventTabRT = eventTabRT.Elem()
	}

	if eventTabRT.PkgPath() == "" || eventTabRT.Name() == "" || !reflect.PointerTo(eventTabRT).Implements(reflect.TypeFor[IEventTab]()) {
		exception.Panicf("%w: unsupported type", ErrEvent)
	}

	hash := fnv.New64a()
	hash.Write(types.String2Bytes(types.FullNameRT(eventTabRT)))
	return hash.Sum64() << 16
}

// GenEventTabIdT 生成事件表Id
func GenEventTabIdT[T any]() uint64 {
	return GenEventTabId(types.ZeroT[T]())
}

// GenEventId 生成事件Id
func GenEventId(eventTab any, pos int) uint64 {
	if pos < 0 || pos > math.MaxUint16 {
		exception.Panicf("%w: %w: pos out of bounds [0,%d]", ErrEvent, exception.ErrArgs, math.MaxUint16)
	}
	return GenEventTabId(eventTab) + uint64(pos)
}

// GenEventIdT 生成事件Id
func GenEventIdT[T any](pos int) uint64 {
	return GenEventId(types.ZeroT[T](), pos)
}

var (
	declareEventTabs = &sync.Map{}
	declareEvents    = &sync.Map{}
)

// DeclareEventTabId 声明事件表Id
func DeclareEventTabId(eventTab any) uint64 {
	id := GenEventTabId(eventTab)

	eventTabRT, ok := eventTab.(reflect.Type)
	if !ok {
		eventTabRT = reflect.ValueOf(eventTab).Type()
	}

	for eventTabRT.Kind() == reflect.Pointer {
		eventTabRT = eventTabRT.Elem()
	}

	info := types.FullNameRT(eventTabRT)

	if exists, loaded := declareEventTabs.LoadOrStore(id, info); loaded {
		exception.Panicf("%w: event tab %q id %d conflict with %q, rename required", ErrEvent, info, id, exists)
	}

	return id
}

// DeclareEventTabIdT 声明事件表Id
func DeclareEventTabIdT[T any]() uint64 {
	return DeclareEventTabId(types.ZeroT[T]())
}

// DeclareEventId 声明事件Id
func DeclareEventId(eventTab any, pos int) uint64 {
	id := GenEventId(eventTab, pos)

	eventTabRT, ok := eventTab.(reflect.Type)
	if !ok {
		eventTabRT = reflect.ValueOf(eventTab).Type()
	}

	for eventTabRT.Kind() == reflect.Pointer {
		eventTabRT = eventTabRT.Elem()
	}

	info := fmt.Sprintf("%s[%d]", types.FullNameRT(eventTabRT), pos)

	if exists, loaded := declareEvents.LoadOrStore(id, info); loaded {
		exception.Panicf("%w: event tab %q id %d conflict with %q, rename required", ErrEvent, info, id, exists)
	}

	return id
}

// DeclareEventIdT 声明事件Id
func DeclareEventIdT[T any](pos int) uint64 {
	return DeclareEventId(types.ZeroT[T](), pos)
}

// SplitEventId 分解事件Id
func SplitEventId(eventId uint64) (eventTabId uint64, pos int) {
	return eventId & 0xFFFFFFFFFFFF0000, int(eventId & 0xFFFF)
}
