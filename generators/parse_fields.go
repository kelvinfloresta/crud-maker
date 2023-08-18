package generators

import (
	"crud-maker/config"
	"crud-maker/utils"
	"strings"
)

type parseTemplateInput struct {
	Template   string
	Name       string
	Type       string
	NamePlural string
}

func parseTemplate(input parseTemplateInput) string {
	parsed := strings.ReplaceAll(input.Template, "{{name}}", strings.ToLower(input.Name))
	parsed = strings.ReplaceAll(parsed, "{{name_capitalized}}", strings.Title(input.Name))
	parsed = strings.ReplaceAll(parsed, "{{method_capitalized}}", strings.Title(input.Type))
	parsed = strings.ReplaceAll(parsed, "{{name_plural}}", utils.ToSnakeCase(input.NamePlural))
	parsed = strings.ReplaceAll(parsed, "{{project_name}}", config.ProjectName)

	return parsed
}
