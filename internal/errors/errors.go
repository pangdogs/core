package errors

import "errors"

var (
	ErrGolaxy   = errors.New("golaxy")   // GOLAXY框架错误
	ErrPanicked = errors.New("panicked") // panic错误
	ErrArgs     = errors.New("args")     // 参数错误
)
