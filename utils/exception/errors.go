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

func TraceStack(err error) error {
	stackBuf := make([]byte, 4096)
	n := runtime.Stack(stackBuf, false)
	return fmt.Errorf("error:\n%w\nstack:\n%s\n", err, stackBuf[:n])
}

func Error(v any) error {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Errorf("%v (%s:%d)", v, file, line)
}

func Errorf(format string, args ...any) error {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Errorf(format+" (%s:%d)", append(args, file, line)...)
}

func Panic(v any) {
	panic(Error(v))
}

func Panicf(format string, args ...any) {
	panic(Errorf(format, args...))
}
