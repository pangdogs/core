package core

import (
	"unsafe"
)

// IFaceCache 接口存储器，因为Golang原生的接口转换性能较差，所以在某些性能要求较高的场景（例如瞬时演算）下，需要尽量较少接口转换。
//目前可用的优化方案是，在编码时明确知晓接口类型时，可以将接口转换为[2]unsafe.Pointer保存接口，使用时再转换为接口即可。
//为了使用方便，本包提供了一套方法用于支持此类需求，注意不安全，在明确了解此功能时再使用。
type IFaceCache [2]unsafe.Pointer

// NilIFaceCache 空接口存储器
var NilIFaceCache IFaceCache

// Cache2IFace 接口存储器转换为接口
func Cache2IFace[T any](fi IFaceCache) T {
	return *(*T)(unsafe.Pointer(&fi))
}

// IFace2Cache 接口转换为接口存储器
func IFace2Cache[T any](i T) IFaceCache {
	return *(*IFaceCache)(unsafe.Pointer(&i))
}

// NewFace 创建面
func NewFace[T any](iFace T) Face[T] {
	return Face[T]{
		IFace: iFace,
		Cache: IFace2Cache[T](iFace),
	}
}

// Face 面，用于存储接口与其存储器，因为在接口转换为存储器之后，存储器并不能向GC标记接口指向的对象被引用，
//其随时会被GC释放掉，所以在需要长期使用接口存储器时，请使用Face同时持有接口与存储器，防止被GC释放。
type Face[T any] struct {
	IFace T
	Cache IFaceCache
}

// IsNil 是否为空
func (f *Face[T]) IsNil() bool {
	return IFace2Cache[T](f.IFace) == NilIFaceCache || f.Cache == NilIFaceCache
}

// FaceAny interface{}面
type FaceAny = Face[interface{}]
