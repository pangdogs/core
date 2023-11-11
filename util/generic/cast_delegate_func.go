package generic

func CastDelegateFunc0[R any, F ~func() R](fs ...F) DelegateFunc0[R] {
	d := make(DelegateFunc0[R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc0(fs[i]))
	}
	return d
}

func CastDelegateFunc1[A1, R any, F ~func(A1) R](fs ...F) DelegateFunc1[A1, R] {
	d := make(DelegateFunc1[A1, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc1(fs[i]))
	}
	return d
}

func CastDelegateFunc2[A1, A2, R any, F ~func(A1, A2) R](fs ...F) DelegateFunc2[A1, A2, R] {
	d := make(DelegateFunc2[A1, A2, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc2(fs[i]))
	}
	return d
}

func CastDelegateFunc3[A1, A2, A3, R any, F ~func(A1, A2, A3) R](fs ...F) DelegateFunc3[A1, A2, A3, R] {
	d := make(DelegateFunc3[A1, A2, A3, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc3(fs[i]))
	}
	return d
}

func CastDelegateFunc4[A1, A2, A3, A4, R any, F ~func(A1, A2, A3, A4) R](fs ...F) DelegateFunc4[A1, A2, A3, A4, R] {
	d := make(DelegateFunc4[A1, A2, A3, A4, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc4(fs[i]))
	}
	return d
}

func CastDelegateFunc5[A1, A2, A3, A4, A5, R any, F ~func(
	A1, A2, A3, A4, A5,
) R](fs ...F) DelegateFunc5[A1, A2, A3, A4, A5, R] {
	d := make(DelegateFunc5[A1, A2, A3, A4, A5, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc5(fs[i]))
	}
	return d
}

func CastDelegateFunc6[A1, A2, A3, A4, A5, A6, R any, F ~func(
	A1, A2, A3, A4, A5, A6,
) R](fs ...F) DelegateFunc6[A1, A2, A3, A4, A5, A6, R] {
	d := make(DelegateFunc6[A1, A2, A3, A4, A5, A6, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc6(fs[i]))
	}
	return d
}

func CastDelegateFunc7[A1, A2, A3, A4, A5, A6, A7, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7,
) R](fs ...F) DelegateFunc7[A1, A2, A3, A4, A5, A6, A7, R] {
	d := make(DelegateFunc7[A1, A2, A3, A4, A5, A6, A7, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc7(fs[i]))
	}
	return d
}

func CastDelegateFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8,
) R](fs ...F) DelegateFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	d := make(DelegateFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc8(fs[i]))
	}
	return d
}

func CastDelegateFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
) R](fs ...F) DelegateFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	d := make(DelegateFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc9(fs[i]))
	}
	return d
}

func CastDelegateFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
) R](fs ...F) DelegateFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {
	d := make(DelegateFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc10(fs[i]))
	}
	return d
}

func CastDelegateFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
) R](fs ...F) DelegateFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {
	d := make(DelegateFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc11(fs[i]))
	}
	return d
}

func CastDelegateFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
) R](fs ...F) DelegateFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {
	d := make(DelegateFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc12(fs[i]))
	}
	return d
}

func CastDelegateFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
) R](fs ...F) DelegateFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {
	d := make(DelegateFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc13(fs[i]))
	}
	return d
}

func CastDelegateFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
) R](fs ...F) DelegateFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {
	d := make(DelegateFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc14(fs[i]))
	}
	return d
}

func CastDelegateFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
) R](fs ...F) DelegateFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R] {
	d := make(DelegateFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc15(fs[i]))
	}
	return d
}

func CastDelegateFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
) R](fs ...F) DelegateFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R] {
	d := make(DelegateFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R], 0, len(fs))
	for i := range fs {
		d = append(d, CastFunc16(fs[i]))
	}
	return d
}
