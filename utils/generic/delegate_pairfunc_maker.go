package generic

func MakeDelegatePairFunc0[R1, R2 any, F ~func() (R1, R2)](fs ...F) DelegatePairFunc0[R1, R2] {
	d := make(DelegatePairFunc0[R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc0(fs[i]))
	}
	return d
}

func MakeDelegatePairFunc1[A1, R1, R2 any, F ~func(A1) (R1, R2)](fs ...F) DelegatePairFunc1[A1, R1, R2] {
	d := make(DelegatePairFunc1[A1, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc1(fs[i]))
	}
	return d
}

func MakeDelegatePairFunc2[A1, A2, R1, R2 any, F ~func(A1, A2) (R1, R2)](fs ...F) DelegatePairFunc2[A1, A2, R1, R2] {
	d := make(DelegatePairFunc2[A1, A2, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc2(fs[i]))
	}
	return d
}

func MakeDelegatePairFunc3[A1, A2, A3, R1, R2 any, F ~func(A1, A2, A3) (R1, R2)](fs ...F) DelegatePairFunc3[A1, A2, A3, R1, R2] {
	d := make(DelegatePairFunc3[A1, A2, A3, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc3(fs[i]))
	}
	return d
}

func MakeDelegatePairFunc4[A1, A2, A3, A4, R1, R2 any, F ~func(A1, A2, A3, A4) (R1, R2)](fs ...F) DelegatePairFunc4[A1, A2, A3, A4, R1, R2] {
	d := make(DelegatePairFunc4[A1, A2, A3, A4, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc4(fs[i]))
	}
	return d
}

func MakeDelegatePairFunc5[A1, A2, A3, A4, A5, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5,
) (R1, R2)](fs ...F) DelegatePairFunc5[A1, A2, A3, A4, A5, R1, R2] {
	d := make(DelegatePairFunc5[A1, A2, A3, A4, A5, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc5(fs[i]))
	}
	return d
}

func MakeDelegatePairFunc6[A1, A2, A3, A4, A5, A6, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6,
) (R1, R2)](fs ...F) DelegatePairFunc6[A1, A2, A3, A4, A5, A6, R1, R2] {
	d := make(DelegatePairFunc6[A1, A2, A3, A4, A5, A6, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc6(fs[i]))
	}
	return d
}

func MakeDelegatePairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7,
) (R1, R2)](fs ...F) DelegatePairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2] {
	d := make(DelegatePairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc7(fs[i]))
	}
	return d
}

func MakeDelegatePairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8,
) (R1, R2)](fs ...F) DelegatePairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2] {
	d := make(DelegatePairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc8(fs[i]))
	}
	return d
}

func MakeDelegatePairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
) (R1, R2)](fs ...F) DelegatePairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2] {
	d := make(DelegatePairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc9(fs[i]))
	}
	return d
}

func MakeDelegatePairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
) (R1, R2)](fs ...F) DelegatePairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2] {
	d := make(DelegatePairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc10(fs[i]))
	}
	return d
}

func MakeDelegatePairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
) (R1, R2)](fs ...F) DelegatePairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2] {
	d := make(DelegatePairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc11(fs[i]))
	}
	return d
}

func MakeDelegatePairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
) (R1, R2)](fs ...F) DelegatePairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2] {
	d := make(DelegatePairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc12(fs[i]))
	}
	return d
}

func MakeDelegatePairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
) (R1, R2)](fs ...F) DelegatePairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2] {
	d := make(DelegatePairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc13(fs[i]))
	}
	return d
}

func MakeDelegatePairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
) (R1, R2)](fs ...F) DelegatePairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2] {
	d := make(DelegatePairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc14(fs[i]))
	}
	return d
}

func MakeDelegatePairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
) (R1, R2)](fs ...F) DelegatePairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2] {
	d := make(DelegatePairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc15(fs[i]))
	}
	return d
}

func MakeDelegatePairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
) (R1, R2)](fs ...F) DelegatePairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2] {
	d := make(DelegatePairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2], 0, len(fs))
	for i := range fs {
		d = append(d, MakePairFunc16(fs[i]))
	}
	return d
}
