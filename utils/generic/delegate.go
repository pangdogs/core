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

type Delegate0[R any] []Func0[R]

func (d Delegate0[R]) Exec(interrupt Func2[R, error, bool]) (r R) {
	r, _ = d.Call(false, nil, interrupt)
	return
}

func (d Delegate0[R]) Invoke(interrupt Func2[R, error, bool]) (r R, panicErr error) {
	return d.Call(true, nil, interrupt)
}

func (d Delegate0[R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool]) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate0[R]) ToFunc(interrupt Func2[R, error, bool]) Func0[R] {
	return func() R { return d.Exec(interrupt) }
}

func (d Delegate0[R]) Combine(f ...Func0[R]) Delegate0[R] {
	return append(d, f...)
}

type Delegate1[A1, R any] []Func1[A1, R]

func (d Delegate1[A1, R]) Exec(interrupt Func2[R, error, bool], a1 A1) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1)
	return
}

func (d Delegate1[A1, R]) Invoke(interrupt Func2[R, error, bool], a1 A1) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1)
}

func (d Delegate1[A1, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], a1 A1) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate1[A1, R]) ToFunc(interrupt Func2[R, error, bool]) Func1[A1, R] {
	return func(a1 A1) R { return d.Exec(interrupt, a1) }
}

func (d Delegate1[A1, R]) Combine(f ...Func1[A1, R]) Delegate1[A1, R] {
	return append(d, f...)
}

type Delegate2[A1, A2, R any] []Func2[A1, A2, R]

func (d Delegate2[A1, A2, R]) Exec(interrupt Func2[R, error, bool], a1 A1, a2 A2) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2)
	return
}

func (d Delegate2[A1, A2, R]) Invoke(interrupt Func2[R, error, bool], a1 A1, a2 A2) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2)
}

func (d Delegate2[A1, A2, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], a1 A1, a2 A2) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate2[A1, A2, R]) ToFunc(interrupt Func2[R, error, bool]) Func2[A1, A2, R] {
	return func(a1 A1, a2 A2) R { return d.Exec(interrupt, a1, a2) }
}

type Delegate3[A1, A2, A3, R any] []Func3[A1, A2, A3, R]

func (d Delegate3[A1, A2, A3, R]) Exec(interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3)
	return
}

func (d Delegate3[A1, A2, A3, R]) Invoke(interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3)
}

func (d Delegate3[A1, A2, A3, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate3[A1, A2, A3, R]) ToFunc(interrupt Func2[R, error, bool]) Func3[A1, A2, A3, R] {
	return func(a1 A1, a2 A2, a3 A3) R { return d.Exec(interrupt, a1, a2, a3) }
}

func (d Delegate3[A1, A2, A3, R]) Combine(f ...Func3[A1, A2, A3, R]) Delegate3[A1, A2, A3, R] {
	return append(d, f...)
}

type Delegate4[A1, A2, A3, A4, R any] []Func4[A1, A2, A3, A4, R]

func (d Delegate4[A1, A2, A3, A4, R]) Exec(interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, a4 A4) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4)
	return
}

func (d Delegate4[A1, A2, A3, A4, R]) Invoke(interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, a4 A4) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4)
}

func (d Delegate4[A1, A2, A3, A4, R]) Call(autoRecover bool, reportError chan error, interrupt Func2[R, error, bool], a1 A1, a2 A2, a3 A3, a4 A4) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate4[A1, A2, A3, A4, R]) ToFunc(interrupt Func2[R, error, bool]) Func4[A1, A2, A3, A4, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) R { return d.Exec(interrupt, a1, a2, a3, a4) }
}

func (d Delegate4[A1, A2, A3, A4, R]) Combine(f ...Func4[A1, A2, A3, A4, R]) Delegate4[A1, A2, A3, A4, R] {
	return append(d, f...)
}

type Delegate5[A1, A2, A3, A4, A5, R any] []Func5[
	A1, A2, A3, A4, A5, R,
]

func (d Delegate5[A1, A2, A3, A4, A5, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5)
	return
}

func (d Delegate5[A1, A2, A3, A4, A5, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5)
}

func (d Delegate5[A1, A2, A3, A4, A5, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate5[A1, A2, A3, A4, A5, R]) ToFunc(interrupt Func2[R, error, bool]) Func5[A1, A2, A3, A4, A5, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R { return d.Exec(interrupt, a1, a2, a3, a4, a5) }
}

func (d Delegate5[A1, A2, A3, A4, A5, R]) Combine(f ...Func5[A1, A2, A3, A4, A5, R]) Delegate5[A1, A2, A3, A4, A5, R] {
	return append(d, f...)
}

type Delegate6[A1, A2, A3, A4, A5, A6, R any] []Func6[
	A1, A2, A3, A4, A5, A6, R,
]

func (d Delegate6[A1, A2, A3, A4, A5, A6, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6)
	return
}

func (d Delegate6[A1, A2, A3, A4, A5, A6, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6)
}

func (d Delegate6[A1, A2, A3, A4, A5, A6, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate6[A1, A2, A3, A4, A5, A6, R]) ToFunc(interrupt Func2[R, error, bool]) Func6[A1, A2, A3, A4, A5, A6, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R { return d.Exec(interrupt, a1, a2, a3, a4, a5, a6) }
}

func (d Delegate6[A1, A2, A3, A4, A5, A6, R]) Combine(f ...Func6[A1, A2, A3, A4, A5, A6, R]) Delegate6[A1, A2, A3, A4, A5, A6, R] {
	return append(d, f...)
}

type Delegate7[A1, A2, A3, A4, A5, A6, A7, R any] []Func7[
	A1, A2, A3, A4, A5, A6, A7, R,
]

func (d Delegate7[A1, A2, A3, A4, A5, A6, A7, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7)
	return
}

func (d Delegate7[A1, A2, A3, A4, A5, A6, A7, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7)
}

func (d Delegate7[A1, A2, A3, A4, A5, A6, A7, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate7[A1, A2, A3, A4, A5, A6, A7, R]) ToFunc(interrupt Func2[R, error, bool]) Func7[A1, A2, A3, A4, A5, A6, A7, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7)
	}
}

func (d Delegate7[A1, A2, A3, A4, A5, A6, A7, R]) Combine(f ...Func7[A1, A2, A3, A4, A5, A6, A7, R]) Delegate7[A1, A2, A3, A4, A5, A6, A7, R] {
	return append(d, f...)
}

type Delegate8[A1, A2, A3, A4, A5, A6, A7, A8, R any] []Func8[
	A1, A2, A3, A4, A5, A6, A7, A8, R,
]

func (d Delegate8[A1, A2, A3, A4, A5, A6, A7, A8, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8)
	return
}

func (d Delegate8[A1, A2, A3, A4, A5, A6, A7, A8, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8)
}

func (d Delegate8[A1, A2, A3, A4, A5, A6, A7, A8, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ToFunc(interrupt Func2[R, error, bool]) Func8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (d Delegate8[A1, A2, A3, A4, A5, A6, A7, A8, R]) Combine(f ...Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) Delegate8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return append(d, f...)
}

type Delegate9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any] []Func9[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, R,
]

func (d Delegate9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9)
	return
}

func (d Delegate9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9)
}

func (d Delegate9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ToFunc(interrupt Func2[R, error, bool]) Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (d Delegate9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Combine(f ...Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Delegate9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return append(d, f...)
}

type Delegate10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any] []Func10[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R,
]

func (d Delegate10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	return
}

func (d Delegate10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
}

func (d Delegate10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) ToFunc(interrupt Func2[R, error, bool]) Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	}
}

func (d Delegate10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) Combine(f ...Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) Delegate10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {
	return append(d, f...)
}

type Delegate11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any] []Func11[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R,
]

func (d Delegate11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	return
}

func (d Delegate11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
}

func (d Delegate11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) ToFunc(interrupt Func2[R, error, bool]) Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	}
}

func (d Delegate11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) Combine(f ...Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) Delegate11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {
	return append(d, f...)
}

type Delegate12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any] []Func12[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R,
]

func (d Delegate12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	return
}

func (d Delegate12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
}

func (d Delegate12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) ToFunc(interrupt Func2[R, error, bool]) Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	}
}

func (d Delegate12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) Combine(f ...Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) Delegate12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {
	return append(d, f...)
}

type Delegate13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any] []Func13[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R,
]

func (d Delegate13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	return
}

func (d Delegate13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
}

func (d Delegate13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) ToFunc(interrupt Func2[R, error, bool]) Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	}
}

func (d Delegate13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) Combine(f ...Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) Delegate13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {
	return append(d, f...)
}

type Delegate14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any] []Func14[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R,
]

func (d Delegate14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	return
}

func (d Delegate14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
}

func (d Delegate14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) ToFunc(interrupt Func2[R, error, bool]) Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	}
}

func (d Delegate14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) Combine(f ...Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) Delegate14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {
	return append(d, f...)
}

type Delegate15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R any] []Func15[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R,
]

func (d Delegate15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	return
}

func (d Delegate15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
}

func (d Delegate15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) ToFunc(interrupt Func2[R, error, bool]) Func15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	}
}

func (d Delegate15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) Combine(f ...Func15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R]) Delegate15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R] {
	return append(d, f...)
}

type Delegate16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R any] []Func16[
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R,
]

func (d Delegate16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) Exec(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (r R) {
	r, _ = d.Call(false, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	return
}

func (d Delegate16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) Invoke(interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (r R, panicErr error) {
	return d.Call(true, nil, interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
}

func (d Delegate16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) Call(
	autoRecover bool, reportError chan error, interrupt Func2[R, error, bool],
	a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16,
) (r R, panicErr error) {
	if len(d) <= 0 {
		return
	}

	for i := range d {
		r, panicErr = d[i].Call(autoRecover, reportError, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
		if interrupt.Exec(r, panicErr) {
			return
		}
	}

	return
}

func (d Delegate16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) ToFunc(interrupt Func2[R, error, bool]) Func16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a15 A15, a16 A16) R {
		return d.Exec(interrupt, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16)
	}
}

func (d Delegate16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) Combine(f ...Func16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R]) Delegate16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R] {
	return append(d, f...)
}
