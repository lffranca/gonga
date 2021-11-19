package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lffranca/gonga/pkg/server/presenter"
	"log"
	"net/http"
)

type RouteService service

func (server *RouteService) listGET(c *gin.Context) {
	var gatewayString presenter.GatewayString
	if err := c.ShouldBindQuery(&gatewayString); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var optionQuery presenter.OptionQuery
	if err := c.ShouldBindQuery(&optionQuery); err != nil {
		log.Println("[WARNING] c.ShouldBindQuery(&optionQuery): ", err)
	}

	items, options, err := server.Server.routeRepository.List(
		c.Request.Context(),
		gatewayString.ID,
		optionQuery.Size,
		optionQuery.Offset,
		optionQuery.Tags,
		optionQuery.MatchAllTags)
	if err != nil {
		log.Println("proxy.Route.List: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    items,
		"options": options,
	})
}
