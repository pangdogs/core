package internal

type Serializer[T any] interface {
	MarshalText(t T) ([]byte, error)
	UnmarshalText(t T, b []byte) error
	MarshalBinary(t T) ([]byte, error)
	UnmarshalBinary(t T, b []byte) error
}
