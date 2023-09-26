package types

import (
	"reflect"
	"strings"
)

// Zero 创建零值
func Zero[T any]() T {
	var zero T
	return zero
}

// New 创建并初始化
func New[T any](v T) *T {
	var rv T
	rv = v
	return &rv
}

// Name 类型名
func Name[T any]() string {
	return reflect.TypeOf((*T)(nil)).Elem().Name()
}

// FullName 类型全名
func FullName[T any]() string {
	t := reflect.TypeOf((*T)(nil)).Elem()
	return t.PkgPath() + "." + t.Name()
}

// AnyName 类型名
func AnyName(i any) string {
	t, ok := i.(reflect.Type)
	if !ok {
		t = reflect.TypeOf(i)
	}
	return t.Name()
}

// AnyFullName 类型全名
func AnyFullName(i any) string {
	t, ok := i.(reflect.Type)
	if !ok {
		t = reflect.TypeOf(i)
	}
	return t.PkgPath() + "." + t.Name()
}

// TypeName 类型名
func TypeName(t reflect.Type) string {
	return t.Name()
}

// TypeFullName 类型全名
func TypeFullName(t reflect.Type) string {
	return t.PkgPath() + "." + t.Name()
}

// WriteFullName 写入类型全名
func WriteFullName[T any](sb *strings.Builder) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	sb.WriteString(t.PkgPath())
	sb.WriteString(".")
	sb.WriteString(t.Name())
}

// WriteAnyFullName 写入类型全名
func WriteAnyFullName(sb *strings.Builder, i any) {
	t, ok := i.(reflect.Type)
	if !ok {
		t = reflect.TypeOf(i)
	}
	sb.WriteString(t.PkgPath())
	sb.WriteString(".")
	sb.WriteString(t.Name())
}

// WriteTypeFullName 写入类型全名
func WriteTypeFullName(sb *strings.Builder, t reflect.Type) {
	sb.WriteString(t.PkgPath())
	sb.WriteString(".")
	sb.WriteString(t.Name())
}
