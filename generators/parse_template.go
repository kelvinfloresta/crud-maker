package generators

import (
	"crud-maker/config"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type ParseTemplateInput struct {
	Template     string
	Name         string
	MethodName   string
	MethodInput  string
	MethodOutput string
	NamePlural   string
	Fields       map[string]Field
}

func ParseTemplate(input ParseTemplateInput) string {
	var (
		fields         = ""
		fieldsModel    = ""
		adaptInput     = ""
		fieldsOptional = ""
	)

	for fieldName, field := range input.Fields {
		fields += parseField(fieldName, field)
		fieldsModel += parseModel(fieldName, field)
		adaptInput += fmt.Sprintf("%s: input.%s,\n", fieldName, fieldName)
		fieldsOptional += fmt.Sprintf("%s *%s\n", fieldName, field.Type)
	}

	template := strings.ReplaceAll(input.Template, "{{adapt_input}}", adaptInput)
	template = strings.ReplaceAll(template, "{{fields}}", fields)
	template = strings.ReplaceAll(template, "{{fields_model}}", fieldsModel)
	template = strings.ReplaceAll(template, "{{fields_optional}}", fieldsOptional)
	template = strings.ReplaceAll(template, "{{method_input}}", input.MethodInput)
	template = strings.ReplaceAll(template, "{{method_output}}", input.MethodOutput)
	template = strings.ReplaceAll(template, "{{name}}", strings.ToLower(input.Name))
	template = strings.ReplaceAll(template, "{{name_capitalized}}", strings.Title(input.Name))
	template = strings.ReplaceAll(template, "{{method_capitalized}}", strings.Title(input.MethodName))
	template = strings.ReplaceAll(template, "{{name_plural}}", utils.ToSnakeCase(input.NamePlural))
	template = strings.ReplaceAll(template, "{{project_name}}", config.ProjectName)
	return template
}

func parseField(fieldName string, field Field) string {
	pointer := "*"
	if field.IsRequired {
		pointer = ""
	}

	return fmt.Sprintf("%s %s%s\n", fieldName, pointer, field.Type)
}

func parseModel(fieldName string, field Field) string {
	gormNull := ""
	if field.IsRequired {
		gormNull = "`gorm:\"not null\"`"
	}

	return fmt.Sprintf("%s %s %s\n", fieldName, field.Type, gormNull)
}
