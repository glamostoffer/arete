package server

import "github.com/gin-gonic/gin"

type httpHandler interface {
	StartSignUp(c *gin.Context)
	ConfirmEmail(c *gin.Context)
	SignIn(c *gin.Context)
	VerifyCredentials(c *gin.Context)
	RefreshSession(c *gin.Context)
	GetUserInfo(c *gin.Context)
}
