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

func CastFuncPairVar0[VA, R1, R2 any, F ~func(...VA) (R1, R2)](f F) FuncPairVar0[VA, R1, R2] {
	return FuncPairVar0[VA, R1, R2](f)
}

func CastFuncPairVar1[A1, VA, R1, R2 any, F ~func(A1, ...VA) (R1, R2)](f F) FuncPairVar1[A1, VA, R1, R2] {
	return FuncPairVar1[A1, VA, R1, R2](f)
}

func CastFuncPairVar2[A1, A2, VA, R1, R2 any, F ~func(A1, A2, ...VA) (R1, R2)](f F) FuncPairVar2[A1, A2, VA, R1, R2] {
	return FuncPairVar2[A1, A2, VA, R1, R2](f)
}

func CastFuncPairVar3[A1, A2, A3, VA, R1, R2 any, F ~func(A1, A2, A3, ...VA) (R1, R2)](f F) FuncPairVar3[A1, A2, A3, VA, R1, R2] {
	return FuncPairVar3[A1, A2, A3, VA, R1, R2](f)
}

func CastFuncPairVar4[A1, A2, A3, A4, VA, R1, R2 any, F ~func(A1, A2, A3, A4, ...VA) (R1, R2)](f F) FuncPairVar4[A1, A2, A3, A4, VA, R1, R2] {
	return FuncPairVar4[A1, A2, A3, A4, VA, R1, R2](f)
}

func CastFuncPairVar5[A1, A2, A3, A4, A5, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, ...VA,
) (R1, R2)](f F) FuncPairVar5[A1, A2, A3, A4, A5, VA, R1, R2] {
	return FuncPairVar5[A1, A2, A3, A4, A5, VA, R1, R2](f)
}

func CastFuncPairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, ...VA,
) (R1, R2)](f F) FuncPairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2] {
	return FuncPairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2](f)
}

func CastFuncPairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
) (R1, R2)](f F) FuncPairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2] {
	return FuncPairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2](f)
}

func CastFuncPairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
) (R1, R2)](f F) FuncPairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2] {
	return FuncPairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2](f)
}

func CastFuncPairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
) (R1, R2)](f F) FuncPairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2] {
	return FuncPairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2](f)
}

func CastFuncPairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
) (R1, R2)](f F) FuncPairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2] {
	return FuncPairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2](f)
}

func CastFuncPairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
) (R1, R2)](f F) FuncPairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2] {
	return FuncPairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2](f)
}

func CastFuncPairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
) (R1, R2)](f F) FuncPairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2] {
	return FuncPairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2](f)
}

func CastFuncPairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
) (R1, R2)](f F) FuncPairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2] {
	return FuncPairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2](f)
}

func CastFuncPairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
) (R1, R2)](f F) FuncPairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2] {
	return FuncPairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2](f)
}

func CastFuncPairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
) (R1, R2)](f F) FuncPairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2] {
	return FuncPairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2](f)
}

func CastFuncPairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
) (R1, R2)](f F) FuncPairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2] {
	return FuncPairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2](f)
}
