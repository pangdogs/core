package internal

import "database/sql/driver"

type Serializer[T any] interface {
	MarshalText(t T) ([]byte, error)
	UnmarshalText(t T, b []byte) error
	MarshalBinary(t T) ([]byte, error)
	UnmarshalBinary(t T, b []byte) error
	Value(t T) (driver.Value, error)
	Scan(t T, src interface{}) error
}
