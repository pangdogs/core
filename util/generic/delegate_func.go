package generic

type DelegateFunc0[R any] []Func0[R]

func (d DelegateFunc0[R]) Exec(interrupt Func2[R, error, bool]) (r R) {
	r, _ = d.Call(false, nil, interrupt)
	return
}

func (d DelegateFunc0[R]) Invoke(interrupt Func2[R, error, bool]) (r R, panicErr error) {
	return d.Call(true, nil, interrupt)
}

func (d DelegateFunc0[R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool]) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc0[R]) CastFunc(interrupt Func2[R, error, bool]) Func0[R] {
	return func() R { return d.Exec(interrupt) }
}

type DelegateFunc1[A1, R any] []Func1[A1, R]

func (d DelegateFunc1[A1, R]) Exec(interrupt Func2[R, error, bool], a1 A1) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1)
	return
}

func (d DelegateFunc1[A1, R]) Invoke(interrupt Func2[R, error, bool], a1 A1) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1)
}

func (d DelegateFunc1[A1, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], a1 A1) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc1[A1, R]) CastFunc(interrupt Func2[R, error, bool]) Func1[A1, R] {
	return func(a1 A1) R { return d.Exec(interrupt, a1) }
}

type DelegateFunc2[A1, A2, R any] []Func2[A1, A2, R]

func (d DelegateFunc2[A1, A2, R]) Exec(interrupt Func2[R, error, bool], a1 A1, a2 A2) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2)
	return
}

func (d DelegateFunc2[A1, A2, R]) Invoke(interrupt Func2[R, error, bool], a1 A1, a2 A2) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2)
}

func (d DelegateFunc2[A1, A2, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], a1 A1, a2 A2) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc2[A1, A2, R]) CastFunc(interrupt Func2[R, error, bool]) Func2[A1, A2, R] {
	return func(a1 A1, a2 A2) R { return d.Exec(interrupt, a1, a2) }
}

type DelegateFunc3[A1, A2, A3, R any] []Func3[A1, A2, A3, R]

func (d DelegateFunc3[A1, A2, A3, R]) Exec(interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3)
	return
}

func (d DelegateFunc3[A1, A2, A3, R]) Invoke(interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3)
}

func (d DelegateFunc3[A1, A2, A3, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc3[A1, A2, A3, R]) CastFunc(interrupt Func2[R, error, bool]) Func3[A1, A2, A3, R] {
	return func(a1 A1, a2 A2, a3 A3) R { return d.Exec(interrupt, a1, a2, a3) }
}

type DelegateFunc4[A1, A2, A3, A4, R any] []Func4[A1, A2, A3, A4, R]

func (d DelegateFunc4[A1, A2, A3, A4, R]) Exec(interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, a4 A4) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4)
	return
}

func (d DelegateFunc4[A1, A2, A3, A4, R]) Invoke(interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, a4 A4) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4)
}

func (d DelegateFunc4[A1, A2, A3, A4, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, a4 A4) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc4[A1, A2, A3, A4, R]) CastFunc(interrupt Func2[R, error, bool]) Func4[A1, A2, A3, A4, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) R { return d.Exec(interrupt, a1, a2, a3, a4) }
}

type DelegateFunc5[A1, A2, A3, A4, A5, R any] []Func5[
	A1, A2, A3, A4, A5, R,
]

func (d DelegateFunc5[A1, A2, A3, A4, A5, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5)
	return
}

func (d DelegateFunc5[A1, A2, A3, A4, A5, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5)
}

func (d DelegateFunc5[A1, A2, A3, A4, A5, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc5[A1, A2, A3, A4, A5, R]) CastFunc(interrupt Func2[R, error, bool]) Func5[A1, A2, A3, A4, A5, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R { return d.Exec(interrupt, a1, a2, a3, a4, a5) }
}

type DelegateFunc6[A1, A2, A3, A4, A5, A6, R any] []Func6[
	A1, A2, A3, A4, A5, A6, R,
]

func (d DelegateFunc6[A1, A2, A3, A4, A5, A6, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6)
	return
}

func (d DelegateFunc6[A1, A2, A3, A4, A5, A6, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6)
}

func (d DelegateFunc6[A1, A2, A3, A4, A5, A6, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc6[A1, A2, A3, A4, A5, A6, R]) CastFunc(interrupt Func2[R, error, bool]) Func6[A1, A2, A3, A4, A5, A6, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R { return d.Exec(interrupt, a1, a2, a3, a4, a5, a6) }
}

type DelegateFunc7[A1, A2, A3, A4, A5, A6, A7, R any] []Func7[
	A1, A2, A3, A4, A5, A6, A7, R,
]

func (d DelegateFunc7[A1, A2, A3, A4, A5, A6, A7, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7)
	return
}

func (d DelegateFunc7[A1, A2, A3, A4, A5, A6, A7, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7)
}

func (d DelegateFunc7[A1, A2, A3, A4, A5, A6, A7, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc7[A1, A2, A3, A4, A5, A6, A7, R]) CastFunc(interrupt Func2[R, error, bool]) Func7[A1, A2, A3, A4, A5, A6, A7, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7)
	}
}

type DelegateFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R any] []Func8[
	A1, A2, A3, A4, A5, A6, A7, A8, R,
]

func (d DelegateFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8)
	return
}

func (d DelegateFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8)
}

func (d DelegateFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R]) CastFunc(interrupt Func2[R, error, bool]) Func8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

type DelegateFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any] []Func9[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, R,
]

func (d DelegateFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9)
	return
}

func (d DelegateFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9)
}

func (d DelegateFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) CastFunc(interrupt Func2[R, error, bool]) Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

type DelegateFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any] []Func10[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R,
]

func (d DelegateFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	return
}

func (d DelegateFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
}

func (d DelegateFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) CastFunc(interrupt Func2[R, error, bool]) Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	}
}

type DelegateFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any] []Func11[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R,
]

func (d DelegateFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	return
}

func (d DelegateFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
}

func (d DelegateFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) CastFunc(interrupt Func2[R, error, bool]) Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	}
}

type DelegateFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any] []Func12[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R,
]

func (d DelegateFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	return
}

func (d DelegateFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
}

func (d DelegateFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) CastFunc(interrupt Func2[R, error, bool]) Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	}
}

type DelegateFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any] []Func13[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R,
]

func (d DelegateFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	return
}

func (d DelegateFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
}

func (d DelegateFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) CastFunc(interrupt Func2[R, error, bool]) Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	}
}

type DelegateFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any] []Func14[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R,
]

func (d DelegateFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	return
}

func (d DelegateFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
}

func (d DelegateFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) CastFunc(interrupt Func2[R, error, bool]) Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	}
}

type DelegateFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R any] []Func15[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R,
]

func (d DelegateFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	return
}

func (d DelegateFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
}

func (d DelegateFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) CastFunc(interrupt Func2[R, error, bool]) Func15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	}
}

type DelegateFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R any] []Func16[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R,
]

func (d DelegateFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	return
}

func (d DelegateFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
}

func (d DelegateFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) CastFunc(interrupt Func2[R, error, bool]) Func16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	}
}
