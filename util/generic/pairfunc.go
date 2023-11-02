package generic

import (
	"fmt"
	"kit.golaxy.org/golaxy/internal"
	"kit.golaxy.org/golaxy/util/types"
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f()
	return
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1)
	return
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2)
	return
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3)
	return
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4)
	return
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5)
	return
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6)
	return
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7)
	return
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8)
	return
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	return
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	return
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	return
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	return
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	return
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	return
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	return
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
		return types.Zero[R1](), types.Zero[R2](), nil
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

	r1, r2 = f(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	return
}
