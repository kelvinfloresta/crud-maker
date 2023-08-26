package create

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type CreateCase struct {
	name       string
	namePlural string
	outputPath string
	fields     map[string]generators.Field
}

func NewCase(name, namePlural string, fields map[string]generators.Field) CreateCase {
	return CreateCase{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputPath: fmt.Sprintf("usecases/%s_case/create.go", strings.ToLower(name)),
	}
}

func (c CreateCase) Generate() {
	template := utils.ReadTemplate("case_create.template")

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		MethodName:   "Create",
		Fields:       c.fields,
		MethodInput:  "input CreateInput",
		MethodOutput: "(string, error)",
	})

	utils.WriteTemplate(template, c.outputPath)
}