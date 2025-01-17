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

type DelegatePairVar0[VA, R1, R2 any] []FuncPairVar0[VA, R1, R2]

func (d DelegatePairVar0[VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool], args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, args...)
	return
}

func (d DelegatePairVar0[VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool], args ...VA) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, args...)
}

func (d DelegatePairVar0[VA, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], args ...VA) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar0[VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar0[VA, R1, R2] {
	return func(args ...VA) (R1, R2) { return d.UnsafeCall(interrupt, args...) }
}

func (d DelegatePairVar0[VA, R1, R2]) Combine(f ...FuncPairVar0[VA, R1, R2]) DelegatePairVar0[VA, R1, R2] {
	return append(d, f...)
}

type DelegatePairVar1[A1, VA, R1, R2 any] []FuncPairVar1[A1, VA, R1, R2]

func (d DelegatePairVar1[A1, VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool], a1 A1, args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, args...)
	return
}

func (d DelegatePairVar1[A1, VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool], a1 A1, args ...VA) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, args...)
}

func (d DelegatePairVar1[A1, VA, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, args ...VA) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar1[A1, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar1[A1, VA, R1, R2] {
	return func(a1 A1, args ...VA) (R1, R2) { return d.UnsafeCall(interrupt, a1, args...) }
}

func (d DelegatePairVar1[A1, VA, R1, R2]) Combine(f ...FuncPairVar1[A1, VA, R1, R2]) DelegatePairVar1[A1, VA, R1, R2] {
	return append(d, f...)
}

type DelegatePairVar2[A1, A2, VA, R1, R2 any] []FuncPairVar2[A1, A2, VA, R1, R2]

func (d DelegatePairVar2[A1, A2, VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, args...)
	return
}

func (d DelegatePairVar2[A1, A2, VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, args ...VA) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, args...)
}

func (d DelegatePairVar2[A1, A2, VA, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, args ...VA) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar2[A1, A2, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar2[A1, A2, VA, R1, R2] {
	return func(a1 A1, a2 A2, args ...VA) (R1, R2) { return d.UnsafeCall(interrupt, a1, a2, args...) }
}

func (d DelegatePairVar2[A1, A2, VA, R1, R2]) Combine(f ...FuncPairVar2[A1, A2, VA, R1, R2]) DelegatePairVar2[A1, A2, VA, R1, R2] {
	return append(d, f...)
}

type DelegatePairVar3[A1, A2, A3, VA, R1, R2 any] []FuncPairVar3[A1, A2, A3, VA, R1, R2]

func (d DelegatePairVar3[A1, A2, A3, VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, args...)
	return
}

func (d DelegatePairVar3[A1, A2, A3, VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, args ...VA) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, args...)
}

func (d DelegatePairVar3[A1, A2, A3, VA, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, args ...VA) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar3[A1, A2, A3, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar3[A1, A2, A3, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, args ...VA) (R1, R2) { return d.UnsafeCall(interrupt, a1, a2, a3, args...) }
}

func (d DelegatePairVar3[A1, A2, A3, VA, R1, R2]) Combine(f ...FuncPairVar3[A1, A2, A3, VA, R1, R2]) DelegatePairVar3[A1, A2, A3, VA, R1, R2] {
	return append(d, f...)
}

type DelegatePairVar4[A1, A2, A3, A4, VA, R1, R2 any] []FuncPairVar4[A1, A2, A3, A4, VA, R1, R2]

func (d DelegatePairVar4[A1, A2, A3, A4, VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, args...)
	return
}

func (d DelegatePairVar4[A1, A2, A3, A4, VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, args...)
}

func (d DelegatePairVar4[A1, A2, A3, A4, VA, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar4[A1, A2, A3, A4, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar4[A1, A2, A3, A4, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, args ...VA) (R1, R2) {
		return d.UnsafeCall(interrupt, a1, a2, a3, a4, args...)
	}
}

func (d DelegatePairVar4[A1, A2, A3, A4, VA, R1, R2]) Combine(f ...FuncPairVar4[A1, A2, A3, A4, VA, R1, R2]) DelegatePairVar4[A1, A2, A3, A4, VA, R1, R2] {
	return append(d, f...)
}

type DelegatePairVar5[A1, A2, A3, A4, A5, VA, R1, R2 any] []FuncPairVar5[
	A1, A2, A3, A4, A5, VA, R1, R2,
]

func (d DelegatePairVar5[A1, A2, A3, A4, A5, VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, args...)
	return
}

func (d DelegatePairVar5[A1, A2, A3, A4, A5, VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, args...)
}

func (d DelegatePairVar5[A1, A2, A3, A4, A5, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar5[A1, A2, A3, A4, A5, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar5[A1, A2, A3, A4, A5, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, args ...VA) (R1, R2) {
		return d.UnsafeCall(interrupt, a1, a2, a3, a4, a5, args...)
	}
}

func (d DelegatePairVar5[A1, A2, A3, A4, A5, VA, R1, R2]) Combine(f ...FuncPairVar5[A1, A2, A3, A4, A5, VA, R1, R2]) DelegatePairVar5[A1, A2, A3, A4, A5, VA, R1, R2] {
	return append(d, f...)
}

type DelegatePairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2 any] []FuncPairVar6[
	A1, A2, A3, A4, A5, A6, VA, R1, R2,
]

func (d DelegatePairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, args...)
	return
}

func (d DelegatePairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, args...)
}

func (d DelegatePairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, args ...VA) (R1, R2) {
		return d.UnsafeCall(interrupt, a1, a2, a3, a4, a5, a6, args...)
	}
}

func (d DelegatePairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) Combine(f ...FuncPairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2]) DelegatePairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2] {
	return append(d, f...)
}

type DelegatePairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2 any] []FuncPairVar7[
	A1, A2, A3, A4, A5, A6, A7, VA, R1, R2,
]

func (d DelegatePairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, args...)
	return
}

func (d DelegatePairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, args...)
}

func (d DelegatePairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, args ...VA) (R1, R2) {
		return d.UnsafeCall(interrupt, a1, a2, a3, a4, a5, a6, a7, args...)
	}
}

func (d DelegatePairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) Combine(f ...FuncPairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2]) DelegatePairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2] {
	return append(d, f...)
}

type DelegatePairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2 any] []FuncPairVar8[
	A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2,
]

func (d DelegatePairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, args...)
	return
}

func (d DelegatePairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, args...)
}

func (d DelegatePairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, args ...VA) (R1, R2) {
		return d.UnsafeCall(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, args...)
	}
}

func (d DelegatePairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) Combine(f ...FuncPairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2]) DelegatePairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2] {
	return append(d, f...)
}

type DelegatePairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2 any] []FuncPairVar9[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2,
]

func (d DelegatePairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
	return
}

func (d DelegatePairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
}

func (d DelegatePairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, args ...VA) (R1, R2) {
		return d.UnsafeCall(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, args...)
	}
}

func (d DelegatePairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) Combine(f ...FuncPairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2]) DelegatePairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2] {
	return append(d, f...)
}

type DelegatePairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2 any] []FuncPairVar10[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2,
]

func (d DelegatePairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
	return
}

func (d DelegatePairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
}

func (d DelegatePairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, args ...VA) (R1, R2) {
		return d.UnsafeCall(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, args...)
	}
}

func (d DelegatePairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) Combine(f ...FuncPairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2]) DelegatePairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2] {
	return append(d, f...)
}

type DelegatePairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2 any] []FuncPairVar11[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2,
]

func (d DelegatePairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
	return
}

func (d DelegatePairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
}

func (d DelegatePairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, args ...VA) (R1, R2) {
		return d.UnsafeCall(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, args...)
	}
}

func (d DelegatePairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) Combine(f ...FuncPairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2]) DelegatePairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2] {
	return append(d, f...)
}

type DelegatePairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2 any] []FuncPairVar12[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2,
]

func (d DelegatePairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
	return
}

func (d DelegatePairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
}

func (d DelegatePairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, args ...VA) (R1, R2) {
		return d.UnsafeCall(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, args...)
	}
}

func (d DelegatePairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) Combine(f ...FuncPairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2]) DelegatePairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2] {
	return append(d, f...)
}

type DelegatePairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2 any] []FuncPairVar13[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2,
]

func (d DelegatePairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
	return
}

func (d DelegatePairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
}

func (d DelegatePairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, args ...VA) (R1, R2) {
		return d.UnsafeCall(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, args...)
	}
}

func (d DelegatePairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) Combine(f ...FuncPairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2]) DelegatePairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2] {
	return append(d, f...)
}

type DelegatePairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2 any] []FuncPairVar14[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2,
]

func (d DelegatePairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
	return
}

func (d DelegatePairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
}

func (d DelegatePairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, args ...VA) (R1, R2) {
		return d.UnsafeCall(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, args...)
	}
}

func (d DelegatePairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) Combine(f ...FuncPairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2]) DelegatePairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2] {
	return append(d, f...)
}

type DelegatePairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2 any] []FuncPairVar15[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2,
]

func (d DelegatePairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
	return
}

func (d DelegatePairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
}

func (d DelegatePairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, args ...VA) (R1, R2) {
		return d.UnsafeCall(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, args...)
	}
}

func (d DelegatePairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) Combine(f ...FuncPairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2]) DelegatePairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2] {
	return append(d, f...)
}

type DelegatePairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2 any] []FuncPairVar16[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2,
]

func (d DelegatePairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) UnsafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
	return
}

func (d DelegatePairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) SafeCall(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
}

func (d DelegatePairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) Call(
	autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA,
) (r1 R1, r2 R2, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r1, r2, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
		if interrupt.UnsafeCall(r1, r2, panicErr) {
			return
		}
	}

	return
}

func (d DelegatePairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16, args ...VA) (R1, R2) {
		return d.UnsafeCall(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, args...)
	}
}

func (d DelegatePairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) Combine(f ...FuncPairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2]) DelegatePairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2] {
	return append(d, f...)
}
