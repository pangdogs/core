package generic

func CastPairFuncVar0[VA, R1, R2 any, F ~func(...VA) (R1, R2)](f F) PairFuncVar0[VA, R1, R2] {
	return PairFuncVar0[VA, R1, R2](f)
}

func CastPairFuncVar1[A1, VA, R1, R2 any, F ~func(A1, ...VA) (R1, R2)](f F) PairFuncVar1[A1, VA, R1, R2] {
	return PairFuncVar1[A1, VA, R1, R2](f)
}

func CastPairFuncVar2[A1, A2, VA, R1, R2 any, F ~func(A1, A2, ...VA) (R1, R2)](f F) PairFuncVar2[A1, A2, VA, R1, R2] {
	return PairFuncVar2[A1, A2, VA, R1, R2](f)
}

func CastPairFuncVar3[A1, A2, A3, VA, R1, R2 any, F ~func(A1, A2, A3, ...VA) (R1, R2)](f F) PairFuncVar3[A1, A2, A3, VA, R1, R2] {
	return PairFuncVar3[A1, A2, A3, VA, R1, R2](f)
}

func CastPairFuncVar4[A1, A2, A3, A4, VA, R1, R2 any, F ~func(A1, A2, A3, A4, ...VA) (R1, R2)](f F) PairFuncVar4[A1, A2, A3, A4, VA, R1, R2] {
	return PairFuncVar4[A1, A2, A3, A4, VA, R1, R2](f)
}

func CastPairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, ...VA,
) (R1, R2)](f F) PairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2] {
	return PairFuncVar5[A1, A2, A3, A4, A5, VA, R1, R2](f)
}

func CastPairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, ...VA,
) (R1, R2)](f F) PairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2] {
	return PairFuncVar6[A1, A2, A3, A4, A5, A6, VA, R1, R2](f)
}

func CastPairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
) (R1, R2)](f F) PairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2] {
	return PairFuncVar7[A1, A2, A3, A4, A5, A6, A7, VA, R1, R2](f)
}

func CastPairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
) (R1, R2)](f F) PairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2] {
	return PairFuncVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA, R1, R2](f)
}

func CastPairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
) (R1, R2)](f F) PairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2] {
	return PairFuncVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA, R1, R2](f)
}

func CastPairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
) (R1, R2)](f F) PairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2] {
	return PairFuncVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA, R1, R2](f)
}

func CastPairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
) (R1, R2)](f F) PairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2] {
	return PairFuncVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA, R1, R2](f)
}

func CastPairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
) (R1, R2)](f F) PairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2] {
	return PairFuncVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA, R1, R2](f)
}

func CastPairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
) (R1, R2)](f F) PairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2] {
	return PairFuncVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA, R1, R2](f)
}

func CastPairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
) (R1, R2)](f F) PairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2] {
	return PairFuncVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA, R1, R2](f)
}

func CastPairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
) (R1, R2)](f F) PairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2] {
	return PairFuncVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA, R1, R2](f)
}

func CastPairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
) (R1, R2)](f F) PairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2] {
	return PairFuncVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA, R1, R2](f)
}
