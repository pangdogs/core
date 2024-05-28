package generic

func MakePairFunc0[R1, R2 any, F ~func() (R1, R2)](f F) PairFunc0[R1, R2] {
	return PairFunc0[R1, R2](f)
}

func MakePairFunc1[A1, R1, R2 any, F ~func(A1) (R1, R2)](f F) PairFunc1[A1, R1, R2] {
	return PairFunc1[A1, R1, R2](f)
}

func MakePairFunc2[A1, A2, R1, R2 any, F ~func(A1, A2) (R1, R2)](f F) PairFunc2[A1, A2, R1, R2] {
	return PairFunc2[A1, A2, R1, R2](f)
}

func MakePairFunc3[A1, A2, A3, R1, R2 any, F ~func(A1, A2, A3) (R1, R2)](f F) PairFunc3[A1, A2, A3, R1, R2] {
	return PairFunc3[A1, A2, A3, R1, R2](f)
}

func MakePairFunc4[A1, A2, A3, A4, R1, R2 any, F ~func(A1, A2, A3, A4) (R1, R2)](f F) PairFunc4[A1, A2, A3, A4, R1, R2] {
	return PairFunc4[A1, A2, A3, A4, R1, R2](f)
}

func MakePairFunc5[A1, A2, A3, A4, A5, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5,
) (R1, R2)](f F) PairFunc5[A1, A2, A3, A4, A5, R1, R2] {
	return PairFunc5[A1, A2, A3, A4, A5, R1, R2](f)
}

func MakePairFunc6[A1, A2, A3, A4, A5, A6, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6,
) (R1, R2)](f F) PairFunc6[A1, A2, A3, A4, A5, A6, R1, R2] {
	return PairFunc6[A1, A2, A3, A4, A5, A6, R1, R2](f)
}

func MakePairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7,
) (R1, R2)](f F) PairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2] {
	return PairFunc7[A1, A2, A3, A4, A5, A6, A7, R1, R2](f)
}

func MakePairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8,
) (R1, R2)](f F) PairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2] {
	return PairFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R1, R2](f)
}

func MakePairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
) (R1, R2)](f F) PairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2] {
	return PairFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R1, R2](f)
}

func MakePairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
) (R1, R2)](f F) PairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2] {
	return PairFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R1, R2](f)
}

func MakePairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
) (R1, R2)](f F) PairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2] {
	return PairFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R1, R2](f)
}

func MakePairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
) (R1, R2)](f F) PairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2] {
	return PairFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R1, R2](f)
}

func MakePairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
) (R1, R2)](f F) PairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2] {
	return PairFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R1, R2](f)
}

func MakePairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
) (R1, R2)](f F) PairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2] {
	return PairFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R1, R2](f)
}

func MakePairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
) (R1, R2)](f F) PairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2] {
	return PairFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R1, R2](f)
}

func MakePairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
) (R1, R2)](f F) PairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2] {
	return PairFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R1, R2](f)
}
