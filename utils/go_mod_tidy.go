package utils

import (
	"fmt"
	"os/exec"
)

func GoModTidy() {
	fmt.Println("Running: go mod tidy")
	cmd := exec.Command("go", "mod", "tidy")
	PanicIfError(cmd.Run())
}
