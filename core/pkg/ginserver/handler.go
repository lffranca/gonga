package ginserver

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(controller Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		queryData := controller.GetBindQuery()
		if queryData != nil {
			if err := c.ShouldBindQuery(queryData); err != nil {
				log.Printf("[ERROR] %v\n", err)
				c.JSON(http.StatusBadRequest, "Invalid URL parameters")
				c.Abort()
				return
			}
		}

		headerData := controller.GetBindHeader()
		if headerData != nil {
			if err := c.ShouldBindHeader(headerData); err != nil {
				log.Printf("[ERROR] %v\n", err)
				c.JSON(http.StatusBadRequest, "Invalid header parameters")
				c.Abort()
				return
			}
		}

		bodyData := controller.GetBindBody()
		if bodyData != nil {
			if err := c.ShouldBindJSON(bodyData); err != nil {
				log.Printf("[ERROR] %v\n", err)
				c.JSON(http.StatusBadRequest, "Invalid body parameters")
				c.Abort()
				return
			}
		}

		response, contentType, errResponse := controller.Handler(c.Request.Context())
		if errResponse != nil {
			log.Printf("[ERROR] %v\n", errResponse)
			c.JSON(http.StatusBadRequest, "Invalid request")
			c.Abort()
			return
		}

		if contentType == nil {
			c.Data(http.StatusOK, "application/json", response)
			return
		}

		c.Data(http.StatusOK, *contentType, response)
	}
}
