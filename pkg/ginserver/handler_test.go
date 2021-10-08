package ginserver

//func newRoutes() []*RouteServer {
//	redirect := "/app"
//	temp := "service/index.tpl"
//
//	return []*RouteServer{
//		{
//			Method:   "GET",
//			Path:     "/",
//			Redirect: &redirect,
//		},
//		{
//			Path:     "/app/*any",
//			Method:   "GET",
//			Template: &temp,
//		},
//	}
//}
//
//func newTemplateFunc() template.FuncMap {
//	return template.FuncMap{
//		"getUnixTime": func() int64 {
//			return time.Now().Unix()
//		},
//	}
//}
//
//func TestNewGinServer(t *testing.T) {
//	pattern := "/home/lucas/go/src/github.com/lffranca/gonga/admin/view/template/**/*.tpl"
//	temp := template.Must(
//		template.New("").
//			Funcs(newTemplateFunc()).
//			ParseGlob(pattern))
//
//	server := NewGinServer(temp, nil)
//
//	for _, route := range newRoutes() {
//		server.AddRoute(route)
//	}
//
//	if err := server.Run(fmt.Sprintf(":%s", "8080")); err != nil {
//		log.Panicf("error server run: %s", err)
//	}
//}
//
//type RouteServer struct {
//	Query       interface{}
//	Header      interface{}
//	Body        interface{}
//	Path        string
//	Method      string
//	Template    *string
//	Redirect    *string
//	HandlerFunc func(ctx context.Context) (body interface{}, contentType *string, err error)
//}
//
//func (route *RouteServer) GetBindQuery() interface{} {
//	return route.Query
//}
//
//func (route *RouteServer) GetBindHeader() interface{} {
//	return route.Header
//}
//
//func (route *RouteServer) GetBindBody() interface{} {
//	return route.Body
//}
//
//func (route *RouteServer) GetRoutePath() string {
//	return route.Path
//}
//
//func (route *RouteServer) GetMethod() string {
//	return route.Method
//}
//
//func (route *RouteServer) GetTemplate() *string {
//	return route.Template
//}
//
//func (route *RouteServer) GetRedirect() *string {
//	return route.Redirect
//}
//
//func (route *RouteServer) Handler(ctx context.Context) (body interface{}, contentType *string, err error) {
//	return route.HandlerFunc(ctx)
//}
