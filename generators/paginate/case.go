package paginate

import (
	"crud-maker/generators"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type PaginateCase struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]generators.Field
}

func NewCase(name, namePlural string, fields map[string]generators.Field) PaginateCase {
	return PaginateCase{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: fmt.Sprintf("usecases/%s_case/paginate.go", strings.ToLower(name)),
	}
}

func (c PaginateCase) Generate() {
	template := utils.ReadTemplate("case_paginate")

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
