package delete

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type DeleteCase struct {
	name       string
	namePlural string
	outputPath string
	fields     map[string]generators.Field
}

func NewCase(name, namePlural string, fields map[string]generators.Field) DeleteCase {
	return DeleteCase{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputPath: fmt.Sprintf("usecases/%s_case/delete.go", strings.ToLower(name)),
	}
}

func (c DeleteCase) Generate() {
	template := utils.ReadTemplate("case_delete.template")

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		MethodName:   "Delete",
		Fields:       c.fields,
		MethodInput:  "id string",
		MethodOutput: "(bool, error)",
	})

	utils.WriteTemplate(template, c.outputPath)
}
