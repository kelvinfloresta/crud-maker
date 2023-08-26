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

func AppendMethodToInterface(template string) string {
	return strings.Replace(
		template,
		"}",
		"{{method_capitalized}}({{method_input}}) {{method_output}}\n}",
		1,
	)
}

func ParseTemplate(input ParseTemplateInput) string {
	var (
		fields         = ""
		fieldsModel    = ""
		fieldsOptional = ""
		adaptInput     = ""
		adaptFilter    = ""
		adaptValues    = ""
		filters        = ""
		fieldsPointer  = ""
		fieldsQuery    = ""
	)

	for fieldName, field := range input.Fields {
		fields += parseField(fieldName, field)
		fieldsModel += parseModel(fieldName, field)
		adaptInput += fmt.Sprintf("%s: input.%s,\n", fieldName, fieldName)
		adaptFilter += fmt.Sprintf("%s: filter.%s,", fieldName, fieldName)
		adaptValues += fmt.Sprintf("%s: values.%s,", fieldName, fieldName)
		fieldsOptional += fmt.Sprintf("%s *%s\n", fieldName, field.Type)
		fieldsQuery += fmt.Sprintf("%s := ctx.Query(\"%s\")\n", fieldName, fieldName)
		fieldsPointer += fmt.Sprintf(`%s: &%s,`, fieldName, fieldName)

		filters += fmt.Sprintf(`
		if filter.%s != nil {
			query = query.Where("%s = ?", filter.%s)
		}
		`, fieldName, utils.ToSnakeCase(fieldName), fieldName)

	}

	template := strings.ReplaceAll(input.Template, "{{adapt_input}}", adaptInput)
	template = strings.ReplaceAll(template, "{{adapt_filter}}", adaptFilter)
	template = strings.ReplaceAll(template, "{{adapt_values}}", adaptValues)
	template = strings.ReplaceAll(template, "{{fields}}", fields)
	template = strings.ReplaceAll(template, "{{filters}}", filters)
	template = strings.ReplaceAll(template, "{{fields_model}}", fieldsModel)
	template = strings.ReplaceAll(template, "{{fields_optional}}", fieldsOptional)
	template = strings.ReplaceAll(template, "{{method_input}}", input.MethodInput)
	template = strings.ReplaceAll(template, "{{method_output}}", input.MethodOutput)
	template = strings.ReplaceAll(template, "{{name}}", strings.ToLower(input.Name))
	template = strings.ReplaceAll(template, "{{name_capitalized}}", strings.Title(input.Name))
	template = strings.ReplaceAll(template, "{{method_capitalized}}", strings.Title(input.MethodName))
	template = strings.ReplaceAll(template, "{{name_plural}}", utils.ToSnakeCase(input.NamePlural))
	template = strings.ReplaceAll(template, "{{fields_pointer}}", fieldsPointer)
	template = strings.ReplaceAll(template, "{{fields_query}}", fieldsQuery)
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
	isNullable := ""
	if field.IsRequired {
		isNullable = "`gorm:\"not null\"`"
	}

	return fmt.Sprintf("%s %s %s\n", fieldName, field.Type, isNullable)
}
