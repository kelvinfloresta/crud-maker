package create

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type CreateGateway struct {
	name       string
	namePlural string
	outputPath string
	fields     map[string]generators.Field
}

func NewGateway(name, namePlural string, fields map[string]generators.Field) *CreateGateway {
	return &CreateGateway{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputPath: fmt.Sprintf("adapters/gateways/%s_gateway/interface.go", strings.ToLower(name)),
	}
}

func (c CreateGateway) Generate() {
	template, fileExist := utils.ReadExistingFile(c.outputPath)
	if fileExist {
		template = generators.AppendMethodToInterface(template)
	} else {
		template = utils.ReadTemplate("gateway_interface.template")
	}

	template = fmt.Sprintf(`%s

		type CreateInput struct {
			{{fields}}
		}`, template)

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
