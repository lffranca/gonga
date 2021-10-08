package ginserver

import "context"

type Route interface {
	GetBindQuery() interface{}
	GetBindHeader() interface{}
	GetBindBody() interface{}
	GetTemplate() *string
	GetRedirect() *string
	GetRoutePath() string
	GetMethod() string
	Handler(ctx context.Context) (body interface{}, contentType *string, err error)
}
