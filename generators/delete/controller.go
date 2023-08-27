package delete

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type DeleteController struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewController(name, namePlural string, fields map[string]generators.Field) *DeleteController {
	return &DeleteController{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: fmt.Sprintf("adapters/controllers/%s_controller/fiber_delete.go", strings.ToLower(name)),
	}
}

func (c DeleteController) Generate() {
	template := utils.ReadTemplate("controller_fiber_delete")

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		MethodName:   "Delete",
		Fields:       c.fields,
		MethodInput:  "",
		MethodOutput: "",
	})

	utils.WriteTemplate(template, c.outputFile)
}
