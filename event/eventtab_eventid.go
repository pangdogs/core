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
	"git.golaxy.org/core/utils/types"
	"hash/fnv"
	"reflect"
	"sync"
)

// MakeEventTabId 创建事件表Id
func MakeEventTabId(eventTab IEventTab) uint64 {
	hash := fnv.New32a()
	rt := reflect.ValueOf(eventTab).Type()
	if rt.PkgPath() == "" || rt.Name() == "" {
		panic("unsupported type")
	}
	hash.Write([]byte(types.FullNameRT(rt)))
	return uint64(hash.Sum32()) << 32
}

// MakeEventTabIdT 创建事件表Id
func MakeEventTabIdT[T any]() uint64 {
	hash := fnv.New32a()
	rt := reflect.TypeFor[T]()
	if rt.PkgPath() == "" || rt.Name() == "" || !reflect.PointerTo(rt).Implements(reflect.TypeFor[IEventTab]()) {
		panic("unsupported type")
	}
	hash.Write([]byte(types.FullNameRT(rt)))
	return uint64(hash.Sum32()) << 32
}

// MakeEventId 创建事件Id
func MakeEventId(eventTab IEventTab, pos int32) uint64 {
	return MakeEventTabId(eventTab) + uint64(pos)
}

// MakeEventIdT 创建事件Id
func MakeEventIdT[T any](pos int32) uint64 {
	return MakeEventTabIdT[T]() + uint64(pos)
}

var (
	declareEventTabs = &sync.Map{}
	declareEvents    = &sync.Map{}
)

// DeclareEventTabId 声明事件表Id
func DeclareEventTabId(eventTab IEventTab) uint64 {
	id := MakeEventTabId(eventTab)
	if name, loaded := declareEventTabs.LoadOrStore(id, types.FullNameRT(reflect.TypeOf(eventTab).Elem())); loaded {
		panic(fmt.Errorf("event_tab(%d) has already been declared by %q", id, name))
	}
	return id
}

// DeclareEventTabIdT 声明事件表Id
func DeclareEventTabIdT[T any]() uint64 {
	id := MakeEventTabIdT[T]()
	if name, loaded := declareEventTabs.LoadOrStore(id, types.FullNameT[T]()); loaded {
		panic(fmt.Errorf("event_tab(%d) has already been declared by %q", id, name))
	}
	return id
}

// DeclareEventId 声明事件Id
func DeclareEventId(eventTab IEventTab, pos int32) uint64 {
	id := MakeEventTabId(eventTab) + uint64(pos)
	if name, loaded := declareEvents.LoadOrStore(id, types.FullNameRT(reflect.TypeOf(eventTab).Elem())); loaded {
		panic(fmt.Errorf("event(%d) has already been declared by %q", id, name))
	}
	return id
}

// DeclareEventIdT 声明事件Id
func DeclareEventIdT[T any](pos int32) uint64 {
	id := MakeEventTabIdT[T]() + uint64(pos)
	if name, loaded := declareEvents.LoadOrStore(id, types.FullNameT[T]()); loaded {
		panic(fmt.Errorf("event(%d) has already been declared by %q", id, name))
	}
	return id
}
