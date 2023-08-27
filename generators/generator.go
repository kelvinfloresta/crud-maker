package generators

import (
	"crud-maker/config"
	"crud-maker/utils"
	"fmt"
	"os"
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

	g.createPath()
	utils.WriteTemplate(parsed, g.outputFile)
}

func (g *Generator) createPath() {
	filePath := strings.Split(g.outputFile, "/")
	onlyFolders := filePath[:len(filePath)-1]
	err := os.MkdirAll(strings.Join(onlyFolders, "/"), config.Permission)
	utils.PanicIfError(err)
}

func GenerateAdapter(name string) {
	g := Generator{
		name:         name,
		namePlural:   "",
		templateName: config.AdapterTemplate,
		outputFile:   fmt.Sprintf("adapters/gateways/%s_gateway/gorm_adapter.go", strings.ToLower(name)),
		fields:       nil,
	}

	g.Generate()
}

func GenerateModel(name, namePlural string, fields map[string]Field) {
	model := Generator{
		name:         name,
		namePlural:   namePlural,
		templateName: config.ModelTemplate,
		outputFile:   fmt.Sprintf("frameworks/database/gorm/models/%s_model.go", strings.ToLower(name)),
		fields:       fields,
	}

	model.Generate()
}

func GenerateController(name, namePlural string, fields map[string]Field) {

	g := Generator{
		name:         name,
		namePlural:   namePlural,
		templateName: config.ControllerTemplate,
		outputFile:   fmt.Sprintf("frameworks/http/fiber/controllers/%s_controller/controller_struct.go", strings.ToLower(name)),
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
