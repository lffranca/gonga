package ginserver

import (
	"log"

	"github.com/gin-gonic/gin"
)

func NewGinServer() *ginServer {
	r := gin.Default()

	return &ginServer{
		engine: r,
	}
}

type ginServer struct {
	engine *gin.Engine
}

func (server *ginServer) AddController(controller Controller) {
	switch controller.GetMethod() {
	case "GET":
		server.engine.GET(controller.GetRoutePath(), handler(controller))
	case "POST":
		server.engine.POST(controller.GetRoutePath(), handler(controller))
	case "PUT":
		server.engine.PUT(controller.GetRoutePath(), handler(controller))
	case "DELETE":
		server.engine.DELETE(controller.GetRoutePath(), handler(controller))
	default:
		log.Printf("invalid method: %s - %s", controller.GetMethod(), controller.GetRoutePath())
	}
}

func (server *ginServer) Run(port string) error {
	return server.engine.Run(port)
}
