/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

package exception

import (
	"errors"
	"fmt"
	"runtime"
)

var (
	ErrCore     = errors.New("core")     // ErrCore 是框架错误的共同根错误。
	ErrPanicked = errors.New("panicked") // ErrPanicked 标识由 panic 恢复得到的错误。
	ErrArgs     = errors.New("args")     // ErrArgs 标识参数错误。
)

// ErrorWithStack 将错误与捕获到的当前协程堆栈组合为一个错误值。
type ErrorWithStack struct {
	Err   error  // Err 是原始错误。
	Stack []byte // Stack 是 runtime.Stack 生成的堆栈文本。
}

// Error 返回原始错误与堆栈文本。
func (e ErrorWithStack) Error() string {
	return fmt.Sprintf("%s\n\n%s\n", e.Err, e.Stack)
}

// TraceStack 捕获当前协程堆栈并将其附加到 err。
func TraceStack(err error) error {
	stackBuf := make([]byte, 4096)
	n := runtime.Stack(stackBuf, false)
	return &ErrorWithStack{
		Err:   err,
		Stack: stackBuf[:n],
	}
}

// Error 将 v 格式化为带直接调用位置的错误。
func Error(v any) error {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Errorf("%v (at %s:%d)", v, file, line)
}

// Errorf 按 format 构造带直接调用位置的错误。
func Errorf(format string, args ...any) error {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Errorf(format+" (at %s:%d)", append(args, file, line)...)
}

// ErrorSkip 将 v 格式化为错误，并使用 runtime.Caller(skip) 指定的位置。
func ErrorSkip(skip int, v any) error {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Errorf("%v (at %s:%d)", v, file, line)
}

// ErrorfSkip 按 format 构造错误，并使用 runtime.Caller(skip) 指定的位置。
func ErrorfSkip(skip int, format string, args ...any) error {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Errorf(format+" (at %s:%d)", append(args, file, line)...)
}

// Panic 以带直接调用位置的 error 触发 panic。
func Panic(v any) {
	_, file, line, _ := runtime.Caller(1)
	panic(fmt.Errorf("%v (at %s:%d)", v, file, line))
}

// Panicf 按 format 构造带直接调用位置的 error 并触发 panic。
func Panicf(format string, args ...any) {
	_, file, line, _ := runtime.Caller(1)
	panic(fmt.Errorf(format+" (at %s:%d)", append(args, file, line)...))
}

// PanicSkip 以 error 触发 panic，并使用 runtime.Caller(skip) 指定的位置。
func PanicSkip(skip int, v any) {
	_, file, line, _ := runtime.Caller(skip)
	panic(fmt.Errorf("%v (at %s:%d)", v, file, line))
}

// PanicfSkip 按 format 构造 error 并触发 panic，位置由 runtime.Caller(skip) 指定。
func PanicfSkip(skip int, format string, args ...any) {
	_, file, line, _ := runtime.Caller(skip)
	panic(fmt.Errorf(format+" (at %s:%d)", append(args, file, line)...))
}
