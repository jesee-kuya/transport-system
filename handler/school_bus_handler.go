package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func (t *Transport) AddBus(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req domain.AddBusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bus, err := t.SchoolService.AddBus(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, bus)
}

func (t *Transport) GetBuses(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	buses, err := t.SchoolService.GetBuses(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, buses)
}

func (t *Transport) GetBus(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	busID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bus id"})
		return
	}
	bus, err := t.SchoolService.GetBus(c.Request.Context(), claims.UserID, busID)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bus)
}

func (t *Transport) EditBus(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	busID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bus id"})
		return
	}
	var req domain.EditBusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bus, err := t.SchoolService.EditBus(c.Request.Context(), claims.UserID, busID, &req)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bus)
}

func (t *Transport) FilterBuses(c *gin.Context) {
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
	buses, err := t.SchoolService.FilterBuses(c.Request.Context(), claims.UserID, isActive)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, buses)
}

func (t *Transport) SearchBuses(c *gin.Context) {
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
	buses, err := t.SchoolService.SearchBuses(c.Request.Context(), claims.UserID, q)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, buses)
}

func (t *Transport) RemoveBus(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	busID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bus id"})
		return
	}
	if err := t.SchoolService.RemoveBus(c.Request.Context(), claims.UserID, busID); err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "bus removed"})
}

func (t *Transport) TrackBus(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	busID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bus id"})
		return
	}
	trip, err := t.SchoolService.TrackBus(c.Request.Context(), claims.UserID, busID)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trip)
}
