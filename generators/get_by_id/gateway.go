package get_by_id

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
)

type GetByIDGateway struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewGateway(name, namePlural string, fields map[string]generators.Field) *GetByIDGateway {
	return &GetByIDGateway{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: "frameworks/database/interface.go",
	}
}

func (c GetByIDGateway) Generate() {
	template, fileExist := utils.ReadExistingFile(c.outputFile)
	if fileExist {
		template = generators.AppendMethodToInterface(template)
	} else {
		template = utils.ReadTemplate("gateway_interface")
	}

	template = fmt.Sprintf(`%s

	type GetByIDOutput struct {
		{{fields}}
	}`, template)

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		MethodName:   "GetByID",
		Fields:       c.fields,
		MethodInput:  "id string",
		MethodOutput: "(*GetByIDOutput, error)",
	})

	utils.OverwriteTemplate(template, c.outputFile)
}
