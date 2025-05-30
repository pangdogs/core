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

func CastFuncVar0[VA, R any, F ~func(...VA) R](f F) FuncVar0[VA, R] {
	return FuncVar0[VA, R](f)
}

func CastFuncVar1[A1, VA, R any, F ~func(A1, ...VA) R](f F) FuncVar1[A1, VA, R] {
	return FuncVar1[A1, VA, R](f)
}

func CastFuncVar2[A1, A2, VA, R any, F ~func(A1, A2, ...VA) R](f F) FuncVar2[A1, A2, VA, R] {
	return FuncVar2[A1, A2, VA, R](f)
}

func CastFuncVar3[A1, A2, A3, VA, R any, F ~func(A1, A2, A3, ...VA) R](f F) FuncVar3[A1, A2, A3, VA, R] {
	return FuncVar3[A1, A2, A3, VA, R](f)
}

func CastFuncVar4[A1, A2, A3, A4, VA, R any, F ~func(A1, A2, A3, A4, ...VA) R](f F) FuncVar4[A1, A2, A3, A4, VA, R] {
	return FuncVar4[A1, A2, A3, A4, VA, R](f)
}

func CastFuncVar5[A1, A2, A3, A4, A5, VA, R any, F ~func(
	A1, A2, A3, A4, A5, ...VA,
) R](f F) FuncVar5[A1, A2, A3, A4, A5, VA, R] {
	return FuncVar5[A1, A2, A3, A4, A5, VA, R](f)
}

func CastFuncVar6[A1, A2, A3, A4, A5, A6, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, ...VA,
) R](f F) FuncVar6[A1, A2, A3, A4, A5, A6, VA, R] {
	return FuncVar6[A1, A2, A3, A4, A5, A6, VA, R](f)
}

func CastFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
) R](f F) FuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R] {
	return FuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R](f)
}

func CastFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
) R](f F) FuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R] {
	return FuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R](f)
}

func CastFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
) R](f F) FuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R] {
	return FuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R](f)
}

func CastFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
) R](f F) FuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R] {
	return FuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R](f)
}

func CastFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
) R](f F) FuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R] {
	return FuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R](f)
}

func CastFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
) R](f F) FuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R] {
	return FuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R](f)
}

func CastFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
) R](f F) FuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R] {
	return FuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R](f)
}

func CastFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
) R](f F) FuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R] {
	return FuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R](f)
}

func CastFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
) R](f F) FuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R] {
	return FuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R](f)
}

func CastFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
) R](f F) FuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R] {
	return FuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R](f)
}
