package utils

func GetEndpoint(method string) string {
	switch method {
	case "GetByID", "Patch", "Delete":
		return `"/:id"`

	case "Create", "Paginate", "List":
		return `"/"`

	}

	return ""
}
