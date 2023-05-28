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
