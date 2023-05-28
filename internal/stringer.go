package internal

type Stringer[T any] interface {
	String(t T) string
}
