package generic

import (
	"fmt"
	"git.golaxy.org/core/internal/exception"
	"git.golaxy.org/core/util/types"
)

type PairFuncVar0[VA, R1, R2 any] func(...VA) (R1, R2)

func (f PairFuncVar0[VA, R1, R2]) Exec(va ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, va...)
	return
}

func (f PairFuncVar0[VA, R1, R2]) Invoke(va ...VA) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, va...)
}

func (f PairFuncVar0[VA, R1, R2]) Call(autoRecover bool, reportError chan error, va ...VA) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(va...)
	return
}

type PairFuncVar1[A1, VA, R1, R2 any] func(A1, ...VA) (R1, R2)

func (f PairFuncVar1[A1, VA, R1, R2]) Exec(a1 A1, va ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, va...)
	return
}

func (f PairFuncVar1[A1, VA, R1, R2]) Invoke(a1 A1, va ...VA) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, va...)
}

func (f PairFuncVar1[A1, VA, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, va ...VA) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, va...)
	return
}

type PairFuncVar2[A1, A2, VA, R1, R2 any] func(A1, A2, ...VA) (R1, R2)

func (f PairFuncVar2[A1, A2, VA, R1, R2]) Exec(a1 A1, a2 A2, va ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, va...)
	return
}

func (f PairFuncVar2[A1, A2, VA, R1, R2]) Invoke(a1 A1, a2 A2, va ...VA) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, va...)
}

func (f PairFuncVar2[A1, A2, VA, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, va ...VA) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, va...)
	return
}

type PairFuncVar3[A1, A2, A3, VA, R1, R2 any] func(A1, A2, A3, ...VA) (R1, R2)

func (f PairFuncVar3[A1, A2, A3, VA, R1, R2]) Exec(a1 A1, a2 A2, a3 A3, va ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, va...)
	return
}

func (f PairFuncVar3[A1, A2, A3, VA, R1, R2]) Invoke(a1 A1, a2 A2, a3 A3, va ...VA) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, va...)
}

func (f PairFuncVar3[A1, A2, A3, VA, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, va ...VA) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, va...)
	return
}

type PairFuncVar4[A1, A2, A3, A4, VA, R1, R2 any] func(A1, A2, A3, A4, ...VA) (R1, R2)

func (f PairFuncVar4[A1, A2, A3, A4, VA, R1, R2]) Exec(a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, va...)
	return
}

func (f PairFuncVar4[A1, A2, A3, A4, VA, R1, R2]) Invoke(a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, va...)
}

func (f PairFuncVar4[A1, A2, A3, A4, VA, R1, R2]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, va...)
	return
}

type PairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, ...VA,
) (R1, R2)

func (f PairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, va...)
	return
}

func (f PairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, va...)
}

func (f PairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, va...)
	return
}

type PairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, ...VA,
) (R1, R2)

func (f PairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, va...)
	return
}

func (f PairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, va...)
}

func (f PairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, va...)
	return
}

type PairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
) (R1, R2)

func (f PairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, va...)
	return
}

func (f PairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, va...)
}

func (f PairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, va...)
	return
}

type PairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
) (R1, R2)

func (f PairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, va...)
	return
}

func (f PairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, va...)
}

func (f PairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, va...)
	return
}

type PairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
) (R1, R2)

func (f PairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
	return
}

func (f PairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
}

func (f PairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
	return
}

type PairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
) (R1, R2)

func (f PairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
	return
}

func (f PairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
}

func (f PairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
	return
}

type PairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
) (R1, R2)

func (f PairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
	return
}

func (f PairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
}

func (f PairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
	return
}

type PairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
) (R1, R2)

func (f PairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
	return
}

func (f PairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
}

func (f PairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
	return
}

type PairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
) (R1, R2)

func (f PairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
	return
}

func (f PairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
}

func (f PairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
	return
}

type PairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
) (R1, R2)

func (f PairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
	return
}

func (f PairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
}

func (f PairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
	return
}

type PairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
) (R1, R2)

func (f PairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
	return
}

func (f PairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
}

func (f PairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
	return
}

type PairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2 any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
) (R1, R2)

func (f PairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
	return
}

func (f PairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
}

func (f PairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if f == nil {
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
	return
}
