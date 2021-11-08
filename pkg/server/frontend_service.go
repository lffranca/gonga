package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FrontendService service

func (pkg *FrontendService) indexGET(c *gin.Context) {
	c.HTML(http.StatusOK, "service/index.tpl", gin.H{})
}
