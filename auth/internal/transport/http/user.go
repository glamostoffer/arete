package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glamostoffer/arete/auth/internal/service/dto"
	"github.com/glamostoffer/arete/auth/pkg/errlist"
)

type handler struct {
	service service
}

func New(service service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) StartSignUp(c *gin.Context) {
	var req dto.StartSignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("%+v", req)

	res, err := h.service.StartSignUp(c.Request.Context(), req)
	if err != nil {
		c.JSON(errlist.GetErrStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (h *handler) ConfirmEmail(c *gin.Context) {
	var req dto.ConfirmEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.ConfirmEmail(c.Request.Context(), req)
	if err != nil {
		c.JSON(errlist.GetErrStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (h *handler) SignIn(c *gin.Context) {
	var req dto.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.SignIn(c.Request.Context(), req)
	if err != nil {
		c.JSON(errlist.GetErrStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (h *handler) VerifyCredentials(c *gin.Context) {
	var req dto.VerifyCredentialsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.VerifyCredentials(c.Request.Context(), req)
	if err != nil {
		c.JSON(errlist.GetErrStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (h *handler) RefreshSession(c *gin.Context) {
	var req dto.RefreshSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.RefreshSession(c.Request.Context(), req)
	if err != nil {
		c.JSON(errlist.GetErrStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (h *handler) GetUserInfo(c *gin.Context) {
	var req dto.GetUserInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.GetUserInfo(c.Request.Context(), req)
	if err != nil {
		c.JSON(errlist.GetErrStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
