package ginserver

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(route Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		queryData := route.GetBindQuery()
		if queryData != nil {
			if err := c.ShouldBindQuery(queryData); err != nil {
				log.Printf("[ERROR] %v\n", err)
				c.JSON(http.StatusBadRequest, "Invalid URL parameters")
				c.Abort()
				return
			}
		}

		headerData := route.GetBindHeader()
		if headerData != nil {
			if err := c.ShouldBindHeader(headerData); err != nil {
				log.Printf("[ERROR] %v\n", err)
				c.JSON(http.StatusBadRequest, "Invalid header parameters")
				c.Abort()
				return
			}
		}

		bodyData := route.GetBindBody()
		if bodyData != nil {
			if err := c.ShouldBindJSON(bodyData); err != nil {
				log.Printf("[ERROR] %v\n", err)
				c.JSON(http.StatusBadRequest, "Invalid body parameters")
				c.Abort()
				return
			}
		}

		resp, contentType, errResponse := route.Handler(c.Request.Context())
		if errResponse != nil {
			log.Printf("[ERROR] %v\n", errResponse)
			c.JSON(http.StatusBadRequest, "Invalid request")
			c.Abort()
			return
		}

		temp := route.GetTemplate()
		if temp != nil {
			c.HTML(http.StatusOK, *temp, resp)
			return
		}

		response, ok := resp.([]byte)
		if !ok {
			log.Println("[ERROR] Response is invalid []byte")
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
