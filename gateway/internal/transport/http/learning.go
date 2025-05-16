package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glamostoffer/arete/auth/pkg/errlist"
	"github.com/glamostoffer/arete/gateway/internal/service/dto"
)

func (h *handler) SetupLearningRoutes(e *gin.Engine) {
	api := e.Group("/api/v1/learn")
	{
		course := api.Group("/course")
		{
			course.GET("/categories", h.VerifyCredentials, h.GetCourseCategories)
			course.POST("", h.VerifyCredentials, h.GetCourses)
			course.POST("/enroll", h.VerifyCredentials, h.EnrollToCourse)
			course.GET("/user-courses", h.VerifyCredentials, h.GetUserCourses)
		}
		lesson := api.Group("/lesson")
		{
			lesson.POST("", h.VerifyCredentials, h.GetCourseLessons)
			lesson.POST("/details", h.VerifyCredentials, h.GetLessonDetails)
		}
	}
}

func (h *handler) GetCourseCategories(c *gin.Context) {
	_, exists := c.Get("userID") // todo make it const
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing userID"})
		return
	}

	res, err := h.learning.GetCourseCategories(c.Request.Context(), dto.GetCourseCategoriesRequest{})
	if err != nil {
		c.JSON(errlist.GetErrStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *handler) GetCourses(c *gin.Context) {
	userID, exists := c.Get("userID") // todo make it const
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing userID"})
		return
	}

	var req dto.GetCoursesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UserID = userID.(int64) // todo

	res, err := h.learning.GetCourses(c.Request.Context(), req)
	if err != nil {
		c.JSON(errlist.GetErrStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *handler) GetCourseLessons(c *gin.Context) {
	_, exists := c.Get("userID") // todo make it const
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing userID"})
		return
	}

	var req dto.GetCourseLessonsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.learning.GetCourseLessons(c.Request.Context(), req)
	if err != nil {
		c.JSON(errlist.GetErrStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *handler) GetLessonDetails(c *gin.Context) {
	_, exists := c.Get("userID") // todo make it const
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing userID"})
		return
	}

	var req dto.GetLessonDetailsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.learning.GetLessonDetails(c.Request.Context(), req)
	if err != nil {
		c.JSON(errlist.GetErrStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *handler) EnrollToCourse(c *gin.Context) {
	_, exists := c.Get("userID") // todo make it const
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing userID"})
		return
	}

	var req dto.EnrollToCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.learning.EnrollToCourse(c.Request.Context(), req)
	if err != nil {
		c.JSON(errlist.GetErrStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *handler) GetUserCourses(c *gin.Context) {
	userID, exists := c.Get("userID") // todo make it const
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing userID"})
		return
	}

	req := dto.GetUserCoursesRequest{
		UserID: userID.(int64),
	}

	res, err := h.learning.GetUserCourses(c.Request.Context(), req)
	if err != nil {
		c.JSON(errlist.GetErrStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
