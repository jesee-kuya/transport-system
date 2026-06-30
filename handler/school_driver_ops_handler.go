package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func schoolDriverErrStatus(err error) int {
	if errors.Is(err, domain.ErrNotFound) {
		return http.StatusNotFound
	}
	if errors.Is(err, domain.ErrTripAlreadyActive) || errors.Is(err, domain.ErrStudentAlreadyBoarded) {
		return http.StatusConflict
	}
	return http.StatusInternalServerError
}

func (t *Transport) StartTrip(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	switch claims.Role {
	case "school_driver":
		var req domain.StartTripRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		trip, err := t.SchoolDriverService.StartTrip(c.Request.Context(), claims.UserID, &req)
		if err != nil {
			c.JSON(schoolDriverErrStatus(err), gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, trip)
	case "private_driver":
		var req domain.StartPrivateTripRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		trip, err := t.PrivateDriverService.StartTrip(c.Request.Context(), claims.UserID, &req)
		if err != nil {
			c.JSON(schoolDriverErrStatus(err), gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, trip)
	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
	}
}

func (t *Transport) EndTrip(c *gin.Context) {
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
	switch claims.Role {
	case "school_driver":
		trip, err := t.SchoolDriverService.EndTrip(c.Request.Context(), claims.UserID, tripID)
		if err != nil {
			c.JSON(schoolDriverErrStatus(err), gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, trip)
	case "private_driver":
		trip, err := t.PrivateDriverService.EndTrip(c.Request.Context(), claims.UserID, tripID)
		if err != nil {
			c.JSON(schoolDriverErrStatus(err), gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, trip)
	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
	}
}

func (t *Transport) UpdateTripStatus(c *gin.Context) {
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
	var req domain.UpdateTripStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	trip, err := t.SchoolDriverService.UpdateTripStatus(c.Request.Context(), claims.UserID, tripID, &req)
	if err != nil {
		c.JSON(schoolDriverErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trip)
}

func (t *Transport) OnboardStudent(c *gin.Context) {
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
	var req domain.OnboardStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	studentID, err := uuid.Parse(req.StudentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}
	ts, err := t.SchoolDriverService.OnboardStudent(c.Request.Context(), claims.UserID, tripID, studentID)
	if err != nil {
		c.JSON(schoolDriverErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ts)
}

func (t *Transport) ViewBoardedStudents(c *gin.Context) {
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
	students, err := t.SchoolDriverService.ViewBoardedStudents(c.Request.Context(), claims.UserID, tripID)
	if err != nil {
		c.JSON(schoolDriverErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}
