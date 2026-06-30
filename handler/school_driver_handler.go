package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func (t *Transport) AddDriver(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req domain.AddDriverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	driver, err := t.SchoolService.AddDriver(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, driver)
}

func (t *Transport) GetDrivers(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	drivers, err := t.SchoolService.GetDrivers(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, drivers)
}

func (t *Transport) GetDriver(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	driverID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid driver id"})
		return
	}
	driver, err := t.SchoolService.GetDriver(c.Request.Context(), claims.UserID, driverID)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, driver)
}

func (t *Transport) EditDriver(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	driverID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid driver id"})
		return
	}
	var req domain.EditDriverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	driver, err := t.SchoolService.EditDriver(c.Request.Context(), claims.UserID, driverID, &req)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, driver)
}

func (t *Transport) FilterDrivers(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var isActive *bool
	if v := c.Query("is_active"); v != "" {
		b := v == "true"
		isActive = &b
	}
	drivers, err := t.SchoolService.FilterDrivers(c.Request.Context(), claims.UserID, isActive)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, drivers)
}

func (t *Transport) SearchDrivers(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	q := c.Query("q")
	if q == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query parameter 'q' is required"})
		return
	}
	drivers, err := t.SchoolService.SearchDrivers(c.Request.Context(), claims.UserID, q)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, drivers)
}

func (t *Transport) RemoveDriver(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	driverID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid driver id"})
		return
	}
	if err := t.SchoolService.RemoveDriver(c.Request.Context(), claims.UserID, driverID); err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "driver removed"})
}

func (t *Transport) ConnectPrivateDriver(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req domain.ConnectPrivateDriverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := t.SchoolService.ConnectPrivateDriver(c.Request.Context(), claims.UserID, &req); err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "driver connected"})
}

func (t *Transport) GetPrivateDrivers(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	drivers, err := t.SchoolService.GetPrivateDrivers(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, drivers)
}

func (t *Transport) GetPrivateDriver(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	driverID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid driver id"})
		return
	}
	driver, err := t.SchoolService.GetPrivateDriver(c.Request.Context(), claims.UserID, driverID)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, driver)
}

func (t *Transport) TrackPrivateTrips(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	trips, err := t.SchoolService.TrackPrivateTrips(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trips)
}

func (t *Transport) GetPrivateTripData(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	tripID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid trip id"})
		return
	}
	trip, err := t.SchoolService.GetPrivateTripData(c.Request.Context(), claims.UserID, tripID)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trip)
}
