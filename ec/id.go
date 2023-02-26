package ec

import (
	"encoding/base64"
	"errors"
	"reflect"
	"unsafe"
)

// ID 唯一ID（160位）
type ID [20]byte

func (id ID) String() string {
	return EncodeIDToString(id)
}

var EncodeIDToString = func(id ID) string {
	return base64.RawURLEncoding.EncodeToString(id[:])
}

var DecodeStringToID = func(str string) (id ID, err error) {
	if base64.RawURLEncoding.DecodedLen(len(str)) > len(id) {
		err = errors.New("string too long")
		return
	}
	base64.RawURLEncoding.Decode(id[:], string2Bytes(str))
	return
}

func string2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
