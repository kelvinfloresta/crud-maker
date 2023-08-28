package generators

import (
	"crud-maker/utils"
)

func GenerateStatic(templateName, path string) {
	template := utils.ReadTemplate(templateName)

	template = ParseTemplate(ParseTemplateInput{
		Template:     template,
		Name:         "",
		MethodName:   "",
		MethodInput:  "",
		MethodOutput: "",
		NamePlural:   "",
		Fields:       map[string]Field{},
	})

	utils.WriteTemplate(template, path)
}
