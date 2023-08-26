package list

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type ListController struct {
	name       string
	namePlural string
	outputPath string
	fields     map[string]generators.Field
}

func NewController(name, namePlural string, fields map[string]generators.Field) *ListController {
	return &ListController{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputPath: fmt.Sprintf("adapters/controllers/%s_controller/fiber_list.go", strings.ToLower(name)),
	}
}

func (c ListController) Generate() {
	template := utils.ReadTemplate("controller_fiber_list.template")

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
