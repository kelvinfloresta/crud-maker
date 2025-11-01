package create

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type CreateGatewayAdapter struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewGatewayAdapter(name, namePlural string, fields map[string]generators.Field) *CreateGatewayAdapter {
	return &CreateGatewayAdapter{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: fmt.Sprintf("libs/database/gateways/%s_gateway/gorm_create.go", strings.ToLower(name)),
	}
}

func (c *CreateGatewayAdapter) Generate() {
	template := utils.ReadTemplate("gateway_gorm_create")

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		Fields:       c.fields,
		MethodName:   "Create",
		MethodInput:  "input CreateInput",
		MethodOutput: "(string, error)",
	})

	utils.WriteTemplate(template, c.outputFile)
}
