package ginserver

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

func NewGinServer(temp *template.Template, staticRoutes map[string]string) *ginServer {
	r := gin.Default()

	if temp != nil {
		r.SetHTMLTemplate(temp)
	}

	for route, path := range staticRoutes {
		r.Static(route, path)
	}

	return &ginServer{
		engine: r,
	}
}

type ginServer struct {
	engine *gin.Engine
}

func (server *ginServer) AddRoute(route Route) {
	switch route.GetMethod() {
	case "GET":
		server.engine.GET(route.GetRoutePath(), handler(route))
	case "POST":
		server.engine.POST(route.GetRoutePath(), handler(route))
	case "PUT":
		server.engine.PUT(route.GetRoutePath(), handler(route))
	case "DELETE":
		server.engine.DELETE(route.GetRoutePath(), handler(route))
	default:
		log.Printf("invalid method: %s - %s", route.GetMethod(), route.GetRoutePath())
	}
}

func (server *ginServer) AddRoutes(routes []Route) {
	for _, route := range routes {
		server.AddRoute(route)
	}
}

func (server *ginServer) Run(port string) error {
	return server.engine.Run(port)
}
