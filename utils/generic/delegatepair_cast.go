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

func CastDelegatePair0[R1, R2 any, F ~func() (R1, R2)](fs ...F) DelegatePair0[R1, R2] {
	d := make(DelegatePair0[R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair0(fs[i]))
	}
	return d
}

func CastDelegatePair1[A1, R1, R2 any, F ~func(A1) (R1, R2)](fs ...F) DelegatePair1[A1, R1, R2] {
	d := make(DelegatePair1[A1, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair1(fs[i]))
	}
	return d
}

func CastDelegatePair2[A1, A2, R1, R2 any, F ~func(A1, A2) (R1, R2)](fs ...F) DelegatePair2[A1, A2, R1, R2] {
	d := make(DelegatePair2[A1, A2, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair2(fs[i]))
	}
	return d
}

func CastDelegatePair3[A1, A2, A3, R1, R2 any, F ~func(A1, A2, A3) (R1, R2)](fs ...F) DelegatePair3[A1, A2, A3, R1, R2] {
	d := make(DelegatePair3[A1, A2, A3, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair3(fs[i]))
	}
	return d
}

func CastDelegatePair4[A1, A2, A3, A4, R1, R2 any, F ~func(A1, A2, A3, A4) (R1, R2)](fs ...F) DelegatePair4[A1, A2, A3, A4, R1, R2] {
	d := make(DelegatePair4[A1, A2, A3, A4, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair4(fs[i]))
	}
	return d
}

func CastDelegatePair5[A1, A2, A3, A4, A5, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5,
) (R1, R2)](fs ...F) DelegatePair5[A1, A2, A3, A4, A5, R1, R2] {
	d := make(DelegatePair5[A1, A2, A3, A4, A5, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair5(fs[i]))
	}
	return d
}

func CastDelegatePair6[A1, A2, A3, A4, A5, A6, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6,
) (R1, R2)](fs ...F) DelegatePair6[A1, A2, A3, A4, A5, A6, R1, R2] {
	d := make(DelegatePair6[A1, A2, A3, A4, A5, A6, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair6(fs[i]))
	}
	return d
}

func CastDelegatePair7[A1, A2, A3, A4, A5, A6, A7, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7,
) (R1, R2)](fs ...F) DelegatePair7[A1, A2, A3, A4, A5, A6, A7, R1, R2] {
	d := make(DelegatePair7[A1, A2, A3, A4, A5, A6, A7, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair7(fs[i]))
	}
	return d
}

func CastDelegatePair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8,
) (R1, R2)](fs ...F) DelegatePair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2] {
	d := make(DelegatePair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair8(fs[i]))
	}
	return d
}

func CastDelegatePair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
) (R1, R2)](fs ...F) DelegatePair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2] {
	d := make(DelegatePair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair9(fs[i]))
	}
	return d
}

func CastDelegatePair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
) (R1, R2)](fs ...F) DelegatePair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2] {
	d := make(DelegatePair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair10(fs[i]))
	}
	return d
}

func CastDelegatePair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
) (R1, R2)](fs ...F) DelegatePair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2] {
	d := make(DelegatePair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair11(fs[i]))
	}
	return d
}

func CastDelegatePair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
) (R1, R2)](fs ...F) DelegatePair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2] {
	d := make(DelegatePair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair12(fs[i]))
	}
	return d
}

func CastDelegatePair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
) (R1, R2)](fs ...F) DelegatePair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2] {
	d := make(DelegatePair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair13(fs[i]))
	}
	return d
}

func CastDelegatePair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
) (R1, R2)](fs ...F) DelegatePair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2] {
	d := make(DelegatePair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair14(fs[i]))
	}
	return d
}

func CastDelegatePair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
) (R1, R2)](fs ...F) DelegatePair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2] {
	d := make(DelegatePair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair15(fs[i]))
	}
	return d
}

func CastDelegatePair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
) (R1, R2)](fs ...F) DelegatePair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2] {
	d := make(DelegatePair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPair16(fs[i]))
	}
	return d
}
