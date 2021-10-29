package server

import (
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/lffranca/gonga/pkg/gkong"
	"sync"
)

func New(gatewayDatabase GatewayDatabase) (*Server, error) {
	if gatewayDatabase == nil {
		return nil, errors.New("gateway database is required")
	}

	server := new(Server)
	server.common.Server = server
	server.proxyMutex = &sync.RWMutex{}
	server.service = (*serviceRoute)(&server.common)
	server.gateway = (*gatewayRoute)(&server.common)
	server.gatewayDatabase = gatewayDatabase

	server.app = gin.Default()

	server.routes()

	return server, nil
}

type service struct {
	Server *Server
}

type Server struct {
	common          service
	app             *gin.Engine
	proxyMutex      *sync.RWMutex
	proxy           *gkong.Client
	service         *serviceRoute
	gateway         *gatewayRoute
	gatewayDatabase GatewayDatabase
}

func (pkg *Server) Run(addr ...string) error {
	return pkg.app.Run(addr...)
}

func (pkg *Server) routes() {
	v1 := pkg.app.Group("/api/v1")
	{
		v1.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
			AllowHeaders:     []string{"Accept", "Origin", "Content-Length", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
		}))

		v1.Use(gzip.Gzip(gzip.DefaultCompression))

		gatewayRoute := v1.Group("/gateway")
		{
			gatewayRoute.POST("", pkg.gateway.createPOST)
			gatewayRoute.GET("/:gateway", pkg.gateway.getGET)
			gatewayRoute.GET("/activate/:gateway", pkg.gateway.activateGET)
			gatewayRoute.GET("/disable/:gateway", pkg.gateway.disableGET)
		}

		serviceRoute := v1.Group("/service")
		{
			serviceRoute.GET("/all", pkg.service.listAllGET)
		}
	}
}
