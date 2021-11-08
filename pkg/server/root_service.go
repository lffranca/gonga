package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RootService service

func (pkg *RootService) redirectToPathGET(path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, path)
	}
}
