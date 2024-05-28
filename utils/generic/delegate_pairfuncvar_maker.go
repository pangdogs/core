package generic

func MakeDelegatePairFuncVar0[VA, R1, R2 any, F ~func(...VA) (R1, R2)](fs ...F) DelegatePairFuncVar0[VA, R1, R2] {
	d := make(DelegatePairFuncVar0[VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar0(fs[i]))
	}
	return d
}

func MakeDelegatePairFuncVar1[A1, VA, R1, R2 any, F ~func(A1, ...VA) (R1, R2)](fs ...F) DelegatePairFuncVar1[A1, VA, R1, R2] {
	d := make(DelegatePairFuncVar1[A1, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar1(fs[i]))
	}
	return d
}

func MakeDelegatePairFuncVar2[A1, A2, VA, R1, R2 any, F ~func(A1, A2, ...VA) (R1, R2)](fs ...F) DelegatePairFuncVar2[A1, A2, VA, R1, R2] {
	d := make(DelegatePairFuncVar2[A1, A2, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar2(fs[i]))
	}
	return d
}

func MakeDelegatePairFuncVar3[A1, A2, A3, VA, R1, R2 any, F ~func(A1, A2, A3, ...VA) (R1, R2)](fs ...F) DelegatePairFuncVar3[A1, A2, A3, VA, R1, R2] {
	d := make(DelegatePairFuncVar3[A1, A2, A3, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar3(fs[i]))
	}
	return d
}

func MakeDelegatePairFuncVar4[A1, A2, A3, A4, VA, R1, R2 any, F ~func(A1, A2, A3, A4, ...VA) (R1, R2)](fs ...F) DelegatePairFuncVar4[A1, A2, A3, A4, VA, R1, R2] {
	d := make(DelegatePairFuncVar4[A1, A2, A3, A4, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar4(fs[i]))
	}
	return d
}

func MakeDelegatePairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, ...VA,
) (R1, R2)](fs ...F) DelegatePairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2] {
	d := make(DelegatePairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar5(fs[i]))
	}
	return d
}

func MakeDelegatePairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, ...VA,
) (R1, R2)](fs ...F) DelegatePairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2] {
	d := make(DelegatePairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar6(fs[i]))
	}
	return d
}

func MakeDelegatePairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
) (R1, R2)](fs ...F) DelegatePairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2] {
	d := make(DelegatePairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar7(fs[i]))
	}
	return d
}

func MakeDelegatePairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
) (R1, R2)](fs ...F) DelegatePairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2] {
	d := make(DelegatePairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar8(fs[i]))
	}
	return d
}

func MakeDelegatePairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
) (R1, R2)](fs ...F) DelegatePairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2] {
	d := make(DelegatePairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar9(fs[i]))
	}
	return d
}

func MakeDelegatePairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
) (R1, R2)](fs ...F) DelegatePairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2] {
	d := make(DelegatePairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar10(fs[i]))
	}
	return d
}

func MakeDelegatePairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
) (R1, R2)](fs ...F) DelegatePairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2] {
	d := make(DelegatePairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar11(fs[i]))
	}
	return d
}

func MakeDelegatePairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
) (R1, R2)](fs ...F) DelegatePairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2] {
	d := make(DelegatePairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar12(fs[i]))
	}
	return d
}

func MakeDelegatePairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
) (R1, R2)](fs ...F) DelegatePairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2] {
	d := make(DelegatePairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar13(fs[i]))
	}
	return d
}

func MakeDelegatePairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
) (R1, R2)](fs ...F) DelegatePairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2] {
	d := make(DelegatePairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar14(fs[i]))
	}
	return d
}

func MakeDelegatePairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
) (R1, R2)](fs ...F) DelegatePairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2] {
	d := make(DelegatePairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar15(fs[i]))
	}
	return d
}

func MakeDelegatePairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
) (R1, R2)](fs ...F) DelegatePairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2] {
	d := make(DelegatePairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFuncVar16(fs[i]))
	}
	return d
}
