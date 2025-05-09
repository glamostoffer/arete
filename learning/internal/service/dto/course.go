package dto

import (
	"github.com/glamostoffer/arete/learning/internal/domain"
	v1 "github.com/glamostoffer/arete/learning/pkg/api/grpc/v1"
)

type GetCoursesListRequest struct {
	UserID     int64
	Categories []string
	Limit      int64
	Offset     int64
}

type GetCoursesListResponse struct {
	Courses []domain.Course
	HasNext bool
}

func (r *GetCoursesListResponse) ToProto() *v1.GetCoursesListResponse {
	pbCoursesList := make([]*v1.Course, 0, len(r.Courses))

	for _, course := range r.Courses {
		pbCourse := &v1.Course{
			Id:          course.ID,
			Title:       course.Title,
			Description: course.Description,
			Duration:    course.Duration,
			Difficulty:  course.Difficulty,
			Category:    course.Category,
			ImageURL:    course.ImageURL,
			IsEnrolled:  course.IsEnrolled,
		}

		pbCoursesList = append(pbCoursesList, pbCourse)
	}

	return &v1.GetCoursesListResponse{
		Courses: pbCoursesList,
		HasNext: r.HasNext,
	}
}

type GetCourseCategoriesRequest struct {
}
type GetCourseCategoriesResponse struct {
	Categories []string
}

type EnrollToCourseRequest struct {
	UserID   int64
	CourseID int64
}
type EnrollToCourseResponse struct {
}

type GetUserCoursesRequest struct {
	UserID int64
}
type GetUserCoursesResponse struct {
	Courses []domain.Course
}

func (r *GetUserCoursesResponse) ToProto() *v1.GetUserCoursesResponse {
	pbCoursesList := make([]*v1.Course, 0, len(r.Courses))

	for _, course := range r.Courses {
		pbCourse := &v1.Course{
			Id:          course.ID,
			Title:       course.Title,
			Description: course.Description,
			Duration:    course.Duration,
			Difficulty:  course.Difficulty,
			Category:    course.Category,
			ImageURL:    course.ImageURL,
			IsEnrolled:  course.IsEnrolled,
		}

		pbCoursesList = append(pbCoursesList, pbCourse)
	}

	return &v1.GetUserCoursesResponse{
		Courses: pbCoursesList,
	}
}
