package ginserver

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handler(route Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		redirect := route.GetRedirect()
		if redirect != nil {
			c.Redirect(http.StatusPermanentRedirect, *redirect)
			return
		}

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

		var resp interface{}
		var contentType *string
		if route.Handler != nil {
			var err error
			resp, contentType, err = route.Handler(c.Request.Context())
			if err != nil {
				log.Printf("[ERROR] %v\n", err)
				c.JSON(http.StatusBadRequest, "Invalid request")
				c.Abort()
				return
			}
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
