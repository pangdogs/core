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
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/types"
	"hash/fnv"
	"reflect"
	"sync"
)

// MakeEventTabId 创建事件表Id
func MakeEventTabId(eventTab any) uint64 {
	if eventTab == nil {
		exception.Panicf("%w: %w: eventTab is nil", ErrEvent, exception.ErrArgs)
	}

	eventTabRT := reflect.ValueOf(eventTab).Type()

	for eventTabRT.Kind() == reflect.Pointer {
		eventTabRT = eventTabRT.Elem()
	}

	if eventTabRT.PkgPath() == "" || eventTabRT.Name() == "" || !reflect.PointerTo(eventTabRT).Implements(reflect.TypeFor[IEventTab]()) {
		exception.Panicf("unsupported type")
	}

	hash := fnv.New32a()
	hash.Write([]byte(types.FullNameRT(eventTabRT)))
	return uint64(hash.Sum32()) << 32
}

// MakeEventTabIdT 创建事件表Id
func MakeEventTabIdT[T any]() uint64 {
	return MakeEventTabId(types.ZeroT[T]())
}

// MakeEventId 创建事件Id
func MakeEventId(eventTab any, pos int32) uint64 {
	return MakeEventTabId(eventTab) + uint64(pos)
}

// MakeEventIdT 创建事件Id
func MakeEventIdT[T any](pos int32) uint64 {
	return MakeEventId(types.ZeroT[T](), pos)
}

var (
	declareEventTabs = &sync.Map{}
	declareEvents    = &sync.Map{}
)

// DeclareEventTabId 声明事件表Id
func DeclareEventTabId(eventTab any) uint64 {
	id := MakeEventTabId(eventTab)
	if name, loaded := declareEventTabs.LoadOrStore(id, types.FullNameRT(reflect.TypeOf(eventTab).Elem())); loaded {
		exception.Panicf("event_tab(%d) has already been declared by %q", id, name)
	}
	return id
}

// DeclareEventTabIdT 声明事件表Id
func DeclareEventTabIdT[T any]() uint64 {
	return DeclareEventTabId(types.ZeroT[T]())
}

// DeclareEventId 声明事件Id
func DeclareEventId(eventTab any, pos int32) uint64 {
	id := MakeEventTabId(eventTab) + uint64(pos)
	if name, loaded := declareEvents.LoadOrStore(id, types.FullNameRT(reflect.TypeOf(eventTab).Elem())); loaded {
		exception.Panicf("event(%d) has already been declared by %q", id, name)
	}
	return id
}

// DeclareEventIdT 声明事件Id
func DeclareEventIdT[T any](pos int32) uint64 {
	return DeclareEventId(types.ZeroT[T](), pos)
}
