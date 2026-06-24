package middleware

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

var allowedRoutes = map[string][]string{
	"/health":       {"GET"},
	"/auth/signup":  {"POST"},
}

func (middleware *MiddlewareStruct) RouteChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedMethods, ok := allowedRoutes[c.FullPath()]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "route not found",
			})
			c.Abort()
			fmt.Println("not found")
			return
		}

		methodAllowed := slices.Contains(allowedMethods, c.Request.Method)

		if !methodAllowed {
			c.JSON(http.StatusMethodNotAllowed, gin.H{
				"error": "method not allowed",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
