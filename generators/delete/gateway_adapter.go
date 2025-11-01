package delete

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type DeleteGatewayAdapter struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewGatewayAdapter(name, namePlural string, fields map[string]generators.Field) *DeleteGatewayAdapter {
	return &DeleteGatewayAdapter{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: fmt.Sprintf("libs/database/gateways/%s_gateway/gorm_delete.go", strings.ToLower(name)),
	}
}

func (c *DeleteGatewayAdapter) Generate() {
	template := utils.ReadTemplate("gateway_gorm_delete")

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		Fields:       c.fields,
		MethodName:   "Delete",
		MethodInput:  "id string",
		MethodOutput: "(bool, error)",
	})

	utils.WriteTemplate(template, c.outputFile)
}
