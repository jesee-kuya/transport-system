package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func guardianErrStatus(err error) int {
	if errors.Is(err, domain.ErrNotFound) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}

func (t *Transport) GetMyStudents(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	students, err := t.GuardianService.GetMyStudents(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(guardianErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (t *Transport) GetMyStudent(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	studentID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}
	student, err := t.GuardianService.GetMyStudent(c.Request.Context(), claims.UserID, studentID)
	if err != nil {
		c.JSON(guardianErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (t *Transport) TrackMyStudent(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	studentID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}
	trip, err := t.GuardianService.TrackMyStudent(c.Request.Context(), claims.UserID, studentID)
	if err != nil {
		c.JSON(guardianErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trip)
}

func (t *Transport) GetMyProfile(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	guardian, err := t.GuardianService.GetMyProfile(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(guardianErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, guardian)
}

func (t *Transport) EditMyProfile(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req domain.EditGuardianProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	guardian, err := t.GuardianService.EditMyProfile(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		c.JSON(guardianErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, guardian)
}
