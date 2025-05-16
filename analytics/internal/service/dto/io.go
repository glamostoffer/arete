package dto

import (
	"github.com/glamostoffer/arete/analytics/internal/domain"
	v1 "github.com/glamostoffer/arete/analytics/pkg/api/grpc/v1"
)

type GetUserStatsRequest struct {
	UserID int64
}

type GetUserStatsResponse struct {
	domain.UserStats
}

func (r *GetUserStatsResponse) ToProto() *v1.GetUserStatsResponse {
	pbStats := &v1.GetUserStatsResponse{}

	pbStats.CourseRating = &v1.CourseRating{
		UserID:      r.CourseRating.UserID,
		CourseID:    r.CourseRating.CourseID,
		Rating:      r.CourseRating.Rating,
		Position:    r.CourseRating.Position,
		LastUpdated: r.CourseRating.LastUpdated.Unix(),
	}

	pbStats.Progress = &v1.UserCourseProgress{
		UserID:               r.UserCourseProgress.UserID,
		CourseID:             r.UserCourseProgress.CourseID,
		CompletionPercentage: r.UserCourseProgress.CompletionPercentage,
		LastUpdated:          r.UserCourseProgress.LastUpdated.Unix(),
		CompletedLessons:     r.UserCourseProgress.CompletedLessons,
		CompletedQuizzes:     r.UserCourseProgress.CompletedQuizzes,
		CompletedTasks:       r.UserCourseProgress.CompletedTasks,
	}

	pbStats.Rating = &v1.GlobalRating{
		UserID:      r.GlobalRating.UserID,
		Rating:      r.GlobalRating.Rating,
		Position:    r.GlobalRating.Position,
		LastUpdated: r.GlobalRating.LastUpdated.Unix(),
	}

	return pbStats
}
