package concurrent

import (
	"sync"
	"sync/atomic"
)

type Map[K comparable, V any] struct {
	mu     sync.Mutex
	read   atomic.Value
	dirty  map[K]*entry[V]
	misses int
}

type readOnly[K comparable, V any] struct {
	m       map[K]*entry[V]
	amended bool
}

func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	read, _ := m.read.Load().(readOnly[K, V])
	e, ok := read.m[key]
	if !ok && read.amended {
		m.mu.Lock()
		read, _ = m.read.Load().(readOnly[K, V])
		e, ok = read.m[key]
		if !ok && read.amended {
			e, ok = m.dirty[key]
			m.missLocked()
		}
		m.mu.Unlock()
	}
	if !ok {
		return zero[V](), false
	}
	return e.load()
}

func (m *Map[K, V]) Store(key K, value V) {
	read, _ := m.read.Load().(readOnly[K, V])
	if e, ok := read.m[key]; ok && e.tryStore(&value) {
		return
	}

	m.mu.Lock()
	read, _ = m.read.Load().(readOnly[K, V])
	if e, ok := read.m[key]; ok {
		if e.unexpungeLocked() {
			m.dirty[key] = e
		}
		e.storeLocked(&value)
	} else if e, ok := m.dirty[key]; ok {
		e.storeLocked(&value)
	} else {
		if !read.amended {
			m.dirtyLocked()
			m.read.Store(readOnly[K, V]{m: read.m, amended: true})
		}
		m.dirty[key] = newEntry[V](value)
	}
	m.mu.Unlock()
}

func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	read, _ := m.read.Load().(readOnly[K, V])
	if e, ok := read.m[key]; ok {
		actual, loaded, ok := e.tryLoadOrStore(value)
		if ok {
			return actual, loaded
		}
	}

	m.mu.Lock()
	read, _ = m.read.Load().(readOnly[K, V])
	if e, ok := read.m[key]; ok {
		if e.unexpungeLocked() {
			m.dirty[key] = e
		}
		actual, loaded, _ = e.tryLoadOrStore(value)
	} else if e, ok := m.dirty[key]; ok {
		actual, loaded, _ = e.tryLoadOrStore(value)
		m.missLocked()
	} else {
		if !read.amended {
			m.dirtyLocked()
			m.read.Store(readOnly[K, V]{m: read.m, amended: true})
		}
		m.dirty[key] = newEntry[V](value)
		actual, loaded = value, false
	}
	m.mu.Unlock()

	return actual, loaded
}

func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	read, _ := m.read.Load().(readOnly[K, V])
	e, ok := read.m[key]
	if !ok && read.amended {
		m.mu.Lock()
		read, _ = m.read.Load().(readOnly[K, V])
		e, ok = read.m[key]
		if !ok && read.amended {
			e, ok = m.dirty[key]
			delete(m.dirty, key)
			m.missLocked()
		}
		m.mu.Unlock()
	}
	if ok {
		return e.delete()
	}
	return zero[V](), false
}

func (m *Map[K, V]) Delete(key K) {
	m.LoadAndDelete(key)
}

func (m *Map[K, V]) TryLoadAndDelete(key K, compareValue func(v V) bool) (value V, loaded bool) {
	read, _ := m.read.Load().(readOnly[K, V])
	e, ok := read.m[key]
	if !ok && read.amended {
		if !func() bool {
			m.mu.Lock()
			defer m.mu.Unlock()
			read, _ = m.read.Load().(readOnly[K, V])
			e, ok = read.m[key]
			if !ok && read.amended {
				e, ok = m.dirty[key]
				if !invokeCompareValue(compareValue, e) {
					return false
				}
				delete(m.dirty, key)
				m.missLocked()
			}
			return true
		}() {
			return zero[V](), false
		}
	}
	if ok {
		if !invokeCompareValue(compareValue, e) {
			return zero[V](), false
		}
		return e.delete()
	}
	return zero[V](), false
}

func (m *Map[K, V]) TryDelete(key K, compareValue func(v V) bool) {
	m.TryLoadAndDelete(key, compareValue)
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	read, _ := m.read.Load().(readOnly[K, V])
	if read.amended {
		m.mu.Lock()
		read, _ = m.read.Load().(readOnly[K, V])
		if read.amended {
			read = readOnly[K, V]{m: m.dirty}
			m.read.Store(read)
			m.dirty = nil
			m.misses = 0
		}
		m.mu.Unlock()
	}

	for k, e := range read.m {
		v, ok := e.load()
		if !ok {
			continue
		}
		if !f(k, v) {
			break
		}
	}
}

func (m *Map[K, V]) missLocked() {
	m.misses++
	if m.misses < len(m.dirty) {
		return
	}
	m.read.Store(readOnly[K, V]{m: m.dirty})
	m.dirty = nil
	m.misses = 0
}

func (m *Map[K, V]) dirtyLocked() {
	if m.dirty != nil {
		return
	}

	read, _ := m.read.Load().(readOnly[K, V])
	m.dirty = make(map[K]*entry[V], len(read.m))
	for k, e := range read.m {
		if !e.tryExpungeLocked() {
			m.dirty[k] = e
		}
	}
}

func invokeCompareValue[V any](compareValue func(v V) bool, e *entry[V]) bool {
	if compareValue == nil || e == nil {
		return true
	}

	v, ok := e.load()
	if !ok {
		return true
	}

	return compareValue(v)
}

func zero[T any]() T {
	var z T
	return z
}
