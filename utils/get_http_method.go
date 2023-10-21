package utils

func GetHTTPMethod(method string) string {
	switch method {
	case "GetByID", "Paginate", "List":
		return "Get"

	case "Create":
		return "Post"

	case "Patch":
		return "Patch"

	case "Delete":
		return "Delete"
	}

	return ""
}
