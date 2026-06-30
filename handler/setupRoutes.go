package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *Transport) SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(t.Middleware.RouteChecker())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	auth := router.Group("/auth")
	{
		auth.POST("/signup", t.SignUp)
		auth.POST("/login", t.Login)
		auth.POST("/forgot-password", t.ForgotPassword)
		auth.POST("/reset-password", t.ResetPassword)
	}

	protected := router.Group("/auth")
	protected.Use(t.Middleware.AuthMiddleware())
	{
		protected.POST("/change-password", t.ChangePassword)
	}

	school := router.Group("/school")
	school.Use(t.Middleware.AuthMiddleware())
	{
		school.POST("", t.AddMySchool)

		// Students — static paths before param paths
		school.POST("/students", t.EnrollStudent)
		school.GET("/students", t.GetStudents)
		school.GET("/students/filter", t.FilterStudents)
		school.GET("/students/search", t.SearchStudents)
		school.GET("/students/:id", t.GetStudent)
		school.PUT("/students/:id", t.EditStudent)
		school.DELETE("/students/:id", t.RemoveStudent)
		school.PATCH("/students/:id/manage", t.ManageStudents)
		school.GET("/students/:id/guardians", t.GetStudentGuardians)
		school.POST("/students/:id/guardians", t.AddStudentGuardian)

		// Buses — static paths before param paths
		school.POST("/buses", t.AddBus)
		school.GET("/buses", t.GetBuses)
		school.GET("/buses/filter", t.FilterBuses)
		school.GET("/buses/search", t.SearchBuses)
		school.GET("/buses/:id", t.GetBus)
		school.PUT("/buses/:id", t.EditBus)
		school.DELETE("/buses/:id", t.RemoveBus)
		school.GET("/buses/:id/track", t.TrackBus)

		// School drivers — static paths before param paths
		school.POST("/drivers", t.AddDriver)
		school.GET("/drivers", t.GetDrivers)
		school.GET("/drivers/filter", t.FilterDrivers)
		school.GET("/drivers/search", t.SearchDrivers)
		school.GET("/drivers/:id", t.GetDriver)
		school.PUT("/drivers/:id", t.EditDriver)
		school.DELETE("/drivers/:id", t.RemoveDriver)

		// Private drivers and trips
		school.POST("/private-drivers", t.ConnectPrivateDriver)
		school.GET("/private-drivers", t.GetPrivateDrivers)
		school.GET("/private-drivers/:id", t.GetPrivateDriver)
		school.GET("/private-trips", t.TrackPrivateTrips)
		school.GET("/private-trips/:id", t.GetPrivateTripData)
	}

	return router
}
