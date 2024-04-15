package generic

type DelegatePairFuncVar0[VA, R1, R2 any] []PairFuncVar0[VA, R1, R2]

func (d DelegatePairFuncVar0[VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], va ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, va...)
	return
}

func (d DelegatePairFuncVar0[VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], va ...VA) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, va...)
}

func (d DelegatePairFuncVar0[VA, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], va ...VA) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar0[VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar0[VA, R1, R2] {
	return func(va ...VA) (R1, R2) { return d.Exec(interrupt, va...) }
}

type DelegatePairFuncVar1[A1, VA, R1, R2 any] []PairFuncVar1[A1, VA, R1, R2]

func (d DelegatePairFuncVar1[A1, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], a1 A1, va ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, va...)
	return
}

func (d DelegatePairFuncVar1[A1, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], a1 A1, va ...VA) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, va...)
}

func (d DelegatePairFuncVar1[A1, VA, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, va ...VA) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar1[A1, VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar1[A1, VA, R1, R2] {
	return func(a1 A1, va ...VA) (R1, R2) { return d.Exec(interrupt, a1, va...) }
}

type DelegatePairFuncVar2[A1, A2, VA, R1, R2 any] []PairFuncVar2[A1, A2, VA, R1, R2]

func (d DelegatePairFuncVar2[A1, A2, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, va ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, va...)
	return
}

func (d DelegatePairFuncVar2[A1, A2, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, va ...VA) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, va...)
}

func (d DelegatePairFuncVar2[A1, A2, VA, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, va ...VA) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar2[A1, A2, VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar2[A1, A2, VA, R1, R2] {
	return func(a1 A1, a2 A2, va ...VA) (R1, R2) { return d.Exec(interrupt, a1, a2, va...) }
}

type DelegatePairFuncVar3[A1, A2, A3, VA, R1, R2 any] []PairFuncVar3[A1, A2, A3, VA, R1, R2]

func (d DelegatePairFuncVar3[A1, A2, A3, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, va ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, va...)
	return
}

func (d DelegatePairFuncVar3[A1, A2, A3, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, va ...VA) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, va...)
}

func (d DelegatePairFuncVar3[A1, A2, A3, VA, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, va ...VA) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar3[A1, A2, A3, VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar3[A1, A2, A3, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, va ...VA) (R1, R2) { return d.Exec(interrupt, a1, a2, a3, va...) }
}

type DelegatePairFuncVar4[A1, A2, A3, A4, VA, R1, R2 any] []PairFuncVar4[A1, A2, A3, A4, VA, R1, R2]

func (d DelegatePairFuncVar4[A1, A2, A3, A4, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, va...)
	return
}

func (d DelegatePairFuncVar4[A1, A2, A3, A4, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, va...)
}

func (d DelegatePairFuncVar4[A1, A2, A3, A4, VA, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar4[A1, A2, A3, A4, VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar4[A1, A2, A3, A4, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (R1, R2) { return d.Exec(interrupt, a1, a2, a3, a4, va...) }
}

type DelegatePairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2 any] []PairFuncVar5[
	A1, A2, A3, A4, A5, VA, R1, R2,
]

func (d DelegatePairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, va...)
	return
}

func (d DelegatePairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, va...)
}

func (d DelegatePairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, va...)
	}
}

type DelegatePairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2 any] []PairFuncVar6[
	A1, A2, A3, A4, A5, A6, VA, R1, R2,
]

func (d DelegatePairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, va...)
	return
}

func (d DelegatePairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, va...)
}

func (d DelegatePairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, va...)
	}
}

type DelegatePairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2 any] []PairFuncVar7[
	A1, A2, A3, A4, A5, A6, A7, VA, R1, R2,
]

func (d DelegatePairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, va...)
	return
}

func (d DelegatePairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, va...)
}

func (d DelegatePairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, va...)
	}
}

type DelegatePairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2 any] []PairFuncVar8[
	A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2,
]

func (d DelegatePairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, va...)
	return
}

func (d DelegatePairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, va...)
}

func (d DelegatePairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, va...)
	}
}

type DelegatePairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2 any] []PairFuncVar9[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2,
]

func (d DelegatePairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
	return
}

func (d DelegatePairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
}

func (d DelegatePairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
	}
}

type DelegatePairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2 any] []PairFuncVar10[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2,
]

func (d DelegatePairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
	return
}

func (d DelegatePairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
}

func (d DelegatePairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
	}
}

type DelegatePairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2 any] []PairFuncVar11[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2,
]

func (d DelegatePairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
	return
}

func (d DelegatePairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
}

func (d DelegatePairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
	}
}

type DelegatePairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2 any] []PairFuncVar12[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2,
]

func (d DelegatePairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
	return
}

func (d DelegatePairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
}

func (d DelegatePairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
	}
}

type DelegatePairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2 any] []PairFuncVar13[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2,
]

func (d DelegatePairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
	return
}

func (d DelegatePairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
}

func (d DelegatePairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
	}
}

type DelegatePairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2 any] []PairFuncVar14[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2,
]

func (d DelegatePairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
	return
}

func (d DelegatePairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
}

func (d DelegatePairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
	}
}

type DelegatePairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2 any] []PairFuncVar15[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2,
]

func (d DelegatePairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
	return
}

func (d DelegatePairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
}

func (d DelegatePairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
	}
}

type DelegatePairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2 any] []PairFuncVar16[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2,
]

func (d DelegatePairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
	return
}

func (d DelegatePairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
}

func (d DelegatePairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) CastFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
	}
}
