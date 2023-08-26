package patch

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type PatchGatewayAdapter struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewGatewayAdapter(name, namePlural string, fields map[string]generators.Field) *PatchGatewayAdapter {
	return &PatchGatewayAdapter{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: fmt.Sprintf("adapters/gateways/%s_gateway/gorm_patch.go", strings.ToLower(name)),
	}
}

func (c *PatchGatewayAdapter) Generate() {
	template := utils.ReadTemplate("gateway_gorm_patch.template")

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		Fields:       c.fields,
		MethodName:   "Patch",
		MethodInput:  "filter PatchFilter, values PatchValues",
		MethodOutput: "(bool, error)",
	})

	utils.WriteTemplate(template, c.outputFile)
}
