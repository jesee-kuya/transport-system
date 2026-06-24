package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesee-kuya/transport-system/domain"
)

func (t *Transport) SignUp(c *gin.Context) {
	var req domain.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := t.AuthService.SignUp(c.Request.Context(), &req)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrEmailInUse), errors.Is(err, domain.ErrUsernameInUse):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, resp)
}
