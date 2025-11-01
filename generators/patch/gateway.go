package patch

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type PatchGateway struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewGateway(name, namePlural string, fields map[string]generators.Field) *PatchGateway {
	return &PatchGateway{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: fmt.Sprintf("libs/database/gateways/%s_gateway/interface.go", strings.ToLower(name)),
	}
}

func (c PatchGateway) Generate() {
	template, fileExist := utils.ReadExistingFile(c.outputFile)
	if fileExist {
		template = generators.AppendMethodToInterface(template)
	} else {
		template = utils.ReadTemplate("gateway_interface")
	}

	template = fmt.Sprintf(`%s

	type PatchFilter struct {
		ID *string
		{{fields_optional}}
	}

	type PatchValues struct {
		{{fields}}
	}`, template)

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		MethodName:   "Patch",
		Fields:       c.fields,
		MethodInput:  "filter PatchFilter, values PatchValues",
		MethodOutput: "(bool, error)",
	})

	utils.OverwriteTemplate(template, c.outputFile)
}
