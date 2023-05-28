package internal

type TextSerialization interface {
	MarshalText() ([]byte, error)
	UnmarshalText(b []byte) error
}

type BinarySerialization interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary(b []byte) error
}
