package generic

func MakeFuncVar0[VA, R any, F ~func(...VA) R](f F) FuncVar0[VA, R] {
	return FuncVar0[VA, R](f)
}

func MakeFuncVar1[A1, VA, R any, F ~func(A1, ...VA) R](f F) FuncVar1[A1, VA, R] {
	return FuncVar1[A1, VA, R](f)
}

func MakeFuncVar2[A1, A2, VA, R any, F ~func(A1, A2, ...VA) R](f F) FuncVar2[A1, A2, VA, R] {
	return FuncVar2[A1, A2, VA, R](f)
}

func MakeFuncVar3[A1, A2, A3, VA, R any, F ~func(A1, A2, A3, ...VA) R](f F) FuncVar3[A1, A2, A3, VA, R] {
	return FuncVar3[A1, A2, A3, VA, R](f)
}

func MakeFuncVar4[A1, A2, A3, A4, VA, R any, F ~func(A1, A2, A3, A4, ...VA) R](f F) FuncVar4[A1, A2, A3, A4, VA, R] {
	return FuncVar4[A1, A2, A3, A4, VA, R](f)
}

func MakeFuncVar5[A1, A2, A3, A4, A5, VA, R any, F ~func(
	A1, A2, A3, A4, A5, ...VA,
) R](f F) FuncVar5[A1, A2, A3, A4, A5, VA, R] {
	return FuncVar5[A1, A2, A3, A4, A5, VA, R](f)
}

func MakeFuncVar6[A1, A2, A3, A4, A5, A6, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, ...VA,
) R](f F) FuncVar6[A1, A2, A3, A4, A5, A6, VA, R] {
	return FuncVar6[A1, A2, A3, A4, A5, A6, VA, R](f)
}

func MakeFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
) R](f F) FuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R] {
	return FuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R](f)
}

func MakeFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
) R](f F) FuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R] {
	return FuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R](f)
}

func MakeFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
) R](f F) FuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R] {
	return FuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R](f)
}

func MakeFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
) R](f F) FuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R] {
	return FuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R](f)
}

func MakeFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
) R](f F) FuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R] {
	return FuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R](f)
}

func MakeFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
) R](f F) FuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R] {
	return FuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R](f)
}

func MakeFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
) R](f F) FuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R] {
	return FuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R](f)
}

func MakeFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
) R](f F) FuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R] {
	return FuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R](f)
}

func MakeFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
) R](f F) FuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R] {
	return FuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R](f)
}

func MakeFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
) R](f F) FuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R] {
	return FuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R](f)
}
