package generic

func CastDelegatePairFunc0[R1, R2 any, F ~func() (R1, R2)](fs ...F) DelegatePairFunc0[R1, R2] {
	d := make(DelegatePairFunc0[R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc0(fs[i]))
	}
	return d
}

func CastDelegatePairFunc1[A1, R1, R2 any, F ~func(A1) (R1, R2)](fs ...F) DelegatePairFunc1[A1, R1, R2] {
	d := make(DelegatePairFunc1[A1, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc1(fs[i]))
	}
	return d
}

func CastDelegatePairFunc2[A1, A2, R1, R2 any, F ~func(A1, A2) (R1, R2)](fs ...F) DelegatePairFunc2[A1, A2, R1, R2] {
	d := make(DelegatePairFunc2[A1, A2, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc2(fs[i]))
	}
	return d
}

func CastDelegatePairFunc3[A1, A2, A3, R1, R2 any, F ~func(A1, A2, A3) (R1, R2)](fs ...F) DelegatePairFunc3[A1, A2, A3, R1, R2] {
	d := make(DelegatePairFunc3[A1, A2, A3, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc3(fs[i]))
	}
	return d
}

func CastDelegatePairFunc4[A1, A2, A3, A4, R1, R2 any, F ~func(A1, A2, A3, A4) (R1, R2)](fs ...F) DelegatePairFunc4[A1, A2, A3, A4, R1, R2] {
	d := make(DelegatePairFunc4[A1, A2, A3, A4, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc4(fs[i]))
	}
	return d
}

func CastDelegatePairFunc5[A1, A2, A3, A4, A5, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5,
) (R1, R2)](fs ...F) DelegatePairFunc5[A1, A2, A3, A4, A5, R1, R2] {
	d := make(DelegatePairFunc5[A1, A2, A3, A4, A5, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc5(fs[i]))
	}
	return d
}

func CastDelegatePairFunc6[A1, A2, A3, A4, A5, A6, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6,
) (R1, R2)](fs ...F) DelegatePairFunc6[A1, A2, A3, A4, A5, A6, R1, R2] {
	d := make(DelegatePairFunc6[A1, A2, A3, A4, A5, A6, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc6(fs[i]))
	}
	return d
}

func CastDelegatePairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7,
) (R1, R2)](fs ...F) DelegatePairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2] {
	d := make(DelegatePairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc7(fs[i]))
	}
	return d
}

func CastDelegatePairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8,
) (R1, R2)](fs ...F) DelegatePairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2] {
	d := make(DelegatePairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc8(fs[i]))
	}
	return d
}

func CastDelegatePairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
) (R1, R2)](fs ...F) DelegatePairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2] {
	d := make(DelegatePairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc9(fs[i]))
	}
	return d
}

func CastDelegatePairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
) (R1, R2)](fs ...F) DelegatePairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2] {
	d := make(DelegatePairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc10(fs[i]))
	}
	return d
}

func CastDelegatePairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
) (R1, R2)](fs ...F) DelegatePairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2] {
	d := make(DelegatePairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc11(fs[i]))
	}
	return d
}

func CastDelegatePairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
) (R1, R2)](fs ...F) DelegatePairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2] {
	d := make(DelegatePairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc12(fs[i]))
	}
	return d
}

func CastDelegatePairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
) (R1, R2)](fs ...F) DelegatePairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2] {
	d := make(DelegatePairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc13(fs[i]))
	}
	return d
}

func CastDelegatePairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
) (R1, R2)](fs ...F) DelegatePairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2] {
	d := make(DelegatePairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc14(fs[i]))
	}
	return d
}

func CastDelegatePairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
) (R1, R2)](fs ...F) DelegatePairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2] {
	d := make(DelegatePairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc15(fs[i]))
	}
	return d
}

func CastDelegatePairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
) (R1, R2)](fs ...F) DelegatePairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2] {
	d := make(DelegatePairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, CastPairFunc16(fs[i]))
	}
	return d
}
