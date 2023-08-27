package config

import (
	"io/fs"
)

var (
	TemplatePath       = "generators/templates/"
	AdapterTemplate    = "gateway_gorm_adapter"
	ModelTemplate      = "gorm_model"
	ControllerTemplate = "controller_fiber"
	Permission         = fs.FileMode(0777)
)
