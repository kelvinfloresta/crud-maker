package main

import (
	"crud-maker/generators"
	"crud-maker/prompts"
	"crud-maker/utils"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	err := os.RemoveAll("output")
	generators.CheckError(err)

	name := prompts.GetName()
	namePlural := prompts.GetPlural(name)
	fields := prompts.GetFields(nil)
	selecteds := prompts.GetMethods(nil)

	generators.GenerateModel(name, namePlural, fields)
	generators.GenerateController(name, namePlural, fields)
	generators.GenerateUsecase(name, namePlural, fields)
	generators.GenerateAdapter(name)

	for _, method := range selecteds {
		g := generators.Generator{
			Name:         name,
			NamePlural:   namePlural,
			Type:         method,
			TemplateName: "gateway_interface",
			OutputName:   "interface",
			Fields:       fields,
		}

		g.Generate()

		methodSnakeCase := utils.ToSnakeCase(method)
		g = generators.Generator{
			Name:         name,
			NamePlural:   namePlural,
			Type:         method,
			TemplateName: fmt.Sprintf("gateway_gorm_%s", methodSnakeCase),
			OutputName:   fmt.Sprintf("gorm_%s", methodSnakeCase),
			Fields:       fields,
		}

		g.Generate()

		g = generators.Generator{
			Name:         name,
			NamePlural:   namePlural,
			Type:         method,
			TemplateName: fmt.Sprintf("controller_fiber_%s", methodSnakeCase),
			OutputName:   fmt.Sprintf("fiber_%s", methodSnakeCase),
			Fields:       fields,
		}

		g.Generate()

		g = generators.Generator{
			Name:         name,
			NamePlural:   namePlural,
			Type:         method,
			TemplateName: fmt.Sprintf("case_%s", methodSnakeCase),
			OutputName:   methodSnakeCase,
			Fields:       fields,
		}

		g.Generate()
	}

	fmt.Println("go fmt ./...")
	cmd := exec.Command("go", "fmt", "./...")
	err = cmd.Run()

	if err != nil {
		fmt.Println("\n** Error **")
		fmt.Println(err)
		return
	}

	fmt.Println("Done!")
}
