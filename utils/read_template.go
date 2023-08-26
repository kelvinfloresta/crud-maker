package utils

import (
	"crud-maker/config"
	"fmt"
	"os"
)

func ReadTemplate(name string) string {
	path := fmt.Sprintf("%s%s.template", config.TemplatePath, name)
	file, err := os.ReadFile(path)
	PanicIfError(err)
	template := string(file)
	return template
}
