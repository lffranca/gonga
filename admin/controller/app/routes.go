package app

import "github.com/lffranca/gonga/core/pkg/model"

func NewRoutes() []*model.RouteServer {
	redirect := "/app"
	temp := "service/index.tpl"

	return []*model.RouteServer{
		{
			Method:   "GET",
			Path:     "/",
			Redirect: &redirect,
		},
		{
			Path:     "/app/*any",
			Method:   "GET",
			Template: &temp,
		},
	}
}
