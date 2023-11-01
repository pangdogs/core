package generic

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/util/types"
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(va...)

	return nil
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, va...)

	return nil
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, va...)

	return nil
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, va...)

	return nil
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, va...)

	return nil
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, va...)

	return nil
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, va...)

	return nil
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, va...)

	return nil
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, va...)

	return nil
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)

	return nil
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)

	return nil
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)

	return nil
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)

	return nil
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)

	return nil
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)

	return nil
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)

	return nil
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
				panicErr = fmt.Errorf("%w: %w", internal.ErrPanicked, panicErr)

				if reportError != nil {
					select {
					case reportError <- internal.PrintStackTrace(panicErr):
					default:
					}
				}
			}
		}()
	}

	f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)

	return nil
}
