package dto

import (
	"github.com/glamostoffer/arete/learning/internal/domain"
	v1 "github.com/glamostoffer/arete/learning/pkg/api/grpc/v1"
)

type GetCourseLessonsRequest struct {
	CourseID int64
	Limit    int64
	Offset   int64
}
type GetCourseLessonsResponse struct {
	Lessons []domain.Lesson
	HasNext bool
}

func (r *GetCourseLessonsResponse) ToProto() *v1.GetCourseLessonsResponse {
	pbLessonsList := make([]*v1.Lesson, 0, len(r.Lessons))

	for _, lesson := range r.Lessons {
		pbLesson := &v1.Lesson{
			Id:          lesson.ID,
			CourseID:    lesson.CourseID,
			Title:       lesson.Title,
			Description: lesson.Description,
			Duration:    lesson.Duration,
		}

		pbLessonsList = append(pbLessonsList, pbLesson)
	}

	return &v1.GetCourseLessonsResponse{
		Lessons: pbLessonsList,
		HasNext: r.HasNext,
	}
}

type GetLessonDetailsRequest struct {
	LessonID int64
}
type GetLessonDetailsResponse struct {
	Lesson domain.Lesson
}

func (r *GetLessonDetailsResponse) ToProto() *v1.GetLessonDetailsResponse {
	return &v1.GetLessonDetailsResponse{
		Lesson: &v1.Lesson{
			Id:          r.Lesson.ID,
			CourseID:    r.Lesson.CourseID,
			Title:       r.Lesson.Title,
			Description: r.Lesson.Description,
			Duration:    r.Lesson.Duration,
			Content:     &r.Lesson.Content,
		},
	}
}
