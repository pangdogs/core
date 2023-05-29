package internal

type TextSerialization interface {
	MarshalText() ([]byte, error)
	UnmarshalText(bs []byte) error
}

type BinarySerialization interface {
	MarshalBinary() ([]byte, error)
	UnmarshalBinary(bs []byte) error
}
