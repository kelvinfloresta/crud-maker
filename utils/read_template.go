package utils

import (
	"crud-maker/config"
	"fmt"
)

func ReadTemplate(name string) string {
	path := fmt.Sprintf("%s%s.template", config.TemplatePath, name)
	file, err := Asset(path)
	PanicIfError(err)
	template := string(file)
	return template
}
