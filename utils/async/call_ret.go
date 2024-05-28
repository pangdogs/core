package async

import "fmt"

type (
	Ret = RetT[any]
)

var (
	MakeRet = MakeRetT[any]
	VoidRet = MakeRet(nil, nil)
)

// MakeRetT 创建调用结果
func MakeRetT[T any](val T, err error) RetT[T] {
	return RetT[T]{
		Value: val,
		Error: err,
	}
}

// CastRetT 转换
func CastRetT[T any](ret Ret) RetT[T] {
	return RetT[T]{
		Value: ret.Value.(T),
		Error: ret.Error,
	}
}

// RetT 调用结果
type RetT[T any] struct {
	Value T     // 返回值
	Error error // error
}

// OK 是否成功
func (ret RetT[T]) OK() bool {
	return ret.Error == nil
}

// String implements fmt.Stringer
func (ret RetT[T]) String() string {
	if ret.Error != nil {
		return ret.Error.Error()
	}
	return fmt.Sprintf("%v", ret.Value)
}
