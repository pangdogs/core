package types

import (
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"
	"reflect"
	"unsafe"
)

// Bool2Int bool转int
func Bool2Int[T constraints.Integer](b bool) T {
	if b {
		return 1
	}
	return 0
}

// Int2Bool int转bool
func Int2Bool[T constraints.Integer](v T) bool {
	if v != 0 {
		return true
	}
	return false
}

// String2Bytes 快速string转bytes
func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// Bytes2String 快速bytes转string
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Panic2Err panic转换为error
func Panic2Err(panicInfo any) error {
	switch info := panicInfo.(type) {
	case nil:
		return nil
	case error:
		return info
	case string:
		return errors.New(info)
	default:
		return fmt.Errorf("%v", panicInfo)
	}
}
