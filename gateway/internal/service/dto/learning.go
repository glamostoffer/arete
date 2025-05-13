package dto

import v1 "github.com/glamostoffer/arete/learning/pkg/api/grpc/v1"

type Course struct {
	ID          int64  `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Duration    string `json:"duration" db:"duration"`
	Difficulty  string `json:"difficulty" db:"difficulty"`
	Category    string `json:"category" db:"category"`
	ImageURL    string `json:"imageURL" db:"image_url"`
	IsEnrolled  bool   `json:"isEnrolled" db:"is_enrolled"`
}

type Lesson struct {
	ID          int64  `json:"id" db:"id"`
	CourseID    int64  `json:"courseID" db:"course_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Duration    string `json:"duration" db:"duration"`
	Content     string `json:"content" db:"content,omitempty"`
}

type GetCourseCategoriesRequest struct{}
type GetCourseCategoriesResponse struct {
	Categories []string `json:"items"`
}

type GetCoursesRequest struct {
	UserID     int64    `json:"-"`
	Categories []string `json:"categories"`
	Limit      int64    `json:"limit"`
	Offset     int64    `json:"offset"`
}
type GetCoursesResponse struct {
	Courses []Course `json:"items"`
	HasNext bool     `json:"hasNext"`
}

func (r GetCoursesResponse) FromProto(in *v1.GetCoursesListResponse) GetCoursesResponse {
	r.Courses = make([]Course, 0, len(in.GetCourses()))
	r.HasNext = in.GetHasNext()

	for _, pbCourse := range in.GetCourses() {
		r.Courses = append(r.Courses, Course{
			ID:          pbCourse.GetId(),
			Title:       pbCourse.GetTitle(),
			Description: pbCourse.GetCategory(),
			Duration:    pbCourse.GetDuration(),
			Difficulty:  pbCourse.GetDifficulty(),
			Category:    pbCourse.GetCategory(),
			ImageURL:    pbCourse.GetImageURL(),
			IsEnrolled:  pbCourse.GetIsEnrolled(),
		})
	}

	return r
}

type GetCourseLessonsRequest struct {
	CourseID int64 `json:"courseID"`
	Limit    int64 `json:"limit"`
	Offset   int64 `json:"offset"`
}
type GetCourseLessonsResponse struct {
	Lessons []Lesson `json:"items"`
	HasNext bool     `json:"hasNext"`
}

func (r GetCourseLessonsResponse) FromProto(in *v1.GetCourseLessonsResponse) GetCourseLessonsResponse {
	r.Lessons = make([]Lesson, 0, len(in.GetLessons()))
	r.HasNext = in.GetHasNext()

	for _, pbLesson := range in.GetLessons() {
		r.Lessons = append(r.Lessons, Lesson{
			ID:          pbLesson.GetId(),
			CourseID:    pbLesson.GetCourseID(),
			Title:       pbLesson.GetTitle(),
			Description: pbLesson.GetDescription(),
			Duration:    pbLesson.GetDuration(),
		})
	}

	return r
}

type GetLessonDetailsRequest struct {
	LessonID int64 `json:"lessonID"`
}
type GetLessonDetailsResponse struct {
	Lesson
}

func (r GetLessonDetailsResponse) FromProto(in *v1.GetLessonDetailsResponse) GetLessonDetailsResponse {
	pbLesson := in.GetLesson()

	r.Lesson = Lesson{
		ID:          pbLesson.GetId(),
		CourseID:    pbLesson.GetCourseID(),
		Title:       pbLesson.GetTitle(),
		Description: pbLesson.GetDescription(),
		Duration:    pbLesson.GetDuration(),
		Content:     pbLesson.GetContent(),
	}

	return r
}

type EnrollToCourseRequest struct {
	UserID   int64 `json:"-"`
	CourseID int64 `json:"courseID"`
}
type EnrollToCourseResponse struct {
}

type GetUserCoursesRequest struct {
	UserID int64 `json:"-"`
}
type GetUserCoursesResponse struct {
	Courses []Course
}

func (r GetUserCoursesResponse) FromProto(in *v1.GetUserCoursesResponse) GetUserCoursesResponse {
	r.Courses = make([]Course, 0, len(in.GetCourses()))

	for _, pbCourse := range in.GetCourses() {
		r.Courses = append(r.Courses, Course{
			ID:          pbCourse.GetId(),
			Title:       pbCourse.GetTitle(),
			Description: pbCourse.GetCategory(),
			Duration:    pbCourse.GetDuration(),
			Difficulty:  pbCourse.GetDifficulty(),
			Category:    pbCourse.GetCategory(),
			ImageURL:    pbCourse.GetImageURL(),
			IsEnrolled:  pbCourse.GetIsEnrolled(),
		})
	}

	return r
}
