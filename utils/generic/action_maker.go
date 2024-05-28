package generic

func MakeAction0[F ~func()](f F) Action0 {
	return Action0(f)
}

func MakeAction1[A1 any, F ~func(A1)](f F) Action1[A1] {
	return Action1[A1](f)
}

func MakeAction2[A1, A2 any, F ~func(A1, A2)](f F) Action2[A1, A2] {
	return Action2[A1, A2](f)
}

func MakeAction3[A1, A2, A3 any, F ~func(A1, A2, A3)](f F) Action3[A1, A2, A3] {
	return Action3[A1, A2, A3](f)
}

func MakeAction4[A1, A2, A3, A4 any, F ~func(A1, A2, A3, A4)](f F) Action4[A1, A2, A3, A4] {
	return Action4[A1, A2, A3, A4](f)
}

func MakeAction5[A1, A2, A3, A4, A5 any, F ~func(
	A1, A2, A3, A4, A5,
)](f F) Action5[A1, A2, A3, A4, A5] {
	return Action5[A1, A2, A3, A4, A5](f)
}

func MakeAction6[A1, A2, A3, A4, A5, A6 any, F ~func(
	A1, A2, A3, A4, A5, A6,
)](f F) Action6[A1, A2, A3, A4, A5, A6] {
	return Action6[A1, A2, A3, A4, A5, A6](f)
}

func MakeAction7[A1, A2, A3, A4, A5, A6, A7 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7,
)](f F) Action7[A1, A2, A3, A4, A5, A6, A7] {
	return Action7[A1, A2, A3, A4, A5, A6, A7](f)
}

func MakeAction8[A1, A2, A3, A4, A5, A6, A7, A8 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8,
)](f F) Action8[A1, A2, A3, A4, A5, A6, A7, A8] {
	return Action8[A1, A2, A3, A4, A5, A6, A7, A8](f)
}

func MakeAction9[A1, A2, A3, A4, A5, A6, A7, A8, A9 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9,
)](f F) Action9[A1, A2, A3, A4, A5, A6, A7, A8, A9] {
	return Action9[A1, A2, A3, A4, A5, A6, A7, A8, A9](f)
}

func MakeAction10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10,
)](f F) Action10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10] {
	return Action10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10](f)
}

func MakeAction11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11,
)](f F) Action11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11] {
	return Action11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11](f)
}

func MakeAction12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12,
)](f F) Action12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12] {
	return Action12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12](f)
}

func MakeAction13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13,
)](f F) Action13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13] {
	return Action13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13](f)
}

func MakeAction14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14,
)](f F) Action14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14] {
	return Action14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14](f)
}

func MakeAction15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15,
)](f F) Action15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15] {
	return Action15[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15](f)
}

func MakeAction16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16 any, F ~func(
	A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16,
)](f F) Action16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16] {
	return Action16[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A15, A16](f)
}
