package golaxy

import (
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/runtime"
)

// NewCallStack 创建调用栈
func NewCallStack[T any](runtimeCtx runtime.Context, variables T) CallStack[T] {
	if runtimeCtx == nil {
		panic("nil runtimeCtx")
	}

	return CallStack[T]{
		ctxList:   []runtime.Context{runtimeCtx},
		Variables: variables,
	}
}

// CallStack 调用栈，用于在多个运行时之间安全的互相同步调用
type CallStack[T any] struct {
	ctxList   []runtime.Context
	Variables T
}

// Call 同步调用
func (stack CallStack[T]) Call(runtimeCtx runtime.Context, segment func(stack CallStack[T]) runtime.Ret) runtime.Ret {
	if runtimeCtx == nil {
		panic("nil runtimeCtx")
	}

	if segment == nil {
		panic("nil segment")
	}

	if stack.Exist(runtimeCtx) {
		stack.ctxList = append(stack.ctxList, runtimeCtx)
		return segment(stack)
	}

	stack.ctxList = append(stack.ctxList, runtimeCtx)
	return runtimeCtx.SyncCall(func() internal.Ret {
		return segment(stack)
	})
}

// CallNoRet 同步调用，无返回值
func (stack CallStack[T]) CallNoRet(runtimeCtx runtime.Context, segment func(stack CallStack[T])) {
	if runtimeCtx == nil {
		panic("nil runtimeCtx")
	}

	if segment == nil {
		panic("nil segment")
	}

	if stack.Exist(runtimeCtx) {
		stack.ctxList = append(stack.ctxList, runtimeCtx)
		segment(stack)
		return
	}

	stack.ctxList = append(stack.ctxList, runtimeCtx)
	runtimeCtx.SyncCallNoRet(func() {
		segment(stack)
	})
}

// Exist 运行时是否存在
func (stack *CallStack[T]) Exist(runtimeCtx runtime.Context) bool {
	for i := range stack.ctxList {
		if stack.ctxList[i] == runtimeCtx {
			return true
		}
	}
	return false
}
