package generators

import (
	"crud-maker/utils"
	"fmt"
	"strings"
)

type RouteGenerator struct {
	name       string
	namePlural string
	outputFile string
	fields     map[string]Field
}

func NewRoute(name, namePlural string, fields map[string]Field) *RouteGenerator {
	return &RouteGenerator{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
		outputFile: fmt.Sprintf("frameworks/http/fiber/routes/%s_routes.go", strings.ToLower(name)),
	}
}

func (c RouteGenerator) Generate(methodName string) {
	template, fileExist := utils.ReadExistingFile(c.outputFile)
	if fileExist {
		template = AppendRoute(template)
	} else {
		template = utils.ReadTemplate("routes_fiber")
	}

	template = ParseTemplate(ParseTemplateInput{
		Template:     template,
		Name:         c.name,
		NamePlural:   c.namePlural,
		MethodName:   methodName,
		Fields:       c.fields,
		MethodInput:  "",
		MethodOutput: "",
	})

	utils.OverwriteTemplate(template, c.outputFile)
}
