package generic

func CastDelegateActionVar0[VA any, F ~func(...VA)](fs ...F) DelegateActionVar0[VA] {
	d := make(DelegateActionVar0[VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar0(fs[i]))
	}
	return d
}

func CastDelegateActionVar1[A1, VA, F ~func(A1, ...VA)](fs ...F) DelegateActionVar1[A1, VA] {
	d := make(DelegateActionVar1[A1, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar1(fs[i]))
	}
	return d
}

func CastDelegateActionVar2[A1, A2, VA any, F ~func(A1, A2, ...VA)](fs ...F) DelegateActionVar2[A1, A2, VA] {
	d := make(DelegateActionVar2[A1, A2, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar2(fs[i]))
	}
	return d
}

func CastDelegateActionVar3[A1, A2, A3, VA any, F ~func(A1, A2, A3, ...VA)](fs ...F) DelegateActionVar3[A1, A2, A3, VA] {
	d := make(DelegateActionVar3[A1, A2, A3, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar3(fs[i]))
	}
	return d
}

func CastDelegateActionVar4[A1, A2, A3, A4, VA any, F ~func(A1, A2, A3, A4, ...VA)](fs ...F) DelegateActionVar4[A1, A2, A3, A4, VA] {
	d := make(DelegateActionVar4[A1, A2, A3, A4, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar4(fs[i]))
	}
	return d
}

func CastDelegateActionVar5[A1, A2, A3, A4, A5, VA any, F ~func(
	A1, A2, A3, A4, A5, ...VA,
)](fs ...F) DelegateActionVar5[A1, A2, A3, A4, A5, VA] {
	d := make(DelegateActionVar5[A1, A2, A3, A4, A5, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar5(fs[i]))
	}
	return d
}

func CastDelegateActionVar6[A1, A2, A3, A4, A5, A6, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, ...VA,
)](fs ...F) DelegateActionVar6[A1, A2, A3, A4, A5, A6, VA] {
	d := make(DelegateActionVar6[A1, A2, A3, A4, A5, A6, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar6(fs[i]))
	}
	return d
}

func CastDelegateActionVar7[A1, A2, A3, A4, A5, A6, A7, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
)](fs ...F) DelegateActionVar7[A1, A2, A3, A4, A5, A6, A7, VA] {
	d := make(DelegateActionVar7[A1, A2, A3, A4, A5, A6, A7, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar7(fs[i]))
	}
	return d
}

func CastDelegateActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
)](fs ...F) DelegateActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA] {
	d := make(DelegateActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar8(fs[i]))
	}
	return d
}

func CastDelegateActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
)](fs ...F) DelegateActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA] {
	d := make(DelegateActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar9(fs[i]))
	}
	return d
}

func CastDelegateActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
)](fs ...F) DelegateActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA] {
	d := make(DelegateActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar10(fs[i]))
	}
	return d
}

func CastDelegateActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
)](fs ...F) DelegateActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA] {
	d := make(DelegateActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar11(fs[i]))
	}
	return d
}

func CastDelegateActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
)](fs ...F) DelegateActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA] {
	d := make(DelegateActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar12(fs[i]))
	}
	return d
}

func CastDelegateActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
)](fs ...F) DelegateActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA] {
	d := make(DelegateActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar13(fs[i]))
	}
	return d
}

func CastDelegateActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
)](fs ...F) DelegateActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA] {
	d := make(DelegateActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar14(fs[i]))
	}
	return d
}

func CastDelegateActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
)](fs ...F) DelegateActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA] {
	d := make(DelegateActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar15(fs[i]))
	}
	return d
}

func CastDelegateActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
)](fs ...F) DelegateActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA] {
	d := make(DelegateActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA], 0, len(fs))
	for i := range fs {
		d = append(d, CastActionVar16(fs[i]))
	}
	return d
}
