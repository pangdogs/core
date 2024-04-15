package generic

func MakePairFuncVar0[VA, R1, R2 any, F ~func(...VA) (R1, R2)](f F) PairFuncVar0[VA, R1, R2] {
	return PairFuncVar0[VA, R1, R2](f)
}

func MakePairFuncVar1[A1, VA, R1, R2 any, F ~func(A1, ...VA) (R1, R2)](f F) PairFuncVar1[A1, VA, R1, R2] {
	return PairFuncVar1[A1, VA, R1, R2](f)
}

func MakePairFuncVar2[A1, A2, VA, R1, R2 any, F ~func(A1, A2, ...VA) (R1, R2)](f F) PairFuncVar2[A1, A2, VA, R1, R2] {
	return PairFuncVar2[A1, A2, VA, R1, R2](f)
}

func MakePairFuncVar3[A1, A2, A3, VA, R1, R2 any, F ~func(A1, A2, A3, ...VA) (R1, R2)](f F) PairFuncVar3[A1, A2, A3, VA, R1, R2] {
	return PairFuncVar3[A1, A2, A3, VA, R1, R2](f)
}

func MakePairFuncVar4[A1, A2, A3, A4, VA, R1, R2 any, F ~func(A1, A2, A3, A4, ...VA) (R1, R2)](f F) PairFuncVar4[A1, A2, A3, A4, VA, R1, R2] {
	return PairFuncVar4[A1, A2, A3, A4, VA, R1, R2](f)
}

func MakePairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, ...VA,
) (R1, R2)](f F) PairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2] {
	return PairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2](f)
}

func MakePairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, ...VA,
) (R1, R2)](f F) PairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2] {
	return PairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2](f)
}

func MakePairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
) (R1, R2)](f F) PairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2] {
	return PairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2](f)
}

func MakePairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
) (R1, R2)](f F) PairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2] {
	return PairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2](f)
}

func MakePairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
) (R1, R2)](f F) PairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2] {
	return PairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2](f)
}

func MakePairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
) (R1, R2)](f F) PairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2] {
	return PairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2](f)
}

func MakePairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
) (R1, R2)](f F) PairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2] {
	return PairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2](f)
}

func MakePairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
) (R1, R2)](f F) PairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2] {
	return PairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2](f)
}

func MakePairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
) (R1, R2)](f F) PairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2] {
	return PairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2](f)
}

func MakePairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
) (R1, R2)](f F) PairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2] {
	return PairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2](f)
}

func MakePairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
) (R1, R2)](f F) PairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2] {
	return PairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2](f)
}

func MakePairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
) (R1, R2)](f F) PairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2] {
	return PairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2](f)
}
