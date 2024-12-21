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

func CastDelegateVar0[VA, R any, F ~func(...VA) R](fs ...F) DelegateVar0[VA, R] {
	d := make(DelegateVar0[VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar0(fs[i]))
	}
	return d
}

func CastDelegateVar1[A1, VA, R any, F ~func(A1, ...VA) R](fs ...F) DelegateVar1[A1, VA, R] {
	d := make(DelegateVar1[A1, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar1(fs[i]))
	}
	return d
}

func CastDelegateVar2[A1, A2, VA, R any, F ~func(A1, A2, ...VA) R](fs ...F) DelegateVar2[A1, A2, VA, R] {
	d := make(DelegateVar2[A1, A2, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar2(fs[i]))
	}
	return d
}

func CastDelegateVar3[A1, A2, A3, VA, R any, F ~func(A1, A2, A3, ...VA) R](fs ...F) DelegateVar3[A1, A2, A3, VA, R] {
	d := make(DelegateVar3[A1, A2, A3, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar3(fs[i]))
	}
	return d
}

func CastDelegateVar4[A1, A2, A3, A4, VA, R any, F ~func(A1, A2, A3, A4, ...VA) R](fs ...F) DelegateVar4[A1, A2, A3, A4, VA, R] {
	d := make(DelegateVar4[A1, A2, A3, A4, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar4(fs[i]))
	}
	return d
}

func CastDelegateVar5[A1, A2, A3, A4, A5, VA, R any, F ~func(
	A1, A2, A3, A4, A5, ...VA,
) R](fs ...F) DelegateVar5[A1, A2, A3, A4, A5, VA, R] {
	d := make(DelegateVar5[A1, A2, A3, A4, A5, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar5(fs[i]))
	}
	return d
}

func CastDelegateVar6[A1, A2, A3, A4, A5, A6, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, ...VA,
) R](fs ...F) DelegateVar6[A1, A2, A3, A4, A5, A6, VA, R] {
	d := make(DelegateVar6[A1, A2, A3, A4, A5, A6, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar6(fs[i]))
	}
	return d
}

func CastDelegateVar7[A1, A2, A3, A4, A5, A6, A7, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
) R](fs ...F) DelegateVar7[A1, A2, A3, A4, A5, A6, A7, VA, R] {
	d := make(DelegateVar7[A1, A2, A3, A4, A5, A6, A7, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar7(fs[i]))
	}
	return d
}

func CastDelegateVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
) R](fs ...F) DelegateVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R] {
	d := make(DelegateVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar8(fs[i]))
	}
	return d
}

func CastDelegateVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
) R](fs ...F) DelegateVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R] {
	d := make(DelegateVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar9(fs[i]))
	}
	return d
}

func CastDelegateVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
) R](fs ...F) DelegateVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R] {
	d := make(DelegateVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar10(fs[i]))
	}
	return d
}

func CastDelegateVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
) R](fs ...F) DelegateVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R] {
	d := make(DelegateVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar11(fs[i]))
	}
	return d
}

func CastDelegateVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
) R](fs ...F) DelegateVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R] {
	d := make(DelegateVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar12(fs[i]))
	}
	return d
}

func CastDelegateVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
) R](fs ...F) DelegateVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R] {
	d := make(DelegateVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar13(fs[i]))
	}
	return d
}

func CastDelegateVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
) R](fs ...F) DelegateVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R] {
	d := make(DelegateVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar14(fs[i]))
	}
	return d
}

func CastDelegateVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
) R](fs ...F) DelegateVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R] {
	d := make(DelegateVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar15(fs[i]))
	}
	return d
}

func CastDelegateVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
) R](fs ...F) DelegateVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R] {
	d := make(DelegateVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncVar16(fs[i]))
	}
	return d
}
