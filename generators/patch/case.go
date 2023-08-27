package patch

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type PatchCase struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewCase(name, namePlural string, fields map[string]generators.Field) PatchCase {
	return PatchCase{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: fmt.Sprintf("usecases/%s_case/patch.go", strings.ToLower(name)),
	}
}

func (c PatchCase) Generate() {
	template := utils.ReadTemplate("case_patch")

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		MethodName:   "Patch",
		Fields:       c.fields,
		MethodInput:  "filter PatchFilter, values PatchValues",
		MethodOutput: "(bool, error)",
	})

	utils.WriteTemplate(template, c.outputFile)
}
