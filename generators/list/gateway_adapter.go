package list

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type ListGatewayAdapter struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewGatewayAdapter(name, namePlural string, fields map[string]generators.Field) *ListGatewayAdapter {
	return &ListGatewayAdapter{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: fmt.Sprintf("adapters/gateways/%s_gateway/gorm_list.go", strings.ToLower(name)),
	}
}

func (c *ListGatewayAdapter) Generate() {
	template := utils.ReadTemplate("gateway_gorm_list.template")

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		Fields:       c.fields,
		MethodName:   "List",
		MethodInput:  "input ListInput",
		MethodOutput: "([]ListOutput, error)",
	})

	utils.WriteTemplate(template, c.outputFile)
}
