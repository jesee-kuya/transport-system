package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jesee-kuya/transport-system/domain"
	"github.com/jesee-kuya/transport-system/middleware"
)

func (t *Transport) adminClaims(c *gin.Context) (*domain.Claims, bool) {
	claims, ok := c.MustGet(middleware.ClaimsKey).(*domain.Claims)
	return claims, ok
}

func schoolErrStatus(err error) int {
	if errors.Is(err, domain.ErrNotFound) {
		return http.StatusNotFound
	}
	if errors.Is(err, domain.ErrSchoolAlreadyExists) || errors.Is(err, domain.ErrPlateNumberInUse) || errors.Is(err, domain.ErrDriverAlreadyConnected) {
		return http.StatusConflict
	}
	return http.StatusInternalServerError
}

// School setup

func (t *Transport) AddMySchool(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req domain.AddSchoolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	school, err := t.SchoolService.AddMySchool(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, school)
}

// Students

func (t *Transport) EnrollStudent(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var req domain.EnrollStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	student, err := t.SchoolService.EnrollStudent(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, student)
}

func (t *Transport) GetStudents(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	students, err := t.SchoolService.GetStudents(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (t *Transport) GetStudent(c *gin.Context) {
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
	student, err := t.SchoolService.GetStudent(c.Request.Context(), claims.UserID, studentID)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (t *Transport) GetStudentGuardians(c *gin.Context) {
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
	guardians, err := t.SchoolService.GetStudentGuardians(c.Request.Context(), claims.UserID, studentID)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, guardians)
}

func (t *Transport) AddStudentGuardian(c *gin.Context) {
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
	var req domain.AddStudentGuardianRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sg, err := t.SchoolService.AddStudentGuardian(c.Request.Context(), claims.UserID, studentID, &req)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, sg)
}

func (t *Transport) ManageStudents(c *gin.Context) {
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
	var req domain.ManageStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	student, err := t.SchoolService.ManageStudent(c.Request.Context(), claims.UserID, studentID, &req)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (t *Transport) FilterStudents(c *gin.Context) {
	claims, ok := t.adminClaims(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	grade := c.Query("grade")
	var isActive *bool
	if v := c.Query("is_active"); v != "" {
		b := v == "true"
		isActive = &b
	}
	students, err := t.SchoolService.FilterStudents(c.Request.Context(), claims.UserID, grade, isActive)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (t *Transport) SearchStudents(c *gin.Context) {
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
	students, err := t.SchoolService.SearchStudents(c.Request.Context(), claims.UserID, q)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (t *Transport) RemoveStudent(c *gin.Context) {
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
	if err := t.SchoolService.RemoveStudent(c.Request.Context(), claims.UserID, studentID); err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "student removed"})
}

func (t *Transport) EditStudent(c *gin.Context) {
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
	var req domain.EditStudentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	student, err := t.SchoolService.EditStudent(c.Request.Context(), claims.UserID, studentID, &req)
	if err != nil {
		c.JSON(schoolErrStatus(err), gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}
