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

type PairFuncVar0[VA, R1, R2 any] func(...VA) (R1, R2)

func (f PairFuncVar0[VA, R1, R2]) Exec(args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, args...)
	return
}

func (f PairFuncVar0[VA, R1, R2]) Invoke(args ...VA) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, args...)
}

func (f PairFuncVar0[VA, R1, R2]) Call(autoRecover bool, reportError chan error, args ...VA) (r1 R1, r2 R2, panicErr error) {
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

	r1, r2 = f(args...)
	return
}

func (f PairFuncVar0[VA, R1, R2]) CastDelegate() DelegatePairFuncVar0[VA, R1, R2] {
	return []PairFuncVar0[VA, R1, R2]{f}
}

type PairFuncVar1[A1, VA, R1, R2 any] func(A1, ...VA) (R1, R2)

func (f PairFuncVar1[A1, VA, R1, R2]) Exec(a1 A1, args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, args...)
	return
}

func (f PairFuncVar1[A1, VA, R1, R2]) Invoke(a1 A1, args ...VA) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, args...)
}

func (f PairFuncVar1[A1, VA, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, args ...VA) (r1 R1, r2 R2, panicErr error) {
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

	r1, r2 = f(a1, args...)
	return
}

func (f PairFuncVar1[A1, VA, R1, R2]) CastDelegate() DelegatePairFuncVar1[A1, VA, R1, R2] {
	return []PairFuncVar1[A1, VA, R1, R2]{f}
}

type PairFuncVar2[A1, A2, VA, R1, R2 any] func(A1, A2, ...VA) (R1, R2)

func (f PairFuncVar2[A1, A2, VA, R1, R2]) Exec(a1 A1, a2 A2, args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, args...)
	return
}

func (f PairFuncVar2[A1, A2, VA, R1, R2]) Invoke(a1 A1, a2 A2, args ...VA) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, args...)
}

func (f PairFuncVar2[A1, A2, VA, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, args ...VA) (r1 R1, r2 R2, panicErr error) {
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

	r1, r2 = f(a1, a2, args...)
	return
}

func (f PairFuncVar2[A1, A2, VA, R1, R2]) CastDelegate() DelegatePairFuncVar2[A1, A2, VA, R1, R2] {
	return []PairFuncVar2[A1, A2, VA, R1, R2]{f}
}

type PairFuncVar3[A1, A2, A3, VA, R1, R2 any] func(A1, A2, A3, ...VA) (R1, R2)

func (f PairFuncVar3[A1, A2, A3, VA, R1, R2]) Exec(a1 A1, a2 A2, a3 A3, args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, args...)
	return
}

func (f PairFuncVar3[A1, A2, A3, VA, R1, R2]) Invoke(a1 A1, a2 A2, a3 A3, args ...VA) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, args...)
}

func (f PairFuncVar3[A1, A2, A3, VA, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, args ...VA) (r1 R1, r2 R2, panicErr error) {
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

	r1, r2 = f(a1, a2, a3, args...)
	return
}

func (f PairFuncVar3[A1, A2, A3, VA, R1, R2]) CastDelegate() DelegatePairFuncVar3[A1, A2, A3, VA, R1, R2] {
	return []PairFuncVar3[A1, A2, A3, VA, R1, R2]{f}
}

type PairFuncVar4[A1, A2, A3, A4, VA, R1, R2 any] func(A1, A2, A3, A4, ...VA) (R1, R2)

func (f PairFuncVar4[A1, A2, A3, A4, VA, R1, R2]) Exec(a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, args...)
	return
}

func (f PairFuncVar4[A1, A2, A3, A4, VA, R1, R2]) Invoke(a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, args...)
}

func (f PairFuncVar4[A1, A2, A3, A4, VA, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r1 R1, r2 R2, panicErr error) {
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

	r1, r2 = f(a1, a2, a3, a4, args...)
	return
}

func (f PairFuncVar4[A1, A2, A3, A4, VA, R1, R2]) CastDelegate() DelegatePairFuncVar4[A1, A2, A3, A4, VA, R1, R2] {
	return []PairFuncVar4[A1, A2, A3, A4, VA, R1, R2]{f}
}

type PairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, ...VA,
) (R1, R2)

func (f PairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, args...)
	return
}

func (f PairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, args...)
}

func (f PairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
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

	r1, r2 = f(a1, a2, a3, a4, a5, args...)
	return
}

func (f PairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2]) CastDelegate() DelegatePairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2] {
	return []PairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2]{f}
}

type PairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, ...VA,
) (R1, R2)

func (f PairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, args...)
	return
}

func (f PairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, args...)
}

func (f PairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, args...)
	return
}

func (f PairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) CastDelegate() DelegatePairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2] {
	return []PairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]{f}
}

type PairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
) (R1, R2)

func (f PairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, args...)
	return
}

func (f PairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, args...)
}

func (f PairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, args...)
	return
}

func (f PairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) CastDelegate() DelegatePairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2] {
	return []PairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]{f}
}

type PairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
) (R1, R2)

func (f PairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, args...)
	return
}

func (f PairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, args...)
}

func (f PairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, args...)
	return
}

func (f PairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) CastDelegate() DelegatePairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2] {
	return []PairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]{f}
}

type PairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
) (R1, R2)

func (f PairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
	return
}

func (f PairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
}

func (f PairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
	return
}

func (f PairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) CastDelegate() DelegatePairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2] {
	return []PairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]{f}
}

type PairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
) (R1, R2)

func (f PairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
	return
}

func (f PairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
}

func (f PairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
	return
}

func (f PairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) CastDelegate() DelegatePairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2] {
	return []PairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]{f}
}

type PairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
) (R1, R2)

func (f PairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
	return
}

func (f PairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
}

func (f PairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
	return
}

func (f PairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) CastDelegate() DelegatePairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2] {
	return []PairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]{f}
}

type PairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
) (R1, R2)

func (f PairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
	return
}

func (f PairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
}

func (f PairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
	return
}

func (f PairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) CastDelegate() DelegatePairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2] {
	return []PairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]{f}
}

type PairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
) (R1, R2)

func (f PairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
	return
}

func (f PairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
}

func (f PairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
	return
}

func (f PairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) CastDelegate() DelegatePairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2] {
	return []PairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]{f}
}

type PairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
) (R1, R2)

func (f PairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
	return
}

func (f PairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
}

func (f PairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
	return
}

func (f PairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) CastDelegate() DelegatePairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2] {
	return []PairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]{f}
}

type PairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
) (R1, R2)

func (f PairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
	return
}

func (f PairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
}

func (f PairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
	return
}

func (f PairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) CastDelegate() DelegatePairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2] {
	return []PairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]{f}
}

type PairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
) (R1, R2)

func (f PairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
	return
}

func (f PairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
}

func (f PairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
	return
}

func (f PairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) CastDelegate() DelegatePairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2] {
	return []PairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]{f}
}
