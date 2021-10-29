package main

import (
	"fmt"
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

	app, err := server.New(db.Gateway)
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
