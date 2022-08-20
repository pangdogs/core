package core

import (
	"unsafe"
)

type IFaceCache [2]unsafe.Pointer

var NilIFaceCache IFaceCache

func Cache2IFace[T any](fi IFaceCache) T {
	return *(*T)(unsafe.Pointer(&fi))
}

func IFace2Cache[T any](i T) IFaceCache {
	return *(*IFaceCache)(unsafe.Pointer(&i))
}

func NewFace[T any](iFace T) Face[T] {
	return Face[T]{
		IFace: iFace,
		Cache: IFace2Cache[T](iFace),
	}
}

type Face[T any] struct {
	IFace T
	Cache IFaceCache
}

func (f *Face[T]) IsNil() bool {
	return IFace2Cache[T](f.IFace) == NilIFaceCache || f.Cache == NilIFaceCache
}

type FaceAny = Face[interface{}]
