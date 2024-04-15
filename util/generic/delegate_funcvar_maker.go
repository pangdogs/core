package generic

func MakeDelegateFuncVar0[VA, R any, F ~func(...VA) R](fs ...F) DelegateFuncVar0[VA, R] {
	d := make(DelegateFuncVar0[VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar0(fs[i]))
	}
	return d
}

func MakeDelegateFuncVar1[A1, VA, R any, F ~func(A1, ...VA) R](fs ...F) DelegateFuncVar1[A1, VA, R] {
	d := make(DelegateFuncVar1[A1, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar1(fs[i]))
	}
	return d
}

func MakeDelegateFuncVar2[A1, A2, VA, R any, F ~func(A1, A2, ...VA) R](fs ...F) DelegateFuncVar2[A1, A2, VA, R] {
	d := make(DelegateFuncVar2[A1, A2, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar2(fs[i]))
	}
	return d
}

func MakeDelegateFuncVar3[A1, A2, A3, VA, R any, F ~func(A1, A2, A3, ...VA) R](fs ...F) DelegateFuncVar3[A1, A2, A3, VA, R] {
	d := make(DelegateFuncVar3[A1, A2, A3, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar3(fs[i]))
	}
	return d
}

func MakeDelegateFuncVar4[A1, A2, A3, A4, VA, R any, F ~func(A1, A2, A3, A4, ...VA) R](fs ...F) DelegateFuncVar4[A1, A2, A3, A4, VA, R] {
	d := make(DelegateFuncVar4[A1, A2, A3, A4, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar4(fs[i]))
	}
	return d
}

func MakeDelegateFuncVar5[A1, A2, A3, A4, A5, VA, R any, F ~func(
	A1, A2, A3, A4, A5, ...VA,
) R](fs ...F) DelegateFuncVar5[A1, A2, A3, A4, A5, VA, R] {
	d := make(DelegateFuncVar5[A1, A2, A3, A4, A5, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar5(fs[i]))
	}
	return d
}

func MakeDelegateFuncVar6[A1, A2, A3, A4, A5, A6, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, ...VA,
) R](fs ...F) DelegateFuncVar6[A1, A2, A3, A4, A5, A6, VA, R] {
	d := make(DelegateFuncVar6[A1, A2, A3, A4, A5, A6, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar6(fs[i]))
	}
	return d
}

func MakeDelegateFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
) R](fs ...F) DelegateFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R] {
	d := make(DelegateFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar7(fs[i]))
	}
	return d
}

func MakeDelegateFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
) R](fs ...F) DelegateFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R] {
	d := make(DelegateFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar8(fs[i]))
	}
	return d
}

func MakeDelegateFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
) R](fs ...F) DelegateFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R] {
	d := make(DelegateFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar9(fs[i]))
	}
	return d
}

func MakeDelegateFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
) R](fs ...F) DelegateFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R] {
	d := make(DelegateFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar10(fs[i]))
	}
	return d
}

func MakeDelegateFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
) R](fs ...F) DelegateFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R] {
	d := make(DelegateFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar11(fs[i]))
	}
	return d
}

func MakeDelegateFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
) R](fs ...F) DelegateFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R] {
	d := make(DelegateFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar12(fs[i]))
	}
	return d
}

func MakeDelegateFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
) R](fs ...F) DelegateFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R] {
	d := make(DelegateFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar13(fs[i]))
	}
	return d
}

func MakeDelegateFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
) R](fs ...F) DelegateFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R] {
	d := make(DelegateFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar14(fs[i]))
	}
	return d
}

func MakeDelegateFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
) R](fs ...F) DelegateFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R] {
	d := make(DelegateFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar15(fs[i]))
	}
	return d
}

func MakeDelegateFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
) R](fs ...F) DelegateFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R] {
	d := make(DelegateFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFuncVar16(fs[i]))
	}
	return d
}
