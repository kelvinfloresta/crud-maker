package paginate

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type PaginateGateway struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewGateway(name, namePlural string, fields map[string]generators.Field) *PaginateGateway {
	return &PaginateGateway{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: fmt.Sprintf("adapters/gateways/%s_gateway/interface.go", strings.ToLower(name)),
	}
}

func (c PaginateGateway) Generate() {
	template, fileExist := utils.ReadExistingFile(c.outputFile)
	if fileExist {
		template = generators.AppendMethodToInterface(template)
	} else {
		template = utils.ReadTemplate("gateway_interface")
	}

	template = fmt.Sprintf(`%s

	type PaginateFilter struct {
		{{fields_optional}}
	}

	type PaginateData struct {
		{{fields}}
	}

	type PaginateOutput struct {
		Data     []PaginateData
		MaxPages int
	}`, template)

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		MethodName:   "Paginate",
		Fields:       c.fields,
		MethodInput:  "filter PaginateFilter, paginate database.PaginateInput",
		MethodOutput: "(*PaginateOutput, error)",
	})

	utils.OverwriteTemplate(template, c.outputFile)
}
