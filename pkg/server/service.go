package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type serviceRoute service

func (server *serviceRoute) listAllGET(c *gin.Context) {
	server.Server.proxyMutex.RLock()
	defer server.Server.proxyMutex.RUnlock()

	if server.Server.proxy == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "proxy is not selected"})
		return
	}

	items, err := server.Server.proxy.Service.ListAll(c.Request.Context())
	if err != nil {
		log.Println("server.Server.proxy.Service.ListAll: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}
