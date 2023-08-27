package config

import "io/fs"

var (
	ProjectName        = "crud-maker"
	TemplatePath       = "generators/templates/"
	AdapterTemplate    = "gateway_gorm_adapter"
	ModelTemplate      = "gorm_model"
	ControllerTemplate = "controller_fiber"
	Permission         = fs.FileMode(0777)
)
