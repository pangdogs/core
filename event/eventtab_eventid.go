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
