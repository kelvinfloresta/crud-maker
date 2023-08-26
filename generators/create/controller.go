package create

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type CreateController struct {
	name       string
	namePlural string
	outputPath string
	fields     map[string]generators.Field
}

func NewController(name, namePlural string, fields map[string]generators.Field) *CreateController {
	return &CreateController{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputPath: fmt.Sprintf("adapters/controllers/%s_controller/fiber_create.go", strings.ToLower(name)),
	}
}

func (c CreateController) Generate() {
	template := utils.ReadTemplate("controller_fiber_create")

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		MethodName:   "Create",
		Fields:       c.fields,
		MethodInput:  "",
		MethodOutput: "",
	})

	utils.WriteTemplate(template, c.outputPath)
}
