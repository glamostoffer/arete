package http

import "github.com/gin-gonic/gin"

type handler struct {
	service service
}

func New(service service) *handler {
	return &handler{
		service,
	}
}

func (h *handler) SetupRoutes(e *gin.Engine) {
	h.SetupAuthRoutes(e)

	return
}
