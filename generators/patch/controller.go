package patch

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type PatchController struct {
	name       string
	namePlural string
	outputPath string
	fields     map[string]generators.Field
}

func NewController(name, namePlural string, fields map[string]generators.Field) *PatchController {
	return &PatchController{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputPath: fmt.Sprintf("adapters/controllers/%s_controller/fiber_patch.go", strings.ToLower(name)),
	}
}

func (c PatchController) Generate() {
	template := utils.ReadTemplate("controller_fiber_patch.template")

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		MethodName:   "Patch",
		Fields:       c.fields,
		MethodInput:  "filter PatchFilter, values PatchValues",
		MethodOutput: "(bool, error)",
	})

	utils.WriteTemplate(template, c.outputPath)
}