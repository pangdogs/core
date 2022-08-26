package core

import "reflect"

// Zero ...
func Zero[T any]() T {
	var zero T
	return zero
}

// IFaceName ...
func IFaceName[T any]() string {
	v := reflect.TypeOf((*T)(nil)).Elem()
	return v.PkgPath() + "/" + v.Name()
}
