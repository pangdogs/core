package generic

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal/exception"
	"kit.golaxy.org/golaxy/util/types"
)

type FuncVar0[VA, R any] func(...VA) R

func (f FuncVar0[VA, R]) Exec(va ...VA) (r R) {
	r, _ = f.Call(false, nil, va...)
	return
}

func (f FuncVar0[VA, R]) Invoke(va ...VA) (r R, panicErr error) {
	return f.Call(true, nil, va...)
}

func (f FuncVar0[VA, R]) Call(autoRecover bool, reportError chan error, va ...VA) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(va...), nil
}

type FuncVar1[A1, VA, R any] func(A1, ...VA) R

func (f FuncVar1[A1, VA, R]) Exec(a1 A1, va ...VA) (r R) {
	r, _ = f.Call(false, nil, a1, va...)
	return
}

func (f FuncVar1[A1, VA, R]) Invoke(a1 A1, va ...VA) (r R, panicErr error) {
	return f.Call(true, nil, a1, va...)
}

func (f FuncVar1[A1, VA, R]) Call(autoRecover bool, reportError chan error, a1 A1, va ...VA) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(a1, va...), nil
}

type FuncVar2[A1, A2, VA, R any] func(A1, A2, ...VA) R

func (f FuncVar2[A1, A2, VA, R]) Exec(a1 A1, a2 A2, va ...VA) (r R) {
	r, _ = f.Call(false, nil, a1, a2, va...)
	return
}

func (f FuncVar2[A1, A2, VA, R]) Invoke(a1 A1, a2 A2, va ...VA) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, va...)
}

func (f FuncVar2[A1, A2, VA, R]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, va ...VA) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(a1, a2, va...), nil
}

type FuncVar3[A1, A2, A3, VA, R any] func(A1, A2, A3, ...VA) R

func (f FuncVar3[A1, A2, A3, VA, R]) Exec(a1 A1, a2 A2, a3 A3, va ...VA) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, va...)
	return
}

func (f FuncVar3[A1, A2, A3, VA, R]) Invoke(a1 A1, a2 A2, a3 A3, va ...VA) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, va...)
}

func (f FuncVar3[A1, A2, A3, VA, R]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, va ...VA) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(a1, a2, a3, va...), nil
}

type FuncVar4[A1, A2, A3, A4, VA, R any] func(A1, A2, A3, A4, ...VA) R

func (f FuncVar4[A1, A2, A3, A4, VA, R]) Exec(a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, va...)
	return
}

func (f FuncVar4[A1, A2, A3, A4, VA, R]) Invoke(a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, va...)
}

func (f FuncVar4[A1, A2, A3, A4, VA, R]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(a1, a2, a3, a4, va...), nil
}

type FuncVar5[A1, A2, A3, A4, A5, VA, R any] func(
	A1, A2, A3, A4, A5, ...VA,
) R

func (f FuncVar5[A1, A2, A3, A4, A5, VA, R]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, va...)
	return
}

func (f FuncVar5[A1, A2, A3, A4, A5, VA, R]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, va...)
}

func (f FuncVar5[A1, A2, A3, A4, A5, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(a1, a2, a3, a4, a5, va...), nil
}

type FuncVar6[A1, A2, A3, A4, A5, A6, VA, R any] func(
	A1, A2, A3, A4, A5, A6, ...VA,
) R

func (f FuncVar6[A1, A2, A3, A4, A5, A6, VA, R]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, va...)
	return
}

func (f FuncVar6[A1, A2, A3, A4, A5, A6, VA, R]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, va...)
}

func (f FuncVar6[A1, A2, A3, A4, A5, A6, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(a1, a2, a3, a4, a5, a6, va...), nil
}

type FuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
) R

func (f FuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, va...)
	return
}

func (f FuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, va...)
}

func (f FuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(a1, a2, a3, a4, a5, a6, a7, va...), nil
}

type FuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
) R

func (f FuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, va...)
	return
}

func (f FuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, va...)
}

func (f FuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(a1, a2, a3, a4, a5, a6, a7, a8, va...), nil
}

type FuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
) R

func (f FuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
	return
}

func (f FuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
}

func (f FuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, va...), nil
}

type FuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
) R

func (f FuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
	return
}

func (f FuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
}

func (f FuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...), nil
}

type FuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
) R

func (f FuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
	return
}

func (f FuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
}

func (f FuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...), nil
}

type FuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
) R

func (f FuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
	return
}

func (f FuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
}

func (f FuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...), nil
}

type FuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
) R

func (f FuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
	return
}

func (f FuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
}

func (f FuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...), nil
}

type FuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
) R

func (f FuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
	return
}

func (f FuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
}

func (f FuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...), nil
}

type FuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
) R

func (f FuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
	return
}

func (f FuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
}

func (f FuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...), nil
}

type FuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
) R

func (f FuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) (r R) {
	r, _ = f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
	return
}

func (f FuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) (r R, panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
}

func (f FuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) (r R, panicErr error) {
	if f == nil {
		return types.Zero[R](), nil
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

	return f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...), nil
}
