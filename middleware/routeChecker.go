package middleware

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

var allowedRoutes = map[string][]string{
	"/health":                              {"GET"},
	"/auth/signup":                         {"POST"},
	"/auth/login":                          {"POST"},
	"/auth/forgot-password":                {"POST"},
	"/auth/reset-password":                 {"POST"},
	"/auth/change-password":                {"POST"},
	"/school":                              {"POST"},
	"/school/students":                     {"POST", "GET"},
	"/school/students/filter":              {"GET"},
	"/school/students/search":              {"GET"},
	"/school/students/:id":                 {"GET", "PUT", "DELETE"},
	"/school/students/:id/manage":          {"PATCH"},
	"/school/students/:id/guardians":       {"GET", "POST"},
	"/school/buses":                        {"POST", "GET"},
	"/school/buses/filter":                 {"GET"},
	"/school/buses/search":                 {"GET"},
	"/school/buses/:id":                    {"GET", "PUT", "DELETE"},
	"/school/buses/:id/track":              {"GET"},
	"/school/drivers":                      {"POST", "GET"},
	"/school/drivers/filter":               {"GET"},
	"/school/drivers/search":               {"GET"},
	"/school/drivers/:id":                  {"GET", "PUT", "DELETE"},
	"/school/private-drivers":              {"POST", "GET"},
	"/school/private-drivers/:id":          {"GET"},
	"/school/private-trips":                {"GET"},
	"/school/private-trips/:id":            {"GET"},
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
