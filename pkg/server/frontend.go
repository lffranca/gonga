package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type frontendRoute service

func (pkg *frontendRoute) indexGET(c *gin.Context) {
	c.HTML(http.StatusOK, "service/index.tpl", gin.H{})
}
