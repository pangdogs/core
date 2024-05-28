package generic

func MakeDelegateFunc0[R any, F ~func() R](fs ...F) DelegateFunc0[R] {
	d := make(DelegateFunc0[R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc0(fs[i]))
	}
	return d
}

func MakeDelegateFunc1[A1, R any, F ~func(A1) R](fs ...F) DelegateFunc1[A1, R] {
	d := make(DelegateFunc1[A1, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc1(fs[i]))
	}
	return d
}

func MakeDelegateFunc2[A1, A2, R any, F ~func(A1, A2) R](fs ...F) DelegateFunc2[A1, A2, R] {
	d := make(DelegateFunc2[A1, A2, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc2(fs[i]))
	}
	return d
}

func MakeDelegateFunc3[A1, A2, A3, R any, F ~func(A1, A2, A3) R](fs ...F) DelegateFunc3[A1, A2, A3, R] {
	d := make(DelegateFunc3[A1, A2, A3, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc3(fs[i]))
	}
	return d
}

func MakeDelegateFunc4[A1, A2, A3, A4, R any, F ~func(A1, A2, A3, A4) R](fs ...F) DelegateFunc4[A1, A2, A3, A4, R] {
	d := make(DelegateFunc4[A1, A2, A3, A4, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc4(fs[i]))
	}
	return d
}

func MakeDelegateFunc5[A1, A2, A3, A4, A5, R any, F ~func(
	A1, A2, A3, A4, A5,
) R](fs ...F) DelegateFunc5[A1, A2, A3, A4, A5, R] {
	d := make(DelegateFunc5[A1, A2, A3, A4, A5, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc5(fs[i]))
	}
	return d
}

func MakeDelegateFunc6[A1, A2, A3, A4, A5, A6, R any, F ~func(
	A1, A2, A3, A4, A5, A6,
) R](fs ...F) DelegateFunc6[A1, A2, A3, A4, A5, A6, R] {
	d := make(DelegateFunc6[A1, A2, A3, A4, A5, A6, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc6(fs[i]))
	}
	return d
}

func MakeDelegateFunc7[A1, A2, A3, A4, A5, A6, A7, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7,
) R](fs ...F) DelegateFunc7[A1, A2, A3, A4, A5, A6, A7, R] {
	d := make(DelegateFunc7[A1, A2, A3, A4, A5, A6, A7, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc7(fs[i]))
	}
	return d
}

func MakeDelegateFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8,
) R](fs ...F) DelegateFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	d := make(DelegateFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc8(fs[i]))
	}
	return d
}

func MakeDelegateFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
) R](fs ...F) DelegateFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	d := make(DelegateFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc9(fs[i]))
	}
	return d
}

func MakeDelegateFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
) R](fs ...F) DelegateFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {
	d := make(DelegateFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc10(fs[i]))
	}
	return d
}

func MakeDelegateFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
) R](fs ...F) DelegateFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {
	d := make(DelegateFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc11(fs[i]))
	}
	return d
}

func MakeDelegateFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
) R](fs ...F) DelegateFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {
	d := make(DelegateFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc12(fs[i]))
	}
	return d
}

func MakeDelegateFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
) R](fs ...F) DelegateFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {
	d := make(DelegateFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc13(fs[i]))
	}
	return d
}

func MakeDelegateFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
) R](fs ...F) DelegateFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {
	d := make(DelegateFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc14(fs[i]))
	}
	return d
}

func MakeDelegateFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
) R](fs ...F) DelegateFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R] {
	d := make(DelegateFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc15(fs[i]))
	}
	return d
}

func MakeDelegateFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
) R](fs ...F) DelegateFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R] {
	d := make(DelegateFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R], 0, len(fs))
	for i := range fs {
		d = append(d, MakeFunc16(fs[i]))
	}
	return d
}
