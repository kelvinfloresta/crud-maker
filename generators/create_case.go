package generators

import (
	"crud-maker/config"
	"fmt"
	"os"
	"strings"
)

type CreateCase struct {
	name       string
	namePlural string
	outputPath string
	fields     map[string]Field
}

func NewCreateCase(name, namePlural string, fields map[string]Field) CreateCase {
	return CreateCase{
		name:       name,
		namePlural: namePlural,
		fields:     fields,
	}
}

func (c CreateCase) Generate() {
	c.outputPath = fmt.Sprintf("usecases/%s_case/", strings.ToLower(c.name))

	c.createFolder()

	content := c.readTemplate()

	content = parseTemplate(parseTemplateInput{
		Template:   content,
		Name:       c.name,
		NamePlural: c.namePlural,
		Type:       "Usecase",
	})

	content = c.parseMethods(content)

	c.writeTemplate(content)
}

func (c CreateCase) createFolder() {
	err := os.MkdirAll(c.outputPath, config.Permission)
	CheckError(err)
}

func (c CreateCase) writeTemplate(parsed string) {
	dest := fmt.Sprintf("%s/%s.go", c.outputPath, strings.ToLower(c.name))
	err := os.WriteFile(dest, []byte(parsed), config.Permission)
	CheckError(err)
}

func (c CreateCase) readTemplate() string {
	path := fmt.Sprintf("%s/usecase.template", config.TemplatePath)
	file, err := os.ReadFile(path)
	CheckError(err)
	template := string(file)
	return template
}

func (c CreateCase) parseMethods(content string) string {
	fields := make([]string, 0, len(c.fields))
	for key, field := range c.fields {
		pointer := "*"
		if field.IsRequired {
			pointer = ""
		}

		fields = append(fields, fmt.Sprintf("%s %s%s", key, pointer, field.Type))
	}

	adaptInput := make([]string, 0, len(c.fields))
	for k := range c.fields {
		adaptInput = append(adaptInput, fmt.Sprintf("%s: input.%s,", k, k))
	}

	result := strings.ReplaceAll(content, "{{adapt_input}}", strings.Join(adaptInput, "\n"))
	result = strings.ReplaceAll(result, "{{fields}}", strings.Join(fields, "\n"))
	result = strings.ReplaceAll(result, "{{method_input}}", "input CreateInput")
	result = strings.ReplaceAll(result, "{{method_output}}", "(string, error)")

	return result
}
