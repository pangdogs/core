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

func CastFunc0[R any, F ~func() R](f F) Func0[R] {
	return Func0[R](f)
}

func CastFunc1[A1, R any, F ~func(A1) R](f F) Func1[A1, R] {
	return Func1[A1, R](f)
}

func CastFunc2[A1, A2, R any, F ~func(A1, A2) R](f F) Func2[A1, A2, R] {
	return Func2[A1, A2, R](f)
}

func CastFunc3[A1, A2, A3, R any, F ~func(A1, A2, A3) R](f F) Func3[A1, A2, A3, R] {
	return Func3[A1, A2, A3, R](f)
}

func CastFunc4[A1, A2, A3, A4, R any, F ~func(A1, A2, A3, A4) R](f F) Func4[A1, A2, A3, A4, R] {
	return Func4[A1, A2, A3, A4, R](f)
}

func CastFunc5[A1, A2, A3, A4, A5, R any, F ~func(
	A1, A2, A3, A4, A5,
) R](f F) Func5[A1, A2, A3, A4, A5, R] {
	return Func5[A1, A2, A3, A4, A5, R](f)
}

func CastFunc6[A1, A2, A3, A4, A5, A6, R any, F ~func(
	A1, A2, A3, A4, A5, A6,
) R](f F) Func6[A1, A2, A3, A4, A5, A6, R] {
	return Func6[A1, A2, A3, A4, A5, A6, R](f)
}

func CastFunc7[A1, A2, A3, A4, A5, A6, A7, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7,
) R](f F) Func7[A1, A2, A3, A4, A5, A6, A7, R] {
	return Func7[A1, A2, A3, A4, A5, A6, A7, R](f)
}

func CastFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8,
) R](f F) Func8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return Func8[A1, A2, A3, A4, A5, A6, A7, A8, R](f)
}

func CastFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
) R](f F) Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R](f)
}

func CastFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
) R](f F) Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {
	return Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R](f)
}

func CastFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
) R](f F) Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {
	return Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R](f)
}

func CastFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
) R](f F) Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {
	return Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R](f)
}

func CastFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
) R](f F) Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {
	return Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R](f)
}

func CastFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
) R](f F) Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {
	return Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R](f)
}

func CastFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
) R](f F) Func15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R] {
	return Func15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R](f)
}

func CastFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
) R](f F) Func16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R] {
	return Func16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R](f)
}
