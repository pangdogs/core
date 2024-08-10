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

type DelegateFuncVar0[VA, R any] []FuncVar0[VA, R]

func (d DelegateFuncVar0[VA, R]) Exec(interrupt Func2[R, error, bool], args ...VA) (r R) {
	r, _ = d.Call(false, nil, interrupt, args...)
	return
}

func (d DelegateFuncVar0[VA, R]) Invoke(interrupt Func2[R, error, bool], args ...VA) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, args...)
}

func (d DelegateFuncVar0[VA, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], args ...VA) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar0[VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar0[VA, R] {
	return func(args ...VA) R { return d.Exec(interrupt, args...) }
}

type DelegateFuncVar1[A1, VA, R any] []FuncVar1[A1, VA, R]

func (d DelegateFuncVar1[A1, VA, R]) Exec(interrupt Func2[R, error, bool], a1 A1, args ...VA) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, args...)
	return
}

func (d DelegateFuncVar1[A1, VA, R]) Invoke(interrupt Func2[R, error, bool], a1 A1, args ...VA) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, args...)
}

func (d DelegateFuncVar1[A1, VA, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], a1 A1, args ...VA) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar1[A1, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar1[A1, VA, R] {
	return func(a1 A1, args ...VA) R { return d.Exec(interrupt, a1, args...) }
}

type DelegateFuncVar2[A1, A2, VA, R any] []FuncVar2[A1, A2, VA, R]

func (d DelegateFuncVar2[A1, A2, VA, R]) Exec(interrupt Func2[R, error, bool], a1 A1, a2 A2, args ...VA) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, args...)
	return
}

func (d DelegateFuncVar2[A1, A2, VA, R]) Invoke(interrupt Func2[R, error, bool], a1 A1, a2 A2, args ...VA) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, args...)
}

func (d DelegateFuncVar2[A1, A2, VA, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], a1 A1, a2 A2, args ...VA) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar2[A1, A2, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar2[A1, A2, VA, R] {
	return func(a1 A1, a2 A2, args ...VA) R { return d.Exec(interrupt, a1, a2, args...) }
}

type DelegateFuncVar3[A1, A2, A3, VA, R any] []FuncVar3[A1, A2, A3, VA, R]

func (d DelegateFuncVar3[A1, A2, A3, VA, R]) Exec(interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, args ...VA) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, args...)
	return
}

func (d DelegateFuncVar3[A1, A2, A3, VA, R]) Invoke(interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, args ...VA) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, args...)
}

func (d DelegateFuncVar3[A1, A2, A3, VA, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, args ...VA) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar3[A1, A2, A3, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar3[A1, A2, A3, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, args ...VA) R { return d.Exec(interrupt, a1, a2, a3, args...) }
}

type DelegateFuncVar4[A1, A2, A3, A4, VA, R any] []FuncVar4[A1, A2, A3, A4, VA, R]

func (d DelegateFuncVar4[A1, A2, A3, A4, VA, R]) Exec(interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, args...)
	return
}

func (d DelegateFuncVar4[A1, A2, A3, A4, VA, R]) Invoke(interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, args...)
}

func (d DelegateFuncVar4[A1, A2, A3, A4, VA, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar4[A1, A2, A3, A4, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar4[A1, A2, A3, A4, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) R { return d.Exec(interrupt, a1, a2, a3, a4, args...) }
}

type DelegateFuncVar5[A1, A2, A3, A4, A5, VA, R any] []FuncVar5[
	A1, A2, A3, A4, A5, VA, R,
]

func (d DelegateFuncVar5[A1, A2, A3, A4, A5, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, args...)
	return
}

func (d DelegateFuncVar5[A1, A2, A3, A4, A5, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, args...)
}

func (d DelegateFuncVar5[A1, A2, A3, A4, A5, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar5[A1, A2, A3, A4, A5, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar5[A1, A2, A3, A4, A5, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, args...)
	}
}

type DelegateFuncVar6[A1, A2, A3, A4, A5, A6, VA, R any] []FuncVar6[
	A1, A2, A3, A4, A5, A6, VA, R,
]

func (d DelegateFuncVar6[A1, A2, A3, A4, A5, A6, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, args...)
	return
}

func (d DelegateFuncVar6[A1, A2, A3, A4, A5, A6, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, args...)
}

func (d DelegateFuncVar6[A1, A2, A3, A4, A5, A6, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar6[A1, A2, A3, A4, A5, A6, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar6[A1, A2, A3, A4, A5, A6, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, args...)
	}
}

type DelegateFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R any] []FuncVar7[
	A1, A2, A3, A4, A5, A6, A7, VA, R,
]

func (d DelegateFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, args...)
	return
}

func (d DelegateFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, args...)
}

func (d DelegateFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, args...)
	}
}

type DelegateFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R any] []FuncVar8[
	A1, A2, A3, A4, A5, A6, A7, A8, VA, R,
]

func (d DelegateFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, args...)
	return
}

func (d DelegateFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, args...)
}

func (d DelegateFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, args...)
	}
}

type DelegateFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R any] []FuncVar9[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R,
]

func (d DelegateFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
	return
}

func (d DelegateFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
}

func (d DelegateFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
	}
}

type DelegateFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R any] []FuncVar10[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R,
]

func (d DelegateFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
	return
}

func (d DelegateFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
}

func (d DelegateFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
	}
}

type DelegateFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R any] []FuncVar11[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R,
]

func (d DelegateFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
	return
}

func (d DelegateFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
}

func (d DelegateFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
	}
}

type DelegateFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R any] []FuncVar12[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R,
]

func (d DelegateFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
	return
}

func (d DelegateFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
}

func (d DelegateFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
	}
}

type DelegateFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R any] []FuncVar13[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R,
]

func (d DelegateFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
	return
}

func (d DelegateFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
}

func (d DelegateFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
	}
}

type DelegateFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R any] []FuncVar14[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R,
]

func (d DelegateFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
	return
}

func (d DelegateFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
}

func (d DelegateFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
	}
}

type DelegateFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R any] []FuncVar15[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R,
]

func (d DelegateFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
	return
}

func (d DelegateFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
}

func (d DelegateFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
	}
}

type DelegateFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R any] []FuncVar16[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R,
]

func (d DelegateFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
	return
}

func (d DelegateFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
}

func (d DelegateFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d DelegateFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R]) CastFunc(interrupt Func2[R, error, bool]) FuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
	}
}
