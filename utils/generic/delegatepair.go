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

type DelegatePair0[R1, R2 any] []FuncPair0[R1, R2]

func (d DelegatePair0[R1, R2]) Exec(interrupt Func3[R1, R2, error, bool]) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt)
	return
}

func (d DelegatePair0[R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool]) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt)
}

func (d DelegatePair0[R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool]) (r1 R1, r2 R2, panicErr error) {
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

func (d DelegatePair0[R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair0[R1, R2] {
	return func() (R1, R2) { return d.Exec(interrupt) }
}

func (d DelegatePair0[R1, R2]) Combine(f ...FuncPair0[R1, R2]) DelegatePair0[R1, R2] {
	return append(d, f...)
}

type DelegatePair1[A1, R1, R2 any] []FuncPair1[A1, R1, R2]

func (d DelegatePair1[A1, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], a1 A1) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1)
	return
}

func (d DelegatePair1[A1, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], a1 A1) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1)
}

func (d DelegatePair1[A1, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1) (r1 R1, r2 R2, panicErr error) {
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

func (d DelegatePair1[A1, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair1[A1, R1, R2] {
	return func(a1 A1) (R1, R2) { return d.Exec(interrupt, a1) }
}

func (d DelegatePair1[A1, R1, R2]) Combine(f ...FuncPair1[A1, R1, R2]) DelegatePair1[A1, R1, R2] {
	return append(d, f...)
}

type DelegatePair2[A1, A2, R1, R2 any] []FuncPair2[A1, A2, R1, R2]

func (d DelegatePair2[A1, A2, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2)
	return
}

func (d DelegatePair2[A1, A2, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2)
}

func (d DelegatePair2[A1, A2, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2) (r1 R1, r2 R2, panicErr error) {
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

func (d DelegatePair2[A1, A2, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair2[A1, A2, R1, R2] {
	return func(a1 A1, a2 A2) (R1, R2) { return d.Exec(interrupt, a1, a2) }
}

func (d DelegatePair2[A1, A2, R1, R2]) Combine(f ...FuncPair2[A1, A2, R1, R2]) DelegatePair2[A1, A2, R1, R2] {
	return append(d, f...)
}

type DelegatePair3[A1, A2, A3, R1, R2 any] []FuncPair3[A1, A2, A3, R1, R2]

func (d DelegatePair3[A1, A2, A3, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3)
	return
}

func (d DelegatePair3[A1, A2, A3, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3)
}

func (d DelegatePair3[A1, A2, A3, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3) (r1 R1, r2 R2, panicErr error) {
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

func (d DelegatePair3[A1, A2, A3, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair3[A1, A2, A3, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3) (R1, R2) { return d.Exec(interrupt, a1, a2, a3) }
}

func (d DelegatePair3[A1, A2, A3, R1, R2]) Combine(f ...FuncPair3[A1, A2, A3, R1, R2]) DelegatePair3[A1, A2, A3, R1, R2] {
	return append(d, f...)
}

type DelegatePair4[A1, A2, A3, A4, R1, R2 any] []FuncPair4[A1, A2, A3, A4, R1, R2]

func (d DelegatePair4[A1, A2, A3, A4, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4)
	return
}

func (d DelegatePair4[A1, A2, A3, A4, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4)
}

func (d DelegatePair4[A1, A2, A3, A4, R1, R2]) Call(autoRecover bool, reportError chan error, interrupt Func3[R1, R2, error, bool], a1 A1, a2 A2, a3 A3, a4 A4) (r1 R1, r2 R2, panicErr error) {
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

func (d DelegatePair4[A1, A2, A3, A4, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair4[A1, A2, A3, A4, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) (R1, R2) { return d.Exec(interrupt, a1, a2, a3, a4) }
}

func (d DelegatePair4[A1, A2, A3, A4, R1, R2]) Combine(f ...FuncPair4[A1, A2, A3, A4, R1, R2]) DelegatePair4[A1, A2, A3, A4, R1, R2] {
	return append(d, f...)
}

type DelegatePair5[A1, A2, A3, A4, A5, R1, R2 any] []FuncPair5[
	A1, A2, A3, A4, A5, R1, R2,
]

func (d DelegatePair5[A1, A2, A3, A4, A5, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5)
	return
}

func (d DelegatePair5[A1, A2, A3, A4, A5, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5)
}

func (d DelegatePair5[A1, A2, A3, A4, A5, R1, R2]) Call(
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

func (d DelegatePair5[A1, A2, A3, A4, A5, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair5[A1, A2, A3, A4, A5, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) (R1, R2) { return d.Exec(interrupt, a1, a2, a3, a4, a5) }
}

func (d DelegatePair5[A1, A2, A3, A4, A5, R1, R2]) Combine(f ...FuncPair5[A1, A2, A3, A4, A5, R1, R2]) DelegatePair5[A1, A2, A3, A4, A5, R1, R2] {
	return append(d, f...)
}

type DelegatePair6[A1, A2, A3, A4, A5, A6, R1, R2 any] []FuncPair6[
	A1, A2, A3, A4, A5, A6, R1, R2,
]

func (d DelegatePair6[A1, A2, A3, A4, A5, A6, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6)
	return
}

func (d DelegatePair6[A1, A2, A3, A4, A5, A6, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6)
}

func (d DelegatePair6[A1, A2, A3, A4, A5, A6, R1, R2]) Call(
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

func (d DelegatePair6[A1, A2, A3, A4, A5, A6, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair6[A1, A2, A3, A4, A5, A6, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6)
	}
}

func (d DelegatePair6[A1, A2, A3, A4, A5, A6, R1, R2]) Combine(f ...FuncPair6[A1, A2, A3, A4, A5, A6, R1, R2]) DelegatePair6[A1, A2, A3, A4, A5, A6, R1, R2] {
	return append(d, f...)
}

type DelegatePair7[A1, A2, A3, A4, A5, A6, A7, R1, R2 any] []FuncPair7[
	A1, A2, A3, A4, A5, A6, A7, R1, R2,
]

func (d DelegatePair7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7)
	return
}

func (d DelegatePair7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7)
}

func (d DelegatePair7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) Call(
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

func (d DelegatePair7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair7[A1, A2, A3, A4, A5, A6, A7, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7)
	}
}

func (d DelegatePair7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) Combine(f ...FuncPair7[A1, A2, A3, A4, A5, A6, A7, R1, R2]) DelegatePair7[A1, A2, A3, A4, A5, A6, A7, R1, R2] {
	return append(d, f...)
}

type DelegatePair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2 any] []FuncPair8[
	A1, A2, A3, A4, A5, A6, A7, A8, R1, R2,
]

func (d DelegatePair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8)
	return
}

func (d DelegatePair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8)
}

func (d DelegatePair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) Call(
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

func (d DelegatePair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (d DelegatePair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) Combine(f ...FuncPair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2]) DelegatePair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2] {
	return append(d, f...)
}

type DelegatePair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2 any] []FuncPair9[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2,
]

func (d DelegatePair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9)
	return
}

func (d DelegatePair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9)
}

func (d DelegatePair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) Call(
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

func (d DelegatePair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (d DelegatePair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) Combine(f ...FuncPair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2]) DelegatePair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2] {
	return append(d, f...)
}

type DelegatePair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2 any] []FuncPair10[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2,
]

func (d DelegatePair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	return
}

func (d DelegatePair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
}

func (d DelegatePair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) Call(
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

func (d DelegatePair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	}
}

func (d DelegatePair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) Combine(f ...FuncPair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2]) DelegatePair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2] {
	return append(d, f...)
}

type DelegatePair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2 any] []FuncPair11[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2,
]

func (d DelegatePair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	return
}

func (d DelegatePair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
}

func (d DelegatePair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) Call(
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

func (d DelegatePair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	}
}

func (d DelegatePair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) Combine(f ...FuncPair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2]) DelegatePair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2] {
	return append(d, f...)
}

type DelegatePair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2 any] []FuncPair12[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2,
]

func (d DelegatePair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	return
}

func (d DelegatePair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
}

func (d DelegatePair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) Call(
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

func (d DelegatePair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	}
}

func (d DelegatePair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) Combine(f ...FuncPair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2]) DelegatePair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2] {
	return append(d, f...)
}

type DelegatePair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2 any] []FuncPair13[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2,
]

func (d DelegatePair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	return
}

func (d DelegatePair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
}

func (d DelegatePair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) Call(
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

func (d DelegatePair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	}
}

func (d DelegatePair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) Combine(f ...FuncPair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2]) DelegatePair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2] {
	return append(d, f...)
}

type DelegatePair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2 any] []FuncPair14[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2,
]

func (d DelegatePair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	return
}

func (d DelegatePair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
}

func (d DelegatePair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) Call(
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

func (d DelegatePair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	}
}

func (d DelegatePair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) Combine(f ...FuncPair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2]) DelegatePair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2] {
	return append(d, f...)
}

type DelegatePair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2 any] []FuncPair15[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2,
]

func (d DelegatePair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	return
}

func (d DelegatePair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
}

func (d DelegatePair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) Call(
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

func (d DelegatePair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	}
}

func (d DelegatePair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) Combine(f ...FuncPair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2]) DelegatePair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2] {
	return append(d, f...)
}

type DelegatePair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2 any] []FuncPair16[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2,
]

func (d DelegatePair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) Exec(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (r1 R1, r2 R2) {
	r1, r2, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	return
}

func (d DelegatePair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) Invoke(interrupt Func3[R1, R2, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (r1 R1, r2 R2, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
}

func (d DelegatePair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) Call(
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

func (d DelegatePair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) ToFunc(interrupt Func3[R1, R2, error, bool]) FuncPair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) (R1, R2) {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	}
}

func (d DelegatePair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) Combine(f ...FuncPair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2]) DelegatePair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2] {
	return append(d, f...)
}
