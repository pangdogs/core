package generic

type DelegateFuncVar0[VA, R any] []FuncVar0[VA, R]

func (d DelegateFuncVar0[VA, R]) Exec(interrupt Func2[R, error, bool], va ...VA) (r R) {
	r, _ = d.Call(false, nil, interrupt, va...)
	return
}

func (d DelegateFuncVar0[VA, R]) Invoke(interrupt Func2[R, error, bool], va ...VA) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, va...)
}

func (d DelegateFuncVar0[VA, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], va ...VA) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar0[VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar0[VA, R] {
	return func(va ...VA) R { return d.Exec(interrupt, va...) }
}

type DelegateFuncVar1[A1, VA, R any] []FuncVar1[A1, VA, R]

func (d DelegateFuncVar1[A1, VA, R]) Exec(interrupt Func2[R, error, bool], a1 A1, va ...VA) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, va...)
	return
}

func (d DelegateFuncVar1[A1, VA, R]) Invoke(interrupt Func2[R, error, bool], a1 A1, va ...VA) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, va...)
}

func (d DelegateFuncVar1[A1, VA, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], a1 A1, va ...VA) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar1[A1, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar1[A1, VA, R] {
	return func(a1 A1, va ...VA) R { return d.Exec(interrupt, a1, va...) }
}

type DelegateFuncVar2[A1, A2, VA, R any] []FuncVar2[A1, A2, VA, R]

func (d DelegateFuncVar2[A1, A2, VA, R]) Exec(interrupt Func2[R, error, bool], a1 A1, a2 A2, va ...VA) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, va...)
	return
}

func (d DelegateFuncVar2[A1, A2, VA, R]) Invoke(interrupt Func2[R, error, bool], a1 A1, a2 A2, va ...VA) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, va...)
}

func (d DelegateFuncVar2[A1, A2, VA, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], a1 A1, a2 A2, va ...VA) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar2[A1, A2, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar2[A1, A2, VA, R] {
	return func(a1 A1, a2 A2, va ...VA) R { return d.Exec(interrupt, a1, a2, va...) }
}

type DelegateFuncVar3[A1, A2, A3, VA, R any] []FuncVar3[A1, A2, A3, VA, R]

func (d DelegateFuncVar3[A1, A2, A3, VA, R]) Exec(interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, va ...VA) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, va...)
	return
}

func (d DelegateFuncVar3[A1, A2, A3, VA, R]) Invoke(interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, va ...VA) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, va...)
}

func (d DelegateFuncVar3[A1, A2, A3, VA, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, va ...VA) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar3[A1, A2, A3, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar3[A1, A2, A3, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, va ...VA) R { return d.Exec(interrupt, a1, a2, a3, va...) }
}

type DelegateFuncVar4[A1, A2, A3, A4, VA, R any] []FuncVar4[A1, A2, A3, A4, VA, R]

func (d DelegateFuncVar4[A1, A2, A3, A4, VA, R]) Exec(interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, va...)
	return
}

func (d DelegateFuncVar4[A1, A2, A3, A4, VA, R]) Invoke(interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, va...)
}

func (d DelegateFuncVar4[A1, A2, A3, A4, VA, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar4[A1, A2, A3, A4, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar4[A1, A2, A3, A4, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) R { return d.Exec(interrupt, a1, a2, a3, a4, va...) }
}

type DelegateFuncVar5[A1, A2, A3, A4, A5, VA, R any] []FuncVar5[
	A1, A2, A3, A4, A5, VA, R,
]

func (d DelegateFuncVar5[A1, A2, A3, A4, A5, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, va...)
	return
}

func (d DelegateFuncVar5[A1, A2, A3, A4, A5, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, va...)
}

func (d DelegateFuncVar5[A1, A2, A3, A4, A5, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar5[A1, A2, A3, A4, A5, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar5[A1, A2, A3, A4, A5, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, va...)
	}
}

type DelegateFuncVar6[A1, A2, A3, A4, A5, A6, VA, R any] []FuncVar6[
	A1, A2, A3, A4, A5, A6, VA, R,
]

func (d DelegateFuncVar6[A1, A2, A3, A4, A5, A6, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, va...)
	return
}

func (d DelegateFuncVar6[A1, A2, A3, A4, A5, A6, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, va...)
}

func (d DelegateFuncVar6[A1, A2, A3, A4, A5, A6, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar6[A1, A2, A3, A4, A5, A6, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar6[A1, A2, A3, A4, A5, A6, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, va...)
	}
}

type DelegateFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R any] []FuncVar7[
	A1, A2, A3, A4, A5, A6, A7, VA, R,
]

func (d DelegateFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, va...)
	return
}

func (d DelegateFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, va...)
}

func (d DelegateFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, va...)
	}
}

type DelegateFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R any] []FuncVar8[
	A1, A2, A3, A4, A5, A6, A7, A8, VA, R,
]

func (d DelegateFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, va...)
	return
}

func (d DelegateFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, va...)
}

func (d DelegateFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, va...)
	}
}

type DelegateFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R any] []FuncVar9[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R,
]

func (d DelegateFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
	return
}

func (d DelegateFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
}

func (d DelegateFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
	}
}

type DelegateFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R any] []FuncVar10[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R,
]

func (d DelegateFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
	return
}

func (d DelegateFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
}

func (d DelegateFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
	}
}

type DelegateFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R any] []FuncVar11[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R,
]

func (d DelegateFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
	return
}

func (d DelegateFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
}

func (d DelegateFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
	}
}

type DelegateFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R any] []FuncVar12[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R,
]

func (d DelegateFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
	return
}

func (d DelegateFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
}

func (d DelegateFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
	}
}

type DelegateFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R any] []FuncVar13[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R,
]

func (d DelegateFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
	return
}

func (d DelegateFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
}

func (d DelegateFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
	}
}

type DelegateFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R any] []FuncVar14[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R,
]

func (d DelegateFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
	return
}

func (d DelegateFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
}

func (d DelegateFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
	}
}

type DelegateFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R any] []FuncVar15[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R,
]

func (d DelegateFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
	return
}

func (d DelegateFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
}

func (d DelegateFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
	}
}

type DelegateFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R any] []FuncVar16[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R,
]

func (d DelegateFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
	return
}

func (d DelegateFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
}

func (d DelegateFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
	}
}
