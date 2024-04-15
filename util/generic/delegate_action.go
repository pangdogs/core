package generic

type DelegateAction0 []Action0

func (d DelegateAction0) Exec(interrupt Func1[error, bool]) {
	d.Call(false, nil, interrupt)
}

func (d DelegateAction0) Invoke(interrupt Func1[error, bool]) (panicErr error) {
	return d.Call(true, nil, interrupt)
}

func (d DelegateAction0) Call(autoRecover bool, reportError chan error, interrupt Func1[error, bool]) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction0) CastAction(interrupt Func1[error, bool]) Action0 {
	return func() { d.Exec(interrupt) }
}

type DelegateAction1[A1 any] []Action1[A1]

func (d DelegateAction1[A1]) Exec(interrupt Func1[error, bool], a1 A1) {
	d.Call(false, nil, interrupt, a1)
}

func (d DelegateAction1[A1]) Invoke(interrupt Func1[error, bool], a1 A1) (panicErr error) {
	return d.Call(true, nil, interrupt, a1)
}

func (d DelegateAction1[A1]) Call(autoRecover bool, reportError chan error, interrupt Func1[error, bool], a1 A1) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction1[A1]) CastAction(interrupt Func1[error, bool]) Action1[A1] {
	return func(a1 A1) { d.Exec(interrupt, a1) }
}

type DelegateAction2[A1, A2 any] []Action2[A1, A2]

func (d DelegateAction2[A1, A2]) Exec(interrupt Func1[error, bool], a1 A1, a2 A2) {
	d.Call(false, nil, interrupt, a1, a2)
}

func (d DelegateAction2[A1, A2]) Invoke(interrupt Func1[error, bool], a1 A1, a2 A2) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2)
}

func (d DelegateAction2[A1, A2]) Call(autoRecover bool, reportError chan error, interrupt Func1[error, bool], a1 A1, a2 A2) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction2[A1, A2]) CastAction(interrupt Func1[error, bool]) Action2[A1, A2] {
	return func(a1 A1, a2 A2) { d.Exec(interrupt, a1, a2) }
}

type DelegateAction3[A1, A2, A3 any] []Action3[A1, A2, A3]

func (d DelegateAction3[A1, A2, A3]) Exec(interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3) {
	d.Call(false, nil, interrupt, a1, a2, a3)
}

func (d DelegateAction3[A1, A2, A3]) Invoke(interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3)
}

func (d DelegateAction3[A1, A2, A3]) Call(autoRecover bool, reportError chan error, interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction3[A1, A2, A3]) CastAction(interrupt Func1[error, bool]) Action3[A1, A2, A3] {
	return func(a1 A1, a2 A2, a3 A3) { d.Exec(interrupt, a1, a2, a3) }
}

type DelegateAction4[A1, A2, A3, A4 any] []Action4[A1, A2, A3, A4]

func (d DelegateAction4[A1, A2, A3, A4]) Exec(interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3, a4 A4) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4)
}

func (d DelegateAction4[A1, A2, A3, A4]) Invoke(interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3, a4 A4) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4)
}

func (d DelegateAction4[A1, A2, A3, A4]) Call(autoRecover bool, reportError chan error, interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3, a4 A4) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction4[A1, A2, A3, A4]) CastAction(interrupt Func1[error, bool]) Action4[A1, A2, A3, A4] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) { d.Exec(interrupt, a1, a2, a3, a4) }
}

type DelegateAction5[A1, A2, A3, A4, A5 any] []Action5[
	A1, A2, A3, A4, A5,
]

func (d DelegateAction5[A1, A2, A3, A4, A5]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5)
}

func (d DelegateAction5[A1, A2, A3, A4, A5]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5)
}

func (d DelegateAction5[A1, A2, A3, A4, A5]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction5[A1, A2, A3, A4, A5]) CastAction(interrupt Func1[error, bool]) Action5[A1, A2, A3, A4, A5] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) { d.Exec(interrupt, a1, a2, a3, a4, a5) }
}

type DelegateAction6[A1, A2, A3, A4, A5, A6 any] []Action6[
	A1, A2, A3, A4, A5, A6,
]

func (d DelegateAction6[A1, A2, A3, A4, A5, A6]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6)
}

func (d DelegateAction6[A1, A2, A3, A4, A5, A6]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6)
}

func (d DelegateAction6[A1, A2, A3, A4, A5, A6]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction6[A1, A2, A3, A4, A5, A6]) CastAction(interrupt Func1[error, bool]) Action6[A1, A2, A3, A4, A5, A6] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) { d.Exec(interrupt, a1, a2, a3, a4, a5, a6) }
}

type DelegateAction7[A1, A2, A3, A4, A5, A6, A7 any] []Action7[
	A1, A2, A3, A4, A5, A6, A7,
]

func (d DelegateAction7[A1, A2, A3, A4, A5, A6, A7]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7)
}

func (d DelegateAction7[A1, A2, A3, A4, A5, A6, A7]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7)
}

func (d DelegateAction7[A1, A2, A3, A4, A5, A6, A7]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction7[A1, A2, A3, A4, A5, A6, A7]) CastAction(interrupt Func1[error, bool]) Action7[A1, A2, A3, A4, A5, A6, A7] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) { d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7) }
}

type DelegateAction8[A1, A2, A3, A4, A5, A6, A7, A8 any] []Action8[
	A1, A2, A3, A4, A5, A6, A7, A8,
]

func (d DelegateAction8[A1, A2, A3, A4, A5, A6, A7, A8]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8)
}

func (d DelegateAction8[A1, A2, A3, A4, A5, A6, A7, A8]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8)
}

func (d DelegateAction8[A1, A2, A3, A4, A5, A6, A7, A8]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction8[A1, A2, A3, A4, A5, A6, A7, A8]) CastAction(interrupt Func1[error, bool]) Action8[A1, A2, A3, A4, A5, A6, A7, A8] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

type DelegateAction9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any] []Action9[
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
]

func (d DelegateAction9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9)
}

func (d DelegateAction9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9)
}

func (d DelegateAction9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction9[A1, A2, A3, A4, A5, A6, A7, A8, A9]) CastAction(interrupt Func1[error, bool]) Action9[A1, A2, A3, A4, A5, A6, A7, A8, A9] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

type DelegateAction10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any] []Action10[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
]

func (d DelegateAction10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
}

func (d DelegateAction10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
}

func (d DelegateAction10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10]) CastAction(interrupt Func1[error, bool]) Action10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	}
}

type DelegateAction11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any] []Action11[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
]

func (d DelegateAction11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
}

func (d DelegateAction11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
}

func (d DelegateAction11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11]) CastAction(interrupt Func1[error, bool]) Action11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	}
}

type DelegateAction12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any] []Action12[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
]

func (d DelegateAction12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
}

func (d DelegateAction12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
}

func (d DelegateAction12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12]) CastAction(interrupt Func1[error, bool]) Action12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	}
}

type DelegateAction13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any] []Action13[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
]

func (d DelegateAction13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
}

func (d DelegateAction13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
}

func (d DelegateAction13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13]) CastAction(interrupt Func1[error, bool]) Action13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	}
}

type DelegateAction14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any] []Action14[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
]

func (d DelegateAction14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
}

func (d DelegateAction14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
}

func (d DelegateAction14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14]) CastAction(interrupt Func1[error, bool]) Action14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	}
}

type DelegateAction15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any] []Action15[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
]

func (d DelegateAction15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
}

func (d DelegateAction15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
}

func (d DelegateAction15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15]) CastAction(interrupt Func1[error, bool]) Action15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	}
}

type DelegateAction16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any] []Action16[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
]

func (d DelegateAction16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
}

func (d DelegateAction16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
}

func (d DelegateAction16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateAction16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16]) CastAction(interrupt Func1[error, bool]) Action16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	}
}
