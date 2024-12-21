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

func CastDelegate0[R any, F ~func() R](fs ...F) Delegate0[R] {
	d := make(Delegate0[R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc0(fs[i]))
	}
	return d
}

func CastDelegate1[A1, R any, F ~func(A1) R](fs ...F) Delegate1[A1, R] {
	d := make(Delegate1[A1, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc1(fs[i]))
	}
	return d
}

func CastDelegate2[A1, A2, R any, F ~func(A1, A2) R](fs ...F) Delegate2[A1, A2, R] {
	d := make(Delegate2[A1, A2, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc2(fs[i]))
	}
	return d
}

func CastDelegate3[A1, A2, A3, R any, F ~func(A1, A2, A3) R](fs ...F) Delegate3[A1, A2, A3, R] {
	d := make(Delegate3[A1, A2, A3, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc3(fs[i]))
	}
	return d
}

func CastDelegate4[A1, A2, A3, A4, R any, F ~func(A1, A2, A3, A4) R](fs ...F) Delegate4[A1, A2, A3, A4, R] {
	d := make(Delegate4[A1, A2, A3, A4, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc4(fs[i]))
	}
	return d
}

func CastDelegate5[A1, A2, A3, A4, A5, R any, F ~func(
	A1, A2, A3, A4, A5,
) R](fs ...F) Delegate5[A1, A2, A3, A4, A5, R] {
	d := make(Delegate5[A1, A2, A3, A4, A5, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc5(fs[i]))
	}
	return d
}

func CastDelegate6[A1, A2, A3, A4, A5, A6, R any, F ~func(
	A1, A2, A3, A4, A5, A6,
) R](fs ...F) Delegate6[A1, A2, A3, A4, A5, A6, R] {
	d := make(Delegate6[A1, A2, A3, A4, A5, A6, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc6(fs[i]))
	}
	return d
}

func CastDelegate7[A1, A2, A3, A4, A5, A6, A7, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7,
) R](fs ...F) Delegate7[A1, A2, A3, A4, A5, A6, A7, R] {
	d := make(Delegate7[A1, A2, A3, A4, A5, A6, A7, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc7(fs[i]))
	}
	return d
}

func CastDelegate8[A1, A2, A3, A4, A5, A6, A7, A8, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8,
) R](fs ...F) Delegate8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	d := make(Delegate8[A1, A2, A3, A4, A5, A6, A7, A8, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc8(fs[i]))
	}
	return d
}

func CastDelegate9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
) R](fs ...F) Delegate9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	d := make(Delegate9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc9(fs[i]))
	}
	return d
}

func CastDelegate10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
) R](fs ...F) Delegate10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {
	d := make(Delegate10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc10(fs[i]))
	}
	return d
}

func CastDelegate11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
) R](fs ...F) Delegate11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {
	d := make(Delegate11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc11(fs[i]))
	}
	return d
}

func CastDelegate12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
) R](fs ...F) Delegate12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {
	d := make(Delegate12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc12(fs[i]))
	}
	return d
}

func CastDelegate13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
) R](fs ...F) Delegate13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {
	d := make(Delegate13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc13(fs[i]))
	}
	return d
}

func CastDelegate14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
) R](fs ...F) Delegate14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {
	d := make(Delegate14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc14(fs[i]))
	}
	return d
}

func CastDelegate15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
) R](fs ...F) Delegate15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R] {
	d := make(Delegate15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc15(fs[i]))
	}
	return d
}

func CastDelegate16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
) R](fs ...F) Delegate16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R] {
	d := make(Delegate16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc16(fs[i]))
	}
	return d
}
