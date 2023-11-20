package iface

// MakeFace 创建Face
func MakeFace[T any](iface T) Face[T] {
	return Face[T]{
		Iface: iface,
		Cache: Iface2Cache[T](iface),
	}
}

// MakeFaceAny 创建FaceAny
func MakeFaceAny[T any](iface T) Face[any] {
	return Face[any]{
		Iface: iface,
		Cache: Iface2Cache[T](iface),
	}
}

// Face 用于存储接口与Cache，接口可用于断言转换类型，存储器可用于转换为接口
type Face[T any] struct {
	Iface T
	Cache Cache
}

// IsNil 是否为空
func (f *Face[T]) IsNil() bool {
	return Iface2Cache[T](f.Iface) == NilCache || f.Cache == NilCache
}

// FaceAny any face
type FaceAny = Face[any]
