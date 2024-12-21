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

func CastDelegateVoid0[F ~func()](fs ...F) DelegateVoid0 {
	d := make(DelegateVoid0, 0, len(fs))
	for i := range fs {
		d = append(d, CastAction0(fs[i]))
	}
	return d
}

func CastDelegateVoid1[A1 any, F ~func(A1)](fs ...F) DelegateVoid1[A1] {
	d := make(DelegateVoid1[A1], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction1(fs[i]))
	}
	return d
}

func CastDelegateVoid2[A1, A2 any, F ~func(A1, A2)](fs ...F) DelegateVoid2[A1, A2] {
	d := make(DelegateVoid2[A1, A2], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction2(fs[i]))
	}
	return d
}

func CastDelegateVoid3[A1, A2, A3 any, F ~func(A1, A2, A3)](fs ...F) DelegateVoid3[A1, A2, A3] {
	d := make(DelegateVoid3[A1, A2, A3], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction3(fs[i]))
	}
	return d
}

func CastDelegateVoid4[A1, A2, A3, A4 any, F ~func(A1, A2, A3, A4)](fs ...F) DelegateVoid4[A1, A2, A3, A4] {
	d := make(DelegateVoid4[A1, A2, A3, A4], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction4(fs[i]))
	}
	return d
}

func CastDelegateVoid5[A1, A2, A3, A4, A5 any, F ~func(
	A1, A2, A3, A4, A5,
)](fs ...F) DelegateVoid5[A1, A2, A3, A4, A5] {
	d := make(DelegateVoid5[A1, A2, A3, A4, A5], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction5(fs[i]))
	}
	return d
}

func CastDelegateVoid6[A1, A2, A3, A4, A5, A6 any, F ~func(
	A1, A2, A3, A4, A5, A6,
)](fs ...F) DelegateVoid6[A1, A2, A3, A4, A5, A6] {
	d := make(DelegateVoid6[A1, A2, A3, A4, A5, A6], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction6(fs[i]))
	}
	return d
}

func CastDelegateVoid7[A1, A2, A3, A4, A5, A6, A7 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7,
)](fs ...F) DelegateVoid7[A1, A2, A3, A4, A5, A6, A7] {
	d := make(DelegateVoid7[A1, A2, A3, A4, A5, A6, A7], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction7(fs[i]))
	}
	return d
}

func CastDelegateVoid8[A1, A2, A3, A4, A5, A6, A7, A8 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8,
)](fs ...F) DelegateVoid8[A1, A2, A3, A4, A5, A6, A7, A8] {
	d := make(DelegateVoid8[A1, A2, A3, A4, A5, A6, A7, A8], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction8(fs[i]))
	}
	return d
}

func CastDelegateVoid9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
)](fs ...F) DelegateVoid9[A1, A2, A3, A4, A5, A6, A7, A8, A9] {
	d := make(DelegateVoid9[A1, A2, A3, A4, A5, A6, A7, A8, A9], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction9(fs[i]))
	}
	return d
}

func CastDelegateVoid10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
)](fs ...F) DelegateVoid10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10] {
	d := make(DelegateVoid10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction10(fs[i]))
	}
	return d
}

func CastDelegateVoid11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
)](fs ...F) DelegateVoid11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11] {
	d := make(DelegateVoid11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction11(fs[i]))
	}
	return d
}

func CastDelegateVoid12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
)](fs ...F) DelegateVoid12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12] {
	d := make(DelegateVoid12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction12(fs[i]))
	}
	return d
}

func CastDelegateVoid13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
)](fs ...F) DelegateVoid13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13] {
	d := make(DelegateVoid13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction13(fs[i]))
	}
	return d
}

func CastDelegateVoid14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
)](fs ...F) DelegateVoid14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14] {
	d := make(DelegateVoid14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction14(fs[i]))
	}
	return d
}

func CastDelegateVoid15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
)](fs ...F) DelegateVoid15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15] {
	d := make(DelegateVoid15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction15(fs[i]))
	}
	return d
}

func CastDelegateVoid16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
)](fs ...F) DelegateVoid16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16] {
	d := make(DelegateVoid16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction16(fs[i]))
	}
	return d
}
