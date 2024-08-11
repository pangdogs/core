/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

package generic

type DelegateActionVar0[VA any] []ActionVar0[VA]

func (d DelegateActionVar0[VA]) Exec(interrupt Func1[error, bool], args ...VA) {
	d.Call(false, nil, interrupt, args...)
}

func (d DelegateActionVar0[VA]) Invoke(interrupt Func1[error, bool], args ...VA) (panicErr error) {
	return d.Call(true, nil, interrupt, args...)
}

func (d DelegateActionVar0[VA]) Call(autoRecover bool, reportError chan error, interrupt Func1[error, bool], args ...VA) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar0[VA]) ToAction(interrupt Func1[error, bool]) ActionVar0[VA] {
	return func(args ...VA) { d.Exec(interrupt, args...) }
}

type DelegateActionVar1[A1, VA any] []ActionVar1[A1, VA]

func (d DelegateActionVar1[A1, VA]) Exec(interrupt Func1[error, bool], a1 A1, args ...VA) {
	d.Call(false, nil, interrupt, a1, args...)
}

func (d DelegateActionVar1[A1, VA]) Invoke(interrupt Func1[error, bool], a1 A1, args ...VA) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, args...)
}

func (d DelegateActionVar1[A1, VA]) Call(autoRecover bool, reportError chan error, interrupt Func1[error, bool], a1 A1, args ...VA) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar1[A1, VA]) ToAction(interrupt Func1[error, bool]) ActionVar1[A1, VA] {
	return func(a1 A1, args ...VA) { d.Exec(interrupt, a1, args...) }
}

type DelegateActionVar2[A1, A2, VA any] []ActionVar2[A1, A2, VA]

func (d DelegateActionVar2[A1, A2, VA]) Exec(interrupt Func1[error, bool], a1 A1, a2 A2, args ...VA) {
	d.Call(false, nil, interrupt, a1, a2, args...)
}

func (d DelegateActionVar2[A1, A2, VA]) Invoke(interrupt Func1[error, bool], a1 A1, a2 A2, args ...VA) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, args...)
}

func (d DelegateActionVar2[A1, A2, VA]) Call(autoRecover bool, reportError chan error, interrupt Func1[error, bool], a1 A1, a2 A2, args ...VA) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar2[A1, A2, VA]) ToAction(interrupt Func1[error, bool]) ActionVar2[A1, A2, VA] {
	return func(a1 A1, a2 A2, args ...VA) { d.Exec(interrupt, a1, a2, args...) }
}

type DelegateActionVar3[A1, A2, A3, VA any] []ActionVar3[A1, A2, A3, VA]

func (d DelegateActionVar3[A1, A2, A3, VA]) Exec(interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3, args ...VA) {
	d.Call(false, nil, interrupt, a1, a2, a3, args...)
}

func (d DelegateActionVar3[A1, A2, A3, VA]) Invoke(interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3, args ...VA) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, args...)
}

func (d DelegateActionVar3[A1, A2, A3, VA]) Call(autoRecover bool, reportError chan error, interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3, args ...VA) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar3[A1, A2, A3, VA]) ToAction(interrupt Func1[error, bool]) ActionVar3[A1, A2, A3, VA] {
	return func(a1 A1, a2 A2, a3 A3, args ...VA) { d.Exec(interrupt, a1, a2, a3, args...) }
}

type DelegateActionVar4[A1, A2, A3, A4, VA any] []ActionVar4[A1, A2, A3, A4, VA]

func (d DelegateActionVar4[A1, A2, A3, A4, VA]) Exec(interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, args...)
}

func (d DelegateActionVar4[A1, A2, A3, A4, VA]) Invoke(interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, args...)
}

func (d DelegateActionVar4[A1, A2, A3, A4, VA]) Call(autoRecover bool, reportError chan error, interrupt Func1[error, bool], a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar4[A1, A2, A3, A4, VA]) ToAction(interrupt Func1[error, bool]) ActionVar4[A1, A2, A3, A4, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) { d.Exec(interrupt, a1, a2, a3, a4, args...) }
}

type DelegateActionVar5[A1, A2, A3, A4, A5, VA any] []ActionVar5[
	A1, A2, A3, A4, A5, VA,
]

func (d DelegateActionVar5[A1, A2, A3, A4, A5, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, args...)
}

func (d DelegateActionVar5[A1, A2, A3, A4, A5, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, args...)
}

func (d DelegateActionVar5[A1, A2, A3, A4, A5, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar5[A1, A2, A3, A4, A5, VA]) ToAction(interrupt Func1[error, bool]) ActionVar5[A1, A2, A3, A4, A5, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA) { d.Exec(interrupt, a1, a2, a3, a4, a5, args...) }
}

type DelegateActionVar6[A1, A2, A3, A4, A5, A6, VA any] []ActionVar6[
	A1, A2, A3, A4, A5, A6, VA,
]

func (d DelegateActionVar6[A1, A2, A3, A4, A5, A6, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, args...)
}

func (d DelegateActionVar6[A1, A2, A3, A4, A5, A6, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, args...)
}

func (d DelegateActionVar6[A1, A2, A3, A4, A5, A6, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar6[A1, A2, A3, A4, A5, A6, VA]) ToAction(interrupt Func1[error, bool]) ActionVar6[A1, A2, A3, A4, A5, A6, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, args...)
	}
}

type DelegateActionVar7[A1, A2, A3, A4, A5, A6, A7, VA any] []ActionVar7[
	A1, A2, A3, A4, A5, A6, A7, VA,
]

func (d DelegateActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, args...)
}

func (d DelegateActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, args...)
}

func (d DelegateActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar7[A1, A2, A3, A4, A5, A6, A7, VA]) ToAction(interrupt Func1[error, bool]) ActionVar7[A1, A2, A3, A4, A5, A6, A7, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, args...)
	}
}

type DelegateActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA any] []ActionVar8[
	A1, A2, A3, A4, A5, A6, A7, A8, VA,
]

func (d DelegateActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, args...)
}

func (d DelegateActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, args...)
}

func (d DelegateActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA]) ToAction(interrupt Func1[error, bool]) ActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, args...)
	}
}

type DelegateActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA any] []ActionVar9[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, VA,
]

func (d DelegateActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
}

func (d DelegateActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
}

func (d DelegateActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA]) ToAction(interrupt Func1[error, bool]) ActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
	}
}

type DelegateActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA any] []ActionVar10[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA,
]

func (d DelegateActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
}

func (d DelegateActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
}

func (d DelegateActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA]) ToAction(interrupt Func1[error, bool]) ActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
	}
}

type DelegateActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA any] []ActionVar11[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA,
]

func (d DelegateActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
}

func (d DelegateActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
}

func (d DelegateActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA]) ToAction(interrupt Func1[error, bool]) ActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
	}
}

type DelegateActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA any] []ActionVar12[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA,
]

func (d DelegateActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
}

func (d DelegateActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
}

func (d DelegateActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA]) ToAction(interrupt Func1[error, bool]) ActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
	}
}

type DelegateActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA any] []ActionVar13[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA,
]

func (d DelegateActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
}

func (d DelegateActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
}

func (d DelegateActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA]) ToAction(interrupt Func1[error, bool]) ActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
	}
}

type DelegateActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA any] []ActionVar14[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA,
]

func (d DelegateActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
}

func (d DelegateActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
}

func (d DelegateActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA]) ToAction(interrupt Func1[error, bool]) ActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
	}
}

type DelegateActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA any] []ActionVar15[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA,
]

func (d DelegateActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
}

func (d DelegateActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
}

func (d DelegateActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA]) ToAction(interrupt Func1[error, bool]) ActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
	}
}

type DelegateActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA any] []ActionVar16[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA,
]

func (d DelegateActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]) Exec(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) {
	d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
}

func (d DelegateActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]) Invoke(interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
}

func (d DelegateActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]) Call(
	autoRecover bool, reportError chan error, interrupt Func1[error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
		if interrupt.Exec(panicErr) {
			return
		}
	}

	return
}

func (d DelegateActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA]) ToAction(interrupt Func1[error, bool]) ActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA) {
		d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
	}
}
