package main

import (
	"fmt"
	"github.com/lffranca/gonga/domain/route"
	"github.com/lffranca/gonga/domain/service"
	"github.com/lffranca/gonga/pkg/gkong"
	"github.com/lffranca/gonga/pkg/gmongo"
	"github.com/lffranca/gonga/pkg/server"
	"log"
	"os"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	db, err := gmongo.New(&mongoURI)
	if err != nil {
		log.Panicln(err)
	}

	kongClient, err := gkong.New()
	if err != nil {
		log.Panicln(err)
	}

	serviceMod, err := service.New(&service.Options{
		GatewayRepository: db.Gateway,
		ServiceRepository: kongClient.Service,
	})

	routeMod, err := route.New(&route.Options{
		GatewayRepository:      db.Gateway,
		RouteGatewayRepository: kongClient.Route,
	})

	templatePath := os.Getenv("TEMPLATE_PATH")
	staticJSPath := os.Getenv("STATIC_JS_PATH")
	app, err := server.New(&server.Options{
		TemplatePath:      &templatePath,
		StaticJSPath:      &staticJSPath,
		GatewayRepository: db.Gateway,
		ServiceRepository: serviceMod,
		RouteRepository:   routeMod,
	})
	if err != nil {
		log.Panicln(err)
	}

	if err := app.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT"))); err != nil {
		log.Panicln(err)
	}
}

func init() {
	for _, envVar := range []string{
		"API_PORT",
		"MONGO_URI",
		"TEMPLATE_PATH",
		"STATIC_JS_PATH",
	} {
		if _, ok := os.LookupEnv(envVar); !ok {
			log.Panicf("Required enviroment variable not set: %s\n", envVar)
		}
	}
}
