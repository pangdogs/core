package types

import (
	"reflect"
	"strings"
)

// ZeroT 创建零值
func ZeroT[T any]() T {
	var zero T
	return zero
}

// NewT 新建零值
func NewT[T any]() *T {
	var zero T
	return &zero
}

// NewCopiedT 新建拷贝值
func NewCopiedT[T any](src T) *T {
	return &src
}

// Name 类型名
func Name(i any) string {
	t, ok := i.(reflect.Type)
	if !ok {
		t = reflect.TypeOf(i)
	}
	return t.Name()
}

// NameRT 类型名
func NameRT(t reflect.Type) string {
	return t.Name()
}

// NameT 类型名
func NameT[T any]() string {
	return reflect.TypeFor[T]().Name()
}

// FullName 类型全名
func FullName(i any) string {
	t, ok := i.(reflect.Type)
	if !ok {
		t = reflect.TypeOf(i)
	}
	return FullNameRT(t)
}

// FullNameRT 类型全名
func FullNameRT(t reflect.Type) string {
	return t.PkgPath() + "." + t.Name()
}

// FullNameT 类型全名
func FullNameT[T any]() string {
	return FullNameRT(reflect.TypeFor[T]())
}

// WriteAnyFullName 写入类型全名
func WriteAnyFullName(sb *strings.Builder, i any) {
	t, ok := i.(reflect.Type)
	if !ok {
		t = reflect.TypeOf(i)
	}
	WriteFullNameRT(sb, t)
}

// WriteFullNameRT 写入类型全名
func WriteFullNameRT(sb *strings.Builder, t reflect.Type) {
	sb.WriteString(t.PkgPath())
	sb.WriteString(".")
	sb.WriteString(t.Name())
}

// WriteFullNameT 写入类型全名
func WriteFullNameT[T any](sb *strings.Builder) {
	WriteFullNameRT(sb, reflect.TypeFor[T]())
}
