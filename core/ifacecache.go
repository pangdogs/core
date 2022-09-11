package core

import (
	"unsafe"
)

// IfaceCache 接口存储器，因为Golang原生的接口转换性能较差，所以在某些性能要求较高的场景（例如瞬时演算）下，需要尽量较少接口转换。
//目前可用的优化方案是，在编码时明确知晓接口类型时，可以将接口转换为[2]unsafe.Pointer保存接口，使用时再转换为接口即可。
//为了使用方便，本包提供了一套方法用于支持此类需求，注意不安全，在明确了解此功能时再使用。
type IfaceCache [2]unsafe.Pointer

// NilIfaceCache 空接口存储器
var NilIfaceCache IfaceCache

// Cache2Iface 接口存储器转换为接口
func Cache2Iface[T any](fi IfaceCache) T {
	return *(*T)(unsafe.Pointer(&fi))
}

// Iface2Cache 接口转换为接口存储器
func Iface2Cache[T any](i T) IfaceCache {
	return *(*IfaceCache)(unsafe.Pointer(&i))
}

// NewFace 创建面
func NewFace[T any](iface T) Face[T] {
	return Face[T]{
		Iface: iface,
		Cache: Iface2Cache[T](iface),
	}
}

// Face 面，用于存储接口与其存储器，因为在接口转换为存储器之后，存储器并不能向GC标记接口指向的对象被引用，
//其随时会被GC释放掉，所以在需要长期使用接口存储器时，请使用Face同时持有接口与存储器，防止被GC释放。
type Face[T any] struct {
	Iface T
	Cache IfaceCache
}

// IsNil 是否为空
func (f *Face[T]) IsNil() bool {
	return Iface2Cache[T](f.Iface) == NilIfaceCache || f.Cache == NilIfaceCache
}

// FaceAny interface{}面
type FaceAny = Face[interface{}]
