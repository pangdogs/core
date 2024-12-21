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

func CastDelegateVoidVar0[VA any, F ~func(...VA)](fs ...F) DelegateVoidVar0[VA] {
	d := make(DelegateVoidVar0[VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar0(fs[i]))
	}
	return d
}

func CastDelegateVoidVar1[A1, VA, F ~func(A1, ...VA)](fs ...F) DelegateVoidVar1[A1, VA] {
	d := make(DelegateVoidVar1[A1, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar1(fs[i]))
	}
	return d
}

func CastDelegateVoidVar2[A1, A2, VA any, F ~func(A1, A2, ...VA)](fs ...F) DelegateVoidVar2[A1, A2, VA] {
	d := make(DelegateVoidVar2[A1, A2, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar2(fs[i]))
	}
	return d
}

func CastDelegateVoidVar3[A1, A2, A3, VA any, F ~func(A1, A2, A3, ...VA)](fs ...F) DelegateVoidVar3[A1, A2, A3, VA] {
	d := make(DelegateVoidVar3[A1, A2, A3, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar3(fs[i]))
	}
	return d
}

func CastDelegateVoidVar4[A1, A2, A3, A4, VA any, F ~func(A1, A2, A3, A4, ...VA)](fs ...F) DelegateVoidVar4[A1, A2, A3, A4, VA] {
	d := make(DelegateVoidVar4[A1, A2, A3, A4, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar4(fs[i]))
	}
	return d
}

func CastDelegateVoidVar5[A1, A2, A3, A4, A5, VA any, F ~func(
	A1, A2, A3, A4, A5, ...VA,
)](fs ...F) DelegateVoidVar5[A1, A2, A3, A4, A5, VA] {
	d := make(DelegateVoidVar5[A1, A2, A3, A4, A5, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar5(fs[i]))
	}
	return d
}

func CastDelegateVoidVar6[A1, A2, A3, A4, A5, A6, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, ...VA,
)](fs ...F) DelegateVoidVar6[A1, A2, A3, A4, A5, A6, VA] {
	d := make(DelegateVoidVar6[A1, A2, A3, A4, A5, A6, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar6(fs[i]))
	}
	return d
}

func CastDelegateVoidVar7[A1, A2, A3, A4, A5, A6, A7, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
)](fs ...F) DelegateVoidVar7[A1, A2, A3, A4, A5, A6, A7, VA] {
	d := make(DelegateVoidVar7[A1, A2, A3, A4, A5, A6, A7, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar7(fs[i]))
	}
	return d
}

func CastDelegateVoidVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
)](fs ...F) DelegateVoidVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA] {
	d := make(DelegateVoidVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar8(fs[i]))
	}
	return d
}

func CastDelegateVoidVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
)](fs ...F) DelegateVoidVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA] {
	d := make(DelegateVoidVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar9(fs[i]))
	}
	return d
}

func CastDelegateVoidVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
)](fs ...F) DelegateVoidVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA] {
	d := make(DelegateVoidVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar10(fs[i]))
	}
	return d
}

func CastDelegateVoidVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
)](fs ...F) DelegateVoidVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA] {
	d := make(DelegateVoidVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar11(fs[i]))
	}
	return d
}

func CastDelegateVoidVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
)](fs ...F) DelegateVoidVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA] {
	d := make(DelegateVoidVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar12(fs[i]))
	}
	return d
}

func CastDelegateVoidVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
)](fs ...F) DelegateVoidVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA] {
	d := make(DelegateVoidVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar13(fs[i]))
	}
	return d
}

func CastDelegateVoidVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
)](fs ...F) DelegateVoidVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA] {
	d := make(DelegateVoidVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar14(fs[i]))
	}
	return d
}

func CastDelegateVoidVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
)](fs ...F) DelegateVoidVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA] {
	d := make(DelegateVoidVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar15(fs[i]))
	}
	return d
}

func CastDelegateVoidVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
)](fs ...F) DelegateVoidVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA] {
	d := make(DelegateVoidVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar16(fs[i]))
	}
	return d
}
