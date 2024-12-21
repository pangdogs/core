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

func CastDelegatePairVar0[VA, R1, R2 any, F ~func(...VA) (R1, R2)](fs ...F) DelegatePairVar0[VA, R1, R2] {
	d := make(DelegatePairVar0[VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar0(fs[i]))
	}
	return d
}

func CastDelegatePairVar1[A1, VA, R1, R2 any, F ~func(A1, ...VA) (R1, R2)](fs ...F) DelegatePairVar1[A1, VA, R1, R2] {
	d := make(DelegatePairVar1[A1, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar1(fs[i]))
	}
	return d
}

func CastDelegatePairVar2[A1, A2, VA, R1, R2 any, F ~func(A1, A2, ...VA) (R1, R2)](fs ...F) DelegatePairVar2[A1, A2, VA, R1, R2] {
	d := make(DelegatePairVar2[A1, A2, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar2(fs[i]))
	}
	return d
}

func CastDelegatePairVar3[A1, A2, A3, VA, R1, R2 any, F ~func(A1, A2, A3, ...VA) (R1, R2)](fs ...F) DelegatePairVar3[A1, A2, A3, VA, R1, R2] {
	d := make(DelegatePairVar3[A1, A2, A3, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar3(fs[i]))
	}
	return d
}

func CastDelegatePairVar4[A1, A2, A3, A4, VA, R1, R2 any, F ~func(A1, A2, A3, A4, ...VA) (R1, R2)](fs ...F) DelegatePairVar4[A1, A2, A3, A4, VA, R1, R2] {
	d := make(DelegatePairVar4[A1, A2, A3, A4, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar4(fs[i]))
	}
	return d
}

func CastDelegatePairVar5[A1, A2, A3, A4, A5, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, ...VA,
) (R1, R2)](fs ...F) DelegatePairVar5[A1, A2, A3, A4, A5, VA, R1, R2] {
	d := make(DelegatePairVar5[A1, A2, A3, A4, A5, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar5(fs[i]))
	}
	return d
}

func CastDelegatePairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, ...VA,
) (R1, R2)](fs ...F) DelegatePairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2] {
	d := make(DelegatePairVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar6(fs[i]))
	}
	return d
}

func CastDelegatePairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
) (R1, R2)](fs ...F) DelegatePairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2] {
	d := make(DelegatePairVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar7(fs[i]))
	}
	return d
}

func CastDelegatePairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
) (R1, R2)](fs ...F) DelegatePairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2] {
	d := make(DelegatePairVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar8(fs[i]))
	}
	return d
}

func CastDelegatePairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
) (R1, R2)](fs ...F) DelegatePairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2] {
	d := make(DelegatePairVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar9(fs[i]))
	}
	return d
}

func CastDelegatePairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
) (R1, R2)](fs ...F) DelegatePairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2] {
	d := make(DelegatePairVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar10(fs[i]))
	}
	return d
}

func CastDelegatePairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
) (R1, R2)](fs ...F) DelegatePairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2] {
	d := make(DelegatePairVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar11(fs[i]))
	}
	return d
}

func CastDelegatePairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
) (R1, R2)](fs ...F) DelegatePairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2] {
	d := make(DelegatePairVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar12(fs[i]))
	}
	return d
}

func CastDelegatePairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
) (R1, R2)](fs ...F) DelegatePairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2] {
	d := make(DelegatePairVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar13(fs[i]))
	}
	return d
}

func CastDelegatePairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
) (R1, R2)](fs ...F) DelegatePairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2] {
	d := make(DelegatePairVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar14(fs[i]))
	}
	return d
}

func CastDelegatePairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
) (R1, R2)](fs ...F) DelegatePairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2] {
	d := make(DelegatePairVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar15(fs[i]))
	}
	return d
}

func CastDelegatePairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
) (R1, R2)](fs ...F) DelegatePairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2] {
	d := make(DelegatePairVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastFuncPairVar16(fs[i]))
	}
	return d
}
