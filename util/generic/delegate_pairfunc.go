package generic

type DelegatePairFunc0[R1, R2 any] []PairFunc0[R1, R2]

func (d DelegatePairFunc0[R1, R2]) Exec(interrupt Func3[R1, R2, error, bool]) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt)
	return
}

func (d DelegatePairFunc0[R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool]) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt)
}

func (d DelegatePairFunc0[R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool]) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc0[R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc0[R1, R2] {
	return func() (R1, R2) { return d.Exec(interrupt) }
}

type DelegatePairFunc1[A1, R1, R2 any] []PairFunc1[A1, R1, R2]

func (d DelegatePairFunc1[A1, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], a1 A1) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1)
	return
}

func (d DelegatePairFunc1[A1, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], a1 A1) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1)
}

func (d DelegatePairFunc1[A1, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc1[A1, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc1[A1, R1, R2] {
	return func(a1 A1) (R1, R2) { return d.Exec(interrupt, a1) }
}

type DelegatePairFunc2[A1, A2, R1, R2 any] []PairFunc2[A1, A2, R1, R2]

func (d DelegatePairFunc2[A1, A2, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2)
	return
}

func (d DelegatePairFunc2[A1, A2, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2)
}

func (d DelegatePairFunc2[A1, A2, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc2[A1, A2, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc2[A1, A2, R1, R2] {
	return func(a1 A1, a2 A2) (R1, R2) { return d.Exec(interrupt, a1, a2) }
}

type DelegatePairFunc3[A1, A2, A3, R1, R2 any] []PairFunc3[A1, A2, A3, R1, R2]

func (d DelegatePairFunc3[A1, A2, A3, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3)
	return
}

func (d DelegatePairFunc3[A1, A2, A3, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3)
}

func (d DelegatePairFunc3[A1, A2, A3, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc3[A1, A2, A3, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc3[A1, A2, A3, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3) (R1, R2) { return d.Exec(interrupt, a1, a2, a3) }
}

type DelegatePairFunc4[A1, A2, A3, A4, R1, R2 any] []PairFunc4[A1, A2, A3, A4, R1, R2]

func (d DelegatePairFunc4[A1, A2, A3, A4, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4)
	return
}

func (d DelegatePairFunc4[A1, A2, A3, A4, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4)
}

func (d DelegatePairFunc4[A1, A2, A3, A4, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc4[A1, A2, A3, A4, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc4[A1, A2, A3, A4, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2) { return d.Exec(interrupt, a1, a2, a3, a4) }
}

type DelegatePairFunc5[A1, A2, A3, A4, A5, R1, R2 any] []PairFunc5[
	A1, A2, A3, A4, A5, R1, R2,
]

func (d DelegatePairFunc5[A1, A2, A3, A4, A5, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5)
	return
}

func (d DelegatePairFunc5[A1, A2, A3, A4, A5, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5)
}

func (d DelegatePairFunc5[A1, A2, A3, A4, A5, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc5[A1, A2, A3, A4, A5, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc5[A1, A2, A3, A4, A5, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2) { return d.Exec(interrupt, a1, a2, a3, a4, a5) }
}

type DelegatePairFunc6[A1, A2, A3, A4, A5, A6, R1, R2 any] []PairFunc6[
	A1, A2, A3, A4, A5, A6, R1, R2,
]

func (d DelegatePairFunc6[A1, A2, A3, A4, A5, A6, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6)
	return
}

func (d DelegatePairFunc6[A1, A2, A3, A4, A5, A6, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6)
}

func (d DelegatePairFunc6[A1, A2, A3, A4, A5, A6, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc6[A1, A2, A3, A4, A5, A6, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc6[A1, A2, A3, A4, A5, A6, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6)
	}
}

type DelegatePairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2 any] []PairFunc7[
	A1, A2, A3, A4, A5, A6, A7, R1, R2,
]

func (d DelegatePairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7)
	return
}

func (d DelegatePairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7)
}

func (d DelegatePairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7)
	}
}

type DelegatePairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2 any] []PairFunc8[
	A1, A2, A3, A4, A5, A6, A7, A8, R1, R2,
]

func (d DelegatePairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8)
	return
}

func (d DelegatePairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8)
}

func (d DelegatePairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

type DelegatePairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2 any] []PairFunc9[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2,
]

func (d DelegatePairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9)
	return
}

func (d DelegatePairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9)
}

func (d DelegatePairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

type DelegatePairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2 any] []PairFunc10[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2,
]

func (d DelegatePairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	return
}

func (d DelegatePairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
}

func (d DelegatePairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	}
}

type DelegatePairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2 any] []PairFunc11[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2,
]

func (d DelegatePairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	return
}

func (d DelegatePairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
}

func (d DelegatePairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	}
}

type DelegatePairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2 any] []PairFunc12[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2,
]

func (d DelegatePairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	return
}

func (d DelegatePairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
}

func (d DelegatePairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	}
}

type DelegatePairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2 any] []PairFunc13[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2,
]

func (d DelegatePairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	return
}

func (d DelegatePairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
}

func (d DelegatePairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	}
}

type DelegatePairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2 any] []PairFunc14[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2,
]

func (d DelegatePairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	return
}

func (d DelegatePairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
}

func (d DelegatePairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	}
}

type DelegatePairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2 any] []PairFunc15[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2,
]

func (d DelegatePairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	return
}

func (d DelegatePairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
}

func (d DelegatePairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	}
}

type DelegatePairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2 any] []PairFunc16[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2,
]

func (d DelegatePairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	return
}

func (d DelegatePairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
}

func (d DelegatePairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	}
}
