package server

import "github.com/gin-gonic/gin"

type httpHandler interface {
	SetupRoutes(e *gin.Engine)
}
