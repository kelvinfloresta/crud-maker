package get_by_id

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type GetByIDController struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewController(name, namePlural string, fields map[string]generators.Field) *GetByIDController {
	return &GetByIDController{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: fmt.Sprintf("adapters/controllers/%s_controller/fiber_get_by_id.go", strings.ToLower(name)),
	}
}

func (c GetByIDController) Generate() {
	template := utils.ReadTemplate("controller_fiber_get_by_id")

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
