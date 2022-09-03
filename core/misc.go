package core

import "reflect"

// Zero 创建零值
func Zero[T any]() T {
	var zero T
	return zero
}

// IFaceName 获取接口名
func IFaceName[T any]() string {
	v := reflect.TypeOf((*T)(nil)).Elem()
	return v.PkgPath() + "/" + v.Name()
}
