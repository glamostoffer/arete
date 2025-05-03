package domain

type Lesson struct {
	ID          int64  `json:"id" db:"id"`
	CourseID    int64  `json:"courseID" db:"course_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Duration    string `json:"duration" db:"duration"`
	Content     string `json:"content" db:"content"`
}
