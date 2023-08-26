package list

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type ListCase struct {
	name       string
	namePlural string
	outputPath string
	fields     map[string]generators.Field
}

func NewCase(name, namePlural string, fields map[string]generators.Field) ListCase {
	return ListCase{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputPath: fmt.Sprintf("usecases/%s_case/list.go", strings.ToLower(name)),
	}
}

func (c ListCase) Generate() {
	template := utils.ReadTemplate("case_list.template")

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		MethodName:   "List",
		Fields:       c.fields,
		MethodInput:  "input ListInput",
		MethodOutput: "([]ListOutput, error)",
	})

	utils.WriteTemplate(template, c.outputPath)
}
