package generators

import (
	"crud-maker/config"
	"crud-maker/utils"
	"fmt"
	"strings"
)

type Field struct {
	IsRequired bool
	Type       string
}

type Generator struct {
	name         string
	namePlural   string
	templateName string
	outputFile   string
	fields       map[string]Field
}

func (g *Generator) Generate() {
	template := utils.ReadTemplate(g.templateName)

	parsed := ParseTemplate(ParseTemplateInput{
		Template:     template,
		Name:         g.name,
		MethodName:   "",
		MethodInput:  "",
		MethodOutput: "",
		NamePlural:   g.namePlural,
		Fields:       g.fields,
	})

	utils.WriteTemplate(parsed, g.outputFile)
}

func GenerateDatabaseAdapter(name string) {
	g := Generator{
		name:         name,
		namePlural:   "",
		templateName: config.AdapterTemplate,
		outputFile:   fmt.Sprintf("frameworks/database/gateways/%s_gateway/%s_adapter.go", strings.ToLower(name), config.DatabaseFramework),
		fields:       nil,
	}

	g.Generate()
}

func GenerateModel(name, namePlural string, fields map[string]Field) {
	model := Generator{
		name:         name,
		namePlural:   namePlural,
		templateName: config.ModelTemplate,
		outputFile:   fmt.Sprintf("frameworks/database/%s_adapter/models/%s_model.go", config.DatabaseFramework, strings.ToLower(name)),
		fields:       fields,
	}

	model.Generate()
}

func GenerateController(name, namePlural string, fields map[string]Field) {

	g := Generator{
		name:         name,
		namePlural:   namePlural,
		templateName: config.ControllerTemplate,
		outputFile:   fmt.Sprintf("frameworks/http/controllers/%s_controller/controller_struct.go", strings.ToLower(name)),
		fields:       fields,
	}

	g.Generate()
}

func GenerateUseCase(name string, namePlural string, fields map[string]Field) {
	g := Generator{
		name:         name,
		namePlural:   namePlural,
		templateName: "usecase",
		outputFile:   fmt.Sprintf("usecases/%s_case/case_struct.go", strings.ToLower(name)),
		fields:       fields,
	}

	g.Generate()
}

func GenerateHTTPAdapter(name string, namePlural string, fields map[string]Field) {
	g := Generator{
		name:         name,
		namePlural:   namePlural,
		templateName: fmt.Sprintf("adapter_%s", config.HTTPFramework),
		outputFile:   fmt.Sprintf("frameworks/http/%s_adapter/%s_adapter.go", config.HTTPFramework, config.HTTPFramework),
		fields:       fields,
	}

	g.Generate()
}

func GenerateFactory(name string, namePlural string, fields map[string]Field) {
	g := Generator{
		name:         name,
		namePlural:   namePlural,
		templateName: "factory",
		outputFile:   fmt.Sprintf("factories/%s_controller.go", strings.ToLower(name)),
		fields:       fields,
	}

	g.Generate()
}
