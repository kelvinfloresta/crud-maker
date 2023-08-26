package utils

import (
	"crud-maker/config"
	"os"
)

func WriteTemplate(content, path string) {
	err := os.WriteFile(path, []byte(content), config.Permission)
	PanicIfError(err)
}
