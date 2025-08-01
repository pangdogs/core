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
	ErrCore     = errors.New("core")     // 内核错误
	ErrPanicked = errors.New("panicked") // panic错误
	ErrArgs     = errors.New("args")     // 参数错误
)

type StackedError struct {
	Err   error
	Stack []byte
}

func (e StackedError) Error() string {
	return fmt.Sprintf("%s\n\n%s\n", e.Err, e.Stack)
}

func TraceStack(err error) error {
	stackBuf := make([]byte, 4096)
	n := runtime.Stack(stackBuf, false)
	return &StackedError{
		Err:   err,
		Stack: stackBuf[:n],
	}
}

func Error(v any) error {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Errorf("%v (at %s:%d)", v, file, line)
}

func Errorf(format string, args ...any) error {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Errorf(format+" (at %s:%d)", append(args, file, line)...)
}

func ErrorSkip(skip int, v any) error {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Errorf("%v (at %s:%d)", v, file, line)
}

func ErrorfSkip(skip int, format string, args ...any) error {
	_, file, line, _ := runtime.Caller(skip)
	return fmt.Errorf(format+" (at %s:%d)", append(args, file, line)...)
}

func Panic(v any) {
	_, file, line, _ := runtime.Caller(1)
	panic(fmt.Errorf("%v (at %s:%d)", v, file, line))
}

func Panicf(format string, args ...any) {
	_, file, line, _ := runtime.Caller(1)
	panic(fmt.Errorf(format+" (at %s:%d)", append(args, file, line)...))
}

func PanicSkip(skip int, v any) {
	_, file, line, _ := runtime.Caller(skip)
	panic(fmt.Errorf("%v (at %s:%d)", v, file, line))
}

func PanicfSkip(skip int, format string, args ...any) {
	_, file, line, _ := runtime.Caller(skip)
	panic(fmt.Errorf(format+" (at %s:%d)", append(args, file, line)...))
}
