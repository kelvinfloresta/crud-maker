package get_by_id

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type GetByIDGatewayAdapter struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewGatewayAdapter(name, namePlural string, fields map[string]generators.Field) *GetByIDGatewayAdapter {
	return &GetByIDGatewayAdapter{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: fmt.Sprintf("libs/database/gateways/%s_gateway/gorm_get_by_id.go", strings.ToLower(name)),
	}
}

func (c *GetByIDGatewayAdapter) Generate() {
	template := utils.ReadTemplate("gateway_gorm_get_by_id")

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		Fields:       c.fields,
		MethodName:   "GetByID",
		MethodInput:  "id string",
		MethodOutput: "(*GetByIDOutput, error)",
	})

	utils.WriteTemplate(template, c.outputFile)
}
