package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lffranca/gonga/pkg/gkong"
	"github.com/lffranca/gonga/pkg/server/presenter"
	"log"
	"net/http"
)

type gatewayRoute service

func (server gatewayRoute) listGET(c *gin.Context) {
	items, err := server.Server.gatewayDatabase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, presenter.GatewaysFromEntities(items))
}

func (server gatewayRoute) getGET(c *gin.Context) {
	var gatewayURI presenter.GatewayURI
	if err := c.ShouldBindUri(&gatewayURI); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	item, err := server.Server.gatewayDatabase.Get(c.Request.Context(), gatewayURI.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, presenter.GatewayFromEntity(item))
}

func (server gatewayRoute) createPOST(c *gin.Context) {
	var item presenter.GatewayBody
	if err := c.ShouldBindJSON(&item); err != nil {
		log.Println("c.ShouldBindJSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := server.Server.gatewayDatabase.Save(c.Request.Context(), item.Entity())
	if err != nil {
		log.Println("server.Server.gatewayDatabase.Save: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, presenter.GatewayFromEntity(result))
}

func (server gatewayRoute) activateGET(c *gin.Context) {
	var gatewayURI presenter.GatewayURI
	if err := c.ShouldBindUri(&gatewayURI); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	item, err := server.Server.gatewayDatabase.Get(c.Request.Context(), gatewayURI.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	server.Server.proxyMutex.Lock()
	defer server.Server.proxyMutex.Unlock()

	proxy, err := gkong.New(item.Host, nil)
	if err != nil {
		log.Println("gkong.New: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	server.Server.proxy = proxy

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (server gatewayRoute) disableGET(c *gin.Context) {
	server.Server.proxyMutex.Lock()
	defer server.Server.proxyMutex.Unlock()

	server.Server.proxy = nil

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
