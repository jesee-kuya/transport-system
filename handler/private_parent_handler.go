package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
)

func privateParentErrStatus(err error) int {
	if errors.Is(err, domain.ErrNotFound) {
		return http.StatusNotFound
	}
	if errors.Is(err, domain.ErrKYCAlreadySubmitted) || errors.Is(err, domain.ErrMatchAlreadyExists) || errors.Is(err, domain.ErrSchoolAlreadyConnected) {
		return http.StatusConflict
	}
	return http.StatusInternalServerError
}

func (t *Transport) CollectKYC(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req domain.CollectKYCRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parent, err := t.PrivateParentService.CollectKYC(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, parent)
}

func (t *Transport) GetPrivateProfile(c gin.Context) {
	claims, ok := t.adminClaims(&c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	parent, err := t.PrivateParentService.GetPrivateProfile(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, parent)
}

func (t *Transport) AddMyChild(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req domain.AddPrivateChildRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	child, err := t.PrivateParentService.AddMyChild(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, child)
}

func (t *Transport) GetMyChildren(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	children, err := t.PrivateParentService.GetMyChildren(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, children)
}

func (t *Transport) EditMyPrivateProfile(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req domain.EditPrivateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parent, err := t.PrivateParentService.EditMyPrivateProfile(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, parent)
}

func (t *Transport) EditMyChild(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	childID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid child id"})
		return
	}
	var req domain.EditPrivateChildRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	child, err := t.PrivateParentService.EditMyChild(c.Request.Context(), claims.UserID, childID, &req)
	if err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, child)
}

func (t *Transport) DeleteMyAccount(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	if err := t.PrivateParentService.DeleteMyAccount(c.Request.Context(), claims.UserID); err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "account deleted"})
}

func (t *Transport) MatchWithDriver(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req domain.MatchWithDriverRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	match, err := t.PrivateParentService.MatchWithDriver(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, match)
}

func (t *Transport) ConnectWithSchool(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req domain.ConnectWithSchoolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conn, err := t.PrivateParentService.ConnectWithSchool(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, conn)
}

func (t *Transport) GetSchools(c *gin.Context) {
	schools, err := t.PrivateParentService.GetSchools(c.Request.Context())
	if err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schools)
}

func (t *Transport) SearchSchools(c *gin.Context) {
	q := c.Query("q")
	if q == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query parameter 'q' is required"})
		return
	}
	schools, err := t.PrivateParentService.SearchSchools(c.Request.Context(), q)
	if err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schools)
}

func (t *Transport) FilterSchools(c *gin.Context) {
	address := c.Query("address")
	schools, err := t.PrivateParentService.FilterSchools(c.Request.Context(), address)
	if err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schools)
}

func (t *Transport) GetSchool(c *gin.Context) {
	schoolID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid school id"})
		return
	}
	school, err := t.PrivateParentService.GetSchool(c.Request.Context(), schoolID)
	if err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, school)
}

func (t *Transport) ReceiveStudent(c *gin.Context) {
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
	trip, err := t.PrivateParentService.ReceiveStudent(c.Request.Context(), claims.UserID, tripID)
	if err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trip)
}

func (t *Transport) ConfirmStudentBoarding(c *gin.Context) {
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
	trip, err := t.PrivateParentService.ConfirmStudentBoarding(c.Request.Context(), claims.UserID, tripID)
	if err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trip)
}

func (t *Transport) GetTrips(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	trips, err := t.PrivateParentService.GetTrips(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trips)
}

func (t *Transport) TrackTrip(c *gin.Context) {
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
	trip, err := t.PrivateParentService.TrackTrip(c.Request.Context(), claims.UserID, tripID)
	if err != nil {
		c.JSON(privateParentErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trip)
}
