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

// GenEventTabId 根据事件表具名类型的完整名称生成稳定 ID。
//
// eventTab 可以是值、指针或 reflect.Type，且其指针类型必须实现 IEventTab；否则 panic。
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

// GenEventTabIdT 根据事件表类型 T 的完整名称生成稳定 ID。
func GenEventTabIdT[T any]() uint64 {
	return GenEventTabId(types.Zero[T]())
}

// GenEventId 将事件表 ID 与 16 位位置 pos 组合为事件 ID；pos 越界时 panic。
func GenEventId(eventTab any, pos int) uint64 {
	if pos < 0 || pos > math.MaxUint16 {
		exception.Panicf("%w: %w: pos out of bounds [0,%d]", ErrEvent, exception.ErrArgs, math.MaxUint16)
	}
	return GenEventTabId(eventTab) + uint64(pos)
}

// GenEventIdT 将事件表类型 T 的 ID 与 16 位位置 pos 组合为事件 ID。
func GenEventIdT[T any](pos int) uint64 {
	return GenEventId(types.Zero[T](), pos)
}

var (
	declareEventTabs = &sync.Map{}
	declareEvents    = &sync.Map{}
)

// DeclareEventTabId 生成并登记事件表 ID，用于尽早检测哈希冲突或重复声明。
// 同一进程中每个事件表类型只能声明一次，冲突时 panic。
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

// DeclareEventTabIdT 生成并登记事件表类型 T 的 ID。
func DeclareEventTabIdT[T any]() uint64 {
	return DeclareEventTabId(types.Zero[T]())
}

// DeclareEventId 生成并登记事件 ID，用于检测哈希冲突或重复声明。
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

// DeclareEventIdT 生成并登记事件表类型 T 中指定位置的事件 ID。
func DeclareEventIdT[T any](pos int) uint64 {
	return DeclareEventId(types.Zero[T](), pos)
}

// SplitEventId 将事件 ID 分解为事件表 ID 与表内位置。
func SplitEventId(eventId uint64) (eventTabId uint64, pos int) {
	return eventId & 0xFFFFFFFFFFFF0000, int(eventId & 0xFFFF)
}
