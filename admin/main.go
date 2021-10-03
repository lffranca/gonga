package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lffranca/gonga/admin/util"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	temp := template.Must(
		template.New("").
			Funcs(util.NewTemplateFunc()).
			ParseGlob("view/template/**/*.tpl"))
	r.SetHTMLTemplate(temp)

	r.Static("/js", os.Getenv("STATIC_JS_PATH"))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "service/index.tpl", gin.H{
			"content": "Service",
		})
	})

	if err := r.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT"))); err != nil {
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
