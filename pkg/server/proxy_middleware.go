package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lffranca/gonga/pkg/gkong"
	"github.com/lffranca/gonga/pkg/server/presenter"
	"log"
	"net/http"
)

type ProxyMiddleware service

func (pkg *ProxyMiddleware) validate(c *gin.Context) {
	var gatewayString presenter.GatewayString
	if err := c.ShouldBindQuery(&gatewayString); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	gateway, err := pkg.Server.gatewayRepository.Get(c.Request.Context(), gatewayString.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	proxy, err := gkong.New(gateway.Host, nil)
	if err != nil {
		log.Println("gkong.New: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.Set("proxy", proxy)
	c.Next()
}
