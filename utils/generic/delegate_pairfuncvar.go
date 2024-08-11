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

type DelegatePairFuncVar0[VA, R1, R2 any] []PairFuncVar0[VA, R1, R2]

func (d DelegatePairFuncVar0[VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, args...)
	return
}

func (d DelegatePairFuncVar0[VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], args ...VA) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, args...)
}

func (d DelegatePairFuncVar0[VA, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], args ...VA) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar0[VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar0[VA, R1, R2] {
	return func(args ...VA) (R1, R2) { return d.Exec(interrupt, args...) }
}

type DelegatePairFuncVar1[A1, VA, R1, R2 any] []PairFuncVar1[A1, VA, R1, R2]

func (d DelegatePairFuncVar1[A1, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], a1 A1, args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, args...)
	return
}

func (d DelegatePairFuncVar1[A1, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], a1 A1, args ...VA) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, args...)
}

func (d DelegatePairFuncVar1[A1, VA, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, args ...VA) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar1[A1, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar1[A1, VA, R1, R2] {
	return func(a1 A1, args ...VA) (R1, R2) { return d.Exec(interrupt, a1, args...) }
}

type DelegatePairFuncVar2[A1, A2, VA, R1, R2 any] []PairFuncVar2[A1, A2, VA, R1, R2]

func (d DelegatePairFuncVar2[A1, A2, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, args...)
	return
}

func (d DelegatePairFuncVar2[A1, A2, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, args ...VA) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, args...)
}

func (d DelegatePairFuncVar2[A1, A2, VA, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, args ...VA) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar2[A1, A2, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar2[A1, A2, VA, R1, R2] {
	return func(a1 A1, a2 A2, args ...VA) (R1, R2) { return d.Exec(interrupt, a1, a2, args...) }
}

type DelegatePairFuncVar3[A1, A2, A3, VA, R1, R2 any] []PairFuncVar3[A1, A2, A3, VA, R1, R2]

func (d DelegatePairFuncVar3[A1, A2, A3, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, args...)
	return
}

func (d DelegatePairFuncVar3[A1, A2, A3, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, args ...VA) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, args...)
}

func (d DelegatePairFuncVar3[A1, A2, A3, VA, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, args ...VA) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar3[A1, A2, A3, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar3[A1, A2, A3, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, args ...VA) (R1, R2) { return d.Exec(interrupt, a1, a2, a3, args...) }
}

type DelegatePairFuncVar4[A1, A2, A3, A4, VA, R1, R2 any] []PairFuncVar4[A1, A2, A3, A4, VA, R1, R2]

func (d DelegatePairFuncVar4[A1, A2, A3, A4, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, args...)
	return
}

func (d DelegatePairFuncVar4[A1, A2, A3, A4, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, args...)
}

func (d DelegatePairFuncVar4[A1, A2, A3, A4, VA, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar4[A1, A2, A3, A4, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar4[A1, A2, A3, A4, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, args...)
	}
}

type DelegatePairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2 any] []PairFuncVar5[
	A1, A2, A3, A4, A5, VA, R1, R2,
]

func (d DelegatePairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, args...)
	return
}

func (d DelegatePairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, args...)
}

func (d DelegatePairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, args...)
	}
}

type DelegatePairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2 any] []PairFuncVar6[
	A1, A2, A3, A4, A5, A6, VA, R1, R2,
]

func (d DelegatePairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, args...)
	return
}

func (d DelegatePairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, args...)
}

func (d DelegatePairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, args...)
	}
}

type DelegatePairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2 any] []PairFuncVar7[
	A1, A2, A3, A4, A5, A6, A7, VA, R1, R2,
]

func (d DelegatePairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, args...)
	return
}

func (d DelegatePairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, args...)
}

func (d DelegatePairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, args...)
	}
}

type DelegatePairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2 any] []PairFuncVar8[
	A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2,
]

func (d DelegatePairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, args...)
	return
}

func (d DelegatePairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, args...)
}

func (d DelegatePairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, args...)
	}
}

type DelegatePairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2 any] []PairFuncVar9[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2,
]

func (d DelegatePairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
	return
}

func (d DelegatePairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
}

func (d DelegatePairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
	}
}

type DelegatePairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2 any] []PairFuncVar10[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2,
]

func (d DelegatePairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
	return
}

func (d DelegatePairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
}

func (d DelegatePairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
	}
}

type DelegatePairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2 any] []PairFuncVar11[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2,
]

func (d DelegatePairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
	return
}

func (d DelegatePairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
}

func (d DelegatePairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
	}
}

type DelegatePairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2 any] []PairFuncVar12[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2,
]

func (d DelegatePairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
	return
}

func (d DelegatePairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
}

func (d DelegatePairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
	}
}

type DelegatePairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2 any] []PairFuncVar13[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2,
]

func (d DelegatePairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
	return
}

func (d DelegatePairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
}

func (d DelegatePairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
	}
}

type DelegatePairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2 any] []PairFuncVar14[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2,
]

func (d DelegatePairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
	return
}

func (d DelegatePairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
}

func (d DelegatePairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
	}
}

type DelegatePairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2 any] []PairFuncVar15[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2,
]

func (d DelegatePairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
	return
}

func (d DelegatePairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
}

func (d DelegatePairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
	}
}

type DelegatePairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2 any] []PairFuncVar16[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2,
]

func (d DelegatePairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
	return
}

func (d DelegatePairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
}

func (d DelegatePairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
		if interrupt.Exec(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) PairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
	}
}
