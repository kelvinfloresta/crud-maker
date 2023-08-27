package utils

import (
	"fmt"
	"os"
	"strings"
)

var ProjectName = getProjectName()

func getProjectName() string {
	if _, err := os.Stat("go.mod"); err != nil {
		fmt.Println("unable to open the go.mod due:")
		panic(err)
	}

	file, err := os.ReadFile("go.mod")
	PanicIfError(err)

	firstLine := strings.Split(string(file), "\n")[0]
	projectName := strings.ReplaceAll(firstLine, "module ", "")
	projectName = strings.ReplaceAll(projectName, "\n", "")
	projectName = strings.ReplaceAll(projectName, "\r", "")
	return projectName
}
