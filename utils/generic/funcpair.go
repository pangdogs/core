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

package generic

import (
	"fmt"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/types"
)

type FuncPair0[R1, R2 any] func() (R1, R2)

func (f FuncPair0[R1, R2]) Exec() (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil)
	return
}

func (f FuncPair0[R1, R2]) Invoke() (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil)
}

func (f FuncPair0[R1, R2]) Call(autoRecover bool, reportError chan error) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f()
	return
}

func (f FuncPair0[R1, R2]) ToDelegate() DelegatePair0[R1, R2] {
	return []FuncPair0[R1, R2]{f}
}

type FuncPair1[A1, R1, R2 any] func(A1) (R1, R2)

func (f FuncPair1[A1, R1, R2]) Exec(a1 A1) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1)
	return
}

func (f FuncPair1[A1, R1, R2]) Invoke(a1 A1) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1)
}

func (f FuncPair1[A1, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1)
	return
}

func (f FuncPair1[A1, R1, R2]) ToDelegate() DelegatePair1[A1, R1, R2] {
	return []FuncPair1[A1, R1, R2]{f}
}

type FuncPair2[A1, A2, R1, R2 any] func(A1, A2) (R1, R2)

func (f FuncPair2[A1, A2, R1, R2]) Exec(a1 A1, a2 A2) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2)
	return
}

func (f FuncPair2[A1, A2, R1, R2]) Invoke(a1 A1, a2 A2) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2)
}

func (f FuncPair2[A1, A2, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2)
	return
}

func (f FuncPair2[A1, A2, R1, R2]) ToDelegate() DelegatePair2[A1, A2, R1, R2] {
	return []FuncPair2[A1, A2, R1, R2]{f}
}

type FuncPair3[A1, A2, A3, R1, R2 any] func(A1, A2, A3) (R1, R2)

func (f FuncPair3[A1, A2, A3, R1, R2]) Exec(a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3)
	return
}

func (f FuncPair3[A1, A2, A3, R1, R2]) Invoke(a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3)
}

func (f FuncPair3[A1, A2, A3, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3)
	return
}

func (f FuncPair3[A1, A2, A3, R1, R2]) ToDelegate() DelegatePair3[A1, A2, A3, R1, R2] {
	return []FuncPair3[A1, A2, A3, R1, R2]{f}
}

type FuncPair4[A1, A2, A3, A4, R1, R2 any] func(A1, A2, A3, A4) (R1, R2)

func (f FuncPair4[A1, A2, A3, A4, R1, R2]) Exec(a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4)
	return
}

func (f FuncPair4[A1, A2, A3, A4, R1, R2]) Invoke(a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4)
}

func (f FuncPair4[A1, A2, A3, A4, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4)
	return
}

func (f FuncPair4[A1, A2, A3, A4, R1, R2]) ToDelegate() DelegatePair4[A1, A2, A3, A4, R1, R2] {
	return []FuncPair4[A1, A2, A3, A4, R1, R2]{f}
}

type FuncPair5[A1, A2, A3, A4, A5, R1, R2 any] func(
	A1, A2, A3, A4, A5,
) (R1, R2)

func (f FuncPair5[A1, A2, A3, A4, A5, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5)
	return
}

func (f FuncPair5[A1, A2, A3, A4, A5, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5)
}

func (f FuncPair5[A1, A2, A3, A4, A5, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5)
	return
}

func (f FuncPair5[A1, A2, A3, A4, A5, R1, R2]) ToDelegate() DelegatePair5[A1, A2, A3, A4, A5, R1, R2] {
	return []FuncPair5[A1, A2, A3, A4, A5, R1, R2]{f}
}

type FuncPair6[A1, A2, A3, A4, A5, A6, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6,
) (R1, R2)

func (f FuncPair6[A1, A2, A3, A4, A5, A6, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6)
	return
}

func (f FuncPair6[A1, A2, A3, A4, A5, A6, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6)
}

func (f FuncPair6[A1, A2, A3, A4, A5, A6, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6)
	return
}

func (f FuncPair6[A1, A2, A3, A4, A5, A6, R1, R2]) ToDelegate() DelegatePair6[A1, A2, A3, A4, A5, A6, R1, R2] {
	return []FuncPair6[A1, A2, A3, A4, A5, A6, R1, R2]{f}
}

type FuncPair7[A1, A2, A3, A4, A5, A6, A7, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7,
) (R1, R2)

func (f FuncPair7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7)
	return
}

func (f FuncPair7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7)
}

func (f FuncPair7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7)
	return
}

func (f FuncPair7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) ToDelegate() DelegatePair7[A1, A2, A3, A4, A5, A6, A7, R1, R2] {
	return []FuncPair7[A1, A2, A3, A4, A5, A6, A7, R1, R2]{f}
}

type FuncPair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8,
) (R1, R2)

func (f FuncPair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8)
	return
}

func (f FuncPair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8)
}

func (f FuncPair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8)
	return
}

func (f FuncPair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) ToDelegate() DelegatePair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2] {
	return []FuncPair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]{f}
}

type FuncPair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
) (R1, R2)

func (f FuncPair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9)
	return
}

func (f FuncPair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9)
}

func (f FuncPair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	return
}

func (f FuncPair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) ToDelegate() DelegatePair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2] {
	return []FuncPair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]{f}
}

type FuncPair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
) (R1, R2)

func (f FuncPair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	return
}

func (f FuncPair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
}

func (f FuncPair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	return
}

func (f FuncPair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) ToDelegate() DelegatePair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2] {
	return []FuncPair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]{f}
}

type FuncPair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
) (R1, R2)

func (f FuncPair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	return
}

func (f FuncPair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
}

func (f FuncPair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	return
}

func (f FuncPair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) ToDelegate() DelegatePair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2] {
	return []FuncPair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]{f}
}

type FuncPair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
) (R1, R2)

func (f FuncPair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	return
}

func (f FuncPair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
}

func (f FuncPair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	return
}

func (f FuncPair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) ToDelegate() DelegatePair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2] {
	return []FuncPair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]{f}
}

type FuncPair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
) (R1, R2)

func (f FuncPair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	return
}

func (f FuncPair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
}

func (f FuncPair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	return
}

func (f FuncPair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) ToDelegate() DelegatePair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2] {
	return []FuncPair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]{f}
}

type FuncPair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
) (R1, R2)

func (f FuncPair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	return
}

func (f FuncPair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
}

func (f FuncPair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	return
}

func (f FuncPair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) ToDelegate() DelegatePair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2] {
	return []FuncPair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]{f}
}

type FuncPair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
) (R1, R2)

func (f FuncPair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	return
}

func (f FuncPair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
}

func (f FuncPair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	return
}

func (f FuncPair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) ToDelegate() DelegatePair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2] {
	return []FuncPair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]{f}
}

type FuncPair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
) (R1, R2)

func (f FuncPair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	return
}

func (f FuncPair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
}

func (f FuncPair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	return
}

func (f FuncPair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) ToDelegate() DelegatePair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2] {
	return []FuncPair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]{f}
}
