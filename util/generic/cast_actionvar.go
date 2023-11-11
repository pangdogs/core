package generic

func CastActionVar0[VA any, F ~func(...VA)](f F) ActionVar0[VA] {
	return ActionVar0[VA](f)
}

func CastActionVar1[A1, VA, F ~func(A1, ...VA)](f F) ActionVar1[A1, VA] {
	return ActionVar1[A1, VA](f)
}

func CastActionVar2[A1, A2, VA any, F ~func(A1, A2, ...VA)](f F) ActionVar2[A1, A2, VA] {
	return ActionVar2[A1, A2, VA](f)
}

func CastActionVar3[A1, A2, A3, VA any, F ~func(A1, A2, A3, ...VA)](f F) ActionVar3[A1, A2, A3, VA] {
	return ActionVar3[A1, A2, A3, VA](f)
}

func CastActionVar4[A1, A2, A3, A4, VA any, F ~func(A1, A2, A3, A4, ...VA)](f F) ActionVar4[A1, A2, A3, A4, VA] {
	return ActionVar4[A1, A2, A3, A4, VA](f)
}

func CastActionVar5[A1, A2, A3, A4, A5, VA any, F ~func(
	A1, A2, A3, A4, A5, ...VA,
)](f F) ActionVar5[A1, A2, A3, A4, A5, VA] {
	return ActionVar5[A1, A2, A3, A4, A5, VA](f)
}

func CastActionVar6[A1, A2, A3, A4, A5, A6, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, ...VA,
)](f F) ActionVar6[A1, A2, A3, A4, A5, A6, VA] {
	return ActionVar6[A1, A2, A3, A4, A5, A6, VA](f)
}

func CastActionVar7[A1, A2, A3, A4, A5, A6, A7, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, ...VA,
)](f F) ActionVar7[A1, A2, A3, A4, A5, A6, A7, VA] {
	return ActionVar7[A1, A2, A3, A4, A5, A6, A7, VA](f)
}

func CastActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, ...VA,
)](f F) ActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA] {
	return ActionVar8[A1, A2, A3, A4, A5, A6, A7, A8, VA](f)
}

func CastActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, ...VA,
)](f F) ActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA] {
	return ActionVar9[A1, A2, A3, A4, A5, A6, A7, A8, A9, VA](f)
}

func CastActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, ...VA,
)](f F) ActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA] {
	return ActionVar10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, VA](f)
}

func CastActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, ...VA,
)](f F) ActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA] {
	return ActionVar11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, VA](f)
}

func CastActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, ...VA,
)](f F) ActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA] {
	return ActionVar12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, VA](f)
}

func CastActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, ...VA,
)](f F) ActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA] {
	return ActionVar13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, VA](f)
}

func CastActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, ...VA,
)](f F) ActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA] {
	return ActionVar14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, VA](f)
}

func CastActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, ...VA,
)](f F) ActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA] {
	return ActionVar15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, VA](f)
}

func CastActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, ...VA,
)](f F) ActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA] {
	return ActionVar16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16, VA](f)
}
