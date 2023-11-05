package generic

func FuncError(err, panicErr error) error {
	if panicErr != nil {
		return panicErr
	}
	if err != nil {
		return err
	}
	return nil
}

func PairFuncError[T any](r T, err, panicErr error) (T, error) {
	if panicErr != nil {
		return r, panicErr
	}
	if err != nil {
		return r, err
	}
	return r, nil
}
