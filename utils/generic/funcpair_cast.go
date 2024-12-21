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

func CastFuncPair0[R1, R2 any, F ~func() (R1, R2)](f F) FuncPair0[R1, R2] {
	return FuncPair0[R1, R2](f)
}

func CastFuncPair1[A1, R1, R2 any, F ~func(A1) (R1, R2)](f F) FuncPair1[A1, R1, R2] {
	return FuncPair1[A1, R1, R2](f)
}

func CastFuncPair2[A1, A2, R1, R2 any, F ~func(A1, A2) (R1, R2)](f F) FuncPair2[A1, A2, R1, R2] {
	return FuncPair2[A1, A2, R1, R2](f)
}

func CastFuncPair3[A1, A2, A3, R1, R2 any, F ~func(A1, A2, A3) (R1, R2)](f F) FuncPair3[A1, A2, A3, R1, R2] {
	return FuncPair3[A1, A2, A3, R1, R2](f)
}

func CastFuncPair4[A1, A2, A3, A4, R1, R2 any, F ~func(A1, A2, A3, A4) (R1, R2)](f F) FuncPair4[A1, A2, A3, A4, R1, R2] {
	return FuncPair4[A1, A2, A3, A4, R1, R2](f)
}

func CastFuncPair5[A1, A2, A3, A4, A5, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5,
) (R1, R2)](f F) FuncPair5[A1, A2, A3, A4, A5, R1, R2] {
	return FuncPair5[A1, A2, A3, A4, A5, R1, R2](f)
}

func CastFuncPair6[A1, A2, A3, A4, A5, A6, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6,
) (R1, R2)](f F) FuncPair6[A1, A2, A3, A4, A5, A6, R1, R2] {
	return FuncPair6[A1, A2, A3, A4, A5, A6, R1, R2](f)
}

func CastFuncPair7[A1, A2, A3, A4, A5, A6, A7, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7,
) (R1, R2)](f F) FuncPair7[A1, A2, A3, A4, A5, A6, A7, R1, R2] {
	return FuncPair7[A1, A2, A3, A4, A5, A6, A7, R1, R2](f)
}

func CastFuncPair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8,
) (R1, R2)](f F) FuncPair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2] {
	return FuncPair8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2](f)
}

func CastFuncPair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
) (R1, R2)](f F) FuncPair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2] {
	return FuncPair9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2](f)
}

func CastFuncPair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
) (R1, R2)](f F) FuncPair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2] {
	return FuncPair10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2](f)
}

func CastFuncPair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
) (R1, R2)](f F) FuncPair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2] {
	return FuncPair11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2](f)
}

func CastFuncPair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
) (R1, R2)](f F) FuncPair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2] {
	return FuncPair12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2](f)
}

func CastFuncPair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
) (R1, R2)](f F) FuncPair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2] {
	return FuncPair13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2](f)
}

func CastFuncPair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
) (R1, R2)](f F) FuncPair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2] {
	return FuncPair14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2](f)
}

func CastFuncPair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
) (R1, R2)](f F) FuncPair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2] {
	return FuncPair15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2](f)
}

func CastFuncPair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
) (R1, R2)](f F) FuncPair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2] {
	return FuncPair16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2](f)
}
