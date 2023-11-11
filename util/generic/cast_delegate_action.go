package generic

func CastDelegateAction0[F ~func()](fs ...F) DelegateAction0 {
	d := make(DelegateAction0, 0, len(fs))
	for i := range fs {
		d = append(d, CastAction0(fs[i]))
	}
	return d
}

func CastDelegateAction1[A1 any, F ~func(A1)](fs ...F) DelegateAction1[A1] {
	d := make(DelegateAction1[A1], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction1(fs[i]))
	}
	return d
}

func CastDelegateAction2[A1, A2 any, F ~func(A1, A2)](fs ...F) DelegateAction2[A1, A2] {
	d := make(DelegateAction2[A1, A2], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction2(fs[i]))
	}
	return d
}

func CastDelegateAction3[A1, A2, A3 any, F ~func(A1, A2, A3)](fs ...F) DelegateAction3[A1, A2, A3] {
	d := make(DelegateAction3[A1, A2, A3], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction3(fs[i]))
	}
	return d
}

func CastDelegateAction4[A1, A2, A3, A4 any, F ~func(A1, A2, A3, A4)](fs ...F) DelegateAction4[A1, A2, A3, A4] {
	d := make(DelegateAction4[A1, A2, A3, A4], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction4(fs[i]))
	}
	return d
}

func CastDelegateAction5[A1, A2, A3, A4, A5 any, F ~func(
	A1, A2, A3, A4, A5,
)](fs ...F) DelegateAction5[A1, A2, A3, A4, A5] {
	d := make(DelegateAction5[A1, A2, A3, A4, A5], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction5(fs[i]))
	}
	return d
}

func CastDelegateAction6[A1, A2, A3, A4, A5, A6 any, F ~func(
	A1, A2, A3, A4, A5, A6,
)](fs ...F) DelegateAction6[A1, A2, A3, A4, A5, A6] {
	d := make(DelegateAction6[A1, A2, A3, A4, A5, A6], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction6(fs[i]))
	}
	return d
}

func CastDelegateAction7[A1, A2, A3, A4, A5, A6, A7 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7,
)](fs ...F) DelegateAction7[A1, A2, A3, A4, A5, A6, A7] {
	d := make(DelegateAction7[A1, A2, A3, A4, A5, A6, A7], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction7(fs[i]))
	}
	return d
}

func CastDelegateAction8[A1, A2, A3, A4, A5, A6, A7, A8 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8,
)](fs ...F) DelegateAction8[A1, A2, A3, A4, A5, A6, A7, A8] {
	d := make(DelegateAction8[A1, A2, A3, A4, A5, A6, A7, A8], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction8(fs[i]))
	}
	return d
}

func CastDelegateAction9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
)](fs ...F) DelegateAction9[A1, A2, A3, A4, A5, A6, A7, A8, A9] {
	d := make(DelegateAction9[A1, A2, A3, A4, A5, A6, A7, A8, A9], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction9(fs[i]))
	}
	return d
}

func CastDelegateAction10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
)](fs ...F) DelegateAction10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10] {
	d := make(DelegateAction10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction10(fs[i]))
	}
	return d
}

func CastDelegateAction11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
)](fs ...F) DelegateAction11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11] {
	d := make(DelegateAction11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction11(fs[i]))
	}
	return d
}

func CastDelegateAction12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
)](fs ...F) DelegateAction12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12] {
	d := make(DelegateAction12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction12(fs[i]))
	}
	return d
}

func CastDelegateAction13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
)](fs ...F) DelegateAction13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13] {
	d := make(DelegateAction13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction13(fs[i]))
	}
	return d
}

func CastDelegateAction14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
)](fs ...F) DelegateAction14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14] {
	d := make(DelegateAction14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction14(fs[i]))
	}
	return d
}

func CastDelegateAction15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
)](fs ...F) DelegateAction15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15] {
	d := make(DelegateAction15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction15(fs[i]))
	}
	return d
}

func CastDelegateAction16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
)](fs ...F) DelegateAction16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16] {
	d := make(DelegateAction16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16], 0, len(fs))
	for i := range fs {
		d = append(d, CastAction16(fs[i]))
	}
	return d
}
