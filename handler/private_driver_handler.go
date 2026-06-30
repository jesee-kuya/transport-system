package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func privateDriverErrStatus(err error) int {
	if errors.Is(err, domain.ErrNotFound) {
		return http.StatusNotFound
	}
	if errors.Is(err, domain.ErrKYCAlreadySubmitted) || errors.Is(err, domain.ErrTripAlreadyActive) || errors.Is(err, domain.ErrStudentAlreadyBoarded) {
		return http.StatusConflict
	}
	if errors.Is(err, domain.ErrMatchNotPending) {
		return http.StatusUnprocessableEntity
	}
	return http.StatusInternalServerError
}

func (t *Transport) KYCDriver(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req domain.KYCDriverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	driver, err := t.PrivateDriverService.KYCDriver(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		c.JSON(privateDriverErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, driver)
}

func (t *Transport) MatchWithParent(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	matchID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid match id"})
		return
	}
	var req domain.RespondToMatchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	match, err := t.PrivateDriverService.MatchWithParent(c.Request.Context(), claims.UserID, matchID, &req)
	if err != nil {
		c.JSON(privateDriverErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, match)
}

func (t *Transport) OnboardPrivateStudent(c *gin.Context) {
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
	var req domain.OnboardPrivateStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ptc, err := t.PrivateDriverService.OnboardPrivateStudent(c.Request.Context(), claims.UserID, tripID, &req)
	if err != nil {
		c.JSON(privateDriverErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ptc)
}

func (t *Transport) UpdatePrivateTripStatus(c *gin.Context) {
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
	var req domain.UpdatePrivateTripStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	trip, err := t.PrivateDriverService.UpdatePrivateTripStatus(c.Request.Context(), claims.UserID, tripID, &req)
	if err != nil {
		c.JSON(privateDriverErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trip)
}
