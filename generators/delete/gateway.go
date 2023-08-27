package delete

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type DeleteGateway struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewGateway(name, namePlural string, fields map[string]generators.Field) *DeleteGateway {
	return &DeleteGateway{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: fmt.Sprintf("adapters/gateways/%s_gateway/interface.go", strings.ToLower(name)),
	}
}

func (c DeleteGateway) Generate() {
	template, fileExist := utils.ReadExistingFile(c.outputFile)
	if fileExist {
		template = generators.AppendMethodToInterface(template)
	} else {
		template = utils.ReadTemplate("gateway_interface")
	}

	template = fmt.Sprintf(`%s

		type DeleteInput struct {
			{{fields}}
		}`, template)

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		MethodName:   "Delete",
		Fields:       c.fields,
		MethodInput:  "id string",
		MethodOutput: "(bool, error)",
	})

	utils.WriteTemplate(template, c.outputFile)
}
