package server

import (
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/lffranca/gonga/pkg/server/function"
	"html/template"
)

func New(templatePath *string, staticJSPath *string, gatewayDatabase GatewayDatabase) (*Server, error) {
	if templatePath == nil || staticJSPath == nil {
		return nil, errors.New("required parameters not passed")
	}

	if gatewayDatabase == nil {
		return nil, errors.New("gateway database is required")
	}

	server := new(Server)
	server.common.Server = server
	server.staticJSPath = staticJSPath
	server.root = (*RootService)(&server.common)
	server.frontend = (*FrontendService)(&server.common)
	server.service = (*ServiceService)(&server.common)
	server.gateway = (*GatewayService)(&server.common)
	server.proxyMiddleware = (*ProxyMiddleware)(&server.common)
	server.gatewayDatabase = gatewayDatabase

	server.app = gin.Default()

	server.app.Static("/js", *staticJSPath)

	temp := template.Must(
		template.New("").
			Funcs(function.New()).
			ParseGlob(*templatePath))

	server.app.SetHTMLTemplate(temp)

	server.routes()

	return server, nil
}

type service struct {
	Server *Server
}

type Server struct {
	common          service
	staticJSPath    *string
	app             *gin.Engine
	root            *RootService
	frontend        *FrontendService
	service         *ServiceService
	gateway         *GatewayService
	proxyMiddleware *ProxyMiddleware
	gatewayDatabase GatewayDatabase
}

func (pkg *Server) Run(addr ...string) error {
	return pkg.app.Run(addr...)
}

func (pkg *Server) routes() {
	pkg.app.GET("", pkg.root.redirectToPathGET("/app"))

	app := pkg.app.Group("/app/*any")
	{
		app.GET("", pkg.frontend.indexGET)
	}

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
			gatewayRoute.GET("", pkg.gateway.listGET)
			gatewayRoute.GET("/:gateway", pkg.gateway.getGET)
		}

		proxyRequired := v1.Group("", pkg.proxyMiddleware.validate)
		{
			serviceRoute := proxyRequired.Group("/service")
			{
				serviceRoute.GET("/", pkg.service.listGET)
				serviceRoute.GET("/all", pkg.service.listAllGET)
			}
		}
	}
}
