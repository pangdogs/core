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

type FuncVar0[VA, R any] func(...VA) R

func (f FuncVar0[VA, R]) UnsafeCall(args ...VA) (r R) {
	r, _ = f.Call(false, nil, args...)
	return
}

func (f FuncVar0[VA, R]) SafeCall(args ...VA) (r R, panicErr error) {
	return f.Call(true, nil, args...)
}

func (f FuncVar0[VA, R]) Call(autoRecover bool, reportError chan error, args ...VA) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				select {
				case reportError <- exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)):
				default:
				}
			}
		}()
	}

	return f(args...), nil
}

func (f FuncVar0[VA, R]) ToDelegate() DelegateVar0[VA, R] {
	return []FuncVar0[VA, R]{f}
}

type FuncVar1[A1, VA, R any] func(A1, ...VA) R

func (f FuncVar1[A1, VA, R]) UnsafeCall(a1 A1, args ...VA) (r R) {
	r, _ = f.Call(false, nil, a1, args...)
	return
}

func (f FuncVar1[A1, VA, R]) SafeCall(a1 A1, args ...VA) (r R, panicErr error) {
	return f.Call(true, nil, a1, args...)
}

func (f FuncVar1[A1, VA, R]) Call(autoRecover bool, reportError chan error, a1 A1, args ...VA) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				select {
				case reportError <- exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)):
				default:
				}
			}
		}()
	}

	return f(a1, args...), nil
}

func (f FuncVar1[A1, VA, R]) ToDelegate() DelegateVar1[A1, VA, R] {
	return []FuncVar1[A1, VA, R]{f}
}

type FuncVar2[A1, A2, VA, R any] func(A1, A2, ...VA) R

func (f FuncVar2[A1, A2, VA, R]) UnsafeCall(a1 A1, a2 A2, args ...VA) (r R) {
	r, _ = f.Call(false, nil, a1, a2, args...)
	return
}

func (f FuncVar2[A1, A2, VA, R]) SafeCall(a1 A1, a2 A2, args ...VA) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, args...)
}

func (f FuncVar2[A1, A2, VA, R]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, args ...VA) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				select {
				case reportError <- exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)):
				default:
				}
			}
		}()
	}

	return f(a1, a2, args...), nil
}

func (f FuncVar2[A1, A2, VA, R]) ToDelegate() DelegateVar2[A1, A2, VA, R] {
	return []FuncVar2[A1, A2, VA, R]{f}
}

type FuncVar3[A1, A2, A3, VA, R any] func(A1, A2, A3, ...VA) R

func (f FuncVar3[A1, A2, A3, VA, R]) UnsafeCall(a1 A1, a2 A2, a3 A3, args ...VA) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, args...)
	return
}

func (f FuncVar3[A1, A2, A3, VA, R]) SafeCall(a1 A1, a2 A2, a3 A3, args ...VA) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, args...)
}

func (f FuncVar3[A1, A2, A3, VA, R]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, args ...VA) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				select {
				case reportError <- exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)):
				default:
				}
			}
		}()
	}

	return f(a1, a2, a3, args...), nil
}

func (f FuncVar3[A1, A2, A3, VA, R]) ToDelegate() DelegateVar3[A1, A2, A3, VA, R] {
	return []FuncVar3[A1, A2, A3, VA, R]{f}
}

type FuncVar4[A1, A2, A3, A4, VA, R any] func(A1, A2, A3, A4, ...VA) R

func (f FuncVar4[A1, A2, A3, A4, VA, R]) UnsafeCall(a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, args...)
	return
}

func (f FuncVar4[A1, A2, A3, A4, VA, R]) SafeCall(a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, args...)
}

func (f FuncVar4[A1, A2, A3, A4, VA, R]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				select {
				case reportError <- exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)):
				default:
				}
			}
		}()
	}

	return f(a1, a2, a3, a4, args...), nil
}

func (f FuncVar4[A1, A2, A3, A4, VA, R]) ToDelegate() DelegateVar4[A1, A2, A3, A4, VA, R] {
	return []FuncVar4[A1, A2, A3, A4, VA, R]{f}
}

type FuncVar5[A1, A2, A3, A4, A5, VA, R any] func(
	A1, A2, A3, A4, A5, ...VA,
) R

func (f FuncVar5[A1, A2, A3, A4, A5, VA, R]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, args...)
	return
}

func (f FuncVar5[A1, A2, A3, A4, A5, VA, R]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, args...)
}

func (f FuncVar5[A1, A2, A3, A4, A5, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				select {
				case reportError <- exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)):
				default:
				}
			}
		}()
	}

	return f(a1, a2, a3, a4, a5, args...), nil
}

func (f FuncVar5[A1, A2, A3, A4, A5, VA, R]) ToDelegate() DelegateVar5[A1, A2, A3, A4, A5, VA, R] {
	return []FuncVar5[A1, A2, A3, A4, A5, VA, R]{f}
}

type FuncVar6[A1, A2, A3, A4, A5, A6, VA, R any] func(
	A1, A2, A3, A4, A5, A6, ...VA,
) R

func (f FuncVar6[A1, A2, A3, A4, A5, A6, VA, R]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, args...)
	return
}

func (f FuncVar6[A1, A2, A3, A4, A5, A6, VA, R]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, args...)
}

func (f FuncVar6[A1, A2, A3, A4, A5, A6, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				select {
				case reportError <- exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)):
				default:
				}
			}
		}()
	}

	return f(a1, a2, a3, a4, a5, a6, args...), nil
}

func (f FuncVar6[A1, A2, A3, A4, A5, A6, VA, R]) ToDelegate() DelegateVar6[A1, A2, A3, A4, A5, A6, VA, R] {
	return []FuncVar6[A1, A2, A3, A4, A5, A6, VA, R]{f}
}

type FuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
) R

func (f FuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, args...)
	return
}

func (f FuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, args...)
}

func (f FuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				select {
				case reportError <- exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)):
				default:
				}
			}
		}()
	}

	return f(a1, a2, a3, a4, a5, a6, a7, args...), nil
}

func (f FuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R]) ToDelegate() DelegateVar7[A1, A2, A3, A4, A5, A6, A7, VA, R] {
	return []FuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R]{f}
}

type FuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
) R

func (f FuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, args...)
	return
}

func (f FuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, args...)
}

func (f FuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				select {
				case reportError <- exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)):
				default:
				}
			}
		}()
	}

	return f(a1, a2, a3, a4, a5, a6, a7, a8, args...), nil
}

func (f FuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R]) ToDelegate() DelegateVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R] {
	return []FuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R]{f}
}

type FuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
) R

func (f FuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
	return
}

func (f FuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
}

func (f FuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				select {
				case reportError <- exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)):
				default:
				}
			}
		}()
	}

	return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, args...), nil
}

func (f FuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R]) ToDelegate() DelegateVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R] {
	return []FuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R]{f}
}

type FuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
) R

func (f FuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
	return
}

func (f FuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
}

func (f FuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				select {
				case reportError <- exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)):
				default:
				}
			}
		}()
	}

	return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...), nil
}

func (f FuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R]) ToDelegate() DelegateVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R] {
	return []FuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R]{f}
}

type FuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
) R

func (f FuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
	return
}

func (f FuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
}

func (f FuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				select {
				case reportError <- exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)):
				default:
				}
			}
		}()
	}

	return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...), nil
}

func (f FuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R]) ToDelegate() DelegateVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R] {
	return []FuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R]{f}
}

type FuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
) R

func (f FuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
	return
}

func (f FuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
}

func (f FuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				select {
				case reportError <- exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)):
				default:
				}
			}
		}()
	}

	return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...), nil
}

func (f FuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R]) ToDelegate() DelegateVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R] {
	return []FuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R]{f}
}

type FuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
) R

func (f FuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
	return
}

func (f FuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
}

func (f FuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				select {
				case reportError <- exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)):
				default:
				}
			}
		}()
	}

	return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...), nil
}

func (f FuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R]) ToDelegate() DelegateVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R] {
	return []FuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R]{f}
}

type FuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
) R

func (f FuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
	return
}

func (f FuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
}

func (f FuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				select {
				case reportError <- exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)):
				default:
				}
			}
		}()
	}

	return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...), nil
}

func (f FuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R]) ToDelegate() DelegateVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R] {
	return []FuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R]{f}
}

type FuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
) R

func (f FuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
	return
}

func (f FuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
}

func (f FuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
	}

	if autoRecover {
		defer func() {
			if panicErr = types.Panic2Err(recover()); panicErr != nil {
				select {
				case reportError <- exception.TraceStack(fmt.Errorf("%w: %w", exception.ErrPanicked, panicErr)):
				default:
				}
			}
		}()
	}

	return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...), nil
}

func (f FuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R]) ToDelegate() DelegateVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R] {
	return []FuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R]{f}
}

type FuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
) R

func (f FuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
	return
}

func (f FuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
}

func (f FuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.ZeroT[R](), nil
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

	return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...), nil
}

func (f FuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R]) ToDelegate() DelegateVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R] {
	return []FuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R]{f}
}
