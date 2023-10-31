package list

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
)

type ListGateway struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewGateway(name, namePlural string, fields map[string]generators.Field) *ListGateway {
	return &ListGateway{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: "frameworks/database/interface.go",
	}
}

func (c ListGateway) Generate() {
	template, fileExist := utils.ReadExistingFile(c.outputFile)
	if fileExist {
		template = generators.AppendMethodToInterface(template)
	} else {
		template = utils.ReadTemplate("gateway_interface")
	}

	template = fmt.Sprintf(`%s

	type ListInput struct {
		{{fields_optional}}
	}

	type ListOutput struct {
		{{fields}}
	}`, template)

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		MethodName:   "List",
		Fields:       c.fields,
		MethodInput:  "input ListInput",
		MethodOutput: "([]ListOutput, error)",
	})

	utils.OverwriteTemplate(template, c.outputFile)
}
