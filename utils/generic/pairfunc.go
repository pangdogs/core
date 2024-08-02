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

type PairFunc0[R1, R2 any] func() (R1, R2)

func (f PairFunc0[R1, R2]) Exec() (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil)
	return
}

func (f PairFunc0[R1, R2]) Invoke() (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil)
}

func (f PairFunc0[R1, R2]) Call(autoRecover bool, reportError chan error) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f()
	return
}

func (f PairFunc0[R1, R2]) CastDelegate() DelegatePairFunc0[R1, R2] {
	return []PairFunc0[R1, R2]{f}
}

type PairFunc1[A1, R1, R2 any] func(A1) (R1, R2)

func (f PairFunc1[A1, R1, R2]) Exec(a1 A1) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1)
	return
}

func (f PairFunc1[A1, R1, R2]) Invoke(a1 A1) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1)
}

func (f PairFunc1[A1, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1)
	return
}

func (f PairFunc1[A1, R1, R2]) CastDelegate() DelegatePairFunc1[A1, R1, R2] {
	return []PairFunc1[A1, R1, R2]{f}
}

type PairFunc2[A1, A2, R1, R2 any] func(A1, A2) (R1, R2)

func (f PairFunc2[A1, A2, R1, R2]) Exec(a1 A1, a2 A2) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2)
	return
}

func (f PairFunc2[A1, A2, R1, R2]) Invoke(a1 A1, a2 A2) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2)
}

func (f PairFunc2[A1, A2, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2)
	return
}

func (f PairFunc2[A1, A2, R1, R2]) CastDelegate() DelegatePairFunc2[A1, A2, R1, R2] {
	return []PairFunc2[A1, A2, R1, R2]{f}
}

type PairFunc3[A1, A2, A3, R1, R2 any] func(A1, A2, A3) (R1, R2)

func (f PairFunc3[A1, A2, A3, R1, R2]) Exec(a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3)
	return
}

func (f PairFunc3[A1, A2, A3, R1, R2]) Invoke(a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3)
}

func (f PairFunc3[A1, A2, A3, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3)
	return
}

func (f PairFunc3[A1, A2, A3, R1, R2]) CastDelegate() DelegatePairFunc3[A1, A2, A3, R1, R2] {
	return []PairFunc3[A1, A2, A3, R1, R2]{f}
}

type PairFunc4[A1, A2, A3, A4, R1, R2 any] func(A1, A2, A3, A4) (R1, R2)

func (f PairFunc4[A1, A2, A3, A4, R1, R2]) Exec(a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4)
	return
}

func (f PairFunc4[A1, A2, A3, A4, R1, R2]) Invoke(a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4)
}

func (f PairFunc4[A1, A2, A3, A4, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.ZeroT[R1](), types.ZeroT[R2](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				panicErr = fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4)
	return
}

func (f PairFunc4[A1, A2, A3, A4, R1, R2]) CastDelegate() DelegatePairFunc4[A1, A2, A3, A4, R1, R2] {
	return []PairFunc4[A1, A2, A3, A4, R1, R2]{f}
}

type PairFunc5[A1, A2, A3, A4, A5, R1, R2 any] func(
	A1, A2, A3, A4, A5,
) (R1, R2)

func (f PairFunc5[A1, A2, A3, A4, A5, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5)
	return
}

func (f PairFunc5[A1, A2, A3, A4, A5, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5)
}

func (f PairFunc5[A1, A2, A3, A4, A5, R1, R2]) Call(
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5)
	return
}

func (f PairFunc5[A1, A2, A3, A4, A5, R1, R2]) CastDelegate() DelegatePairFunc5[A1, A2, A3, A4, A5, R1, R2] {
	return []PairFunc5[A1, A2, A3, A4, A5, R1, R2]{f}
}

type PairFunc6[A1, A2, A3, A4, A5, A6, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6,
) (R1, R2)

func (f PairFunc6[A1, A2, A3, A4, A5, A6, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6)
	return
}

func (f PairFunc6[A1, A2, A3, A4, A5, A6, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6)
}

func (f PairFunc6[A1, A2, A3, A4, A5, A6, R1, R2]) Call(
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6)
	return
}

func (f PairFunc6[A1, A2, A3, A4, A5, A6, R1, R2]) CastDelegate() DelegatePairFunc6[A1, A2, A3, A4, A5, A6, R1, R2] {
	return []PairFunc6[A1, A2, A3, A4, A5, A6, R1, R2]{f}
}

type PairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7,
) (R1, R2)

func (f PairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7)
	return
}

func (f PairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7)
}

func (f PairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) Call(
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7)
	return
}

func (f PairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) CastDelegate() DelegatePairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2] {
	return []PairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2]{f}
}

type PairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8,
) (R1, R2)

func (f PairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8)
	return
}

func (f PairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8)
}

func (f PairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) Call(
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8)
	return
}

func (f PairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) CastDelegate() DelegatePairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2] {
	return []PairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]{f}
}

type PairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
) (R1, R2)

func (f PairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9)
	return
}

func (f PairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9)
}

func (f PairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) Call(
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	return
}

func (f PairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) CastDelegate() DelegatePairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2] {
	return []PairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]{f}
}

type PairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
) (R1, R2)

func (f PairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	return
}

func (f PairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
}

func (f PairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) Call(
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	return
}

func (f PairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) CastDelegate() DelegatePairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2] {
	return []PairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]{f}
}

type PairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
) (R1, R2)

func (f PairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	return
}

func (f PairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
}

func (f PairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) Call(
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	return
}

func (f PairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) CastDelegate() DelegatePairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2] {
	return []PairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]{f}
}

type PairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
) (R1, R2)

func (f PairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	return
}

func (f PairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
}

func (f PairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) Call(
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	return
}

func (f PairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) CastDelegate() DelegatePairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2] {
	return []PairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]{f}
}

type PairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
) (R1, R2)

func (f PairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	return
}

func (f PairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
}

func (f PairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) Call(
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	return
}

func (f PairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) CastDelegate() DelegatePairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2] {
	return []PairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]{f}
}

type PairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
) (R1, R2)

func (f PairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	return
}

func (f PairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
}

func (f PairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) Call(
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	return
}

func (f PairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) CastDelegate() DelegatePairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2] {
	return []PairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]{f}
}

type PairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
) (R1, R2)

func (f PairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	return
}

func (f PairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
}

func (f PairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) Call(
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	return
}

func (f PairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) CastDelegate() DelegatePairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2] {
	return []PairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]{f}
}

type PairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
) (R1, R2)

func (f PairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	return
}

func (f PairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
}

func (f PairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) Call(
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	return
}

func (f PairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) CastDelegate() DelegatePairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2] {
	return []PairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]{f}
}
