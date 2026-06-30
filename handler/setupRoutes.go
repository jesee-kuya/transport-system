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

	parent := router.Group("/parent")
	parent.Use(t.Middleware.AuthMiddleware())
	{
		parent.POST("/kyc", t.CollectKYC)
		parent.GET("/profile", func(c *gin.Context) { t.GetPrivateProfile(*c) })
		parent.PUT("/profile", t.EditMyPrivateProfile)
		parent.DELETE("/account", t.DeleteMyAccount)

		parent.POST("/children", t.AddMyChild)
		parent.GET("/children", t.GetMyChildren)
		parent.PUT("/children/:id", t.EditMyChild)

		parent.POST("/drivers/match", t.MatchWithDriver)
		parent.POST("/schools/connect", t.ConnectWithSchool)

		// Static paths before param paths
		parent.GET("/schools", t.GetSchools)
		parent.GET("/schools/search", t.SearchSchools)
		parent.GET("/schools/filter", t.FilterSchools)
		parent.GET("/schools/:id", t.GetSchool)

		parent.GET("/trips", t.GetTrips)
		parent.GET("/trips/:id/track", t.TrackTrip)
		parent.PATCH("/trips/:id/boarding", t.ConfirmStudentBoarding)
		parent.PATCH("/trips/:id/receive", t.ReceiveStudent)
	}

	privateDriver := router.Group("/private-driver")
	privateDriver.Use(t.Middleware.AuthMiddleware())
	{
		privateDriver.POST("/kyc", t.KYCDriver)
		privateDriver.PATCH("/matches/:id", t.MatchWithParent)
		privateDriver.POST("/trips", t.StartTrip)
		privateDriver.PATCH("/trips/:id/end", t.EndTrip)
		privateDriver.POST("/trips/:id/students", t.OnboardPrivateStudent)
		privateDriver.PATCH("/trips/:id/status", t.UpdatePrivateTripStatus)
	}

	schoolDriver := router.Group("/school-driver")
	schoolDriver.Use(t.Middleware.AuthMiddleware())
	{
		schoolDriver.GET("/students/search", t.SearchStudents)
		schoolDriver.POST("/trips", t.StartTrip)
		schoolDriver.PATCH("/trips/:id/end", t.EndTrip)
		schoolDriver.PATCH("/trips/:id/status", t.UpdateTripStatus)
		schoolDriver.POST("/trips/:id/students", t.OnboardStudent)
		schoolDriver.GET("/trips/:id/students", t.ViewBoardedStudents)
	}

	guardian := router.Group("/guardian")
	guardian.Use(t.Middleware.AuthMiddleware())
	{
		guardian.GET("/profile", t.GetMyProfile)
		guardian.PUT("/profile", t.EditMyProfile)
		guardian.GET("/students", t.GetMyStudents)
		guardian.GET("/students/:id", t.GetMyStudent)
		guardian.GET("/students/:id/track", t.TrackMyStudent)
	}

	return router
}
