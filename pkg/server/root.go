package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type rootRoute service

func (pkg *rootRoute) redirectToPathGET(path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, path)
	}
}
