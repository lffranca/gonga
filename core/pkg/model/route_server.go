package model

import "context"

type RouteServer struct {
	Query       interface{}
	Header      interface{}
	Body        interface{}
	Path        string
	Method      string
	Template    *string
	HandlerFunc func(ctx context.Context) (body interface{}, contentType *string, err error)
}

func (route *RouteServer) GetBindQuery() interface{} {
	return route.Query
}

func (route *RouteServer) GetBindHeader() interface{} {
	return route.Header
}

func (route *RouteServer) GetBindBody() interface{} {
	return route.Body
}

func (route *RouteServer) GetRoutePath() string {
	return route.Path
}

func (route *RouteServer) GetMethod() string {
	return route.Method
}

func (route *RouteServer) GetTemplate() *string {
	return route.Template
}

func (route *RouteServer) Handler(ctx context.Context) (body interface{}, contentType *string, err error) {
	return route.HandlerFunc(ctx)
}
