package paginate

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type PaginateController struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewController(name, namePlural string, fields map[string]generators.Field) *PaginateController {
	return &PaginateController{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: fmt.Sprintf("frameworks/http/controllers/%s_controller/fiber_paginate.go", strings.ToLower(name)),
	}
}

func (c PaginateController) Generate() {
	template := utils.ReadTemplate("controller_fiber_paginate")

	template = generators.ParseTemplate(generators.ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		MethodName:   "Paginate",
		Fields:       c.fields,
		MethodInput:  "filter PaginateFilter, paginate database.PaginateInput",
		MethodOutput: "(*PaginateOutput, error)",
	})

	utils.WriteTemplate(template, c.outputFile)
}
