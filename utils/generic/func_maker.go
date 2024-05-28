package generic

func MakeFunc0[R any, F ~func() R](f F) Func0[R] {
	return Func0[R](f)
}

func MakeFunc1[A1, R any, F ~func(A1) R](f F) Func1[A1, R] {
	return Func1[A1, R](f)
}

func MakeFunc2[A1, A2, R any, F ~func(A1, A2) R](f F) Func2[A1, A2, R] {
	return Func2[A1, A2, R](f)
}

func MakeFunc3[A1, A2, A3, R any, F ~func(A1, A2, A3) R](f F) Func3[A1, A2, A3, R] {
	return Func3[A1, A2, A3, R](f)
}

func MakeFunc4[A1, A2, A3, A4, R any, F ~func(A1, A2, A3, A4) R](f F) Func4[A1, A2, A3, A4, R] {
	return Func4[A1, A2, A3, A4, R](f)
}

func MakeFunc5[A1, A2, A3, A4, A5, R any, F ~func(
	A1, A2, A3, A4, A5,
) R](f F) Func5[A1, A2, A3, A4, A5, R] {
	return Func5[A1, A2, A3, A4, A5, R](f)
}

func MakeFunc6[A1, A2, A3, A4, A5, A6, R any, F ~func(
	A1, A2, A3, A4, A5, A6,
) R](f F) Func6[A1, A2, A3, A4, A5, A6, R] {
	return Func6[A1, A2, A3, A4, A5, A6, R](f)
}

func MakeFunc7[A1, A2, A3, A4, A5, A6, A7, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7,
) R](f F) Func7[A1, A2, A3, A4, A5, A6, A7, R] {
	return Func7[A1, A2, A3, A4, A5, A6, A7, R](f)
}

func MakeFunc8[A1, A2, A3, A4, A5, A6, A7, A8, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8,
) R](f F) Func8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return Func8[A1, A2, A3, A4, A5, A6, A7, A8, R](f)
}

func MakeFunc9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
) R](f F) Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R](f)
}

func MakeFunc10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
) R](f F) Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {
	return Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R](f)
}

func MakeFunc11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
) R](f F) Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {
	return Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R](f)
}

func MakeFunc12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
) R](f F) Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {
	return Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R](f)
}

func MakeFunc13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
) R](f F) Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {
	return Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R](f)
}

func MakeFunc14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
) R](f F) Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {
	return Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R](f)
}

func MakeFunc15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
) R](f F) Func15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R] {
	return Func15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, R](f)
}

func MakeFunc16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
) R](f F) Func16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R] {
	return Func16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, R](f)
}
