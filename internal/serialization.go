package internal

import "database/sql/driver"

type TextSerialization interface {
	MarshalText() ([]byte, error)
	UnmarshalText(b []byte) error
}

type BinarySerialization interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary(b []byte) error
}

type SqlValue interface {
	Value() (driver.Value, error)
	Scan(src interface{}) error
}

type Serializer[T any] interface {
	String(t T) string
	MarshalText(t T) ([]byte, error)
	UnmarshalText(t T, b []byte) error
	MarshalBinary(t T) ([]byte, error)
	UnmarshalBinary(t T, b []byte) error
	Value(t T) (driver.Value, error)
	Scan(t T, src interface{}) error
}
