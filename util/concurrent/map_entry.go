package concurrent

import (
	"sync/atomic"
	"unsafe"
)

var expunged = unsafe.Pointer(new(any))

func newEntry[V any](i V) *entry[V] {
	return &entry[V]{p: unsafe.Pointer(&i)}
}

type entry[V any] struct {
	p unsafe.Pointer
}

func (e *entry[V]) load() (value V, ok bool) {
	p := atomic.LoadPointer(&e.p)
	if p == nil || p == expunged {
		return zero[V](), false
	}
	return *(*V)(p), true
}

func (e *entry[V]) tryStore(i *V) bool {
	for {
		p := atomic.LoadPointer(&e.p)
		if p == expunged {
			return false
		}
		if atomic.CompareAndSwapPointer(&e.p, p, unsafe.Pointer(i)) {
			return true
		}
	}
}

func (e *entry[V]) unexpungeLocked() (wasExpunged bool) {
	return atomic.CompareAndSwapPointer(&e.p, expunged, nil)
}

func (e *entry[V]) storeLocked(i *V) {
	atomic.StorePointer(&e.p, unsafe.Pointer(i))
}

func (e *entry[V]) tryLoadOrStore(i V) (actual V, loaded, ok bool) {
	p := atomic.LoadPointer(&e.p)
	if p == expunged {
		return zero[V](), false, false
	}
	if p != nil {
		return *(*V)(p), true, true
	}

	ic := i
	for {
		if atomic.CompareAndSwapPointer(&e.p, nil, unsafe.Pointer(&ic)) {
			return i, false, true
		}
		p = atomic.LoadPointer(&e.p)
		if p == expunged {
			return zero[V](), false, false
		}
		if p != nil {
			return *(*V)(p), true, true
		}
	}
}

func (e *entry[V]) delete() (value V, ok bool) {
	for {
		p := atomic.LoadPointer(&e.p)
		if p == nil || p == expunged {
			return zero[V](), false
		}
		if atomic.CompareAndSwapPointer(&e.p, p, nil) {
			return *(*V)(p), true
		}
	}
}

func (e *entry[V]) tryExpungeLocked() (isExpunged bool) {
	p := atomic.LoadPointer(&e.p)
	for p == nil {
		if atomic.CompareAndSwapPointer(&e.p, nil, expunged) {
			return true
		}
		p = atomic.LoadPointer(&e.p)
	}
	return p == expunged
}
