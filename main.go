package main

import (
	"crud-maker/generators"
	"crud-maker/generators/create"
	delete_pkg "crud-maker/generators/delete"
	"crud-maker/generators/get_by_id"
	"crud-maker/generators/list"
	"crud-maker/generators/paginate"
	"crud-maker/generators/patch"
	"crud-maker/prompts"
	"crud-maker/utils"
	"fmt"
)

func main() {
	name := prompts.GetName()
	namePlural := prompts.GetPlural(name)
	fields := prompts.GetFields(nil)
	selecteds := prompts.GetMethods(nil)

	generators.GenerateModel(name, namePlural, fields)
	generators.GenerateController(name, namePlural, fields)
	generators.GenerateDatabaseAdapter(name)
	generators.GenerateUseCase(name, namePlural, fields)
	generators.GenerateHTTPAdapter(name, namePlural, fields)
	generators.GenerateFactory(name, namePlural, fields)
	generators.GenerateStatic("http_interface", "frameworks/http/interface.go")
	generators.GenerateStatic("parse_body_fiber", "frameworks/http/fiber_adapter/parser/parse_body.go")
	routeGenerator := generators.NewRoute(name, namePlural, fields)

	for _, method := range selecteds {
		routeGenerator.Generate(method)

		if method == "Create" {
			create.NewCase(name, namePlural, fields).Generate()
			create.NewGateway(name, namePlural, fields).Generate()
			create.NewGatewayAdapter(name, namePlural, fields).Generate()
			create.NewController(name, namePlural, fields).Generate()
			continue
		}

		if method == "List" {
			list.NewCase(name, namePlural, fields).Generate()
			list.NewGateway(name, namePlural, fields).Generate()
			list.NewGatewayAdapter(name, namePlural, fields).Generate()
			list.NewController(name, namePlural, fields).Generate()
			continue
		}

		if method == "Patch" {
			patch.NewCase(name, namePlural, fields).Generate()
			patch.NewGateway(name, namePlural, fields).Generate()
			patch.NewGatewayAdapter(name, namePlural, fields).Generate()
			patch.NewController(name, namePlural, fields).Generate()
			continue
		}

		if method == "Paginate" {
			paginate.NewCase(name, namePlural, fields).Generate()
			paginate.NewGateway(name, namePlural, fields).Generate()
			paginate.NewGatewayAdapter(name, namePlural, fields).Generate()
			paginate.NewController(name, namePlural, fields).Generate()
			continue
		}

		if method == "GetByID" {
			get_by_id.NewCase(name, namePlural, fields).Generate()
			get_by_id.NewGateway(name, namePlural, fields).Generate()
			get_by_id.NewGatewayAdapter(name, namePlural, fields).Generate()
			get_by_id.NewController(name, namePlural, fields).Generate()
			continue
		}

		if method == "Delete" {
			delete_pkg.NewCase(name, namePlural, fields).Generate()
			delete_pkg.NewGateway(name, namePlural, fields).Generate()
			delete_pkg.NewGatewayAdapter(name, namePlural, fields).Generate()
			delete_pkg.NewController(name, namePlural, fields).Generate()
			continue
		}
	}

	fmt.Print("\n")
	fmt.Println("Files generated!")
	fmt.Print("\n")

	utils.FormatCode()
	utils.GoModTidy()

	fmt.Println("\nDone!")
}
