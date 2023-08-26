package utils

import "os"

func ReadExistingFile(path string) (string, bool) {
	if _, err := os.Stat(path); err != nil {
		return "", false
	}

	file, err := os.ReadFile(path)
	PanicIfError(err)

	return string(file), true
}
