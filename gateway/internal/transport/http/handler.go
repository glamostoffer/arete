package http

import "github.com/gin-gonic/gin"

type handler struct {
	auth     auth
	learning learning
}

func New(auth auth, learning learning) *handler {
	return &handler{
		auth:     auth,
		learning: learning,
	}
}

func (h *handler) SetupRoutes(e *gin.Engine) {
	h.SetupAuthRoutes(e)
	h.SetupLearningRoutes(e)
}
