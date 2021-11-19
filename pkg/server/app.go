package server

import (
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/lffranca/gonga/pkg/server/function"
	"html/template"
)

func New(options *Options) (*Server, error) {
	if options == nil {
		return nil, errors.New("options param is required")
	}

	if err := options.validate(); err != nil {
		return nil, err
	}

	server := new(Server)
	server.common.Server = server
	server.staticJSPath = options.StaticJSPath
	server.root = (*RootService)(&server.common)
	server.frontend = (*FrontendService)(&server.common)
	server.service = (*ServiceService)(&server.common)
	server.gateway = (*GatewayService)(&server.common)
	server.route = (*RouteService)(&server.common)
	server.gatewayRepository = options.GatewayRepository
	server.serviceRepository = options.ServiceRepository
	server.routeRepository = options.RouteRepository

	server.app = gin.Default()

	server.app.Static("/js", *options.StaticJSPath)

	temp := template.Must(
		template.New("").
			Funcs(function.New()).
			ParseGlob(*options.TemplatePath))

	server.app.SetHTMLTemplate(temp)

	server.routes()

	return server, nil
}

type service struct {
	Server *Server
}

type Server struct {
	common            service
	staticJSPath      *string
	app               *gin.Engine
	root              *RootService
	frontend          *FrontendService
	service           *ServiceService
	gateway           *GatewayService
	route             *RouteService
	gatewayRepository GatewayRepository
	serviceRepository ServiceRepository
	routeRepository   RouteRepository
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

		serviceRoute := v1.Group("/service")
		{
			serviceRoute.GET("/", pkg.service.listGET)
		}

		routeRoute := v1.Group("/route")
		{
			routeRoute.GET("/", pkg.route.listGET)
		}
	}
}
