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
	return id.Encode()
}

func (id ID) Encode() string {
	return base64.RawURLEncoding.EncodeToString(id[:])
}

func (id *ID) Decode(str string) error {
	if base64.RawURLEncoding.DecodedLen(len(str)) > len(id) {
		return errors.New("string too long")
	}
	_, err := base64.RawURLEncoding.Decode(id[:], string2Bytes(str))
	return err
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
