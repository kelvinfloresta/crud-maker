package utils

func PanicIfError(err error) {
	if err == nil {
		return
	}

	panic(err)
}
