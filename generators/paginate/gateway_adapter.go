package paginate

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type PaginateGatewayAdapter struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewGatewayAdapter(name, namePlural string, fields map[string]generators.Field) *PaginateGatewayAdapter {
	return &PaginateGatewayAdapter{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: fmt.Sprintf("adapters/gateways/%s_gateway/gorm_paginate.go", strings.ToLower(name)),
	}
}

func (c *PaginateGatewayAdapter) Generate() {
	template := utils.ReadTemplate("gateway_gorm_paginate")

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		Fields:       c.fields,
		MethodName:   "Paginate",
		MethodInput:  "filter PaginateFilter, paginate database.PaginateInput",
		MethodOutput: "(*PaginateOutput, error)",
	})

	utils.WriteTemplate(template, c.outputFile)
}
