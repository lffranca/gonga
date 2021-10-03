package util

import (
	"html/template"
	"time"
)

func NewTemplateFunc() template.FuncMap {
	return template.FuncMap{
		"getUnixTime": getUnixTime,
	}
}

func getUnixTime() int64 {
	return time.Now().Unix()
}
