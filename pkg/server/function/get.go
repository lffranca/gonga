package function

import "html/template"

func New() template.FuncMap {
	return template.FuncMap{
		"getUnixTime": getUnixTime,
	}
}
