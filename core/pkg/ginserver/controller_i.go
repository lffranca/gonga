package ginserver

import "context"

type Controller interface {
	GetBindQuery() interface{}
	GetBindHeader() interface{}
	GetBindBody() interface{}
	GetRoutePath() string
	GetMethod() string
	Handler(ctx context.Context) (body []byte, contentType *string, err error)
}
