package core

import (
	"reflect"
)

// Zero 创建零值
func Zero[T any]() T {
	var zero T
	return zero
}

// TypeFullName 类型全名
func TypeFullName[T any]() string {
	v := reflect.TypeOf((*T)(nil)).Elem()
	return v.PkgPath() + "/" + v.Name()
}

// TypeName 类型名
func TypeName[T any]() string {
	return reflect.TypeOf((*T)(nil)).Elem().Name()
}
