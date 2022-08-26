package core

type _Callee interface {
	pushCall(segment func())
}

type _SafeCall interface {
	SafeCall(segment func() SafeRet) <-chan SafeRet
	SafeCallNoRet(segment func())
	setCallee(callee _Callee)
}

// SafeRet ...
type SafeRet struct {
	Err error
	Ret interface{}
}
