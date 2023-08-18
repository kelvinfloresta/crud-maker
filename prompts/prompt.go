package prompts

import (
	"crud-maker/generators"
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

func GetName() string {
	promptName := promptui.Prompt{
		Label: "Type the entity name",
		Validate: func(s string) error {
			if len(s) < 3 {
				return errors.New("should have at least 3 characters")
			}

			return nil
		},
	}

	name, err := promptName.Run()
	generators.CheckError(err)

	return name
}

func GetFieldName() string {
	prompt := promptui.Prompt{
		Label: "Type a field name",
		Validate: func(s string) error {
			if len(s) < 1 {
				return errors.New("should have at least 1 character")
			}

			return nil
		},
	}

	name, err := prompt.Run()
	generators.CheckError(err)

	return name
}

func GetPlural(name string) string {
	promptNamePlural := promptui.Prompt{
		Default: fmt.Sprint(name, "s"),
		Label:   "Type the entity plural",
		Validate: func(s string) error {
			if len(s) < 3 {
				return errors.New("should have at least 3 characters")
			}

			return nil
		},
	}

	namePlural, err := promptNamePlural.Run()
	generators.CheckError(err)

	return namePlural
}

var promptMethod = promptui.Select{
	Label: "Select a Method",
	Items: []string{
		"Create",
		"GetByID",
		"Patch",
		"List",
		"Paginate",
		"Delete",
	},
}

func GetMethods(selecteds []string) []string {
	if selecteds == nil {
		selecteds = []string{}
	}

	if len(selecteds) == 1 {
		promptMethod.Items = append([]string{"Finish"}, promptMethod.Items.([]string)...)
	}

	_, method, err := promptMethod.Run()
	generators.CheckError(err)

	if method == "Finish" {
		return selecteds
	}

	newMethods := []string{}
	for _, v := range promptMethod.Items.([]string) {
		if method == v {
			continue
		}
		newMethods = append(newMethods, v)
	}

	promptMethod.Items = newMethods

	return GetMethods(append(selecteds, method))
}

var promptFieldType = promptui.Select{
	Label: "Select a type",
	Items: []string{
		"string",
		"bool",
		"int",
		"int8",
		"int16",
		"int32",
		"uint",
		"uint8",
		"uint16",
		"uint32",
	},
}

func GetFields(fields map[string]generators.Field) map[string]generators.Field {
	if fields == nil {
		fields = make(map[string]generators.Field)
	}

	if len(fields) == 1 {
		promptFieldType.Items = append([]string{"Finish"}, promptFieldType.Items.([]string)...)
	}

	_, fieldType, err := promptFieldType.Run()
	generators.CheckError(err)

	if fieldType == "Finish" {
		return fields
	}

	fieldName := GetFieldName()
	isRequired := GetFieldRequired()

	fields[fieldName] = generators.Field{
		Type:       fieldType,
		IsRequired: isRequired,
	}

	return GetFields(fields)
}

func GetFieldRequired() bool {
	prompt := promptui.Select{
		Label: "Is required?",
		Items: []string{"Yes", "No"},
	}

	_, answer, err := prompt.Run()
	generators.CheckError(err)

	return answer == "Yes"
}
