package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lffranca/gonga/pkg/server/presenter"
	"log"
	"net/http"
)

type GatewayService service

func (server GatewayService) listGET(c *gin.Context) {
	items, err := server.Server.gatewayRepository.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": presenter.GatewaysFromEntities(items),
	})
}

func (server GatewayService) getGET(c *gin.Context) {
	var gatewayURI presenter.GatewayURI
	if err := c.ShouldBindUri(&gatewayURI); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	item, err := server.Server.gatewayRepository.Get(c.Request.Context(), gatewayURI.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, presenter.GatewayFromEntity(item))
}

func (server GatewayService) createPOST(c *gin.Context) {
	var item presenter.GatewayBody
	if err := c.ShouldBindJSON(&item); err != nil {
		log.Println("c.ShouldBindJSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := server.Server.gatewayRepository.Save(c.Request.Context(), item.Entity())
	if err != nil {
		log.Println("server.Server.gatewayDatabase.Save: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, presenter.GatewayFromEntity(result))
}
