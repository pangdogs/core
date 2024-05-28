package generic

func MakeDelegateAction0[F ~func()](fs ...F) DelegateAction0 {
	d := make(DelegateAction0, 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction0(fs[i]))
	}
	return d
}

func MakeDelegateAction1[A1 any, F ~func(A1)](fs ...F) DelegateAction1[A1] {
	d := make(DelegateAction1[A1], 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction1(fs[i]))
	}
	return d
}

func MakeDelegateAction2[A1, A2 any, F ~func(A1, A2)](fs ...F) DelegateAction2[A1, A2] {
	d := make(DelegateAction2[A1, A2], 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction2(fs[i]))
	}
	return d
}

func MakeDelegateAction3[A1, A2, A3 any, F ~func(A1, A2, A3)](fs ...F) DelegateAction3[A1, A2, A3] {
	d := make(DelegateAction3[A1, A2, A3], 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction3(fs[i]))
	}
	return d
}

func MakeDelegateAction4[A1, A2, A3, A4 any, F ~func(A1, A2, A3, A4)](fs ...F) DelegateAction4[A1, A2, A3, A4] {
	d := make(DelegateAction4[A1, A2, A3, A4], 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction4(fs[i]))
	}
	return d
}

func MakeDelegateAction5[A1, A2, A3, A4, A5 any, F ~func(
	A1, A2, A3, A4, A5,
)](fs ...F) DelegateAction5[A1, A2, A3, A4, A5] {
	d := make(DelegateAction5[A1, A2, A3, A4, A5], 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction5(fs[i]))
	}
	return d
}

func MakeDelegateAction6[A1, A2, A3, A4, A5, A6 any, F ~func(
	A1, A2, A3, A4, A5, A6,
)](fs ...F) DelegateAction6[A1, A2, A3, A4, A5, A6] {
	d := make(DelegateAction6[A1, A2, A3, A4, A5, A6], 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction6(fs[i]))
	}
	return d
}

func MakeDelegateAction7[A1, A2, A3, A4, A5, A6, A7 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7,
)](fs ...F) DelegateAction7[A1, A2, A3, A4, A5, A6, A7] {
	d := make(DelegateAction7[A1, A2, A3, A4, A5, A6, A7], 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction7(fs[i]))
	}
	return d
}

func MakeDelegateAction8[A1, A2, A3, A4, A5, A6, A7, A8 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8,
)](fs ...F) DelegateAction8[A1, A2, A3, A4, A5, A6, A7, A8] {
	d := make(DelegateAction8[A1, A2, A3, A4, A5, A6, A7, A8], 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction8(fs[i]))
	}
	return d
}

func MakeDelegateAction9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
)](fs ...F) DelegateAction9[A1, A2, A3, A4, A5, A6, A7, A8, A9] {
	d := make(DelegateAction9[A1, A2, A3, A4, A5, A6, A7, A8, A9], 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction9(fs[i]))
	}
	return d
}

func MakeDelegateAction10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
)](fs ...F) DelegateAction10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10] {
	d := make(DelegateAction10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10], 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction10(fs[i]))
	}
	return d
}

func MakeDelegateAction11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
)](fs ...F) DelegateAction11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11] {
	d := make(DelegateAction11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11], 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction11(fs[i]))
	}
	return d
}

func MakeDelegateAction12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
)](fs ...F) DelegateAction12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12] {
	d := make(DelegateAction12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12], 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction12(fs[i]))
	}
	return d
}

func MakeDelegateAction13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
)](fs ...F) DelegateAction13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13] {
	d := make(DelegateAction13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13], 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction13(fs[i]))
	}
	return d
}

func MakeDelegateAction14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
)](fs ...F) DelegateAction14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14] {
	d := make(DelegateAction14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14], 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction14(fs[i]))
	}
	return d
}

func MakeDelegateAction15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
)](fs ...F) DelegateAction15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15] {
	d := make(DelegateAction15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15], 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction15(fs[i]))
	}
	return d
}

func MakeDelegateAction16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
)](fs ...F) DelegateAction16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16] {
	d := make(DelegateAction16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16], 0, len(fs))
	for i := range fs {
		d = append(d, MakeAction16(fs[i]))
	}
	return d
}
