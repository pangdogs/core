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

type ActionVar0[VA any] func(...VA)

func (f ActionVar0[VA]) UnsafeCall(args ...VA) {
	f.Call(false, nil, args...)
}

func (f ActionVar0[VA]) SafeCall(args ...VA) (panicErr error) {
	return f.Call(true, nil, args...)
}

func (f ActionVar0[VA]) Call(autoRecover bool, reportError chan error, args ...VA) (panicErr error) {
	if f == nil {
		return nil
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

	f(args...)

	return nil
}

func (f ActionVar0[VA]) ToDelegate() DelegateVoidVar0[VA] {
	return []ActionVar0[VA]{f}
}

type ActionVar1[A1, VA any] func(A1, ...VA)

func (f ActionVar1[A1, VA]) UnsafeCall(a1 A1, args ...VA) {
	f.Call(false, nil, a1, args...)
}

func (f ActionVar1[A1, VA]) SafeCall(a1 A1, args ...VA) (panicErr error) {
	return f.Call(true, nil, a1, args...)
}

func (f ActionVar1[A1, VA]) Call(autoRecover bool, reportError chan error, a1 A1, args ...VA) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, args...)

	return nil
}

func (f ActionVar1[A1, VA]) ToDelegate() DelegateVoidVar1[A1, VA] {
	return []ActionVar1[A1, VA]{f}
}

type ActionVar2[A1, A2, VA any] func(A1, A2, ...VA)

func (f ActionVar2[A1, A2, VA]) UnsafeCall(a1 A1, a2 A2, args ...VA) {
	f.Call(false, nil, a1, a2, args...)
}

func (f ActionVar2[A1, A2, VA]) SafeCall(a1 A1, a2 A2, args ...VA) (panicErr error) {
	return f.Call(true, nil, a1, a2, args...)
}

func (f ActionVar2[A1, A2, VA]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, args ...VA) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, args...)

	return nil
}

func (f ActionVar2[A1, A2, VA]) ToDelegate() DelegateVoidVar2[A1, A2, VA] {
	return []ActionVar2[A1, A2, VA]{f}
}

type ActionVar3[A1, A2, A3, VA any] func(A1, A2, A3, ...VA)

func (f ActionVar3[A1, A2, A3, VA]) UnsafeCall(a1 A1, a2 A2, a3 A3, args ...VA) {
	f.Call(false, nil, a1, a2, a3, args...)
}

func (f ActionVar3[A1, A2, A3, VA]) SafeCall(a1 A1, a2 A2, a3 A3, args ...VA) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, args...)
}

func (f ActionVar3[A1, A2, A3, VA]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, args ...VA) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, a3, args...)

	return nil
}

func (f ActionVar3[A1, A2, A3, VA]) ToDelegate() DelegateVoidVar3[A1, A2, A3, VA] {
	return []ActionVar3[A1, A2, A3, VA]{f}
}

type ActionVar4[A1, A2, A3, A4, VA any] func(A1, A2, A3, A4, ...VA)

func (f ActionVar4[A1, A2, A3, A4, VA]) UnsafeCall(a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) {
	f.Call(false, nil, a1, a2, a3, a4, args...)
}

func (f ActionVar4[A1, A2, A3, A4, VA]) SafeCall(a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, args...)
}

func (f ActionVar4[A1, A2, A3, A4, VA]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, a3, a4, args...)

	return nil
}

func (f ActionVar4[A1, A2, A3, A4, VA]) ToDelegate() DelegateVoidVar4[A1, A2, A3, A4, VA] {
	return []ActionVar4[A1, A2, A3, A4, VA]{f}
}

type ActionVar5[A1, A2, A3, A4, A5, VA any] func(
	A1, A2, A3, A4, A5, ...VA,
)

func (f ActionVar5[A1, A2, A3, A4, A5, VA]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, args...)
}

func (f ActionVar5[A1, A2, A3, A4, A5, VA]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, args...)
}

func (f ActionVar5[A1, A2, A3, A4, A5, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, a3, a4, a5, args...)

	return nil
}

func (f ActionVar5[A1, A2, A3, A4, A5, VA]) ToDelegate() DelegateVoidVar5[A1, A2, A3, A4, A5, VA] {
	return []ActionVar5[A1, A2, A3, A4, A5, VA]{f}
}

type ActionVar6[A1, A2, A3, A4, A5, A6, VA any] func(
	A1, A2, A3, A4, A5, A6, ...VA,
)

func (f ActionVar6[A1, A2, A3, A4, A5, A6, VA]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, args...)
}

func (f ActionVar6[A1, A2, A3, A4, A5, A6, VA]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, args...)
}

func (f ActionVar6[A1, A2, A3, A4, A5, A6, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, a3, a4, a5, a6, args...)

	return nil
}

func (f ActionVar6[A1, A2, A3, A4, A5, A6, VA]) ToDelegate() DelegateVoidVar6[A1, A2, A3, A4, A5, A6, VA] {
	return []ActionVar6[A1, A2, A3, A4, A5, A6, VA]{f}
}

type ActionVar7[A1, A2, A3, A4, A5, A6, A7, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
)

func (f ActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, args...)
}

func (f ActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, args...)
}

func (f ActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, a3, a4, a5, a6, a7, args...)

	return nil
}

func (f ActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]) ToDelegate() DelegateVoidVar7[A1, A2, A3, A4, A5, A6, A7, VA] {
	return []ActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]{f}
}

type ActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
)

func (f ActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, args...)
}

func (f ActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, args...)
}

func (f ActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, args...)

	return nil
}

func (f ActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]) ToDelegate() DelegateVoidVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA] {
	return []ActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]{f}
}

type ActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
)

func (f ActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
}

func (f ActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
}

func (f ActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)

	return nil
}

func (f ActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]) ToDelegate() DelegateVoidVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA] {
	return []ActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]{f}
}

type ActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
)

func (f ActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
}

func (f ActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
}

func (f ActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)

	return nil
}

func (f ActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]) ToDelegate() DelegateVoidVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA] {
	return []ActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]{f}
}

type ActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
)

func (f ActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
}

func (f ActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
}

func (f ActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)

	return nil
}

func (f ActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]) ToDelegate() DelegateVoidVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA] {
	return []ActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]{f}
}

type ActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
)

func (f ActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
}

func (f ActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
}

func (f ActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)

	return nil
}

func (f ActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]) ToDelegate() DelegateVoidVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA] {
	return []ActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]{f}
}

type ActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
)

func (f ActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
}

func (f ActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
}

func (f ActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)

	return nil
}

func (f ActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]) ToDelegate() DelegateVoidVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA] {
	return []ActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]{f}
}

type ActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
)

func (f ActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
}

func (f ActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
}

func (f ActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)

	return nil
}

func (f ActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]) ToDelegate() DelegateVoidVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA] {
	return []ActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]{f}
}

type ActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
)

func (f ActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
}

func (f ActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
}

func (f ActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)

	return nil
}

func (f ActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]) ToDelegate() DelegateVoidVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA] {
	return []ActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]{f}
}

type ActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
)

func (f ActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]) UnsafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
}

func (f ActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]) SafeCall(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
}

func (f ActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)

	return nil
}

func (f ActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]) ToDelegate() DelegateVoidVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA] {
	return []ActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]{f}
}
