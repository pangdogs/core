package generic

type DelegateActionVar0[VA any] []ActionVar0[VA]

func (d DelegateActionVar0[VA]) Exec(interrupt Func1[error, bool], va ...VA) {
	d.Call(false, nil, interrupt, va...)
}

func (d DelegateActionVar0[VA]) Invoke(interrupt Func1[error, bool], va ...VA) (panicErr error) {
	return d.Call(true, nil, interrupt, va...)
}

func (d DelegateActionVar0[VA]) Call(autoRecover bool, reportError chan error, interrupt Func1[error, bool], va ...VA) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar0[VA]) CastAction(interrupt Func1[error, bool]) ActionVar0[VA] {
	return func(va ...VA) { d.Exec(interrupt, va...) }
}

type DelegateActionVar1[A1, VA any] []ActionVar1[A1, VA]

func (d DelegateActionVar1[A1, VA]) Exec(interrupt Func1[error, bool], a1 A1, va ...VA) {
	d.Call(false, nil, interrupt, a1, va...)
}

func (d DelegateActionVar1[A1, VA]) Invoke(interrupt Func1[error, bool], a1 A1, va ...VA) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, va...)
}

func (d DelegateActionVar1[A1, VA]) Call(autoRecover bool, reportError chan error, interrupt Func1[error, bool], a1 A1, va ...VA) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar1[A1, VA]) CastAction(interrupt Func1[error, bool]) ActionVar1[A1, VA] {
	return func(a1 A1, va ...VA) { d.Exec(interrupt, a1, va...) }
}

type DelegateActionVar2[A1, A2, VA any] []ActionVar2[A1, A2, VA]

func (d DelegateActionVar2[A1, A2, VA]) Exec(interrupt Func1[error, bool], a1 A1, a2 A2, va ...VA) {
	d.Call(false, nil, interrupt, a1, a2, va...)
}

func (d DelegateActionVar2[A1, A2, VA]) Invoke(interrupt Func1[error, bool], a1 A1, a2 A2, va ...VA) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, va...)
}

func (d DelegateActionVar2[A1, A2, VA]) Call(autoRecover bool, reportError chan error, interrupt Func1[error, bool], a1 A1, a2 A2, va ...VA) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar2[A1, A2, VA]) CastAction(interrupt Func1[error, bool]) ActionVar2[A1, A2, VA] {
	return func(a1 A1, a2 A2, va ...VA) { d.Exec(interrupt, a1, a2, va...) }
}

type DelegateActionVar3[A1, A2, A3, VA any] []ActionVar3[A1, A2, A3, VA]

func (d DelegateActionVar3[A1, A2, A3, VA]) Exec(interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3, va ...VA) {
	d.Call(false, nil, interrupt, a1, a2, a3, va...)
}

func (d DelegateActionVar3[A1, A2, A3, VA]) Invoke(interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3, va ...VA) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, va...)
}

func (d DelegateActionVar3[A1, A2, A3, VA]) Call(autoRecover bool, reportError chan error, interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3, va ...VA) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar3[A1, A2, A3, VA]) CastAction(interrupt Func1[error, bool]) ActionVar3[A1, A2, A3, VA] {
	return func(a1 A1, a2 A2, a3 A3, va ...VA) { d.Exec(interrupt, a1, a2, a3, va...) }
}

type DelegateActionVar4[A1, A2, A3, A4, VA any] []ActionVar4[A1, A2, A3, A4, VA]

func (d DelegateActionVar4[A1, A2, A3, A4, VA]) Exec(interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, va...)
}

func (d DelegateActionVar4[A1, A2, A3, A4, VA]) Invoke(interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, va...)
}

func (d DelegateActionVar4[A1, A2, A3, A4, VA]) Call(autoRecover bool, reportError chan error, interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar4[A1, A2, A3, A4, VA]) CastAction(interrupt Func1[error, bool]) ActionVar4[A1, A2, A3, A4, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, va ...VA) { d.Exec(interrupt, a1, a2, a3, a4, va...) }
}

type DelegateActionVar5[A1, A2, A3, A4, A5, VA any] []ActionVar5[
	A1, A2, A3, A4, A5, VA,
]

func (d DelegateActionVar5[A1, A2, A3, A4, A5, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, va...)
}

func (d DelegateActionVar5[A1, A2, A3, A4, A5, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, va...)
}

func (d DelegateActionVar5[A1, A2, A3, A4, A5, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar5[A1, A2, A3, A4, A5, VA]) CastAction(interrupt Func1[error, bool]) ActionVar5[A1, A2, A3, A4, A5, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, va ...VA) { d.Exec(interrupt, a1, a2, a3, a4, a5, va...) }
}

type DelegateActionVar6[A1, A2, A3, A4, A5, A6, VA any] []ActionVar6[
	A1, A2, A3, A4, A5, A6, VA,
]

func (d DelegateActionVar6[A1, A2, A3, A4, A5, A6, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, va...)
}

func (d DelegateActionVar6[A1, A2, A3, A4, A5, A6, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, va...)
}

func (d DelegateActionVar6[A1, A2, A3, A4, A5, A6, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar6[A1, A2, A3, A4, A5, A6, VA]) CastAction(interrupt Func1[error, bool]) ActionVar6[A1, A2, A3, A4, A5, A6, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, va ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, va...)
	}
}

type DelegateActionVar7[A1, A2, A3, A4, A5, A6, A7, VA any] []ActionVar7[
	A1, A2, A3, A4, A5, A6, A7, VA,
]

func (d DelegateActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, va...)
}

func (d DelegateActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, va...)
}

func (d DelegateActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]) CastAction(interrupt Func1[error, bool]) ActionVar7[A1, A2, A3, A4, A5, A6, A7, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, va ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, va...)
	}
}

type DelegateActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA any] []ActionVar8[
	A1, A2, A3, A4, A5, A6, A7, A8, VA,
]

func (d DelegateActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, va...)
}

func (d DelegateActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, va...)
}

func (d DelegateActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]) CastAction(interrupt Func1[error, bool]) ActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, va ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, va...)
	}
}

type DelegateActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA any] []ActionVar9[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, VA,
]

func (d DelegateActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
}

func (d DelegateActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
}

func (d DelegateActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]) CastAction(interrupt Func1[error, bool]) ActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, va ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, va...)
	}
}

type DelegateActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA any] []ActionVar10[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA,
]

func (d DelegateActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
}

func (d DelegateActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
}

func (d DelegateActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]) CastAction(interrupt Func1[error, bool]) ActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, va ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, va...)
	}
}

type DelegateActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA any] []ActionVar11[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA,
]

func (d DelegateActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
}

func (d DelegateActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
}

func (d DelegateActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]) CastAction(interrupt Func1[error, bool]) ActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, va ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, va...)
	}
}

type DelegateActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA any] []ActionVar12[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA,
]

func (d DelegateActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
}

func (d DelegateActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
}

func (d DelegateActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]) CastAction(interrupt Func1[error, bool]) ActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, va ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, va...)
	}
}

type DelegateActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA any] []ActionVar13[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA,
]

func (d DelegateActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
}

func (d DelegateActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
}

func (d DelegateActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]) CastAction(interrupt Func1[error, bool]) ActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, va ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, va...)
	}
}

type DelegateActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA any] []ActionVar14[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA,
]

func (d DelegateActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
}

func (d DelegateActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
}

func (d DelegateActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]) CastAction(interrupt Func1[error, bool]) ActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, va ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, va...)
	}
}

type DelegateActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA any] []ActionVar15[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA,
]

func (d DelegateActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
}

func (d DelegateActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
}

func (d DelegateActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]) CastAction(interrupt Func1[error, bool]) ActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, va ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, va...)
	}
}

type DelegateActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA any] []ActionVar16[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA,
]

func (d DelegateActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
}

func (d DelegateActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
}

func (d DelegateActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]) CastAction(interrupt Func1[error, bool]) ActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, va ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, va...)
	}
}
