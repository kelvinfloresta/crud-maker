package utils

import (
	"fmt"
	"os/exec"
)

func FormatCode() {
	fmt.Println("Running: go fmt ./...")
	cmd := exec.Command("go", "fmt", "./...")
	PanicIfError(cmd.Run())
}
