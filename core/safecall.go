package core

type _Callee interface {
	pushCall(segment func())
}

type SafeCall interface {
	SafeCall(segment func() SafeRet) <-chan SafeRet
	SafeCallNoRet(segment func())
	setCallee(callee _Callee)
}

type SafeRet struct {
	Err error
	Ret interface{}
}
