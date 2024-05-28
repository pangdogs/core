package generic

import (
	"fmt"
	"git.golaxy.org/core/utils/exception"
	"git.golaxy.org/core/utils/types"
)

type ActionVar0[VA any] func(...VA)

func (f ActionVar0[VA]) Exec(va ...VA) {
	f.Call(false, nil, va...)
}

func (f ActionVar0[VA]) Invoke(va ...VA) (panicErr error) {
	return f.Call(true, nil, va...)
}

func (f ActionVar0[VA]) Call(autoRecover bool, reportError chan error, va ...VA) (panicErr error) {
	if f == nil {
		return nil
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

	f(va...)

	return nil
}

func (f ActionVar0[VA]) CastDelegate() DelegateActionVar0[VA] {
	return []ActionVar0[VA]{f}
}

type ActionVar1[A1, VA any] func(A1, ...VA)

func (f ActionVar1[A1, VA]) Exec(a1 A1, va ...VA) {
	f.Call(false, nil, a1, va...)
}

func (f ActionVar1[A1, VA]) Invoke(a1 A1, va ...VA) (panicErr error) {
	return f.Call(true, nil, a1, va...)
}

func (f ActionVar1[A1, VA]) Call(autoRecover bool, reportError chan error, a1 A1, va ...VA) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, va...)

	return nil
}

func (f ActionVar1[A1, VA]) CastDelegate() DelegateActionVar1[A1, VA] {
	return []ActionVar1[A1, VA]{f}
}

type ActionVar2[A1, A2, VA any] func(A1, A2, ...VA)

func (f ActionVar2[A1, A2, VA]) Exec(a1 A1, a2 A2, va ...VA) {
	f.Call(false, nil, a1, a2, va...)
}

func (f ActionVar2[A1, A2, VA]) Invoke(a1 A1, a2 A2, va ...VA) (panicErr error) {
	return f.Call(true, nil, a1, a2, va...)
}

func (f ActionVar2[A1, A2, VA]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, va ...VA) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, va...)

	return nil
}

func (f ActionVar2[A1, A2, VA]) CastDelegate() DelegateActionVar2[A1, A2, VA] {
	return []ActionVar2[A1, A2, VA]{f}
}

type ActionVar3[A1, A2, A3, VA any] func(A1, A2, A3, ...VA)

func (f ActionVar3[A1, A2, A3, VA]) Exec(a1 A1, a2 A2, a3 A3, va ...VA) {
	f.Call(false, nil, a1, a2, a3, va...)
}

func (f ActionVar3[A1, A2, A3, VA]) Invoke(a1 A1, a2 A2, a3 A3, va ...VA) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, va...)
}

func (f ActionVar3[A1, A2, A3, VA]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, va ...VA) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, a3, va...)

	return nil
}

func (f ActionVar3[A1, A2, A3, VA]) CastDelegate() DelegateActionVar3[A1, A2, A3, VA] {
	return []ActionVar3[A1, A2, A3, VA]{f}
}

type ActionVar4[A1, A2, A3, A4, VA any] func(A1, A2, A3, A4, ...VA)

func (f ActionVar4[A1, A2, A3, A4, VA]) Exec(a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) {
	f.Call(false, nil, a1, a2, a3, a4, va...)
}

func (f ActionVar4[A1, A2, A3, A4, VA]) Invoke(a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, va...)
}

func (f ActionVar4[A1, A2, A3, A4, VA]) Call(autoRecover bool, reportError chan error, a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (panicErr error) {
	if f == nil {
		return nil
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

	f(a1, a2, a3, a4, va...)

	return nil
}

func (f ActionVar4[A1, A2, A3, A4, VA]) CastDelegate() DelegateActionVar4[A1, A2, A3, A4, VA] {
	return []ActionVar4[A1, A2, A3, A4, VA]{f}
}

type ActionVar5[A1, A2, A3, A4, A5, VA any] func(
	A1, A2, A3, A4, A5, ...VA,
)

func (f ActionVar5[A1, A2, A3, A4, A5, VA]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, va...)
}

func (f ActionVar5[A1, A2, A3, A4, A5, VA]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, va...)
}

func (f ActionVar5[A1, A2, A3, A4, A5, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, va...)

	return nil
}

func (f ActionVar5[A1, A2, A3, A4, A5, VA]) CastDelegate() DelegateActionVar5[A1, A2, A3, A4, A5, VA] {
	return []ActionVar5[A1, A2, A3, A4, A5, VA]{f}
}

type ActionVar6[A1, A2, A3, A4, A5, A6, VA any] func(
	A1, A2, A3, A4, A5, A6, ...VA,
)

func (f ActionVar6[A1, A2, A3, A4, A5, A6, VA]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, va...)
}

func (f ActionVar6[A1, A2, A3, A4, A5, A6, VA]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, va...)
}

func (f ActionVar6[A1, A2, A3, A4, A5, A6, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, va...)

	return nil
}

func (f ActionVar6[A1, A2, A3, A4, A5, A6, VA]) CastDelegate() DelegateActionVar6[A1, A2, A3, A4, A5, A6, VA] {
	return []ActionVar6[A1, A2, A3, A4, A5, A6, VA]{f}
}

type ActionVar7[A1, A2, A3, A4, A5, A6, A7, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
)

func (f ActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, va...)
}

func (f ActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, va...)
}

func (f ActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, va...)

	return nil
}

func (f ActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]) CastDelegate() DelegateActionVar7[A1, A2, A3, A4, A5, A6, A7, VA] {
	return []ActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]{f}
}

type ActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
)

func (f ActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, va...)
}

func (f ActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, va...)
}

func (f ActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, va...)

	return nil
}

func (f ActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]) CastDelegate() DelegateActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA] {
	return []ActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]{f}
}

type ActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
)

func (f ActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
}

func (f ActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
}

func (f ActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)

	return nil
}

func (f ActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]) CastDelegate() DelegateActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA] {
	return []ActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]{f}
}

type ActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
)

func (f ActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
}

func (f ActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
}

func (f ActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)

	return nil
}

func (f ActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]) CastDelegate() DelegateActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA] {
	return []ActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]{f}
}

type ActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
)

func (f ActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
}

func (f ActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
}

func (f ActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)

	return nil
}

func (f ActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]) CastDelegate() DelegateActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA] {
	return []ActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]{f}
}

type ActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
)

func (f ActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
}

func (f ActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
}

func (f ActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)

	return nil
}

func (f ActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]) CastDelegate() DelegateActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA] {
	return []ActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]{f}
}

type ActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
)

func (f ActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
}

func (f ActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
}

func (f ActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)

	return nil
}

func (f ActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]) CastDelegate() DelegateActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA] {
	return []ActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]{f}
}

type ActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
)

func (f ActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
}

func (f ActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
}

func (f ActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)

	return nil
}

func (f ActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]) CastDelegate() DelegateActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA] {
	return []ActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]{f}
}

type ActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
)

func (f ActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
}

func (f ActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
}

func (f ActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)

	return nil
}

func (f ActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]) CastDelegate() DelegateActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA] {
	return []ActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]{f}
}

type ActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA any] func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
)

func (f ActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]) Exec(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) {
	f.Call(false, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
}

func (f ActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]) Invoke(
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) (panicErr error) {
	return f.Call(true, nil, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
}

func (f ActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]) Call(
	autoRecover bool, reportError chan error,
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
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
					case reportError <- exception.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)

	return nil
}

func (f ActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]) CastDelegate() DelegateActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA] {
	return []ActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]{f}
}
