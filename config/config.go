package config

import (
	"fmt"
	"io/fs"
)

var (
	HTTPFramework      = "fiber"
	DatabaseFramework  = "gorm"
	TemplatePath       = "generators/templates/"
	AdapterTemplate    = fmt.Sprintf("gateway_%s_adapter", DatabaseFramework)
	ModelTemplate      = fmt.Sprintf("%s_model", DatabaseFramework)
	ControllerTemplate = fmt.Sprintf("controller_%s", HTTPFramework)
	Permission         = fs.FileMode(0777)
)
