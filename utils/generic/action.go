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

type Action0 func()

func (f Action0) Exec() {
	f.Call(false, nil)
}

func (f Action0) Invoke() (panicErr error) {
	return f.Call(true, nil)
}

func (f Action0) Call(autoRecover bool, reportError chan error) (panicErr error) {
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

	f()

	return nil
}

func (f Action0) ToDelegate() DelegateAction0 {
	return []Action0{f}
}

type Action1[A1 any] func(A1)

func (f Action1[A1]) Exec(a1 A1) {
	f.Call(false, nil, a1)
}

func (f Action1[A1]) Invoke(a1 A1) (panicErr error) {
	return f.Call(true, nil, a1)
}

func (f Action1[A1]) Call(autoRecover bool, reportError chan error, a1 A1) (panicErr error) {
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

	f(a1)

	return nil
}

func (f Action1[A1]) ToDelegate() DelegateAction1[A1] {
	return []Action1[A1]{f}
}

type Action2[A1, A2 any] func(A1, A2)

func (f Action2[A1, A2]) Exec(a1 A1, a2 A2) {
	f.Call(false, nil, a1, a2)
}

func (f Action2[A1, A2]) Invoke(a1 A1, a2 A2) (panicErr error) {
	return f.Call(true, nil, a1, a2)
}

func (f Action2[A1, A2]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2) (panicErr error) {
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

	f(a1, a2)

	return nil
}

func (f Action2[A1, A2]) ToDelegate() DelegateAction2[A1, A2] {
	return []Action2[A1, A2]{f}
}

type Action3[A1, A2, A3 any] func(A1, A2, A3)

func (f Action3[A1, A2, A3]) Exec(a1 A1, a2 A2, a3 A3) {
	f.Call(false, nil, a1, a2, a3)
}

func (f Action3[A1, A2, A3]) Invoke(a1 A1, a2 A2, a3 A3) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3)
}

func (f Action3[A1, A2, A3]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3) (panicErr error) {
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

	f(a1, a2, a3)

	return nil
}

func (f Action3[A1, A2, A3]) ToDelegate() DelegateAction3[A1, A2, A3] {
	return []Action3[A1, A2, A3]{f}
}

type Action4[A1, A2, A3, A4 any] func(A1, A2, A3, A4)

func (f Action4[A1, A2, A3, A4]) Exec(a1 A1, a2 A2, a3 A3, a4 A4) {
	f.Call(false, nil, a1, a2, a3, a4)
}

func (f Action4[A1, A2, A3, A4]) Invoke(a1 A1, a2 A2, a3 A3, a4 A4) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4)
}

func (f Action4[A1, A2, A3, A4]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, a4 A4) (panicErr error) {
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

	f(a1, a2, a3, a4)

	return nil
}

func (f Action4[A1, A2, A3, A4]) ToDelegate() DelegateAction4[A1, A2, A3, A4] {
	return []Action4[A1, A2, A3, A4]{f}
}

type Action5[A1, A2, A3, A4, A5 any] func(
	A1, A2, A3, A4, A5,
)

func (f Action5[A1, A2, A3, A4, A5]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5)
}

func (f Action5[A1, A2, A3, A4, A5]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5)
}

func (f Action5[A1, A2, A3, A4, A5]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
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

	f(a1, a2, a3, a4, a5)

	return nil
}

func (f Action5[A1, A2, A3, A4, A5]) ToDelegate() DelegateAction5[A1, A2, A3, A4, A5] {
	return []Action5[A1, A2, A3, A4, A5]{f}
}

type Action6[A1, A2, A3, A4, A5, A6 any] func(
	A1, A2, A3, A4, A5, A6,
)

func (f Action6[A1, A2, A3, A4, A5, A6]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6)
}

func (f Action6[A1, A2, A3, A4, A5, A6]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6)
}

func (f Action6[A1, A2, A3, A4, A5, A6]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
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

	f(a1, a2, a3, a4, a5, a6)

	return nil
}

func (f Action6[A1, A2, A3, A4, A5, A6]) ToDelegate() DelegateAction6[A1, A2, A3, A4, A5, A6] {
	return []Action6[A1, A2, A3, A4, A5, A6]{f}
}

type Action7[A1, A2, A3, A4, A5, A6, A7 any] func(
	A1, A2, A3, A4, A5, A6, A7,
)

func (f Action7[A1, A2, A3, A4, A5, A6, A7]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7)
}

func (f Action7[A1, A2, A3, A4, A5, A6, A7]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7)
}

func (f Action7[A1, A2, A3, A4, A5, A6, A7]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
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

	f(a1, a2, a3, a4, a5, a6, a7)

	return nil
}

func (f Action7[A1, A2, A3, A4, A5, A6, A7]) ToDelegate() DelegateAction7[A1, A2, A3, A4, A5, A6, A7] {
	return []Action7[A1, A2, A3, A4, A5, A6, A7]{f}
}

type Action8[A1, A2, A3, A4, A5, A6, A7, A8 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8,
)

func (f Action8[A1, A2, A3, A4, A5, A6, A7, A8]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8)
}

func (f Action8[A1, A2, A3, A4, A5, A6, A7, A8]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8)
}

func (f Action8[A1, A2, A3, A4, A5, A6, A7, A8]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
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

	f(a1, a2, a3, a4, a5, a6, a7, a8)

	return nil
}

func (f Action8[A1, A2, A3, A4, A5, A6, A7, A8]) ToDelegate() DelegateAction8[A1, A2, A3, A4, A5, A6, A7, A8] {
	return []Action8[A1, A2, A3, A4, A5, A6, A7, A8]{f}
}

type Action9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
)

func (f Action9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9)
}

func (f Action9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9)
}

func (f Action9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9)

	return nil
}

func (f Action9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) ToDelegate() DelegateAction9[A1, A2, A3, A4, A5, A6, A7, A8, A9] {
	return []Action9[A1, A2, A3, A4, A5, A6, A7, A8, A9]{f}
}

type Action10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
)

func (f Action10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
}

func (f Action10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
}

func (f Action10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)

	return nil
}

func (f Action10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) ToDelegate() DelegateAction10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10] {
	return []Action10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]{f}
}

type Action11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
)

func (f Action11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
}

func (f Action11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
}

func (f Action11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)

	return nil
}

func (f Action11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) ToDelegate() DelegateAction11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11] {
	return []Action11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]{f}
}

type Action12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
)

func (f Action12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
}

func (f Action12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
}

func (f Action12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)

	return nil
}

func (f Action12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) ToDelegate() DelegateAction12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12] {
	return []Action12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]{f}
}

type Action13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
)

func (f Action13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
}

func (f Action13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
}

func (f Action13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)

	return nil
}

func (f Action13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) ToDelegate() DelegateAction13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13] {
	return []Action13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]{f}
}

type Action14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
)

func (f Action14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
}

func (f Action14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
}

func (f Action14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)

	return nil
}

func (f Action14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) ToDelegate() DelegateAction14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14] {
	return []Action14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]{f}
}

type Action15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
)

func (f Action15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
}

func (f Action15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
}

func (f Action15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)

	return nil
}

func (f Action15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) ToDelegate() DelegateAction15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15] {
	return []Action15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]{f}
}

type Action16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
)

func (f Action16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
}

func (f Action16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
}

func (f Action16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
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

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)

	return nil
}

func (f Action16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) ToDelegate() DelegateAction16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16] {
	return []Action16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]{f}
}
