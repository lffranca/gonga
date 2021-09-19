package kong

type HTTPMethod string

const (
	HTTPMethodGET     HTTPMethod = "GET"
	HTTPMethodPOST    HTTPMethod = "POST"
	HTTPMethodPATCH   HTTPMethod = "PATCH"
	HTTPMethodPUT     HTTPMethod = "PUT"
	HTTPMethodDELETE  HTTPMethod = "DELETE"
	HTTPMethodOPTIONS HTTPMethod = "OPTIONS"
)
