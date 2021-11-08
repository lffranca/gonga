package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lffranca/gonga/pkg/gkong"
	"github.com/lffranca/gonga/pkg/server/presenter"
	"log"
	"net/http"
)

type ServiceService service

func (server *ServiceService) listGET(c *gin.Context) {
	proxy := c.MustGet("proxy").(*gkong.Client)

	var optionQuery presenter.OptionQuery
	if err := c.ShouldBindQuery(&optionQuery); err != nil {
		log.Println("[WARNING] c.ShouldBindQuery(&optionQuery): ", err)
	}

	items, options, err := proxy.Service.List(
		c.Request.Context(),
		optionQuery.Size,
		optionQuery.Offset,
		optionQuery.Tags,
		optionQuery.MatchAllTags)
	if err != nil {
		log.Println("server.Server.proxy.Service.ListAll: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    items,
		"options": options,
	})
}

func (server *ServiceService) listAllGET(c *gin.Context) {
	proxy := c.MustGet("proxy").(*gkong.Client)

	items, err := proxy.Service.ListAll(c.Request.Context())
	if err != nil {
		log.Println("server.Server.proxy.Service.ListAll: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}
