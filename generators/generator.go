package generators

import (
	"crud-maker/config"
	"crud-maker/utils"
	"fmt"
	"os"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

type Field struct {
	IsRequired bool
	Type       string
}

type Generator struct {
	Name         string
	NamePlural   string
	Type         string
	TemplateName string
	OutputName   string
	outputPath   string
	Fields       map[string]Field
}

func (g Generator) Generate() {
	g.outputPath = g.getOutputPath()
	err := os.MkdirAll(g.outputPath, config.Permission)
	CheckError(err)

	dest := fmt.Sprint(g.outputPath, "/", g.OutputName, ".go")
	var template string

	if _, err := os.Stat(dest); err == nil {
		file, err := os.ReadFile(dest)
		CheckError(err)
		if g.OutputName == "interface" {
			template = strings.Replace(
				string(file),
				"}",
				"{{method_capitalized}}({{method_input}}) {{method_output}}\n}",
				1,
			)
		}
	} else {
		path := fmt.Sprint(config.TemplatePath, "/", g.TemplateName, ".template")
		file, err := os.ReadFile(path)
		CheckError(err)
		template = string(file)
	}

	parsed := g.parseTemplate(template)
	err = os.WriteFile(dest, []byte(parsed), config.Permission)
	CheckError(err)
}

func (g Generator) getOutputPath() string {
	switch {
	case strings.Contains(g.TemplateName, "model"):
		return "frameworks/database/gorm/models"

	case strings.Contains(g.TemplateName, "gateway"):
		return fmt.Sprintf("adapters/gateways/%s_gateway/", strings.ToLower(g.Name))

	case strings.Contains(g.TemplateName, "controller"):
		return fmt.Sprintf("adapters/controllers/%s_controller/", strings.ToLower(g.Name))

	case strings.Contains(g.TemplateName, "case"):
		return fmt.Sprintf("usecases/%s_case/", strings.ToLower(g.Name))

	default:
		panic("output path not found")
	}
}

func (g Generator) parseTemplate(template string) string {
	result := strings.ReplaceAll(template, "{{name}}", strings.ToLower(g.Name))
	result = strings.ReplaceAll(result, "{{name_capitalized}}", strings.Title(g.Name))
	result = strings.ReplaceAll(result, "{{method_capitalized}}", strings.Title(g.Type))
	result = strings.ReplaceAll(result, "{{name_plural}}", utils.ToSnakeCase(g.NamePlural))
	result = strings.ReplaceAll(result, "{{project_name}}", config.ProjectName)

	return g.generateMethods(result)
}

func (g Generator) generateMethods(template string) string {
	var (
		filters        = make([]string, 0, len(g.Fields))
		fields         = make([]string, 0, len(g.Fields))
		fieldsQuery    = make([]string, 0, len(g.Fields))
		fieldsModel    = make([]string, 0, len(g.Fields))
		fieldsOptional = make([]string, 0, len(g.Fields))
		adaptInput     = make([]string, 0, len(g.Fields))
		adaptData      = make([]string, 0, len(g.Fields))
		adaptFilter    = make([]string, 0, len(g.Fields))
		adaptValues    = make([]string, 0, len(g.Fields))
		input          string
		output         string
	)

	for key, field := range g.Fields {
		pointer := "*"
		if field.IsRequired {
			pointer = ""
		}

		fields = append(fields, fmt.Sprintf("%s %s%s", key, pointer, field.Type))

		gormNull := ""
		if field.IsRequired {
			gormNull = "`gorm:\"not null\"`"
		}

		fieldsModel = append(fieldsModel, fmt.Sprintf("%s %s %s", key, field.Type, gormNull))

		fieldsQuery = append(fieldsQuery, fmt.Sprintf(`%s: ctx.Query("%s"),`, key, key))

		fieldsOptional = append(fieldsOptional, fmt.Sprintf("%s *%s", key, field.Type))
		filters = append(filters, fmt.Sprintf(`
		if filter.%s != nil {
			query = query.Where("%s = ?", filter.%s)
		}`, key, utils.ToSnakeCase(key), key))
	}

	if g.OutputName == "interface" {
		switch g.Type {
		case "Patch":
			template = fmt.Sprintf(`%s

			type PatchFilter struct {
				{{fields_optional}}
			}

			type PatchValues struct {
				{{fields}}
			}`, template)

		case "Paginate":
			template = fmt.Sprintf(`%s

			type PaginateFilter struct {
				{{fields_optional}}
			}

			type PaginateData struct {
				{{fields}}
			}

			type PaginateOutput struct {
				Data     []PaginateData
				MaxPages int
			}`, template)

		case "List":
			template = fmt.Sprintf(`%s

			type ListInput struct {
				{{fields_optional}}
			}

			type ListOutput struct {
				{{fields}}
			}`, template)

		case "GetByID":
			template = fmt.Sprintf(`%s

			type GetByIDOutput struct {
				{{fields}}
			}`, template)

		case "Delete":

		default:
			template = fmt.Sprintf(`%s
			
			type %sInput struct {
				{{fields}}
			}`, template, g.Type)
		}
	}

	switch g.Type {
	case "Create":
		input = "input CreateInput"
		output = "(string, error)"
		for k := range g.Fields {
			adaptInput = append(adaptInput, fmt.Sprintf("%s: input.%s,", k, k))
		}

	case "GetByID":
		input = "id string"
		output = "(*GetByIDOutput, error)"
		for k := range g.Fields {
			adaptData = append(adaptData, fmt.Sprintf("%s: data.%s,", k, k))
		}

	case "Delete":
		input = "id string"
		output = "(bool, error)"

	case "Patch":
		input = "filter PatchFilter, values PatchValues"
		output = "(bool, error)"
		for k := range g.Fields {
			adaptFilter = append(adaptFilter, fmt.Sprintf("%s: filter.%s,", k, k))
			adaptValues = append(adaptData, fmt.Sprintf("%s: values.%s,", k, k))
		}

	case "Paginate":
		input = "filter PaginateFilter, paginate database.PaginateInput"
		output = fmt.Sprintf("(*%sOutput, error)", g.Type)

	case "List":
		input = "input ListInput"
		output = "([]ListOutput, error)"

	default:
		input = fmt.Sprintf("input %sInput", g.Type)
		output = fmt.Sprintf("(%sOutput, error)", g.Type)
	}

	result := strings.ReplaceAll(template, "{{adapt_data}}", strings.Join(adaptData, "\n"))
	result = strings.ReplaceAll(result, "{{adapt_input}}", strings.Join(adaptInput, "\n"))
	result = strings.ReplaceAll(result, "{{adapt_filter}}", strings.Join(adaptFilter, "\n"))
	result = strings.ReplaceAll(result, "{{adapt_values}}", strings.Join(adaptValues, "\n"))
	result = strings.ReplaceAll(result, "{{filters}}", strings.Join(filters, "\n"))
	result = strings.ReplaceAll(result, "{{fields}}", strings.Join(fields, "\n"))
	result = strings.ReplaceAll(result, "{{fields_query}}", strings.Join(fieldsQuery, "\n"))
	result = strings.ReplaceAll(result, "{{fields_model}}", strings.Join(fieldsModel, "\n"))
	result = strings.ReplaceAll(result, "{{fields_optional}}", strings.Join(fieldsOptional, "\n"))
	result = strings.ReplaceAll(result, "{{method_input}}", input)
	result = strings.ReplaceAll(result, "{{method_output}}", output)

	return result
}

func GenerateAdapter(name string) {
	g := Generator{
		Name:         name,
		NamePlural:   "",
		Type:         "",
		TemplateName: config.AdapterTemplate,
		OutputName:   "adapter",
		Fields:       nil,
	}

	g.Generate()
}

func GenerateModel(name, namePlural string, fields map[string]Field) {
	model := Generator{
		Name:         name,
		NamePlural:   namePlural,
		Type:         "Model",
		TemplateName: config.ModelTemplate,
		OutputName:   fmt.Sprintf("%s_model", strings.ToLower(name)),
		Fields:       fields,
	}

	model.Generate()
}

func GenerateController(name, namePlural string, fields map[string]Field) {
	g := Generator{
		Name:         name,
		NamePlural:   namePlural,
		Type:         "Controller",
		TemplateName: config.ControllerTemplate,
		OutputName:   strings.ToLower(name),
		Fields:       fields,
	}

	g.Generate()
}

func GenerateUsecase(name, namePlural string, fields map[string]Field) {
	g := Generator{
		Name:         name,
		NamePlural:   namePlural,
		Type:         "Usecase",
		TemplateName: "usecase",
		OutputName:   strings.ToLower(name),
		Fields:       fields,
	}

	g.Generate()
}
