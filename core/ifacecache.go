package core

import (
	"unsafe"
)

// IFaceCache ...
type IFaceCache [2]unsafe.Pointer

// NilIFaceCache ...
var NilIFaceCache IFaceCache

// Cache2IFace ...
func Cache2IFace[T any](fi IFaceCache) T {
	return *(*T)(unsafe.Pointer(&fi))
}

// IFace2Cache ...
func IFace2Cache[T any](i T) IFaceCache {
	return *(*IFaceCache)(unsafe.Pointer(&i))
}

// NewFace ...
func NewFace[T any](iFace T) Face[T] {
	return Face[T]{
		IFace: iFace,
		Cache: IFace2Cache[T](iFace),
	}
}

// Face ...
type Face[T any] struct {
	IFace T
	Cache IFaceCache
}

// IsNil ...
func (f *Face[T]) IsNil() bool {
	return IFace2Cache[T](f.IFace) == NilIFaceCache || f.Cache == NilIFaceCache
}

// FaceAny ...
type FaceAny = Face[interface{}]
