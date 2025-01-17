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

type FuncPairVar0[VA, R1, R2 any] func(...VA) (R1, R2)

func (f FuncPairVar0[VA, R1, R2]) UnsafeCall(args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, args...)
	return
}

func (f FuncPairVar0[VA, R1, R2]) SafeCall(args ...VA) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, args...)
}

func (f FuncPairVar0[VA, R1, R2]) Call(autoRecover bool, reportError chan error, args ...VA) (r1 R1, r2 R2, panicErr error) {
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

	r1, r2 = f(args...)
	return
}

func (f FuncPairVar0[VA, R1, R2]) ToDelegate() DelegatePairVar0[VA, R1, R2] {
	return []FuncPairVar0[VA, R1, R2]{f}
}

type FuncPairVar1[A1, VA, R1, R2 any] func(A1, ...VA) (R1, R2)

func (f FuncPairVar1[A1, VA, R1, R2]) UnsafeCall(a1 A1, args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, args...)
	return
}

func (f FuncPairVar1[A1, VA, R1, R2]) SafeCall(a1 A1, args ...VA) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, args...)
}

func (f FuncPairVar1[A1, VA, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, args ...VA) (r1 R1, r2 R2, panicErr error) {
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

	r1, r2 = f(a1, args...)
	return
}

func (f FuncPairVar1[A1, VA, R1, R2]) ToDelegate() DelegatePairVar1[A1, VA, R1, R2] {
	return []FuncPairVar1[A1, VA, R1, R2]{f}
}

type FuncPairVar2[A1, A2, VA, R1, R2 any] func(A1, A2, ...VA) (R1, R2)

func (f FuncPairVar2[A1, A2, VA, R1, R2]) UnsafeCall(a1 A1, a2 A2, args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, args...)
	return
}

func (f FuncPairVar2[A1, A2, VA, R1, R2]) SafeCall(a1 A1, a2 A2, args ...VA) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, args...)
}

func (f FuncPairVar2[A1, A2, VA, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, args ...VA) (r1 R1, r2 R2, panicErr error) {
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

	r1, r2 = f(a1, a2, args...)
	return
}

func (f FuncPairVar2[A1, A2, VA, R1, R2]) ToDelegate() DelegatePairVar2[A1, A2, VA, R1, R2] {
	return []FuncPairVar2[A1, A2, VA, R1, R2]{f}
}

type FuncPairVar3[A1, A2, A3, VA, R1, R2 any] func(A1, A2, A3, ...VA) (R1, R2)

func (f FuncPairVar3[A1, A2, A3, VA, R1, R2]) UnsafeCall(a1 A1, a2 A2, a3 A3, args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, args...)
	return
}

func (f FuncPairVar3[A1, A2, A3, VA, R1, R2]) SafeCall(a1 A1, a2 A2, a3 A3, args ...VA) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, args...)
}

func (f FuncPairVar3[A1, A2, A3, VA, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, args ...VA) (r1 R1, r2 R2, panicErr error) {
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

	r1, r2 = f(a1, a2, a3, args...)
	return
}

func (f FuncPairVar3[A1, A2, A3, VA, R1, R2]) ToDelegate() DelegatePairVar3[A1, A2, A3, VA, R1, R2] {
	return []FuncPairVar3[A1, A2, A3, VA, R1, R2]{f}
}

type FuncPairVar4[A1, A2, A3, A4, VA, R1, R2 any] func(A1, A2, A3, A4, ...VA) (R1, R2)

func (f FuncPairVar4[A1, A2, A3, A4, VA, R1, R2]) UnsafeCall(a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, args...)
	return
}

func (f FuncPairVar4[A1, A2, A3, A4, VA, R1, R2]) SafeCall(a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, args...)
}

func (f FuncPairVar4[A1, A2, A3, A4, VA, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r1 R1, r2 R2, panicErr error) {
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

	r1, r2 = f(a1, a2, a3, a4, args...)
	return
}

func (f FuncPairVar4[A1, A2, A3, A4, VA, R1, R2]) ToDelegate() DelegatePairVar4[A1, A2, A3, A4, VA, R1, R2] {
	return []FuncPairVar4[A1, A2, A3, A4, VA, R1, R2]{f}
}

type FuncPairVar5[A1, A2, A3, A4, A5, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, ...VA,
) (R1, R2)

func (f FuncPairVar5[A1, A2, A3, A4, A5, VA, R1, R2]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, args...)
	return
}

func (f FuncPairVar5[A1, A2, A3, A4, A5, VA, R1, R2]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, args...)
}

func (f FuncPairVar5[A1, A2, A3, A4, A5, VA, R1, R2]) Call(
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
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, args...)
	return
}

func (f FuncPairVar5[A1, A2, A3, A4, A5, VA, R1, R2]) ToDelegate() DelegatePairVar5[A1, A2, A3, A4, A5, VA, R1, R2] {
	return []FuncPairVar5[A1, A2, A3, A4, A5, VA, R1, R2]{f}
}

type FuncPairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, ...VA,
) (R1, R2)

func (f FuncPairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, args...)
	return
}

func (f FuncPairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, args...)
}

func (f FuncPairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) Call(
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
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, args...)
	return
}

func (f FuncPairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) ToDelegate() DelegatePairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2] {
	return []FuncPairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]{f}
}

type FuncPairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
) (R1, R2)

func (f FuncPairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, args...)
	return
}

func (f FuncPairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, args...)
}

func (f FuncPairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) Call(
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
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, args...)
	return
}

func (f FuncPairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) ToDelegate() DelegatePairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2] {
	return []FuncPairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]{f}
}

type FuncPairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
) (R1, R2)

func (f FuncPairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, args...)
	return
}

func (f FuncPairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, args...)
}

func (f FuncPairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) Call(
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
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, args...)
	return
}

func (f FuncPairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) ToDelegate() DelegatePairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2] {
	return []FuncPairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]{f}
}

type FuncPairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
) (R1, R2)

func (f FuncPairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
	return
}

func (f FuncPairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
}

func (f FuncPairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) Call(
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
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
	return
}

func (f FuncPairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) ToDelegate() DelegatePairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2] {
	return []FuncPairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]{f}
}

type FuncPairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
) (R1, R2)

func (f FuncPairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
	return
}

func (f FuncPairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
}

func (f FuncPairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) Call(
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
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
	return
}

func (f FuncPairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) ToDelegate() DelegatePairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2] {
	return []FuncPairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]{f}
}

type FuncPairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
) (R1, R2)

func (f FuncPairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
	return
}

func (f FuncPairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
}

func (f FuncPairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) Call(
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
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
	return
}

func (f FuncPairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) ToDelegate() DelegatePairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2] {
	return []FuncPairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]{f}
}

type FuncPairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
) (R1, R2)

func (f FuncPairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
	return
}

func (f FuncPairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
}

func (f FuncPairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) Call(
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
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
	return
}

func (f FuncPairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) ToDelegate() DelegatePairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2] {
	return []FuncPairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]{f}
}

type FuncPairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
) (R1, R2)

func (f FuncPairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
	return
}

func (f FuncPairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
}

func (f FuncPairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) Call(
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
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
	return
}

func (f FuncPairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) ToDelegate() DelegatePairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2] {
	return []FuncPairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]{f}
}

type FuncPairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
) (R1, R2)

func (f FuncPairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
	return
}

func (f FuncPairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
}

func (f FuncPairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) Call(
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
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
	return
}

func (f FuncPairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) ToDelegate() DelegatePairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2] {
	return []FuncPairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]{f}
}

type FuncPairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
) (R1, R2)

func (f FuncPairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
	return
}

func (f FuncPairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
}

func (f FuncPairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) Call(
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
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
	return
}

func (f FuncPairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) ToDelegate() DelegatePairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2] {
	return []FuncPairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]{f}
}

type FuncPairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
) (R1, R2)

func (f FuncPairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
	return
}

func (f FuncPairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
}

func (f FuncPairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) Call(
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
					case reportError <- exception.TraceStack(panicErr):
					default:
					}
				}
			}
		}()
	}

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
	return
}

func (f FuncPairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) ToDelegate() DelegatePairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2] {
	return []FuncPairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]{f}
}
