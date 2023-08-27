package get_by_id

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type GetByIDCase struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewCase(name, namePlural string, fields map[string]generators.Field) GetByIDCase {
	return GetByIDCase{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: fmt.Sprintf("usecases/%s_case/get_by_id.go", strings.ToLower(name)),
	}
}

func (c GetByIDCase) Generate() {
	template := utils.ReadTemplate("case_get_by_id")

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		MethodName:   "GetByID",
		Fields:       c.fields,
		MethodInput:  "id string",
		MethodOutput: "(*GetByIDOutput, error)",
	})

	utils.WriteTemplate(template, c.outputFile)
}
