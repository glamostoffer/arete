package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/glamostoffer/arete/auth/pkg/errlist"
	"github.com/glamostoffer/arete/gateway/internal/service/dto"
)

func (h *handler) SetupAuthRoutes(e *gin.Engine) {
	api := e.Group("/api/v1/auth")
	{
		signUp := api.Group("/sign-up")
		{
			signUp.POST("/start", h.StartSignUp)
			signUp.POST("/finalize", h.ConfirmEmail)
		}
		signIn := api.Group("/sign-in")
		{
			signIn.POST("", h.SignIn)
		}
		session := api.Group("/session")
		{
			session.POST("/refresh", h.RefreshSession)
		}
		user := api.Group("/user")
		{
			user.GET("", h.VerifyCredentials, h.GetUserInfo)
		}
	}
}

func (h *handler) StartSignUp(c *gin.Context) {
	var req dto.StartSignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.auth.StartSignUp(c.Request.Context(), req)
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

	res, err := h.auth.ConfirmEmail(c.Request.Context(), req)
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

	res, err := h.auth.SignIn(c.Request.Context(), req)
	if err != nil {
		c.JSON(errlist.GetErrStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (h *handler) VerifyCredentials(c *gin.Context) {
	accessToken := c.GetHeader("X-Access-Token")
	if accessToken == "" {
		c.JSON(http.StatusTeapot, gin.H{"error": "missing X-Access-Token header"})
		return
	}

	res, err := h.auth.VerifyCredentials(c.Request.Context(), dto.VerifyCredentialsRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		c.JSON(http.StatusTeapot, gin.H{"error": err.Error()})
		return
	}

	c.Set("userID", res.UserID) // todo make it const

	return
}

func (h *handler) RefreshSession(c *gin.Context) {
	var req dto.RefreshSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.auth.RefreshSession(c.Request.Context(), req)
	if err != nil {
		c.JSON(errlist.GetErrStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (h *handler) GetUserInfo(c *gin.Context) {
	userID, exists := c.Get("userID") // todo make it const
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing userID"})
		return
	}

	res, err := h.auth.GetUserInfo(c.Request.Context(), dto.GetUserInfoRequest{
		UserID: userID.(int64), // todo опасно
	})
	if err != nil {
		c.JSON(errlist.GetErrStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
