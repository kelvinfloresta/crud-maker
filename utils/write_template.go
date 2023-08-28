package utils

import (
	"crud-maker/config"
	"os"
	"strings"
)

func WriteTemplate(content, path string) {
	createPath(path)
	err := os.WriteFile(path, []byte(content), config.Permission)
	PanicIfError(err)
}

func createPath(path string) {
	filePath := strings.Split(path, "/")
	onlyFolders := filePath[:len(filePath)-1]
	err := os.MkdirAll(strings.Join(onlyFolders, "/"), config.Permission)
	PanicIfError(err)
}
