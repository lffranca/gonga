package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/lffranca/gonga/admin/controller/app"
	"github.com/lffranca/gonga/admin/util"
	"github.com/lffranca/gonga/core/pkg/ginserver"
	"html/template"
	"log"
	"os"
)

func main() {
	temp := template.Must(
		template.New("").
			Funcs(util.NewTemplateFunc()).
			ParseGlob("view/template/**/*.tpl"))

	server := ginserver.NewGinServer(temp, map[string]string{
		"/js": os.Getenv("STATIC_JS_PATH"),
	})

	for _, route := range app.NewRoutes() {
		server.AddRoute(route)
	}

	if err := server.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT"))); err != nil {
		log.Panicf("error server run: %s", err)
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("[INFO] Ignore .env")
	}

	for _, envVar := range []string{
		"API_PORT",
		"STATIC_JS_PATH",
	} {
		if _, ok := os.LookupEnv(envVar); !ok {
			log.Panicf("Required enviroment variable not set: %s\n", envVar)
		}
	}
}
